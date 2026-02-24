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

import . "github.com/dimchat/mkm-go/types"

// Address defines the interface for MKM ID addresses
//
// Represents a network-specific address used to identify entities (User/Group) in the MKM system
type Address interface {
	Stringer

	// Network returns the network/entity type associated with this address
	// Returns: EntityType representing the network type (e.g., User, Group, Bot, ...)
	Network() EntityType
}

// AddressFactory defines the factory interface for Address
type AddressFactory interface {

	// GenerateAddress creates a new Address instance using meta information and network type
	//
	// Parameters:
	//   - meta: Meta information used to generate the address (fingerprint seed)
	//   - network: EntityType specifying the address/network type
	// Returns: Newly generated Address instance
	GenerateAddress(meta Meta, network EntityType) Address

	// ParseAddress converts an address string into an Address instance
	//
	// Parameters
	//   - address: String representation of the address to parse
	// Returns: Parsed Address instance (nil if parsing fails)
	ParseAddress(address string) Address
}

//
//  Factory methods
//

func GenerateAddress(meta Meta, network EntityType) Address {
	helper := GetAddressHelper()
	return helper.GenerateAddress(meta, network)
}

func ParseAddress(address any) Address {
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
