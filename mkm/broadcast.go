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
	IAddress

	_network uint8
}

func (address *BaseAddress) Init(string string, network NetworkType) *BaseAddress {
	if address.ConstantString.Init(string) != nil {
		address._network = uint8(network)
	}
	return address
}

func (address *BaseAddress) self() Address {
	return address.Self().(Address)
}

//-------- IAddress

func (address *BaseAddress) Network() uint8 {
	return address._network
}

func (address *BaseAddress) IsUser() bool {
	return NetworkTypeIsUser(address.self().Network())
}

func (address *BaseAddress) IsGroup() bool {
	return NetworkTypeIsGroup(address.self().Network())
}

func (address *BaseAddress) IsBroadcast() bool {
	return false
}

/**
 *  Broadcast Address
 */
type BroadcastAddress struct {
	BaseAddress
}

func NewBroadcastAddress(address string, network NetworkType) *BroadcastAddress {
	broadcast := new(BroadcastAddress).Init(address, network)
	ObjectRetain(broadcast)
	return broadcast
}

func (address *BroadcastAddress) Init(string string, network NetworkType) *BroadcastAddress {
	if address.BaseAddress.Init(string, network) != nil {
	}
	return address
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

func init() {
	BuildGeneralIDFactory()
	CreateBroadcastAddresses()
	CreateBroadcastIdentifiers()
}
