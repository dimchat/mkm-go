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
 *  ~~~~~~~~~~~~~~~~~~~~~~~~~~
 *
 *      data format: "name@address[/terminal]"
 *
 *      fields:
 *          name     - entity name, the seed of fingerprint to build address
 *          address  - a string to identify an entity
 *          terminal - entity login resource(device), OPTIONAL
 */
type Identifier struct {
	//ID
	ConstantString

	_name     string
	_address  Address
	_terminal string
}

func (id *Identifier) Init(identifier string, name string, address Address, terminal string) ID {
	id.ConstantString.Init(identifier)
	id._name = name
	id._address = address
	id._terminal = terminal
	return id
}

//-------- ID

// Override
func (id *Identifier) Name() string {
	return id._name
}

// Override
func (id *Identifier) Address() Address {
	return id._address
}

// Override
func (id *Identifier) Terminal() string {
	return id._terminal
}

// Override
func (id *Identifier) Type() EntityType {
	return id._address.Network()
}

// Override
func (id *Identifier) IsUser() bool {
	network := id._address.Network()
	return EntityTypeIsUser(network)
}

// Override
func (id *Identifier) IsGroup() bool {
	network := id._address.Network()
	return EntityTypeIsGroup(network)
}

// Override
func (id *Identifier) IsBroadcast() bool {
	network := id._address.Network()
	return EntityTypeIsBroadcast(network)
}

//
//  Creation
//

func NewIdentifier(name string, address Address, terminal string) ID {
	identifier := IdentifierConcat(name, address, terminal)
	id := &Identifier{}
	id.Init(identifier, name, address, terminal)
	return id
}

func IdentifierConcat(name string, address Address, terminal string) string {
	str := address.String()
	if name != "" {
		str = name + "@" + str
	}
	if terminal != "" {
		str = str + "/" + terminal
	}
	return str
}
