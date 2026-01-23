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

/**
 *  Object Coder
 *  ~~~~~~~~~~~~
 *  JsON, XML, ...
 *
 *  1. encode object to string;
 *  2. decode string to object.
 */
type ObjectCoder interface {

	/**
	 *  Encode Map/List object to string
	 *
	 * @param object - Map or List
	 * @return serialized string
	 */
	Encode(object interface{}) string

	/**
	 *  Decode string to Map/List object
	 *
	 * @param string - serialized string
	 * @return Map or List
	 */
	Decode(string string) interface{}
}

//
//  JsON
//

var jsonCoder ObjectCoder = nil

func SetJSONCoder(coder ObjectCoder) {
	jsonCoder = coder
}

func JSONEncode(object interface{}) string {
	return jsonCoder.Encode(object)
}

func JSONDecode(string string) interface{} {
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
