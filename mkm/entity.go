/* license: https://mit-license.org
 *
 *  Ming-Ke-Ming : Decentralized User Identity Authentication
 *
 *                                Written in 2020 by Moky <albert.moky@gmail.com>
 *
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
package mkm

import (
	. "github.com/dimchat/mkm-go/protocol"
	. "github.com/dimchat/mkm-go/types"
)

/**
 *  Entity (User/Group)
 *  ~~~~~~~~~~~~~~~~~~~
 *  Base class of User and Group, ...
 *
 *  properties:
 *      identifier - entity ID
 *      type       - entity type
 *      number     - search number
 *      meta       - meta for generate ID
 *      profile    - entity profile
 *      name       - nickname
 */
type Entity struct {
	Object

	_identifier *ID

	_delegate *EntityDataSource
}

func (entity *Entity) Init(identifier *ID) *Entity {
	entity._identifier = identifier
	return entity
}

func (entity *Entity) Equal(other interface{}) bool {
	ptr, ok := other.(*Entity)
	if !ok {
		obj, ok := other.(Entity)
		if !ok {
			return false
		}
		ptr = &obj
	}
	if *entity == *ptr {
		return true
	}
	// check by ID
	return entity.ID().Equal(ptr.ID())
}

func (entity *Entity) ID() *ID {
	return entity._identifier
}

func (entity *Entity) GetDataSource() *EntityDataSource {
	return entity._delegate
}

func (entity *Entity) SetDataSource(delegate interface{}) {
	ds, ok := delegate.(*EntityDataSource)
	if ok {
		entity._delegate = ds
	} else {
		panic("entity data source error")
	}
}

/**
 *  Get entity type
 *
 * @return ID(address) type as entity type
 */
func (entity *Entity) Type() NetworkType {
	return entity.ID().Type()
}

func (entity *Entity) GetMeta() *Meta {
	delegate := entity.GetDataSource()
	return (*delegate).GetMeta(entity.ID())
}

func (entity *Entity) GetProfile() *Profile {
	delegate := entity.GetDataSource()
	return (*delegate).GetProfile(entity.ID())
}

/**
 *  Get entity name
 *
 * @return name string
 */
func (entity *Entity) GetName() string {
	// get from profile
	profile := entity.GetProfile()
	if profile != nil {
		name := profile.GetName()
		if name != "" {
			return name
		}
	}
	// get ID.name
	return entity.ID().Name()
}
