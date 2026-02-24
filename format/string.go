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

// StringCoder defines the interface for string encoding/decoding between string and binary data
//
//	Supported character encodings include:
//	    UTF-8, UTF-16, GBK, GB2312, ...
type StringCoder interface {

	// Encode converts a local string to binary data (byte slice) using the specified charset
	//
	// Parameters:
	//   - str: Input local string to be encoded
	//
	// Returns: Binary data (byte slice) of the encoded string
	Encode(str string) []byte

	// Decode parses binary data (byte slice) back to a local string using the specified charset
	//
	// Parameters:
	//   - bytes: Input binary data (byte slice) to be decoded
	//
	// Returns: Local string decoded from the binary data
	Decode(bytes []byte) string
}

//
//  UTF-8
//

var utf8Coder StringCoder = nil

func SetUTF8Coder(coder StringCoder) {
	utf8Coder = coder
}

func UTF8Encode(string string) []byte {
	return utf8Coder.Encode(string)
}

func UTF8Decode(bytes []byte) string {
	return utf8Coder.Decode(bytes)
}
