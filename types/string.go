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
)

type Stringer interface {
	Object
	fmt.Stringer
}

func FetchString(str interface{}) string {
	wrapper, ok := str.(fmt.Stringer)
	if ok {
		return wrapper.String()
	} else {
		return str.(string)
	}
}

/**
 *  Constant String Wrapper
 *  ~~~~~~~~~~~~~~~~~~~~~~~
 */
type ConstantString struct {
	BaseObject

	_string string
}

func (str *ConstantString) Init(string string) *ConstantString {
	str._string = string
	return str
}

//-------- fmt.Stringer

func (str *ConstantString) String() string {
	return str._string
}

//-------- IObject

func (str *ConstantString) Equal(other interface{}) bool {
	// compare pointers
	if str == other {
		return true
	}
	// compare inner strings
	wrapper, ok := other.(fmt.Stringer)
	if ok {
		return str._string == wrapper.String()
	}
	text, ok := other.(string)
	if ok {
		return str._string == text
	} else {
		return str._string == other
	}
}
