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

type IMetaDelegate interface {

	Status() int8

	GenerateID(network uint8, terminal string) ID

	MatchID(identifier ID) bool

	MatchKey(key VerifyKey) bool
}

/**
 *  Meta Shadow
 *  ~~~~~~~~~~~
 *
 *  Delegate for handling meta info
 */
type BaseMetaShadow struct {
	IMetaDelegate

	_meta IMeta
}

func (shadow *BaseMetaShadow) Init(meta IMeta) *BaseMetaShadow {
	shadow._meta = meta
	return shadow
}

func (shadow *BaseMetaShadow) Status() int8 {
	return MetaStatus(shadow._meta)
}

func (shadow *BaseMetaShadow) GenerateID(network uint8, terminal string) ID {
	return MetaGenerateID(shadow._meta, network, terminal)
}

func (shadow *BaseMetaShadow) MatchID(identifier ID) bool {
	return MetaMatchID(shadow._meta, identifier)
}

func (shadow *BaseMetaShadow) MatchKey(key VerifyKey) bool {
	return MetaMatchKey(shadow._meta, key)
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

func MetaGenerateID(meta IMeta, network uint8, terminal string) ID {
	address := meta.GenerateAddress(network)
	if address == nil {
		return nil
	} else {
		return IDCreate(meta.Seed(), address, terminal)
	}
}

func MetaMatchID(meta IMeta, identifier ID) bool {
	if meta.IsValid() == false {
		return false
	}
	// check ID.name
	if identifier.Name() != meta.Seed() {
		return false
	}
	// check ID.address
	address := meta.GenerateAddress(identifier.Type())
	return identifier.Address().Equal(address)
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
