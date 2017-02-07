// Code generated by thriftrw v1.1.0
// @generated

package services

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/gen/testdata/exceptions"
	"go.uber.org/thriftrw/gen/testdata/unions"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type KeyValue_GetManyValues_Args struct {
	Range []Key `json:"range"`
}

type _List_Key_ValueList []Key

func (v _List_Key_ValueList) ForEach(f func(wire.Value) error) error {
	for _, x := range v {
		w, err := x.ToWire()
		if err != nil {
			return err
		}
		err = f(w)
		if err != nil {
			return err
		}
	}
	return nil
}

func (v _List_Key_ValueList) Size() int {
	return len(v)
}

func (_List_Key_ValueList) ValueType() wire.Type {
	return wire.TBinary
}

func (_List_Key_ValueList) Close() {
}

func (v *KeyValue_GetManyValues_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Range != nil {
		w, err = wire.NewValueList(_List_Key_ValueList(v.Range)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _List_Key_Read(l wire.ValueList) ([]Key, error) {
	if l.ValueType() != wire.TBinary {
		return nil, nil
	}
	o := make([]Key, 0, l.Size())
	err := l.ForEach(func(x wire.Value) error {
		i, err := _Key_Read(x)
		if err != nil {
			return err
		}
		o = append(o, i)
		return nil
	})
	l.Close()
	return o, err
}

func (v *KeyValue_GetManyValues_Args) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TList {
				v.Range, err = _List_Key_Read(field.Value.GetList())
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (v *KeyValue_GetManyValues_Args) String() string {
	var fields [1]string
	i := 0
	if v.Range != nil {
		fields[i] = fmt.Sprintf("Range: %v", v.Range)
		i++
	}
	return fmt.Sprintf("KeyValue_GetManyValues_Args{%v}", strings.Join(fields[:i], ", "))
}

func _List_Key_Equals(lhs, rhs []Key) bool {
	if len(lhs) != len(rhs) {
		return false
	}
	for i, lv := range lhs {
		rv := rhs[i]
		if !(lv == rv) {
			return false
		}
	}
	return true
}

func (v *KeyValue_GetManyValues_Args) Equals(rhs *KeyValue_GetManyValues_Args) bool {
	if !((v.Range == nil && rhs.Range == nil) || (v.Range != nil && rhs.Range != nil && _List_Key_Equals(v.Range, rhs.Range))) {
		return false
	}
	return true
}

func (v *KeyValue_GetManyValues_Args) MethodName() string {
	return "getManyValues"
}

func (v *KeyValue_GetManyValues_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

var KeyValue_GetManyValues_Helper = struct {
	Args           func(range2 []Key) *KeyValue_GetManyValues_Args
	IsException    func(error) bool
	WrapResponse   func([]*unions.ArbitraryValue, error) (*KeyValue_GetManyValues_Result, error)
	UnwrapResponse func(*KeyValue_GetManyValues_Result) ([]*unions.ArbitraryValue, error)
}{}

func init() {
	KeyValue_GetManyValues_Helper.Args = func(range2 []Key) *KeyValue_GetManyValues_Args {
		return &KeyValue_GetManyValues_Args{Range: range2}
	}
	KeyValue_GetManyValues_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *exceptions.DoesNotExistException:
			return true
		default:
			return false
		}
	}
	KeyValue_GetManyValues_Helper.WrapResponse = func(success []*unions.ArbitraryValue, err error) (*KeyValue_GetManyValues_Result, error) {
		if err == nil {
			return &KeyValue_GetManyValues_Result{Success: success}, nil
		}
		switch e := err.(type) {
		case *exceptions.DoesNotExistException:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for KeyValue_GetManyValues_Result.DoesNotExist")
			}
			return &KeyValue_GetManyValues_Result{DoesNotExist: e}, nil
		}
		return nil, err
	}
	KeyValue_GetManyValues_Helper.UnwrapResponse = func(result *KeyValue_GetManyValues_Result) (success []*unions.ArbitraryValue, err error) {
		if result.DoesNotExist != nil {
			err = result.DoesNotExist
			return
		}
		if result.Success != nil {
			success = result.Success
			return
		}
		err = errors.New("expected a non-void result")
		return
	}
}

type KeyValue_GetManyValues_Result struct {
	Success      []*unions.ArbitraryValue          `json:"success"`
	DoesNotExist *exceptions.DoesNotExistException `json:"doesNotExist,omitempty"`
}

type _List_ArbitraryValue_ValueList []*unions.ArbitraryValue

func (v _List_ArbitraryValue_ValueList) ForEach(f func(wire.Value) error) error {
	for i, x := range v {
		if x == nil {
			return fmt.Errorf("invalid [%v]: value is nil", i)
		}
		w, err := x.ToWire()
		if err != nil {
			return err
		}
		err = f(w)
		if err != nil {
			return err
		}
	}
	return nil
}

func (v _List_ArbitraryValue_ValueList) Size() int {
	return len(v)
}

func (_List_ArbitraryValue_ValueList) ValueType() wire.Type {
	return wire.TStruct
}

func (_List_ArbitraryValue_ValueList) Close() {
}

func (v *KeyValue_GetManyValues_Result) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Success != nil {
		w, err = wire.NewValueList(_List_ArbitraryValue_ValueList(v.Success)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	if v.DoesNotExist != nil {
		w, err = v.DoesNotExist.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if i != 1 {
		return wire.Value{}, fmt.Errorf("KeyValue_GetManyValues_Result should have exactly one field: got %v fields", i)
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _ArbitraryValue_Read(w wire.Value) (*unions.ArbitraryValue, error) {
	var v unions.ArbitraryValue
	err := v.FromWire(w)
	return &v, err
}

func _List_ArbitraryValue_Read(l wire.ValueList) ([]*unions.ArbitraryValue, error) {
	if l.ValueType() != wire.TStruct {
		return nil, nil
	}
	o := make([]*unions.ArbitraryValue, 0, l.Size())
	err := l.ForEach(func(x wire.Value) error {
		i, err := _ArbitraryValue_Read(x)
		if err != nil {
			return err
		}
		o = append(o, i)
		return nil
	})
	l.Close()
	return o, err
}

func (v *KeyValue_GetManyValues_Result) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TList {
				v.Success, err = _List_ArbitraryValue_Read(field.Value.GetList())
				if err != nil {
					return err
				}
			}
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.DoesNotExist, err = _DoesNotExistException_Read(field.Value)
				if err != nil {
					return err
				}
			}
		}
	}
	count := 0
	if v.Success != nil {
		count++
	}
	if v.DoesNotExist != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("KeyValue_GetManyValues_Result should have exactly one field: got %v fields", count)
	}
	return nil
}

func (v *KeyValue_GetManyValues_Result) String() string {
	var fields [2]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}
	if v.DoesNotExist != nil {
		fields[i] = fmt.Sprintf("DoesNotExist: %v", v.DoesNotExist)
		i++
	}
	return fmt.Sprintf("KeyValue_GetManyValues_Result{%v}", strings.Join(fields[:i], ", "))
}

func _List_ArbitraryValue_Equals(lhs, rhs []*unions.ArbitraryValue) bool {
	if len(lhs) != len(rhs) {
		return false
	}
	for i, lv := range lhs {
		rv := rhs[i]
		if !lv.Equals(rv) {
			return false
		}
	}
	return true
}

func (v *KeyValue_GetManyValues_Result) Equals(rhs *KeyValue_GetManyValues_Result) bool {
	if !((v.Success == nil && rhs.Success == nil) || (v.Success != nil && rhs.Success != nil && _List_ArbitraryValue_Equals(v.Success, rhs.Success))) {
		return false
	}
	if !((v.DoesNotExist == nil && rhs.DoesNotExist == nil) || (v.DoesNotExist != nil && rhs.DoesNotExist != nil && v.DoesNotExist.Equals(rhs.DoesNotExist))) {
		return false
	}
	return true
}

func (v *KeyValue_GetManyValues_Result) MethodName() string {
	return "getManyValues"
}

func (v *KeyValue_GetManyValues_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
