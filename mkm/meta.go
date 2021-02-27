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
 * @abstract:
 *      GenerateAddress(network uint8) Address
 */
type BaseMeta struct {
	Dictionary
	IMetaExt

	_type uint8
	_key VerifyKey

	_seed string
	_fingerprint []byte

	_status int8  // 1 for valid, -1 for invalid
}

func (meta *BaseMeta) Init(this Meta, dict map[string]interface{}) *BaseMeta {
	if meta.Dictionary.Init(this, dict) != nil {
		// lazy load
		meta._type = 0
		meta._key = nil
		meta._seed = ""
		meta._fingerprint = nil
		meta._status = 0
	}
	return meta
}

func (meta *BaseMeta) InitWithType(this Map, version uint8, key VerifyKey, seed string, fingerprint []byte) *BaseMeta {
	dict := make(map[string]interface{})
	// meta type
	dict["type"] = version
	// meta key
	dict["key"] = key.GetMap(false)
	// seed
	if seed != "" {
		dict["seed"] = seed
	}
	// fingerprint
	if fingerprint != nil {
		dict["fingerprint"] = Base64Encode(fingerprint)
	}
	if meta.Dictionary.Init(this, dict) != nil {
		// set values
		meta._type = version
		meta._key = key
		meta._seed = seed
		meta._fingerprint = fingerprint
		meta._status = 0
	}
	return meta
}

//-------- IMeta

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
		if MetaTypeHasSeed(meta.Self().(IMeta).Type()) {
			meta._seed = MetaGetSeed(meta.GetMap(false))
		}
	}
	return meta._seed
}

func (meta *BaseMeta) Fingerprint() []byte {
	if meta._fingerprint == nil {
		if MetaTypeHasSeed(meta.Self().(IMeta).Type()) {
			meta._fingerprint = MetaGetFingerprint(meta.GetMap(false))
		}
	}
	return meta._fingerprint
}

func (meta *BaseMeta) IsValid() bool {
	if meta._status == 0 {
		meta._status = MetaStatus(meta.Self().(IMeta))
	}
	return meta._status == 1
}

func (meta *BaseMeta) GenerateID(network uint8, terminal string) ID {
	return MetaGenerateID(meta.Self().(IMetaExt), network, terminal)
}

func (meta *BaseMeta) MatchID(identifier ID) bool {
	return MetaMatchID(meta.Self().(IMetaExt), identifier)
}

func (meta *BaseMeta) MatchKey(key VerifyKey) bool {
	return MetaMatchKey(meta.Self().(IMeta), key)
}

//
//  Functions for handling meta info
//

func MetaStatus(meta IMeta) int8 {
	key := meta.Key()
	if key == nil {
		// meta.key should not be empty
		return -1
	} else if MetaTypeHasSeed(meta.Type()) {
		seed := meta.Seed()
		fingerprint := meta.Fingerprint()
		if seed == "" || fingerprint == nil {
			// seed and fingerprint should not be empty
			return -1
		} else if key.Verify(UTF8Encode(seed), fingerprint) {
			// fingerprint matched
			return 1
		} else {
			// fingerprint not matched
			return -1
		}
	} else {
		return 1
	}
}

func MetaGenerateID(meta IMetaExt, network uint8, terminal string) ID {
	address := meta.GenerateAddress(network)
	if address == nil {
		return nil
	} else {
		return IDCreate(meta.Seed(), address, terminal)
	}
}

func MetaMatchID(meta IMetaExt, identifier ID) bool {
	if meta.IsValid() == false {
		return false
	}
	// check ID.name
	if identifier.Name() != meta.Seed() {
		return false
	}
	// check ID.address
	addr1 := identifier.Address()
	addr2 := meta.GenerateAddress(identifier.Type())
	return addr1.Equal(addr2)
}

func MetaMatchKey(meta IMeta, key VerifyKey) bool {
	if meta.IsValid() == false {
		return false
	}
	// check whether the public key equals to meta.key
	other, ok := key.(Object)
	if ok && other.Equal(meta.Key()) {
		return true
	}
	// check with seed & fingerprint
	if MetaTypeHasSeed(meta.Type()) {
		// check whether keys equal by verifying signature
		seed := meta.Seed()
		fingerprint := meta.Fingerprint()
		return key.Verify(UTF8Encode(seed), fingerprint)
	} else {
		// ID with BTC/ETH address has no username
		// so we can just compare the key.data to check matching
		return false
	}
}
