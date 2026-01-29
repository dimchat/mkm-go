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

import (
	"fmt"
	"reflect"
)

type Stringer interface {
	Object
	fmt.Stringer

	//Length() int

	IsEmpty() bool // Length() == 0
}

/**
 *  Constant String Wrapper
 *  ~~~~~~~~~~~~~~~~~~~~~~~
 */
type ConstantString struct {
	_string string
}

func (str *ConstantString) Init(text string) {
	str._string = text
}

//// Override
//func (str *ConstantString) Length() int {
//	return len(str._string)
//}

// Override
func (str *ConstantString) IsEmpty() bool {
	return len(str._string) == 0
}

//-------- fmt.Stringer

// Override
func (str *ConstantString) String() string {
	return str._string
}

//-------- IObject

// Override
func (str *ConstantString) Equal(other interface{}) bool {
	if other == nil {
		return str._string == ""
	} else if other == str {
		// same object
		return true
	}
	// check targeted value
	target, rv := ObjectReflectValue(other)
	if target == nil {
		return str._string == ""
	}
	// check value types
	switch v := target.(type) {
	case fmt.Stringer:
		other = v.String()
	case string:
		other = v
	default:
		// other types
		switch rv.Kind() {
		case reflect.String:
			other = rv.String()
		default:
			// type not matched
			return false
		}
	}
	return str._string == other
}
