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
 *  ID for entity (User/Group)
 *
 *      data format: "name@address[/terminal]"
 *
 *      fields:
 *          name     - entity name, the seed of fingerprint to build address
 *          address  - a string to identify an entity
 *          terminal - entity login resource(device), OPTIONAL
 */
type ID struct {
	String

	_name string
	_address *Address
	_terminal string
}

func (id *ID) Init(string string, name string, address *Address, terminal string) *ID {
	if id.String.Init(string) != nil {
		id._name = name
		id._address = address
		id._terminal = terminal
	}
	return id
}

func (id *ID) Name() string {
	return id._name
}

func (id *ID) Terminal() string {
	return id._terminal
}

/**
 *  Get ID Address
 *
 * @return address
 */
func (id *ID) Address() *Address {
	return id._address
}

/**
 *  Get Network ID
 *
 * @return address type as network ID
 */
func (id *ID) Type() NetworkType {
	address := id.Address()
	return (*address).Type()
}

func (id *ID) IsUser() bool {
	address := id.Address()
	return AddressIsUser(address)
}

func (id *ID) IsGroup() bool {
	address := id.Address()
	return AddressIsGroup(address)
}

func (id *ID) IsBroadcast() bool {
	address := id.Address()
	return AddressIsBroadcast(address)
}

func (id *ID) Equal(other interface{}) bool {
	//if (*id).String.Equal(other) {
	//	return true
	//}
	ptr, ok := other.(*ID)
	if !ok {
		obj, ok := other.(ID)
		if !ok {
			return false
		}
		ptr = &obj
	}
	if *id == *ptr {
		return true
	}

	// check ID.name
	if id.Name() != ptr.Name() {
		return false
	}
	// check ID.address
	addr1 := id.Address()
	addr2 := ptr.Address()
	return (*addr1).Equal(addr2)
}

func CreateID(name string, address *Address, terminal string) *ID {
	identifier := (*address).String()
	if name != "" {
		identifier = name + "@" + identifier
	}
	if terminal != "" {
		identifier = identifier + "/" + terminal
	}
	return new(ID).Init(identifier, name, address, terminal)
}

/**
 *  ID for broadcast
 */
const (
	anyone = "anyone"
	everyone = "everyone"
)

var ANYONE = CreateID(anyone, ANYWHERE, "")
var EVERYONE = CreateID(everyone, EVERYWHERE, "")
