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

import "encoding/json"

type JSONParser struct {
	DataParser
}

func (parser *JSONParser) Encode(object interface{}) []byte {
	bytes, err := json.Marshal(object)
	if err == nil {
		return bytes
	} else {
		//panic("failed to encode to JsON string")
		return nil
	}
}

func (parser *JSONParser) Decode(bytes []byte) interface{} {
	var dict map[string]interface{}
	err := json.Unmarshal(bytes, &dict)
	if err == nil {
		return dict
	} else {
		return nil
	}
}

var jsonParser DataParser = new(JSONParser)

func SetJSONParser(parser DataParser) {
	jsonParser = parser
}

func JSONMapFromBytes(bytes []byte) map[string]interface{} {
	dict := jsonParser.Decode(bytes)
	res, ok := dict.(map[string]interface{})
	if ok {
		return res
	} else {
		return nil
	}
}

func JSONBytesFromMap(dictionary map[string]interface{}) []byte {
	return jsonParser.Encode(dictionary)
}
