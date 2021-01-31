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
	. "github.com/dimchat/mkm-go/crypto"
	. "github.com/dimchat/mkm-go/format"
	. "github.com/dimchat/mkm-go/types"
)

/**
 *  User/Group Meta data
 *  ~~~~~~~~~~~~~~~~~~~~
 *  This class is used to generate entity ID
 *
 *      data format: {
 *          version: 1,         // algorithm version
 *          key: {public key},  // PK = secp256k1(SK);
 *          seed: "moKy",       // user/group name
 *          fingerprint: "..."  // CT = sign(seed, SK);
 *      }
 *
 *      algorithm:
 *          fingerprint = sign(seed, SK);
 */
type Meta interface {
	Map

	/**
	 *  Meta algorithm version
	 *
	 *      0x01 - username@address
	 *      0x02 - btc_address
	 *      0x03 - username@btc_address
	 */
	Type() MetaType

	/**
	 *  Public key (used for signature)
	 *
	 *      RSA / ECC
	 */
	Key() PublicKey

	/**
	 *  Seed to generate fingerprint
	 *
	 *      Username / Group-X
	 */
	Seed() string

	/**
	 *  Fingerprint to verify ID and public key
	 *
	 *      Build: fingerprint = sign(seed, privateKey)
	 *      Check: verify(seed, fingerprint, publicKey)
	 */
	Fingerprint() []byte

	/**
	 *  Check meta valid
	 *  (must call this when received a new meta from network)
	 *
	 * @return true on valid
	 */
	IsValid() bool

	// comparing
	MatchID(identifier ID) bool
	MatchAddress(address Address) bool
	MatchKey(key PublicKey) bool

	// call 'GenerateAddress'
	GenerateID(network NetworkType) ID

	/**
	 *  Generate address with meta info and address type
	 *
	 * @param network - address network type
	 * @return Address object
	 */
	GenerateAddress(network NetworkType) Address
}

func MetasEqual(meta1, meta2 Meta) bool {
	if MapsEqual(meta1, meta2) {
		return true
	}
	// check by generating ID
	id := meta1.GenerateID(MAIN)
	return MetaMatchID(meta2, id)
}

func MetaGenerateID(meta Meta, network NetworkType) ID {
	name := meta.Seed()
	address := meta.GenerateAddress(network)
	return NewID(name, address, "")
}

func MetaMatchID(meta Meta, identifier ID) bool {
	// check ID.name
	name := identifier.Name()
	if MetaTypeHasSeed(meta.Type()) {
		// this meta has seed, it must be equal to ID.name
		if meta.Seed() == name {
			return false
		}
	} else if name != "" {
		// this meta hasn't seed, so ID.name should be empty
		return false
	}
	// check ID.address
	return MetaMatchAddress(meta, identifier.Address())
}

func MetaMatchAddress(meta Meta, address Address) bool {
	network := address.Type()
	addr := meta.GenerateAddress(network)
	return addr.Equal(address)
}

func MetaMatchKey(meta Meta, key PublicKey) bool {
	// check whether the public key equals to meta.key
	if PublicKeysEqual(meta.Key(), key) {
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
