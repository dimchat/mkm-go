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

type IMeta interface {
	Meta

	/**
	 *  Generate address with network(type)
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
 *  Abstract method:
 *      GenerateAddress(network uint8) Address
 */
type BaseMeta struct {
	Dictionary
	IMeta

	_type uint8
	_key VerifyKey

	_seed string
	_fingerprint []byte

	_status int8  // 1 for valid, -1 for invalid

	_delegate IMetaDelegate
}

func (meta *BaseMeta) Init(dict map[string]interface{}) *BaseMeta {
	if meta.Dictionary.Init(dict) != nil {
		// lazy load
		meta._type = 0
		meta._key = nil
		meta._seed = ""
		meta._fingerprint = nil
		meta._status = 0
		meta._delegate = nil
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
		// set values
		meta._type = version
		meta._key = key
		meta._seed = seed
		meta._fingerprint = fingerprint
		meta._status = 0
		meta._delegate = nil
	}
	return meta
}

func (meta *BaseMeta) Equal(other interface{}) bool {
	return meta.Dictionary.Equal(other)
}

//-------- Map

func (meta *BaseMeta) Get(name string) interface{} {
	return meta.Dictionary.Get(name)
}

func (meta *BaseMeta) Set(name string, value interface{}) {
	meta.Dictionary.Set(name, value)
}

func (meta *BaseMeta) Keys() []string {
	return meta.Dictionary.Keys()
}

func (meta *BaseMeta) GetMap(clone bool) map[string]interface{} {
	return meta.Dictionary.GetMap(clone)
}

//-------- Meta

func (meta *BaseMeta) Type() uint8 {
	if meta._type == 0 {
		meta._type = MetaGetType(meta.GetMap(false))
	}
	return meta._type
}

func (meta *BaseMeta) Key() VerifyKey {
	if meta._key == nil {
		meta._key = MetaGetKey(meta.GetMap(false))
	}
	return meta._key
}

func (meta *BaseMeta) Seed() string {
	if meta._seed == "" {
		if MetaTypeHasSeed(meta.Type()) {
			meta._seed = MetaGetSeed(meta.GetMap(false))
		}
	}
	return meta._seed
}

func (meta *BaseMeta) Fingerprint() []byte {
	if meta._fingerprint == nil {
		if MetaTypeHasSeed(meta.Type()) {
			meta._fingerprint = MetaGetFingerprint(meta.GetMap(false))
		}
	}
	return meta._fingerprint
}

//-------- IMetaDelegate

func (meta *BaseMeta) SetDelegate(delegate IMetaDelegate) {
	meta._delegate = delegate
}
func (meta *BaseMeta) Delegate() IMetaDelegate {
	return meta._delegate
}

func (meta *BaseMeta) IsValid() bool {
	if meta._status == 0 {
		meta._status = meta.Delegate().Status()
	}
	return meta._status == 1
}

func (meta *BaseMeta) GenerateID(network uint8, terminal string) ID {
	return meta.Delegate().GenerateID(network, terminal)
}

func (meta *BaseMeta) MatchID(identifier ID) bool {
	return meta.Delegate().MatchID(identifier)
}

func (meta *BaseMeta) MatchKey(key VerifyKey) bool {
	return meta.Delegate().MatchKey(key)
}
