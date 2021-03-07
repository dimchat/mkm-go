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
 *  Function for creating address from string
 *
 * @param address - address string
 * @return Address
 */
type AddressCreator func(address string) Address

/**
 *  General Address Factory
 *  ~~~~~~~~~~~~~~~~~~~~~~~
 */
type GeneralAddressFactory struct {
	AddressFactory

	CreateAddress AddressCreator

	_addresses map[string]Address
}

func NewGeneralAddressFactory(fn AddressCreator) *GeneralAddressFactory {
	return new(GeneralAddressFactory).Init(fn)
}

func (factory *GeneralAddressFactory) Init(fn AddressCreator) *GeneralAddressFactory {
	factory.CreateAddress = fn
	factory._addresses = make(map[string]Address)
	// cache broadcast addresses
	factory.cacheAddress(ANYWHERE.String(), ANYWHERE)
	factory.cacheAddress(EVERYWHERE.String(), EVERYWHERE)
	return factory
}

func (factory *GeneralAddressFactory) cacheAddress(str string, address Address) {
	factory._addresses[str] = address
}

func (factory *GeneralAddressFactory) ParseAddress(address string) Address {
	addr := factory._addresses[address]
	if addr == nil {
		addr = factory.CreateAddress(address)
		if addr != nil {
			factory.cacheAddress(address, addr)
		}
	}
	return addr
}
