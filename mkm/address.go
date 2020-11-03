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
	"fmt"
	. "github.com/dimchat/mkm-go/protocol"
	. "github.com/dimchat/mkm-go/types"
	"unsafe"
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
	fmt.Stringer
	Object

	/**
	 *  get address type
	 *
	 * @return Network ID
	 */
	Type() NetworkType
}

func AddressIsUser(address *Address) bool {
	network := (*address).Type()
	return NetworkTypeIsUser(network)
}

func AddressIsGroup(address *Address) bool {
	network := (*address).Type()
	return NetworkTypeIsGroup(network)
}

func AddressIsBroadcast(address *Address) bool {
	network := (*address).Type()
	if network == MAIN {
		return (*address).String() == anywhere
	} else if network == GROUP {
		return (*address).String() == everywhere
	}
	return false
}

/**
 *  Address for broadcast
 */
const (
	anywhere = "anywhere"
	everywhere = "everywhere"
)

var ANYWHERE = createBroadcastAddress(anywhere, MAIN)
var EVERYWHERE = createBroadcastAddress(everywhere, GROUP)

func createBroadcastAddress(string string, network NetworkType) *Address {
	address := new(broadcastAddress)
	address.Init(string, network)
	return (*Address)(unsafe.Pointer(address))
}

type broadcastAddress struct {
	String
	Address

	_network NetworkType
}

func (address *broadcastAddress) Init(string string, network NetworkType) *broadcastAddress {
	address.String.Init(string)
	address._network = network
	return address
}

func (address *broadcastAddress) Type() NetworkType {
	return address._network
}
