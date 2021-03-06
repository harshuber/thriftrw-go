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
	"encoding/json"
	"fmt"
	"testing"

	te "go.uber.org/thriftrw/gen/testdata/enums"
	"go.uber.org/thriftrw/wire"

	"github.com/stretchr/testify/assert"
)

func TestValueOfEnumDefault(t *testing.T) {
	tests := []struct {
		e te.EnumDefault
		i int32
	}{
		{te.EnumDefaultFoo, 0},
		{te.EnumDefaultBar, 1},
		{te.EnumDefaultBaz, 2},
	}
	for _, tt := range tests {
		assert.Equal(t, int32(tt.e), tt.i, "Value for %v does not match", tt.e)
	}
}

func TestValueOfEnumWithValues(t *testing.T) {
	tests := []struct {
		e te.EnumWithValues
		i int32
	}{
		{te.EnumWithValuesX, 123},
		{te.EnumWithValuesY, 456},
		{te.EnumWithValuesZ, 789},
	}
	for _, tt := range tests {
		assert.Equal(t, int32(tt.e), tt.i, "Value for %v does not match", tt.e)
	}
}

func TestEnumDefaultWire(t *testing.T) {
	tests := []struct {
		e te.EnumDefault
		v wire.Value
	}{
		{te.EnumDefaultFoo, wire.NewValueI32(0)},
		{te.EnumDefaultBar, wire.NewValueI32(1)},
		{te.EnumDefaultBaz, wire.NewValueI32(2)},
	}

	for _, tt := range tests {
		assertRoundTrip(t, &tt.e, tt.v, "EnumDefault")
	}
}

func TestValueOfEnumWithDuplicateValues(t *testing.T) {
	tests := []struct {
		e te.EnumWithDuplicateValues
		i int32
	}{
		{te.EnumWithDuplicateValuesP, 0},
		{te.EnumWithDuplicateValuesQ, -1},
		{te.EnumWithDuplicateValuesR, 0},
	}
	for _, tt := range tests {
		assert.Equal(t, int32(tt.e), tt.i, "Value for %v does not match", tt.e)
	}
}

func TestEnumWithDuplicateValuesWire(t *testing.T) {
	tests := []struct {
		e te.EnumWithDuplicateValues
		v wire.Value
	}{
		{te.EnumWithDuplicateValuesP, wire.NewValueI32(0)},
		{te.EnumWithDuplicateValuesQ, wire.NewValueI32(-1)},
		{te.EnumWithDuplicateValuesR, wire.NewValueI32(0)},
	}

	for _, tt := range tests {
		assertRoundTrip(t, &tt.e, tt.v, "EnumWithDuplicateValues")
	}
}

func TestUnknownEnumValue(t *testing.T) {
	var e te.EnumDefault
	if assert.NoError(t, e.FromWire(wire.NewValueI32(42))) {
		assert.Equal(t, te.EnumDefault(42), e)
	}
}

func TestOptionalEnum(t *testing.T) {
	foo := te.EnumDefaultFoo

	tests := []struct {
		s te.StructWithOptionalEnum
		v wire.Value
	}{
		{
			te.StructWithOptionalEnum{E: &foo},
			wire.NewValueStruct(wire.Struct{Fields: []wire.Field{
				{ID: 1, Value: wire.NewValueI32(0)},
			}}),
		},
		{
			te.StructWithOptionalEnum{},
			wire.NewValueStruct(wire.Struct{Fields: []wire.Field{}}),
		},
	}

	for _, tt := range tests {
		assertRoundTrip(t, &tt.s, tt.v, "StructWithOptionalEnum")
	}
}

func TestEnumString(t *testing.T) {
	tests := []struct {
		give fmt.Stringer
		want string
	}{
		{
			te.EmptyEnum(42),
			"EmptyEnum(42)",
		},
		{
			te.EnumDefaultFoo,
			"Foo",
		},
		{
			te.EnumDefault(-1),
			"EnumDefault(-1)",
		},
		{
			te.EnumWithDuplicateValuesP,
			"P",
		},
		{
			te.EnumWithDuplicateValuesR,
			"P", // same value as P
		},
		{
			te.EnumWithDuplicateValuesQ,
			"Q",
		},
		{
			te.LowerCaseEnumContaining,
			"containing",
		},
		{
			te.LowerCaseEnumLowerCase,
			"lower_case",
		},
		{
			te.LowerCaseEnumItems,
			"items",
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.want, tt.give.String())
	}
}

func TestJson(t *testing.T) {
	tests := []struct {
		title string
		e     te.EnumWithDuplicateValues
		json  string
	}{
		{`Q <-> "Q"`, te.EnumWithDuplicateValuesQ, `"Q"`},
		{`P <-> "P"`, te.EnumWithDuplicateValuesP, `"P"`},
		{`Q <-> "P"`, te.EnumWithDuplicateValuesR, `"P"`}, // not a typo.
		{`42 <-> 42`, 42, `42`},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			b, err := json.Marshal(tt.e)
			if assert.NoError(t, err, "unabled to marshal enum") {
				assert.Equal(t, b, []byte(tt.json), "json output doesn't match")

				var e te.EnumWithDuplicateValues = 99
				err = json.Unmarshal([]byte(tt.json), &e)
				if assert.NoError(t, err, "unabled to unmarshal enum") {
					assert.Equal(t, tt.e, e, "parsed json doesn't match")
				}
			}
		})
	}
}

func TestInvalidJSON(t *testing.T) {
	tests := []struct {
		title       string
		json        string
		errContains string
	}{
		{`empty`, ``, "unexpected end of JSON input"},
		{`ident`, `P`, "invalid character 'P'"},
		{`not terminated`, `"P`, "unexpected end of JSON input"},
		{`not started`, `P"`, "invalid character 'P'"},
		{`notint`, `42.`, "invalid character"},
		{`toobig`, ` 2147483648`, "overflow"},
		{`toosmall`, `-2147483649`, "underflow"},
		{`justbigenough`, ` 2147483647`, ""},
		{`justsmallenough`, `-2147483648`, ""},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			var e te.EnumWithDuplicateValues
			err := json.Unmarshal([]byte(tt.json), &e)
			if tt.errContains != "" {
				if assert.Error(t, err, "must NOT be able to unmarshal this") {
					assert.Contains(t, err.Error(), tt.errContains)
				}
			} else {
				assert.NoError(t, err, "must be able to unmarshal this")
			}
		})
	}
}
