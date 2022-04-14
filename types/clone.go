/* license: https://mit-license.org
 * ==============================================================================
 * The MIT License (MIT)
 *
 * Copyright (c) 2022 Albert Moky
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

type Cloneable interface {

	Clone() interface{}
}

/**
 *  Shallow Copy
 *  ~~~~~~~~~~~~
 */
func Copy(object interface{}) interface{} {
	if ValueIsNil(object) {
		return nil
	} else if obj, ok := object.(Cloneable); ok {
		return obj.Clone()
	}
	// collections
	if dict, ok := object.(map[string]interface{}); ok {
		return CopyMap(dict)
	} else if arr, ok := object.([]interface{}); ok {
		return CopyList(arr)
	}
	// others
	return object
}

func CopyMap(dictionary map[string]interface{}) map[string]interface{} {
	clone := make(map[string]interface{})
	for key, value := range dictionary {
		clone[key] = value
	}
	return clone
}

func CopyList(array []interface{}) []interface{} {
	clone := make([]interface{}, len(array))
	for key, value := range array {
		clone[key] = value
	}
	return clone
}

/**
 *  Deep Copy
 *  ~~~~~~~~~
 */
func DeepCopy(object interface{}) interface{} {
	if ValueIsNil(object) {
		return nil
	} else if obj, ok := object.(Cloneable); ok {
		return obj.Clone()
	}
	// collections
	if dict, ok := object.(map[string]interface{}); ok {
		return DeepCopyMap(dict)
	} else if arr, ok := object.([]interface{}); ok {
		return DeepCopyList(arr)
	}
	// others
	return object
}

func DeepCopyMap(dictionary map[string]interface{}) map[string]interface{} {
	clone := make(map[string]interface{})
	for key, value := range dictionary {
		clone[key] = DeepCopy(value)
	}
	return clone
}

func DeepCopyList(array []interface{}) []interface{} {
	clone := make([]interface{}, len(array))
	for key, value := range array {
		clone[key] = DeepCopy(value)
	}
	return clone
}
