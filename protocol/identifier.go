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

func IDsEqual(id1, id2 ID) bool {
	// check ID.address
	addr1 := id1.Address()
	addr2 := id2.Address()
	if addr1.Equal(addr2) == false {
		return false
	}
	// check ID.name
	return id1.Name() == id2.Name()
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
	// create broadcast IDs
	CreateBroadcastIdentifiers()
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
	if identifier == nil {
		return nil
	}
	var str string
	value := ObjectValue(identifier)
	switch value.(type) {
	case ID:
		return value.(ID)
	case fmt.Stringer:
		str = value.(fmt.Stringer).String()
	case string:
		str = value.(string)
	default:
		panic(identifier)
	}
	factory := IDGetFactory()
	return factory.ParseID(str)
}

func IDConvert(members []interface{}) []ID {
	res := make([]ID, len(members))
	for index, item := range members {
		res[index] = IDParse(item)
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
