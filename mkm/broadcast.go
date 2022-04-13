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
)

/**
 *  Broadcast Address
 */
type BroadcastAddress struct {
	BaseAddress
}

func NewBroadcastAddress(address string, network NetworkType) Address {
	broadcast := new(BroadcastAddress)
	broadcast.Init(address, network)
	return broadcast
}

//func (address *BroadcastAddress) Init(string string, network NetworkType) Address {
//	if address.BaseAddress.Init(string, network) != nil {
//	}
//	return address
//}

//-------- IAddress

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
