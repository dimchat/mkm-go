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
 *  Address for MKM ID
 *  ~~~~~~~~~~~~~~~~~~
 *  This class is used to build address for ID
 *
 *      properties:
 *          network - address type
 *          number  - search number
 */
type Address interface {
	Stringer

	/**
	 *  get address type
	 *
	 * @return network type
	 */
	Type() NetworkType
}

func AddressesEqual(address1, address2 Address) bool {
	return StringsEqual(address1, address2)
}

func AddressIsUser(address Address) bool {
	network := address.Type()
	return NetworkTypeIsUser(network)
}

func AddressIsGroup(address Address) bool {
	network := address.Type()
	return NetworkTypeIsGroup(network)
}

func AddressIsBroadcast(address Address) bool {
	network := address.Type()
	if network == MAIN {
		return AddressesEqual(address, ANYWHERE)
	} else if network == GROUP {
		return AddressesEqual(address, EVERYWHERE)
	} else {
		return false
	}
}

/**
 *  Address for broadcast
 */
const (
	anywhere = "anywhere"
	everywhere = "everywhere"
)

var ANYWHERE = newBroadcastAddress(anywhere, MAIN)
var EVERYWHERE = newBroadcastAddress(everywhere, GROUP)

func newBroadcastAddress(string string, network NetworkType) Address {
	address := new(broadcastAddress)
	address.Init(string, network)
	return address
}

type broadcastAddress struct {
	ConstantString

	_network NetworkType
}

func (address *broadcastAddress) Init(string string, network NetworkType) *broadcastAddress {
	if address.ConstantString.Init(string) != nil {
		address._network = network
	}
	return address
}

func (address broadcastAddress) Type() NetworkType {
	return address._network
}
