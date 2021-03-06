// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package gen

import "go.uber.org/thriftrw/compile"

// setGenerator generates logic to convert lists of arbitrary Thrift types to
// and from ValueLists.
type setGenerator struct {
	hasReaders
	hasLazyLists
}

// ValueList generates a new ValueList type alias for the given set.
//
// The following is generated:
//
// 	type $valueListName map[$valueType]struct{}
//
// 	func (v $valueListName) ForEach(f func(wire.Value) error) error { ... }
//
// 	func (v $valueListName) Close() { ... }
//
// And $valueListName is returned. This may be used where a ValueList of the
// given type is expected.
func (s *setGenerator) ValueList(g Generator, spec *compile.SetSpec) (string, error) {
	// TODO(abg): Unhashable types
	name := "_" + valueName(spec) + "_ValueList"
	if s.HasLazyList(name) {
		return name, nil
	}

	err := g.DeclareFromTemplate(
		`
			<$wire := import "go.uber.org/thriftrw/wire">
			type <.Name> <typeReference .Spec>

			<$v := newVar "v">
			<$x := newVar "x">
			<$f := newVar "f">
			<$w := newVar "w">
			func (<$v> <.Name>) ForEach(<$f> func(<$wire>.Value) error) error {
				<if isHashable .Spec.ValueSpec>
					for <$x> := range <$v> {
				<else>
					for _, <$x> := range <$v> {
				<end>
						<if not (isPrimitiveType .Spec.ValueSpec)>
							if <$x> == nil {
								return <import "fmt">.Errorf("invalid set item: value is nil")
							}
						<end>

						<$w>, err := <toWire .Spec.ValueSpec $x>
						if err != nil {
							// TODO(abg): nested error "invalid set item: %v"
							return err
						}
						err = <$f>(<$w>)
						if err != nil {
							return err
						}
					}
				return nil
			}

			func (<$v> <.Name>) Size() int {
				return len(<$v>)
			}

			func (<.Name>) ValueType() <$wire>.Type {
				return <typeCode .Spec.ValueSpec>
			}

			func (<.Name>) Close() {}
		`,
		struct {
			Name string
			Spec *compile.SetSpec
		}{Name: name, Spec: spec},
	)

	return name, wrapGenerateError(spec.ThriftName(), err)
}

func (s *setGenerator) Reader(g Generator, spec *compile.SetSpec) (string, error) {
	name := "_" + valueName(spec) + "_Read"
	if s.HasReader(name) {
		return name, nil
	}

	err := g.DeclareFromTemplate(
		`
			<$wire := import "go.uber.org/thriftrw/wire">
			<$setType := typeReference .Spec>

			<$s := newVar "s">
			<$i := newVar "i">
			<$o := newVar "o">
			<$x := newVar "x">
			func <.Name>(<$s> <$wire>.ValueList) (<$setType>, error) {
				if <$s>.ValueType() != <typeCode .Spec.ValueSpec> {
					return nil, nil
				}

				<if isHashable .Spec.ValueSpec>
					<$o> := make(<$setType>, <$s>.Size())
				<else>
					<$o> := make(<$setType>, 0, <$s>.Size())
				<end>
				err := <$s>.ForEach(func(<$x> <$wire>.Value) error {
					<$i>, err := <fromWire .Spec.ValueSpec $x>
					if err != nil {
						return err
					}
					<if isHashable .Spec.ValueSpec>
						<$o>[<$i>] = struct{}{}
					<else>
						<$o> = append(<$o>, <$i>)
					<end>
					return nil
				})
				<$s>.Close()
				return <$o>, err
			}
		`,
		struct {
			Name string
			Spec *compile.SetSpec
		}{Name: name, Spec: spec},
	)

	return name, wrapGenerateError(spec.ThriftName(), err)
}
