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

//  virtual class, need to implement methods:
//      Key() PublicKey
//      GenerateAddress()
//  before using it
type BaseMeta struct {
	Dictionary
	Meta

	_version MetaType

	_key PublicKey

	_seed string
	_fingerprint []byte

	_status int8 // 1 for valid, -1 for invalid
}

func (meta *BaseMeta) Init(dictionary map[string]interface{}) *BaseMeta {
	if meta.Dictionary.Init(dictionary) != nil {
		// lazy load
		meta._version = MetaType(0)
		meta._key = nil
		meta._seed = ""
		meta._fingerprint = nil
		meta._status = 0
	}
	return meta
}

func (meta *BaseMeta) Equal(other interface{}) bool {
	other = ObjectValue(other)
	switch other.(type) {
	case Meta:
		return MetasEqual(meta, other.(Meta))
	default:
		return false
	}
}

/**
 *  Check meta valid
 *  (must call this when received a new meta from network)
 *
 * @return true on valid
 */
func (meta *BaseMeta) IsValid() bool {
	if meta._status == 0 {
		key := meta.Key()
		if key == nil {
			// meta.key should not be empty
			meta._status = -1
		} else if MetaTypeHasSeed(meta.Type()) {
			seed := meta.Seed()
			fingerprint := meta.Fingerprint()
			if seed == "" || fingerprint == nil {
				// seed and fingerprint should not be empty
				meta._status = -1
			} else if key.Verify(UTF8BytesFromString(seed), fingerprint) {
				// fingerprint matched
				meta._status = 1
			} else {
				// fingerprint not matched
				meta._status = -1
			}
		} else {
			meta._status = 1
		}
	}
	return meta._status == 1
}

func (meta *BaseMeta) Type() MetaType {
	if meta._version == 0 {
		version := meta.Get("version")
		meta._version = MetaType(version.(uint8))
	}
	return meta._version
}

func (meta *BaseMeta) Seed() string {
	if meta._seed == "" {
		if MetaTypeHasSeed(meta.Type()) {
			seed := meta.Get("seed")
			meta._seed = seed.(string)
		}
	}
	return meta._seed
}

func (meta *BaseMeta) Fingerprint() []byte {
	if meta._fingerprint == nil {
		if MetaTypeHasSeed(meta.Type()) {
			base64 := meta.Get("fingerprint")
			meta._fingerprint = Base64Decode(base64.(string))
		}
	}
	return meta._fingerprint
}

func (meta *BaseMeta) MatchID(identifier ID) bool {
	return MetaMatchID(meta, identifier)
}

func (meta *BaseMeta) MatchAddress(address Address) bool {
	return MetaMatchAddress(meta, address)
}

func (meta *BaseMeta) MatchKey(key PublicKey) bool {
	return MetaMatchKey(meta, key)
}

func (meta *BaseMeta) GenerateID(network NetworkType) ID {
	var name string
	if MetaTypeHasSeed(meta.Type()) {
		name = meta.Seed()
	} else {
		name = ""
	}
	address := meta.GenerateAddress(network)
	return NewID(name, address, "")
}
