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
 *  Meta Shadow
 *  ~~~~~~~~~~~
 *
 * @abstract:
 *      GenerateAddress(network uint8) Address
 */
type BaseMetaShadow struct {
	IMetaExt

	_meta IMetaExt
}

func (shadow *BaseMetaShadow) Init(meta IMetaExt) *BaseMetaShadow {
	shadow._meta = meta
	return shadow
}

func (shadow *BaseMetaShadow) Meta() IMetaExt {
	return shadow._meta
}

func (shadow *BaseMetaShadow) getMap() map[string]interface{} {
	return shadow.Meta().(Map).GetMap(false)
}

//-------- IMeta

func (shadow *BaseMetaShadow) Type() uint8 {
	return MetaGetType(shadow.getMap())
}

func (shadow *BaseMetaShadow) Key() VerifyKey {
	return MetaGetKey(shadow.getMap())
}

func (shadow *BaseMetaShadow) Seed() string {
	if MetaTypeHasSeed(shadow.Meta().Type()) {
		return MetaGetSeed(shadow.getMap())
	} else {
		return ""
	}
}

func (shadow *BaseMetaShadow) Fingerprint() []byte {
	if MetaTypeHasSeed(shadow.Meta().Type()) {
		return MetaGetFingerprint(shadow.getMap())
	} else {
		return nil
	}
}

func (shadow *BaseMetaShadow) IsValid() bool {
	return shadow.Meta().status() == 1
}

func (shadow *BaseMetaShadow) GenerateID(network uint8, terminal string) ID {
	return MetaGenerateID(shadow.Meta(), network, terminal)
}

func (shadow *BaseMetaShadow) MatchID(identifier ID) bool {
	return MetaMatchID(shadow.Meta(), identifier)
}

func (shadow *BaseMetaShadow) MatchKey(key VerifyKey) bool {
	return MetaMatchKey(shadow.Meta(), key)
}

//-------- IMetaExt

func (shadow *BaseMetaShadow) status() int8 {
	return MetaStatus(shadow.Meta())
}

//
//  Functions for handling meta info
//

func MetaStatus(meta IMetaExt) int8 {
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
	address := meta.GenerateAddress(identifier.Type())
	return identifier.Address().Equal(address)
}

func MetaMatchKey(meta IMetaExt, key VerifyKey) bool {
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
