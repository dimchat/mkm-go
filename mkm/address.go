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
 *  Base Address
 *  ~~~~~~~~~~~~
 *
 *  abstract method:
 *      IsBroadcast()
 */
type BaseAddress struct {
	ConstantString
	IAddress

	_network uint8
}

func (address *BaseAddress) Init(string string, network uint8) *BaseAddress {
	if address.ConstantString.Init(string) != nil {
		address._network = network
	}
	return address
}

//-------- IAddress

func (address *BaseAddress) Network() uint8 {
	return address._network
}

func (address *BaseAddress) IsUser() bool {
	return NetworkTypeIsUser(address._network)
}

func (address *BaseAddress) IsGroup() bool {
	return NetworkTypeIsGroup(address._network)
}

/**
 *  Base Address Factory
 *  ~~~~~~~~~~~~~~~~~~~~
 *
 *  abstract method:
 *      CreateAddress(string)
 */
type BaseAddressFactory struct {
	AddressFactory

	_addresses map[string]Address
}

func (factory *BaseAddressFactory) Init() *BaseAddressFactory {
	factory._addresses = make(map[string]Address)
	// cache broadcast addresses
	factory._addresses[ANYWHERE.String()] = ANYWHERE
	factory._addresses[EVERYWHERE.String()] = EVERYWHERE
	return factory
}

func (factory *BaseAddressFactory) GenerateAddress(meta Meta, network uint8) Address {
	address := meta.GenerateAddress(network)
	if address != nil {
		factory._addresses[address.String()] = address
	}
	return address
}

func (factory *BaseAddressFactory) ParseAddress(address string) Address {
	addr := factory._addresses[address]
	if addr == nil {
		addr = factory.CreateAddress(address)
		if addr != nil {
			factory._addresses[address] = addr
		}
	}
	return addr
}
