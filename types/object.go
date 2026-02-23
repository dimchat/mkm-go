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

// IObject
type Object interface {

	// Equal checks if this object is equal to another object.
	Equal(other any) bool
}

// ObjectType returns the type name of the given object.
// If the object is nil, it returns "<nil>".
func ObjectType(a any) string {
	if a == nil {
		return "<nil>"
	}
	t := reflect.TypeOf(a)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t.String()
}

// ObjectsEqual checks if two objects are equal.
// WARNING: Do not call this method inside Object.Equal().
func ObjectsEqual(a1, a2 any) bool {
	if a1 == nil && a2 == nil {
		return true
	} else if a1 == nil || a2 == nil {
		return false
	} else if p1, ok := a1.(Object); ok {
		return p1.Equal(a2)
	//} else if p2, ok := i2.(Object); ok {
	//	return p2.Equal(i1)
	}
	// check values
	v1 := ObjectTargetValue(a1)
	v2 := ObjectTargetValue(a2)
	if v1 == nil && v2 == nil {
		return true
	} else if v1 == nil || v2 == nil {
		return false
	}
	// other types
	return reflect.DeepEqual(a1, a2)
}

// ObjectTargetValue returns the value pointed to by the given object.
func ObjectTargetValue(a any) any {
	v, _ := ObjectReflectValue(a)
	return v
}

// ObjectReflectValue returns the reflect.Value of the given object and its underlying value.
func ObjectReflectValue(a any) (any, reflect.Value) {
	v := reflect.ValueOf(a)
	if !v.IsValid() {
		return nil, v
	} else if v.Kind() != reflect.Ptr {
		return a, v
	} else if v.IsNil() {
		return nil, v
	}
	return v.Elem().Interface(), v
}

// ObjectReflectPointer returns the reflect.Value of the given object and its pointer.
func ObjectReflectPointer(a any) (any, reflect.Value) {
	v := reflect.ValueOf(a)
	if !v.IsValid() {
		return nil, v
	} else if v.Kind() == reflect.Ptr {
		return a, v
	} else if v.CanAddr() {
		return v.Addr().Interface(), v
	}
	ptr := reflect.New(v.Type())
	ptr.Elem().Set(v)
	return ptr.Interface(), v
}

// ObjectPointer returns the address of the object value.
func ObjectPointer(a any) any {
	p, _ := ObjectReflectPointer(a)
	return p
}

// ObjectIsPointer checks whether the variable is a pointer.
func ObjectIsPointer(a any) bool {
	v := reflect.ValueOf(a)
	if !v.IsValid() {
		return false
	}
	return v.Kind() == reflect.Ptr
}

// ValueIsNil checks whether the value is nil.
func ValueIsNil(a any) bool {
	v := reflect.ValueOf(a)
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
