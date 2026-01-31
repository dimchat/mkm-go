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

func (did *Identifier) Init(identifier string, name string, address Address, terminal string) ID {
	if did.ConstantString.InitWithString(identifier) != nil {
		did._name = name
		did._address = address
		did._terminal = terminal
	}
	return did
}

//-------- ID

// Override
func (did *Identifier) Name() string {
	return did._name
}

// Override
func (did *Identifier) Address() Address {
	return did._address
}

// Override
func (did *Identifier) Terminal() string {
	return did._terminal
}

// Override
func (did *Identifier) Type() EntityType {
	return did._address.Network()
}

// Override
func (did *Identifier) IsUser() bool {
	network := did._address.Network()
	return EntityTypeIsUser(network)
}

// Override
func (did *Identifier) IsGroup() bool {
	network := did._address.Network()
	return EntityTypeIsGroup(network)
}

// Override
func (did *Identifier) IsBroadcast() bool {
	network := did._address.Network()
	return EntityTypeIsBroadcast(network)
}

//
//  Creation
//

func NewID(name string, address Address, terminal string) ID {
	identifier := IDConcat(name, address, terminal)
	did := &Identifier{}
	return did.Init(identifier, name, address, terminal)
}

func IDConcat(name string, address Address, terminal string) string {
	str := address.String()
	if name != "" {
		str = name + "@" + str
	}
	if terminal != "" {
		str = str + "/" + terminal
	}
	return str
}
