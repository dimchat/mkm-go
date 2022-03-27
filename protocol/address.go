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
	IAddress
	Stringer
}
type IAddress interface {

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
	IAddressFactory
}
type IAddressFactory interface {

	/**
	 *  Generate address with meta & network
	 *
	 * @param meta - meta info
	 * @param network - address type
	 @ @return Address
	 */
	GenerateAddress(meta Meta, network uint8) Address

	/**
	 *  Create address from string
	 *
	 * @param address - address string
	 * @return Address
	 */
	CreateAddress(address string) Address

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
//  Factory methods
//
func AddressGenerate(meta Meta, network uint8) Address {
	factory := AddressGetFactory()
	if factory == nil {
		panic("address factory not found")
	}
	return factory.GenerateAddress(meta, network)
}

func AddressCreate(address string) Address {
	factory := AddressGetFactory()
	if factory == nil {
		panic("address factory not found")
	}
	return factory.CreateAddress(address)
}

func AddressParse(address interface{}) Address {
	if ValueIsNil(address) {
		return nil
	}
	value, ok := address.(Address)
	if ok {
		return value
	}
	var addr string
	wrapper, ok := address.(fmt.Stringer)
	if ok {
		addr = wrapper.String()
	} else {
		addr, ok = address.(string)
		if !ok {
			panic(address)
		}
	}
	factory := AddressGetFactory()
	if factory == nil {
		panic("address factory not found")
	}
	return factory.ParseAddress(addr)
}

/**
 *  Address for broadcast
 */
const (
	Anywhere = "anywhere"
	Everywhere = "everywhere"
)

//
//  Broadcast addresses for User/Group
//
var ANYWHERE Address = nil    // "anywhere"
var EVERYWHERE Address = nil  // "everywhere"
