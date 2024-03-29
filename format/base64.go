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
package format

import (
	"encoding/base64"
)

type Base64Coder struct {}

func (coder Base64Coder) Init() DataCoder {
	return coder
}

//-------- IDataCoder

func (coder Base64Coder) Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func (coder Base64Coder) Decode(string string) []byte {
	bytes, err := base64.StdEncoding.DecodeString(string)
	if err != nil {
		panic(err)
	}
	return bytes
}

//
//  Instance of DataCoder
//
var base64Coder DataCoder = new(Base64Coder)

func Base64SetCoder(coder DataCoder) {
	base64Coder = coder
}

func Base64Encode(bytes []byte) string {
	return base64Coder.Encode(bytes)
}

func Base64Decode(b64 string) []byte {
	return base64Coder.Decode(b64)
}
