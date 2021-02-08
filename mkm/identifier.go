/* license: https://mit-license.org
 *
 *  Ming-Ke-Ming : Decentralized User Identity Authentication
 *
 *                                Written in 2021 by Moky <albert.moky@gmail.com>
 *
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
type Identifier struct {
	ConstantString
	ID

	_name string
	_address Address
	_terminal string
}

func NewIdentifier(identifier string, name string, address Address, terminal string) *Identifier {
	return new(Identifier).Init(identifier, name, address, terminal)
}

func (id *Identifier) Init(string string, name string, address Address, terminal string) *Identifier {
	if id.ConstantString.Init(string) != nil {
		id._name = name
		id._address = address
		id._terminal = terminal
	}
	return id
}

func (id *Identifier) String() string {
	return id.ConstantString.String()
}

func (id *Identifier) Equal(other interface{}) bool {
	var identifier = IDParse(other)
	if identifier == nil {
		return false
	} else if id == identifier {
		return true
	}
	// check ID.address & ID.name
	addr1 := id.Address()
	addr2 := identifier.Address()
	return addr1.Equal(addr2) && id.Name() == identifier.Name()
}

func (id *Identifier) Name() string {
	return id._name
}

func (id *Identifier) Address() Address {
	return id._address
}

func (id *Identifier) Terminal() string {
	return id._terminal
}

func (id *Identifier) Type() uint8 {
	return id._address.Network()
}

func (id *Identifier) IsUser() bool {
	return id._address.IsUser()
}

func (id *Identifier) IsGroup() bool {
	return id._address.IsGroup()
}

func (id *Identifier) IsBroadcast() bool {
	return id._address.IsBroadcast()
}
