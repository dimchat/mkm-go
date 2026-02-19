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

	//Length() int

	IsEmpty() bool // Length() == 0
}

/**
 *  Constant String Wrapper
 *  ~~~~~~~~~~~~~~~~~~~~~~~
 */
type ConstantString struct {
	//Stringer

	text string
}

func NewConstantString(s string) *ConstantString {
	return &ConstantString{text: s}
}

//// Override
//func (cs ConstantString) Length() int {
//	return len(cs.text)
//}

// Override
func (cs ConstantString) IsEmpty() bool {
	return len(cs.text) == 0
}

//-------- fmt.Stringer

// Override
func (cs ConstantString) String() string {
	return cs.text
}

//-------- IObject

// Override
func (cs ConstantString) Equal(other interface{}) bool {
	if other == nil {
		return cs.text == ""
	}
	var text string
	switch v := other.(type) {
	case fmt.Stringer:
		text = v.String()
	case string:
		text = v
	default:
		// type not matched
		return false
	}
	return cs.text == text
}
