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
 */
type Address interface {
	Stringer

	/**
	 *  get address type
	 *
	 * @return network type
	 */
	Network() EntityType
}

/**
 *  Address Factory
 *  ~~~~~~~~~~~~~~~
 */
type AddressFactory interface {

	/**
	 *  Generate address with meta & network
	 *
	 * @param meta - meta info
	 * @param network - address type
	 * @return Address
	 */
	GenerateAddress(meta Meta, network EntityType) Address

	/**
	 *  Parse string object to address
	 *
	 * @param address - address string
	 * @return Address
	 */
	ParseAddress(address string) Address
}

//
//  Factory methods
//

func GenerateAddress(meta Meta, network EntityType) Address {
	helper := GetAddressHelper()
	return helper.GenerateAddress(meta, network)
}

func ParseAddress(address interface{}) Address {
	helper := GetAddressHelper()
	return helper.ParseAddress(address)
}

func GetAddressFactory() AddressFactory {
	helper := GetAddressHelper()
	return helper.GetAddressFactory()
}

func SetAddressFactory(factory AddressFactory) {
	helper := GetAddressHelper()
	helper.SetAddressFactory(factory)
}
