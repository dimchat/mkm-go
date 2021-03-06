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
	"encoding/json"
)

type JSONParser struct {
	DataParser
}

func (parser JSONParser) Encode(object interface{}) []byte {
	bytes, err := json.Marshal(object)
	if err == nil {
		return bytes
	} else {
		//panic("failed to encode to JsON string")
		return nil
	}
}

func (parser JSONParser) Decode(bytes []byte) interface{} {
	for _, ch := range bytes {
		if ch == '{' {
			// decode to map
			var dict map[string]interface{}
			err := json.Unmarshal(bytes, &dict)
			if err == nil {
				return dict
			} else {
				return nil
			}
		} else if ch == '[' {
			// decode to array
			var array []interface{}
			err := json.Unmarshal(bytes, &array)
			if err == nil {
				return array
			} else {
				return nil
			}
		} else if ch != ' ' && ch != '\t' {
			// error
			break
		}
	}
	//panic(bytes)
	return nil
}

var jsonParser DataParser = new(JSONParser)

func SetJSONParser(parser DataParser) {
	jsonParser = parser
}

func JSONEncode(object interface{}) []byte {
	return jsonParser.Encode(object)
}

func JSONDecode(bytes []byte) interface{} {
	return jsonParser.Decode(bytes)
}

//
//  JsON <-> Map
//

func JSONEncodeMap(dict map[string]interface{}) []byte {
	return jsonParser.Encode(dict)
}

func JSONDecodeMap(bytes []byte) map[string]interface{} {
	obj, ok := jsonParser.Decode(bytes).(map[string]interface{})
	if ok {
		return obj
	} else {
		return nil
	}
}

//
//  JsON <-> List
//

func JSONEncodeList(array []interface{}) []byte {
	return jsonParser.Encode(array)
}

func JSONDecodeList(bytes []byte) []interface{} {
	obj, ok := jsonParser.Decode(bytes).([]interface{})
	if ok {
		return obj
	} else {
		return nil
	}
}
