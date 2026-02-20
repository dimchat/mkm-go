/* license: https://mit-license.org
 * ==============================================================================
 * The MIT License (MIT)
 *
 * Copyright (c) 2020 Albert Moky
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 * ==============================================================================
 */
package types

import "reflect"

/**
 *  IObject
 */
type Object interface {

	Equal(other interface{}) bool
}

/**
 *  Get object type (class name)
 */
func ObjectType(i interface{}) string {
	if i == nil {
		return "<nil>"
	}
	t := reflect.TypeOf(i)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t.String()
}

/**
 *  Check whether the two objects equal
 *
 *  WARNING:
 *      Don't call this method in Object.Equal()
 */
func ObjectsEqual(i1, i2 interface{}) bool {
	if i1 == nil && i2 == nil {
		return true
	} else if i1 == nil || i2 == nil {
		return false
	} else if p1, ok := i1.(Object); ok {
		return p1.Equal(i2)
	//} else if p2, ok := i2.(Object); ok {
	//	return p2.Equal(i1)
	}
	// check values
	v1 := ObjectTargetValue(i1)
	v2 := ObjectTargetValue(i2)
	if v1 == nil && v2 == nil {
		return true
	} else if v1 == nil || v2 == nil {
		return false
	}
	// other types
	return reflect.DeepEqual(i1, i2)
}

/**
 *  Get targeted value that the object pointer pointing to
 */
func ObjectTargetValue(i interface{}) interface{} {
	v, _ := ObjectReflectValue(i)
	return v
}

func ObjectReflectValue(i interface{}) (interface{}, reflect.Value) {
	v := reflect.ValueOf(i)
	if !v.IsValid() {
		return nil, v
	} else if v.Kind() != reflect.Ptr {
		return i, v
	} else if v.IsNil() {
		return nil, v
	}
	return v.Elem().Interface(), v
}

func ObjectReflectPointer(i interface{}) (interface{}, reflect.Value) {
	v := reflect.ValueOf(i)
	if !v.IsValid() {
		return nil, v
	} else if v.Kind() == reflect.Ptr {
		return i, v
	} else if v.CanAddr() {
		return v.Addr().Interface(), v
	}
	ptr := reflect.New(v.Type())
	ptr.Elem().Set(v)
	return ptr.Interface(), v
}

/**
 *  Get address of the object value
 */
func ObjectPointer(i interface{}) interface{} {
	p, _ := ObjectReflectPointer(i)
	return p
}

/**
 *  Check whether the variable is a pointer
 */
func ObjectIsPointer(i interface{}) bool {
	v := reflect.ValueOf(i)
	if !v.IsValid() {
		return false
	}
	return v.Kind() == reflect.Ptr
}

/**
 *  Check whether the value is nil
 */
func ValueIsNil(i interface{}) bool {
	v := reflect.ValueOf(i)
	if !v.IsValid() {
		return true
	}
	k := v.Kind()
	if k == reflect.Interface {
		if v.IsNil() {
			return true
		}
		// check inner value
		e := v.Elem()
		if !e.IsValid() {
			return true
		}
		return isNil(e, e.Kind())
	}
	return isNil(v, k)
}

func isNil(v reflect.Value, k reflect.Kind) bool {
	switch k {
	case reflect.Ptr, reflect.Map, reflect.Slice, reflect.Func, reflect.Chan, reflect.Interface:
		return v.IsNil()
	default:
		return false
	}
}
