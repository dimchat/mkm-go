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
package protocol

import (
	"fmt"
	. "github.com/dimchat/mkm-go/types"
	"reflect"
)

/**
 *  ID for entity (User/Group)
 *
 *      data format: "name@address[/terminal]"
 *
 *      fields:
 *          name     - entity name, the seed of fingerprint to build address
 *          address  - a string to identify an entity
 *          terminal - entity login resource(device), OPTIONAL
 */
type ID interface {
	Stringer
	IIdentifier
}
type IIdentifier interface {

	Name() string
	Address() Address
	Terminal() string

	/**
	 *  get ID type
	 *
	 * @return network type
	 */
	Type() uint8

	IsUser() bool
	IsGroup() bool
	IsBroadcast() bool
}

/**
 *  ID Factory
 *  ~~~~~~~~~~
 */
type IDFactory interface {

	/**
	 *  Create ID
	 *
	 * @param name     - ID.name
	 * @param address  - ID.address
	 * @param terminal - ID.terminal
	 * @return ID
	 */
	CreateID(name string, address Address, terminal string) ID

	/**
	 *  Parse string object to ID
	 *
	 * @param identifier - ID string
	 * @return ID
	 */
	ParseID(identifier string) ID
}

var idFactory IDFactory = nil

func IDSetFactory(factory IDFactory) {
	idFactory = factory
}

func IDGetFactory() IDFactory {
	return idFactory
}

//
//  Factory methods
//
func IDCreate(name string, address Address, terminal string) ID {
	return idFactory.CreateID(name, address, terminal)
}

func IDParse(identifier interface{}) ID {
	if ValueIsNil(identifier) {
		return nil
	}
	id, ok := identifier.(ID)
	if ok {
		return id
	}
	wrapper, ok := identifier.(fmt.Stringer)
	if ok {
		return IDGetFactory().ParseID(wrapper.String())
	}
	text, ok := identifier.(string)
	if ok {
		return IDGetFactory().ParseID(text)
	}
	panic(identifier)
}

func IDConvert(members interface{}) []ID {
	if reflect.TypeOf(members).Kind() != reflect.Slice {
		panic(members)
		return []ID{}
	}
	values := reflect.ValueOf(members)
	count := values.Len()
	res := make([]ID, 0, count)
	var item ID
	for index := 0; index < count; index++ {
		item = IDParse(values.Index(index).Interface())
		if item == nil {
			continue
		}
		res = append(res, item)
	}
	return res
}

func IDRevert(members []ID) []string {
	res := make([]string, len(members))
	for index, item := range members {
		res[index] = item.String()
	}
	return res
}

/**
 *  ID for broadcast
 */
const (
	Moky = "moky"
	Anyone = "anyone"
	Everyone = "everyone"
)

var FOUNDER ID = nil   // "moky@anywhere"
var ANYONE ID = nil    // "anyone@anywhere"
var EVERYONE ID = nil  // "everyone@everywhere"

func CreateBroadcastIdentifiers() {
	if IDGetFactory() == nil {
		return
	}
	if FOUNDER == nil {
		FOUNDER = IDCreate(Moky, ANYWHERE, "")
	}
	if ANYONE == nil {
		ANYONE = IDCreate(Anyone, ANYWHERE, "")
	}
	if EVERYONE == nil {
		EVERYONE = IDCreate(Everyone, EVERYWHERE, "")
	}
}
