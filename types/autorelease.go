/* license: https://mit-license.org
 * ==============================================================================
 * The MIT License (MIT)
 *
 * Copyright (c) 2021 Albert Moky
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

import "sync"

type AutoreleasePool interface {
	Object
	IAutoreleasePool
}
type IAutoreleasePool interface {

	/**
	 *  Append object to the AutoreleasePool
	 *
	 * @param obj - MRC object
	 */
	Append(obj Object)

	/**
	 *  Release all objects in the AutoreleasePool
	 */
	Purge()
}

type BasePool struct {
	BaseObject
	IAutoreleasePool

	_objects []Object
}

func (pool *BasePool) Init() *BasePool {
	if pool.BaseObject.Init() != nil {
		pool.Purge()
	}
	return pool
}

func (pool *BasePool) Release() int {
	cnt := pool.BaseObject.Release()
	if cnt == 0 {
		// this object is going to be destroyed,
		// release children
		pool.setObjects(nil)
	}
	return cnt
}

func (pool *BasePool) Append(obj Object) {
	pool._objects = append(pool._objects, obj)
}

func (pool *BasePool) Purge() {
	pool.setObjects(make([]Object, 0, 128))
}

func (pool *BasePool) setObjects(objects []Object) {
	if pool._objects != nil {
		for _, item := range pool._objects {
			ObjectRelease(item)
		}
	}
	pool._objects = objects
}

//
//  Pool Stack
//

var poolStack = make([]AutoreleasePool, 0, 16)

var mutexLock sync.Mutex

func AutoreleasePoolLock() {
	mutexLock.Lock()
}
func AutoreleasePoolUnLock() {
	mutexLock.Unlock()
}

// Append an AutoreleasePool on the stack top
func AutoreleasePoolPush(pool AutoreleasePool) {
	AutoreleasePoolLock()
	poolStack = append(poolStack, pool)
	AutoreleasePoolUnLock()
}

// Remove the AutoreleasePool from stack
// if pool is nil, pop the top one
func AutoreleasePoolPop(pool AutoreleasePool) AutoreleasePool {
	AutoreleasePoolLock()
	index := len(poolStack) - 1
	if index < 0 {
		panic("AutoreleasePool stack empty")
	}
	if pool == nil {
		// pop one on the top
		pool = poolStack[index]
		poolStack = poolStack[:index]
	} else if pool == poolStack[index] {
		// found on the top
		poolStack = poolStack[:index]
	} else if pool == poolStack[0] {
		// found under the bottom
		poolStack = poolStack[1:]
	} else {
		index--  // skip the top
		for ; index > 0; index-- {
			if pool == poolStack[index] {
				poolStack = append(poolStack[:index], poolStack[index+1:]...)
				break
			}
		}
	}
	AutoreleasePoolUnLock()
	return pool
}

// Get an AutoreleasePool on the stack top
func AutoreleasePoolTop() AutoreleasePool {
	AutoreleasePoolLock()
	pool := autoreleasePoolTop()
	AutoreleasePoolUnLock()
	return pool
}
func autoreleasePoolTop() AutoreleasePool {
	return poolStack[len(poolStack)-1]
}

// Append an Object to a AutoreleasePool on the stack top synchronously
func AutoreleasePoolAppend(obj Object) Object {
	AutoreleasePoolLock()
	autoreleasePoolTop().Append(obj)
	AutoreleasePoolUnLock()
	return obj
}

// Purge the AutoreleasePool on the stack top synchronously
func AutoreleasePoolPurge() {
	AutoreleasePoolLock()
	autoreleasePoolTop().Purge()
	AutoreleasePoolUnLock()
}

// Purge all AutoreleasePool in the stack synchronously
func AutoreleasePoolPurgeAll() {
	AutoreleasePoolLock()
	index := len(poolStack) - 1
	for ; index >= 0; index-- {
		poolStack[index].Purge()
	}
	AutoreleasePoolUnLock()
}

// Create a new AutoreleasePool and push to the stack top
// the caller should release it manually
func NewAutoreleasePool() AutoreleasePool {
	pool := new(BasePool).Init()
	ObjectRetain(pool)
	AutoreleasePoolPush(pool)
	return pool
}

// Remove the AutoreleasePool from stack and purge it
func DeleteAutoreleasePool(pool AutoreleasePool) {
	AutoreleasePoolPop(pool)
	//pool.Purge()
	ObjectRelease(pool)
}

func init() {
	// create a default AutoreleasePool and keep it under the stack bottom
	NewAutoreleasePool()
}
