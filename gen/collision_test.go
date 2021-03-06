// Copyright (c) 2015 Uber Technologies, Inc.
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
	"bytes"
	"reflect"
	"testing"

	tc "go.uber.org/thriftrw/gen/testdata/collision"
	"go.uber.org/thriftrw/protocol"
	"go.uber.org/thriftrw/ptr"
	"go.uber.org/thriftrw/wire"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStruct(t *testing.T) {
	tests := []struct {
		desc string
		x    thriftType
	}{
		{
			"StructCollision",
			&tc.StructCollision{
				CollisionField:  true,
				CollisionField2: "true",
			},
		},
		{
			"StructCollision2 (struct_collision)",
			&tc.StructCollision2{
				CollisionField:  true,
				CollisionField2: "such trueness",
			},
		},
		{
			"PrimitiveContainers",
			&tc.PrimitiveContainers{
				A: []string{"arbre", "fleur"},
				B: map[string]struct{}{
					"croissant": {},
					"baguette":  {},
				},
				C: map[string]string{
					"voiture": "bleue",
					"cammion": "rouge",
				},
			},
		},
		{
			"StructConstant (struct_constant)",
			tc.StructConstant,
		},
		{
			"UnionCollision .CollisionField",
			&tc.UnionCollision{
				CollisionField: ptr.Bool(true),
			},
		},
		{
			"UnionCollision .collision_field",
			&tc.UnionCollision{
				CollisionField2: ptr.String("so true"),
			},
		},
		{
			"UnionCollision2 (union_collision) .CollisionField",
			&tc.UnionCollision2{
				CollisionField: &[]bool{true}[0],
			},
		},
		{
			"UnionCollision (union_collision) .collision_field",
			&tc.UnionCollision2{
				CollisionField2: ptr.String("true indeed"),
			},
		},
	}
	for _, tt := range tests {
		roundTrip(t, tt.x, tt.desc)
	}
}

func TestConstant(t *testing.T) {
	require.Equal(t, tc.StructCollision2{
		CollisionField2: "false indeed",
	}, *tc.StructConstant)
}

func TestWithDefault(t *testing.T) {
	a := &tc.WithDefault{
		Pouet: &tc.StructCollision2{
			CollisionField2: "false indeed",
		},
	}
	b := &tc.WithDefault{}

	a = roundTrip(t, a, "WithDefault{filled in}").(*tc.WithDefault)
	b = roundTrip(t, b, "WithDefault{}").(*tc.WithDefault)
	if a != nil && b != nil {
		require.Equal(t, a, b)
	}
}

func TestMyEnum(t *testing.T) {
	tests := []struct {
		e tc.MyEnum
		n string
		v int64
	}{
		{tc.MyEnumX, "X", 123},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.n, tt.e.String())
		assert.Equal(t, tt.v, int64(tt.e))
	}
}

func TestMyEnum2(t *testing.T) {
	tests := []struct {
		e tc.MyEnum2
		n string
		v int64
	}{
		{tc.MyEnum2X, "X", 12},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.n, tt.e.String())
		assert.Equal(t, tt.v, int64(tt.e))
	}
}

func roundTrip(t *testing.T, x thriftType, msg string) thriftType {
	v, err := x.ToWire()
	if !assert.NoError(t, err, "failed to serialize: %v", x) {
		return nil
	}

	var buff bytes.Buffer
	if !assert.NoError(t, protocol.Binary.Encode(v, &buff), "%v: failed to serialize", msg) {
		return nil
	}

	newV, err := protocol.Binary.Decode(bytes.NewReader(buff.Bytes()), v.Type())
	if !assert.NoError(t, err, "%v: failed to deserialize", msg) {
		return nil
	}

	if !assert.True(
		t, wire.ValuesAreEqual(newV, v), "%v: deserialize(serialize(%v.ToWire())) != %v", msg, x, v) {
		return nil
	}

	xType := reflect.TypeOf(x)
	if xType.Kind() == reflect.Ptr {
		xType = xType.Elem()
	}

	gotX := reflect.New(xType)
	errval := gotX.MethodByName("FromWire").
		Call([]reflect.Value{reflect.ValueOf(newV)})[0].
		Interface()

	if !assert.Nil(t, errval, "FromWire: %v", msg) {
		return nil
	}

	assert.Equal(t, x, gotX.Interface(), "FromWire: %v", msg)
	return gotX.Interface().(thriftType)
}
