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
	. "github.com/dimchat/mkm-go/ext"
	. "github.com/dimchat/mkm-go/types"
)

/**
 *  ID for entity (User/Group)
 *
 *      data format: "name@address[/terminal]"
 *
 *      fields:
 *          name     - entity name, the seed of fingerprint to build address
 *          address  - a string to identify an entity
 *          terminal - entity login resource(device), OPTIONAL
 */
type ID interface {
	Stringer

	Name() string
	Address() Address
	Terminal() string

	/**
	 *  get ID type
	 *
	 * @return network type
	 */
	Type() EntityType

	IsUser() bool
	IsGroup() bool
	IsBroadcast() bool
}

/**
 *  ID Factory
 *  ~~~~~~~~~~
 */
type IDFactory interface {

	/**
	 *  Generate ID
	 *
	 * @param meta - meta info
	 * @param network - ID.type
	 * @param terminal - ID.terminal
	 * @return ID
	 */
	GenerateID(meta Meta, network EntityType, terminal string) ID

	/**
	 *  Create ID
	 *
	 * @param name     - ID.name
	 * @param address  - ID.address
	 * @param terminal - ID.terminal
	 * @return ID
	 */
	CreateID(name string, address Address, terminal string) ID

	/**
	 *  Parse string object to ID
	 *
	 * @param did - ID string
	 * @return ID
	 */
	ParseID(did string) ID
}

//
//  Factory methods
//

func GenerateID(meta Meta, network EntityType, terminal string) ID {
	helper := GetIDHelper()
	return helper.GenerateID(meta, network, terminal)
}

func CreateID(name string, address Address, terminal string) ID {
	helper := GetIDHelper()
	return helper.CreateID(name, address, terminal)
}

func ParseID(did interface{}) ID {
	helper := GetIDHelper()
	return helper.ParseID(did)
}

func GetIDFactory() IDFactory {
	helper := GetIDHelper()
	return helper.GetIDFactory()
}

func SetIDFactory(factory IDFactory) {
	helper := GetIDHelper()
	helper.SetIDFactory(factory)
}

//
//  Conveniences
//

/**
 *  Convert ID list from string array
 *
 * @param array - string array
 * @return ID list
 */
func IDConvert(array interface{}) []ID {
	members := FetchList(array)
	identifiers := make([]ID, 0, len(members))
	var did ID
	for _, item := range members {
		did = ParseID(item)
		if did == nil {
			continue
		}
		identifiers = append(identifiers, did)
	}
	return identifiers
}

/**
 *  Revert ID list to string array
 *
 * @param identifiers - ID list
 * @return string array
 */
func IDRevert(identifiers []ID) []string {
	array := make([]string, len(identifiers))
	for idx, did := range identifiers {
		array[idx] = did.String()
	}
	return array
}
