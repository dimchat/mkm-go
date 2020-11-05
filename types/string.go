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

import "fmt"

type Stringer interface {

	fmt.Stringer
	Object
}

func StringsEqual(str1, str2 Stringer) bool {
	if ObjectsEqual(str1, str2) {
		return true
	}
	return str1.String() == str2.String()
}

type ConstantString struct {
	Stringer

	_string string
}

func (str *ConstantString) Init(string string) *ConstantString {
	str._string = string
	return str
}

func (str ConstantString) String() string {
	return str._string
}

func (str ConstantString) Equal(other interface{}) bool {
	other = ObjectValue(other)
	switch other.(type) {
	case Stringer:
		return StringsEqual(str, other.(Stringer))
	case string:
		return str._string == other.(string)
	default:
		return false
	}
}
