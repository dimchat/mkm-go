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

/*
 *  @enum MKMNetworkID
 *
 *  @abstract A network type to indicate what kind the entity is.
 *
 *  @discussion An address can identify a person, a group of people,
 *      a team, even a thing.
 *
 *      MKMNetwork_Main indicates this entity is a person's account.
 *      An account should have a public key, which proved by meta data.
 *
 *      MKMNetwork_Group indicates this entity is a group of people,
 *      which should have a founder (also the owner), and some members.
 *
 *      MKMNetwork_Moments indicates a special personal social network,
 *      where the owner can share information and interact with its friends.
 *      The owner is the king here, it can do anything and no one can stop it.
 *
 *      MKMNetwork_Polylogue indicates a virtual (temporary) social network.
 *      It's created to talk with multi-people (but not too much, e.g. less than 100).
 *      Any member can invite people in, but only the founder can expel member.
 *
 *      MKMNetwork_Chatroom indicates a massive (persistent) social network.
 *      It's usually more than 100 people in it, so we need administrators
 *      to help the owner to manage the group.
 *
 *      MKMNetwork_SocialEntity indicates this entity is a social entity.
 *
 *      MKMNetwork_Organization indicates an independent organization.
 *
 *      MKMNetwork_Company indicates this entity is a company.
 *
 *      MKMNetwork_School indicates this entity is a school.
 *
 *      MKMNetwork_Government indicates this entity is a government department.
 *
 *      MKMNetwork_Department indicates this entity is a department.
 *
 *      MKMNetwork_Thing this is reserved for IoT (Internet of Things).
 *
 *  Bits:
 *      0000 0001 - this entity's branch is independent (clear division).
 *      0000 0010 - this entity can contains other group (big organization).
 *      0000 0100 - this entity is top organization.
 *      0000 1000 - (Main) this entity acts like a human.
 *
 *      0001 0000 - this entity contains members (Group)
 *      0010 0000 - this entity needs other administrators (big organization)
 *      0100 0000 - this is an entity in reality.
 *      1000 0000 - (IoT) this entity is a 'Thing'.
 *
 *      (All above are just some advices to help choosing numbers :P)
 */
type NetworkType uint8

const (
	BTCMain      NetworkType = 0x00  // 0000 0000
	//BTCTest    NetworkType = 0x6F  // 0110 1111

	/*
	 *  Person Account
	 */
	MAIN         NetworkType = 0x08  // 0000 1000 (Person)

	/*
	 *  Virtual Groups
	 */
	GROUP        NetworkType = 0x10  // 0001 0000 (Multi-Persons)

	//MOMENTS    NetworkType = 0x18  // 0001 1000 (Twitter)
	POLYLOGUE    NetworkType = 0x10  // 0001 0000 (Multi-Persons Chat, N < 100)
	CHATROOM     NetworkType = 0x30  // 0011 0000 (Multi-Persons Chat, N >= 100)

	/*
	 *  Social Entities in Reality
	 */
	//SOCIAL       NetworkType = 0x50  // 0101 0000

	//ORGANIZATION NetworkType = 0x74  // 0111 0100
	//COMPANY      NetworkType = 0x76  // 0111 0110
	//SCHOOL       NetworkType = 0x77  // 0111 0111
	//GOVERNMENT   NetworkType = 0x73  // 0111 0011
	//DEPARTMENT   NetworkType = 0x52  // 0101 0010

	/*
	 *  Network
	 */
	PROVIDER     NetworkType = 0x76  // 0111 0110 (Service Provider)
	STATION      NetworkType = 0x88  // 1000 1000 (Server Node)

	/*
	 *  Internet of Things
	 */
	THING        NetworkType = 0x80  // 1000 0000 (IoT)
	ROBOT        NetworkType = 0xC8  // 1100 1000
)

func NetworkTypeIsUser(networkType NetworkType) bool {
	return (networkType & MAIN) == MAIN || networkType == BTCMain
}

func NetworkTypeIsGroup(networkType NetworkType) bool {
	return (networkType & GROUP) == GROUP
}

func NetworkTypeParse(network interface{}) NetworkType {
	if ValueIsNil(network) {
		return 0
	}
	return NetworkType(network.(float64))
}

func (network NetworkType) String() string {
	text := NetworkTypeGetAlias(network)
	if text == "" {
		text = fmt.Sprintf("Network(%d)", network)
	}
	return text
}

func NetworkTypeGetAlias(network NetworkType) string {
	return networkNames[network]
}
func NetworkTypeSetAlias(network NetworkType, alias string) {
	networkNames[network] = alias
}

var networkNames = make(map[NetworkType]string, 8)

func init() {
	NetworkTypeSetAlias(BTCMain, "BTCMain")
	//NetworkTypeSetAlias(BTCTest, "BTCTest")

	NetworkTypeSetAlias(MAIN, "MAIN")
	NetworkTypeSetAlias(GROUP, "GROUP")

	//NetworkTypeSetAlias(MOMENTS, "MOMENTS")
	//NetworkTypeSetAlias(POLYLOGUE, "POLYLOGUE")
	NetworkTypeSetAlias(CHATROOM, "CHATROOM")

	//NetworkTypeSetAlias(SOCIAL, "SOCIAL")

	//NetworkTypeSetAlias(ORGANIZATION, "ORGANIZATION")
	//NetworkTypeSetAlias(COMPANY, "COMPANY")
	//NetworkTypeSetAlias(SCHOOL, "SCHOOL")
	//NetworkTypeSetAlias(GOVERNMENT, "GOVERNMENT")
	//NetworkTypeSetAlias(DEPARTMENT, "DEPARTMENT")

	NetworkTypeSetAlias(PROVIDER, "PROVIDER")
	NetworkTypeSetAlias(STATION, "STATION")

	NetworkTypeSetAlias(THING, "THING")
	NetworkTypeSetAlias(ROBOT, "ROBOT")
}
