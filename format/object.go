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

import . "github.com/dimchat/mkm-go/types"

// ObjectCoder defines the interface for object serialization/deserialization
//
//	Supported formats include:
//	    JSON, XML, ...
type ObjectCoder interface {

	// Encode converts a Map or List object to its serialized string form
	//
	// Parameters:
	//   - object: The input Map or List to be encoded
	//
	// Returns: The serialized string representation of the object
	Encode(object any) string

	// Decode parses a serialized string back to a Map or List object
	//
	// Parameters:
	//   - str: The serialized string to be decoded
	//
	// Returns: The reconstructed Map or List object
	Decode(str string) any
}

//
//  JsON
//

var jsonCoder ObjectCoder = nil

func SetJSONCoder(coder ObjectCoder) {
	jsonCoder = coder
}

func JSONEncode(object any) string {
	return jsonCoder.Encode(object)
}

func JSONDecode(string string) any {
	return jsonCoder.Decode(string)
}

//
//  JsON <-> Map
//

func JSONEncodeMap(dict StringKeyMap) string {
	return jsonCoder.Encode(dict)
}

func JSONDecodeMap(str string) StringKeyMap {
	dict := jsonCoder.Decode(str)
	return FetchMap(dict)
}
