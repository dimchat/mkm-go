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

	/**
	 *  Get ID name
	 *
	 * @return ID.name
	 */
	Name() string

	/**
	 *  Get ID address
	 *
	 * @return ID.address
	 */
	Address() Address

	/**
	 *  get ID type
	 *
	 * @return network type
	 */
	Type() NetworkType
}

func IdentifiersEqual(id1, id2 ID) bool {
	if id1 == id2 {
		return true
	}
	// check ID.name
	if id1.Name() != id2.Name() {
		return false
	}
	// check ID.address
	addr1 := id1.Address()
	addr2 := id2.Address()
	return AddressesEqual(addr1, addr2)
}

type Identifier struct {
	ConstantString
	ID

	_name string
	_address Address
	_terminal string
}

func (id *Identifier) Init(string string, name string, address Address, terminal string) *Identifier {
	if id.ConstantString.Init(string) != nil {
		id._name = name
		id._address = address
		id._terminal = terminal
	}
	return id
}

func (id Identifier) Equal(other interface{}) bool {
	id2 := ObjectValue(other).(ID)
	return IdentifiersEqual(id, id2)
}

func (id Identifier) String() string {
	return id.ConstantString.String()
}

func (id Identifier) Name() string {
	return id._name
}

func (id Identifier) Address() Address {
	return id._address
}

/**
 *  Get ID network type
 *
 * @return address type
 */
func (id Identifier) Type() NetworkType {
	address := id.Address()
	return address.Type()
}

/**
 *  Get ID login point
 *
 * @return ID.terminal
 */
func (id Identifier) Terminal() string {
	return id._terminal
}

func (id Identifier) IsUser() bool {
	address := id.Address()
	return AddressIsUser(address)
}

func (id Identifier) IsGroup() bool {
	address := id.Address()
	return AddressIsGroup(address)
}

func (id Identifier) IsBroadcast() bool {
	address := id.Address()
	return AddressIsBroadcast(address)
}

func NewID(name string, address Address, terminal string) ID {
	identifier := address.String()
	if name != "" {
		identifier = name + "@" + identifier
	}
	if terminal != "" {
		identifier = identifier + "/" + terminal
	}
	return new(Identifier).Init(identifier, name, address, terminal)
}

/**
 *  ID for broadcast
 */
const (
	anyone = "anyone"
	everyone = "everyone"
)

var ANYONE = NewID(anyone, ANYWHERE, "")
var EVERYONE = NewID(everyone, EVERYWHERE, "")
