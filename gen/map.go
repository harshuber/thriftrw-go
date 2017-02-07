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

import (
	"fmt"

	"go.uber.org/thriftrw/compile"
)

// mapGenerator generates logic to convert lists of arbitrary Thrift types to
// and from MapItemLists.
type mapGenerator struct {
	hasReaders
	hasLazyLists
}

// MapItemList generates a new MapItemList type alias for the given map.
//
// The following is generated:
//
// 	type $mapItemListName map[$keyType]$valueType
//
// 	func (v $mapItemListName) ForEach(f func(wire.MapItem) error) error { ... }
//
// 	func (v $mapItemListName) Close() { ... }
//
// And $mapItemListName is returned. This may be used where a MapItemList of the
// given type is expected.
func (m *mapGenerator) ItemList(g Generator, spec *compile.MapSpec) (string, error) {
	// TODO(abg): Unhashable types
	name := "_" + valueName(spec) + "_MapItemList"
	if m.HasLazyList(name) {
		return name, nil
	}

	err := g.DeclareFromTemplate(
		`
			<$wire := import "go.uber.org/thriftrw/wire">
			type <.Name> <typeReference .Spec>

			<$m := newVar "m">
			<$f := newVar "f">
			<$k := newVar "k">
			<$v := newVar "v">
			<$i := newVar "i">
			<$kw := newVar "kw">
			<$vw := newVar "vw">
			func (<$m> <.Name>) ForEach(<$f> func(<$wire>.MapItem) error) error {
				<if isHashable .Spec.KeySpec>
					for <$k>, <$v> := range <$m> {
				<else>
					for _, <$i> := range <$m> {
						<$k> := <$i>.Key
						<$v> := <$i>.Value
				<end>
						<if not (isPrimitiveType .Spec.KeySpec)>
							if <$k> == nil {
								return <import "fmt">.Errorf("invalid map key: value is nil")
							}
						<end>

						<if not (isPrimitiveType .Spec.ValueSpec)>
							if <$v> == nil {
								return <import "fmt">.Errorf("invalid [%v]: value is nil", <$k>)
							}
						<end>

						<$kw>, err := <toWire .Spec.KeySpec $k>
						if err != nil {
							// TODO(abg): nested error "invalid map key: %v"
							return err
						}

						<$vw>, err := <toWire .Spec.ValueSpec $v>
						if err != nil {
							// TODO(abg): nested error "invalid [%v]: %v"
							return err
						}
						err = <$f>(<$wire>.MapItem{Key: <$kw>, Value: <$vw>})
						if err != nil {
							return err
						}
					}
				return nil
			}

			func (<$m> <.Name>) Size() int {
				return len(<$m>)
			}

			func (<.Name>) KeyType() <$wire>.Type {
				return <typeCode .Spec.KeySpec>
			}

			func (<.Name>) ValueType() <$wire>.Type {
				return <typeCode .Spec.ValueSpec>
			}

			func (<.Name>) Close() {}
		`,
		struct {
			Name string
			Spec *compile.MapSpec
		}{Name: name, Spec: spec},
	)

	return name, wrapGenerateError(spec.ThriftName(), err)
}

func (m *mapGenerator) Reader(g Generator, spec *compile.MapSpec) (string, error) {
	name := "_" + valueName(spec) + "_Read"
	if m.HasReader(name) {
		return name, nil
	}

	err := g.DeclareFromTemplate(
		`
			<$wire := import "go.uber.org/thriftrw/wire">
			<$mapType := typeReference .Spec>

			<$m := newVar "m">
			<$o := newVar "o">
			<$x := newVar "x">
			<$k := newVar "k">
			<$v := newVar "v">
			func <.Name>(<$m> <$wire>.MapItemList) (<$mapType>, error) {
				if <$m>.KeyType() != <typeCode .Spec.KeySpec> {
					return nil, nil
				}

				if <$m>.ValueType() != <typeCode .Spec.ValueSpec> {
					return nil, nil
				}

				<if isHashable .Spec.KeySpec>
					<$o> := make(<$mapType>, <$m>.Size())
				<else>
					<$o> := make(<$mapType>, 0, <$m>.Size())
				<end>
				err := <$m>.ForEach(func(<$x> <$wire>.MapItem) error {
					<$k>, err := <fromWire .Spec.KeySpec (printf "%s.Key" $x)>
					if err != nil {
						return err
					}

					<$v>, err := <fromWire .Spec.ValueSpec (printf "%s.Value" $x)>
					if err != nil {
						return err
					}

					<if isHashable .Spec.KeySpec>
						<$o>[<$k>] = <$v>
					<else>
						<$o> = append(<$o>, struct {
							Key <typeReference .Spec.KeySpec>
							Value <typeReference .Spec.ValueSpec>
						}{<$k>, <$v>})
					<end>
					return nil
				})
				<$m>.Close()
				return <$o>, err
			}
		`,
		struct {
			Name string
			Spec *compile.MapSpec
		}{Name: name, Spec: spec},
	)

	return name, wrapGenerateError(spec.ThriftName(), err)
}

