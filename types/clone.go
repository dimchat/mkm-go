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

import "reflect"

type Cloneable interface {
	Clone() interface{}
}

/**
 *  Data Copier
 */
type Copier interface {

	//
	//  Shallow Copy
	//

	Copy(object interface{}) interface{}

	CopyMap(dictionary StringKeyMap) StringKeyMap

	CopyList(array []interface{}) []interface{}

	//
	//  Deep Copy
	//

	DeepCopy(object interface{}) interface{}

	DeepCopyMap(dictionary StringKeyMap) StringKeyMap

	DeepCopyList(array []interface{}) []interface{}
}

var sharedCopier Copier = &DataCopier{}

func SetCopier(copier Copier) {
	sharedCopier = copier
}

/**
 *  Shallow Copy
 *  ~~~~~~~~~~~~
 */
func Copy(object interface{}) interface{} {
	return sharedCopier.Copy(object)
}

func CopyMap(dictionary StringKeyMap) StringKeyMap {
	return sharedCopier.CopyMap(dictionary)
}

func CopyList(array []interface{}) []interface{} {
	return sharedCopier.CopyList(array)
}

/**
 *  Deep Copy
 *  ~~~~~~~~~
 */
func DeepCopy(object interface{}) interface{} {
	return sharedCopier.DeepCopy(object)
}

func DeepCopyMap(dictionary StringKeyMap) StringKeyMap {
	return sharedCopier.DeepCopyMap(dictionary)
}

func DeepCopyList(array []interface{}) []interface{} {
	return sharedCopier.DeepCopyList(array)
}

/**
 *  Default Data Copier
 */
type DataCopier struct {
	//Copier
}

// Override
func (cp DataCopier) Copy(object interface{}) interface{} {
	target, rv := ObjectReflectValue(object)
	if target == nil {
		return nil
	}
	switch v := target.(type) {
	case Cloneable:
		return v.Clone()
	case Mapper:
		return CopyMap(v.Map())
	case StringKeyMap:
		return CopyMap(v)
	case []interface{}:
		return CopyList(v)
	}
	// other types
	switch rv.Kind() {
	case reflect.Map:
		return CopyMap(reflectMap(rv))
	case reflect.Array, reflect.Slice:
		return CopyList(reflectList(rv))
	default:
		return target
	}
}

// Override
func (cp DataCopier) DeepCopy(object interface{}) interface{} {
	target, rv := ObjectReflectValue(object)
	if target == nil {
		return nil
	}
	switch v := target.(type) {
	case Cloneable:
		return v.Clone()
	case Mapper:
		return DeepCopyMap(v.Map())
	case StringKeyMap:
		return DeepCopyMap(v)
	case []interface{}:
		return DeepCopyList(v)
	}
	// other types
	switch rv.Kind() {
	case reflect.Map:
		return DeepCopyMap(reflectMap(rv))
	case reflect.Array, reflect.Slice:
		return DeepCopyList(reflectList(rv))
	default:
		return target
	}
}

// Override
func (cp DataCopier) CopyMap(dictionary StringKeyMap) StringKeyMap {
	clone := NewMap()
	for key, value := range dictionary {
		clone[key] = value
	}
	return clone
}

// Override
func (cp DataCopier) DeepCopyMap(dictionary StringKeyMap) StringKeyMap {
	clone := NewMap()
	for key, value := range dictionary {
		clone[key] = DeepCopy(value)
	}
	return clone
}

// Override
func (cp DataCopier) CopyList(array []interface{}) []interface{} {
	clone := make([]interface{}, len(array))
	for key, value := range array {
		clone[key] = value
	}
	return clone
}

// Override
func (cp DataCopier) DeepCopyList(array []interface{}) []interface{} {
	clone := make([]interface{}, len(array))
	for key, value := range array {
		clone[key] = DeepCopy(value)
	}
	return clone
}
