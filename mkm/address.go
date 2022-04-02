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
 */
type BaseAddress struct {
	ConstantString

	_network NetworkType
}

func (address *BaseAddress) Init(string string, network NetworkType) *BaseAddress {
	if address.ConstantString.Init(string) != nil {
		address._network = network
	}
	return address
}

//-------- IAddress

func (address *BaseAddress) Network() NetworkType {
	return address._network
}

func (address *BaseAddress) IsUser() bool {
	return NetworkTypeIsUser(address._network)
}

func (address *BaseAddress) IsGroup() bool {
	return NetworkTypeIsGroup(address._network)
}

func (address *BaseAddress) IsBroadcast() bool {
	//panic("not implemented")
	return false
}

/**
 *  General Address Factory
 *  ~~~~~~~~~~~~~~~~~~~~~~~
 */
type GeneralAddressFactory struct {

	_create AddressCreator

	_addresses map[string]Address
}

type AddressCreator func(address string) Address

func (factory *GeneralAddressFactory) Init(fn AddressCreator) *GeneralAddressFactory {
	factory._create = fn
	factory._addresses = make(map[string]Address)
	// cache broadcast addresses
	factory._addresses[ANYWHERE.String()] = ANYWHERE
	factory._addresses[EVERYWHERE.String()] = EVERYWHERE
	return factory
}

//-------- IAddressFactory

func (factory *GeneralAddressFactory) GenerateAddress(meta Meta, network NetworkType) Address {
	address := meta.GenerateAddress(network)
	if address != nil {
		factory._addresses[address.String()] = address
	}
	return address
}

func (factory *GeneralAddressFactory) CreateAddress(address string) Address {
	return factory._create(address)
}

func (factory *GeneralAddressFactory) ParseAddress(address string) Address {
	addr := factory._addresses[address]
	if addr == nil {
		addr = factory.CreateAddress(address)
		if addr != nil {
			factory._addresses[address] = addr
		}
	}
	return addr
}
