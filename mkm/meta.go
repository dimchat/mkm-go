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
	. "github.com/dimchat/mkm-go/crypto"
	. "github.com/dimchat/mkm-go/format"
	. "github.com/dimchat/mkm-go/protocol"
	. "github.com/dimchat/mkm-go/types"
)

type IMetaExt interface {
	IMeta

	// Check valid status
	status() int8

	/* Generate address with network(type)
	 *
	 * @param network - ID.type
	 * @return Address
	 */
	GenerateAddress(network uint8) Address
}

/**
 *  Base Meta
 *  ~~~~~~~~~
 *
 *  inheritable by BaseMetaShadow
 */
type BaseMeta struct {
	Dictionary
	IMetaExt

	_shadow IMetaExt

	_type uint8
	_key VerifyKey

	_seed string
	_fingerprint []byte

	_status int8  // 1 for valid, -1 for invalid
}

func (meta *BaseMeta) Init(dict map[string]interface{}) *BaseMeta {
	if meta.Dictionary.Init(dict) != nil {
		meta._shadow = nil
		// lazy load
		meta._type = 0
		meta._key = nil
		meta._seed = ""
		meta._fingerprint = nil
		meta._status = 0
	}
	return meta
}

func (meta *BaseMeta) InitWithType(version uint8, key VerifyKey, seed string, fingerprint []byte) *BaseMeta {
	dict := make(map[string]interface{})
	// meta type
	dict["type"] = version
	// meta key
	info, ok := key.(Map)
	if ok {
		dict["key"] = info.GetMap(false)
	} else {
		// TODO: map[string]interface{} ?
		panic(key)
	}
	// seed
	if seed != "" {
		dict["seed"] = seed
	}
	// fingerprint
	if fingerprint != nil {
		dict["fingerprint"] = Base64Encode(fingerprint)
	}
	if meta.Dictionary.Init(dict) != nil {
		meta._shadow = nil
		// set values
		meta._type = version
		meta._key = key
		meta._seed = seed
		meta._fingerprint = fingerprint
	}
	return meta
}

func (meta *BaseMeta) SetShadow(shadow IMetaExt) {
	meta._shadow = shadow
}
func (meta *BaseMeta) Shadow() IMetaExt {
	return meta._shadow
}

//-------- IMeta

func (meta *BaseMeta) Type() uint8 {
	if meta._type == 0 {
		meta._type = meta.Shadow().Type()
	}
	return meta._type
}

func (meta *BaseMeta) Key() VerifyKey {
	if meta._key == nil {
		meta._key = meta.Shadow().Key()
	}
	return meta._key
}

func (meta *BaseMeta) Seed() string {
	if meta._seed == "" {
		meta._seed = meta.Shadow().Seed()
	}
	return meta._seed
}

func (meta *BaseMeta) Fingerprint() []byte {
	if meta._fingerprint == nil {
		meta._fingerprint = meta.Shadow().Fingerprint()
	}
	return meta._fingerprint
}

func (meta *BaseMeta) IsValid() bool {
	return meta.Shadow().IsValid()
}

func (meta *BaseMeta) GenerateID(network uint8, terminal string) ID {
	return meta.Shadow().GenerateID(network, terminal)
}

func (meta *BaseMeta) MatchID(identifier ID) bool {
	return meta.Shadow().MatchID(identifier)
}

func (meta *BaseMeta) MatchKey(key VerifyKey) bool {
	return meta.Shadow().MatchKey(key)
}

//-------- IMetaExt

func (meta *BaseMeta) status() int8 {
	if meta._status == 0 {
		meta._status = meta.Shadow().status()
	}
	return meta._status
}

func (meta *BaseMeta) GenerateAddress(network uint8) Address {
	return meta.Shadow().GenerateAddress(network)
}
