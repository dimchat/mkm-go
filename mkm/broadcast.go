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

const (
	Moky       = "moky"
	Anyone     = "anyone"
	Everyone   = "everyone"

	Anywhere   = "anywhere"
	Everywhere = "everywhere"
)

//
//  Broadcast Address for User/Group
//
var ANYWHERE   = NewBroadcastAddress(Anywhere, ANY)      // "anywhere"
var EVERYWHERE = NewBroadcastAddress(Everywhere, EVERY)  // "everywhere"

//
//  Broadcast ID for User/Group
//
var FOUNDER    = NewID(Moky, ANYWHERE, "")        // "moky@anywhere"
var ANYONE     = NewID(Anyone, ANYWHERE, "")      // "anyone@anywhere"
var EVERYONE   = NewID(Everyone, EVERYWHERE, "")  // "everyone@everywhere"

/**
 *  Broadcast Address
 */
type BroadcastAddress struct {
	//Address
	ConstantString

	_network EntityType
}

func (addr *BroadcastAddress) Init(address string, network EntityType) Address {
	if addr.ConstantString.InitWithString(address) != nil {
		addr._network = network
	}
	return addr
}

//-------- IAddress

// Override
func (addr *BroadcastAddress) Network() EntityType {
	return addr._network
}

//
//  Creation
//

func NewBroadcastAddress(address string, network EntityType) Address {
	broadcast := &BroadcastAddress{}
	return broadcast.Init(address, network)
}
