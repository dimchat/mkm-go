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
 *  Broadcast Address
 */

type BroadcastAddress struct {
	ConstantString
	Address

	_network uint8
}

func NewBroadcastAddress(address string, network NetworkType) *BroadcastAddress {
	return new(BroadcastAddress).Init(address, network)
}

func (address *BroadcastAddress) Init(string string, network NetworkType) *BroadcastAddress {
	if address.ConstantString.Init(string) != nil {
		address._network = uint8(network)
	}
	return address
}

func (address *BroadcastAddress) String() string {
	return address.ConstantString.String()
}

func (address *BroadcastAddress) Equal(other interface{}) bool {
	return address.ConstantString.Equal(other)
}

func (address *BroadcastAddress) Network() uint8 {
	return address._network
}

func (address *BroadcastAddress) IsUser() bool {
	return NetworkTypeIsUser(address._network)
}

func (address *BroadcastAddress) IsGroup() bool {
	return NetworkTypeIsGroup(address._network)
}

func (address *BroadcastAddress) IsBroadcast() bool {
	return true
}

func CreateBroadcastAddresses() {
	if ANYWHERE == nil {
		ANYWHERE = NewBroadcastAddress(Anywhere, MAIN)
	}
	if EVERYWHERE == nil {
		EVERYWHERE = NewBroadcastAddress(Everywhere, GROUP)
	}
}

/**
 *  Address Factory
 */
type IAddressFactory interface {
	AddressFactory

	// override for creating address from string
	CreateAddress(address string) Address
}

type BaseAddressFactory struct {
	IAddressFactory

	_addresses map[string]Address
}

func (factory *BaseAddressFactory) Init() *BaseAddressFactory {
	addresses := make(map[string]Address)
	// cache broadcast addresses
	CreateBroadcastAddresses()
	addresses[ANYWHERE.String()] = ANYWHERE
	addresses[EVERYWHERE.String()] = EVERYWHERE
	factory._addresses = addresses
	return factory
}

func (factory *BaseAddressFactory) ParseAddress(address string) Address {
	add := factory._addresses[address]
	if add == nil {
		add = factory.CreateAddress(address)
		if add != nil {
			factory._addresses[address] = add
		}
	}
	return add
}