// Equals generates a function to compare maps of the given type
//
// 	func $name(lhs, rhs $mapType) bool {
// 		...
// 	}
//
// And returns its name.
func (m *mapGenerator) Equals(g Generator, spec *compile.MapSpec) (string, error) {
	name := "_" + valueName(spec) + "_Equals"

	err := g.EnsureDeclared(
		`
			<$mapType := typeReference .Spec>

			<$lhs := newVar "lhs">
			<$rhs := newVar "rhs">
			func <.Name>(<$lhs>, <$rhs> <$mapType>) bool {
				if len(<$lhs>) != len(<$rhs>) {
					return false
				}

				<$lk := newVar "lk">
				<$lv := newVar "lv">
				<$rv := newVar "rv">
				<$ok := newVar "ok">
				<if isHashable .Spec.KeySpec>
					for <$lk>, <$lv> := range <$lhs> {
						<$rv>, <$ok> := <$rhs>[<$lk>]
						if !<$ok> {
							return false
						}
						if !<equals .Spec.ValueSpec $lv $rv> {
							return false
						}
					}
				<else>
					// Note if keys are not hashable then this is O(n^2) in time complexity.
					<equalsForUnhashableKey .Spec $lhs $rhs>
				<end>
				return true
			}
		`,
		struct {
			Name string
			Spec *compile.MapSpec
		}{Name: name, Spec: spec},
		TemplateFunc("equalsForUnhashableKey", m.equalsForUnhashableKey),
	)

	return name, wrapGenerateError(spec.ThriftName(), err)
}

func (m *mapGenerator) equalsForUnhashableKey(g Generator, spec compile.TypeSpec, lhs, rhs string) (string, error) {
	return g.TextTemplate(
		`
			<$i := newVar "i">
			<$j := newVar "j">
			<$lk := newVar "lk">
			<$lv := newVar "lv">
			<$rk := newVar "rk">
			<$rv := newVar "rv">
			<$ok := newVar "ok">
			for _, <$i> := range <.LHS> {
				<$lk> := <$i>.Key
				<$lv> := <$i>.Value
				<$ok> := false
				for _, <$j> := range <.RHS> {
					<$rk> := <$j>.Key
					<$rv> := <$j>.Value
					if !<equals .Spec.KeySpec $lk $rk> {
						continue
					}

					if !<equals .Spec.ValueSpec $lv $rv> {
						// Caveat: Behavior is undefined if the map contract
						// (only one instance of each key) isn't satisfied.
						return false
					}
					<$ok> = true
					break
				}

				if !<$ok> {
					return false
				}
			}
		`,
		struct {
			Spec compile.TypeSpec
			LHS  string
			RHS  string
		}{Spec: spec, LHS: lhs, RHS: rhs},
	)
}

func valueName(spec compile.TypeSpec) string {
	switch s := spec.(type) {
	case *compile.MapSpec:
		return fmt.Sprintf(
			"Map_%s_%s", valueName(s.KeySpec), valueName(s.ValueSpec),
		)
	case *compile.ListSpec:
		return fmt.Sprintf("List_%s", valueName(s.ValueSpec))
	case *compile.SetSpec:
		return fmt.Sprintf("Set_%s", valueName(s.ValueSpec))
	default:
		return goCase(spec.ThriftName())
	}
}
