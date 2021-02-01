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
	"fmt"
	. "github.com/dimchat/mkm-go/types"
)

/**
 *  Address for MKM ID
 *  ~~~~~~~~~~~~~~~~~~
 *  This class is used to build address for ID
 */
type Address interface {
	Stringer

	/**
	 *  get address type
	 *
	 * @return network type
	 */
	Network() uint8

	IsUser() bool
	IsGroup() bool
	IsBroadcast() bool
}

/**
 *  Address Factory
 *  ~~~~~~~~~~~~~~~
 */
type AddressFactory interface {

	/**
	 *  Parse string object to address
	 *
	 * @param address - address string
	 * @return Address
	 */
	ParseAddress(address string) Address
}

var addressFactory AddressFactory = nil

func AddressSetFactory(factory AddressFactory) {
	addressFactory = factory
}

func AddressGetFactory() AddressFactory {
	return addressFactory
}

//
//  Factory method
//
func AddressParse(address interface{}) Address {
	if address == nil {
		return nil
	}
	var str string
	value := ObjectValue(address)
	switch value.(type) {
	case Address:
		return value.(Address)
	case fmt.Stringer:
		str = value.(fmt.Stringer).String()
	case string:
		str = value.(string)
	default:
		panic(address)
	}
	factory := AddressGetFactory()
	return factory.ParseAddress(str)
}

/**
 *  Address for broadcast
 */
const (
	Anywhere = "anywhere"
	Everywhere = "everywhere"
)

//
//  These addresses will be created when AddressFactory init
//
var ANYWHERE Address = nil    // "anywhere"
var EVERYWHERE Address = nil  // "everywhere"
