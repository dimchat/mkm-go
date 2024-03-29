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

/**
 *  Base Meta
 *  ~~~~~~~~~
 *
 * @abstract:
 *      GenerateAddress(network NetworkType) Address
 */
type BaseMeta struct {
	Dictionary

	_type MetaType
	_key VerifyKey

	_seed string
	_fingerprint []byte
}

func (meta *BaseMeta) Init(dict map[string]interface{}) Meta {
	if meta.Dictionary.Init(dict) != nil {
		// lazy load
		meta._type = 0
		meta._key = nil
		meta._seed = ""
		meta._fingerprint = nil
	}
	return meta
}

func (meta *BaseMeta) InitWithType(version MetaType, key VerifyKey, seed string, fingerprint []byte) Meta {
	dict := make(map[string]interface{})
	// meta type
	dict["type"] = version
	// meta key
	dict["key"] = key.Map()
	// seed
	if seed != "" {
		dict["seed"] = seed
	}
	// fingerprint
	if !ValueIsNil(fingerprint) {
		dict["fingerprint"] = Base64Encode(fingerprint)
	}
	if meta.Dictionary.Init(dict) != nil {
		// set values
		meta._type = version
		meta._key = key
		meta._seed = seed
		meta._fingerprint = fingerprint
	}
	return meta
}

//-------- IMeta

func (meta *BaseMeta) Type() MetaType {
	if meta._type == 0 {
		meta._type = MetaGetType(meta.Map())
	}
	return meta._type
}

func (meta *BaseMeta) Key() VerifyKey {
	if meta._key == nil {
		key := meta.Get("key")
		meta._key = PublicKeyParse(key)
	}
	return meta._key
}

func (meta *BaseMeta) Seed() string {
	if meta._seed == "" {
		if MetaTypeHasSeed(meta.Type()) {
			meta._seed, _ = meta.Get("seed").(string)
		}
	}
	return meta._seed
}

func (meta *BaseMeta) Fingerprint() []byte {
	if meta._fingerprint == nil {
		if MetaTypeHasSeed(meta.Type()) {
			base64, _ := meta.Get("fingerprint").(string)
			meta._fingerprint = Base64Decode(base64)
		}
	}
	return meta._fingerprint
}

func (meta *BaseMeta) GenerateAddress(_ NetworkType) Address {
	panic("BaseMeta::GenerateAddress(network) > override me!")
}
