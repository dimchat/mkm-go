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
	Type() uint8

	/**
	 *  Public key (used for signature)
	 *
	 *      RSA / ECC
	 */
	Key() VerifyKey

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

	/**
	 *  Generate ID with terminal
	 *
	 * @param network - ID.type
	 * @param terminal - ID.terminal
	 * @return ID
	 */
	GenerateID(network uint8, terminal string) ID

	/**
	 *  Check whether meta match with entity ID
	 *  (must call this when received a new meta from network)
	 *
	 * @param identifier - entity ID
	 * @return true on matched
	 */
	MatchID(identifier ID) bool

	/**
	 *  Check whether meta match with public key
	 *
	 * @param pk - public key
	 * @return true on matched
	 */
	MatchKey(key VerifyKey) bool
}

func MetaGetType(meta map[string]interface{}) uint8 {
	version := meta["type"]
	if version == nil {
		// compatible with v1.0
		version = meta["version"]
	}
	return version.(uint8)
}

func MetaGetKey(meta map[string]interface{}) VerifyKey {
	key := meta["key"]
	if key == nil {
		panic("meta key not found: " + UTF8Decode(JSONEncode(meta)))
	}
	return PublicKeyParse(key)
}

func MetaGetSeed(meta map[string]interface{}) string {
	seed := meta["seed"]
	if seed == nil {
		return ""
	}
	return seed.(string)
}

func MetaGetFingerprint(meta map[string]interface{}) []byte {
	base64 := meta["fingerprint"]
	if base64 == nil {
		return nil
	}
	return Base64Decode(base64.(string))
}

/**
 *  Meta Factory
 *  ~~~~~~~~~~~~
 */
type MetaFactory interface {

	/**
	 *  Create meta
	 *
	 * @param key         - public key
	 * @param seed        - ID.name
	 * @param fingerprint - sKey.sign(seed)
	 * @return Meta
	 */
	CreateMeta(key VerifyKey, seed string, fingerprint []byte) Meta

	/**
	 *  Generate meta
	 *
	 * @param sKey    - private key
	 * @param seed    - ID.name
	 * @return Meta
	 */
	GenerateMeta(sKey SignKey, seed string) Meta

	/**
	 *  Parse map object to meta
	 *
	 * @param meta - meta info
	 * @return Meta
	 */
	ParseMeta(meta map[string]interface{}) Meta
}

var metaFactory = make(map[uint8]MetaFactory)

func MetaRegister(version uint8, factory MetaFactory) {
	metaFactory[version] = factory
}

func MetaGetFactory(version uint8) MetaFactory {
	return metaFactory[version]
}

//
//  Factory methods
//
func MetaCreate(version uint8, key VerifyKey, seed string, fingerprint []byte) Meta {
	factory := MetaGetFactory(version)
	if factory == nil {
		panic("meta type not found: " + string(version))
	}
	return factory.CreateMeta(key, seed, fingerprint)
}

func MetaGenerate(version uint8, sKey SignKey, seed string) Meta {
	factory := MetaGetFactory(version)
	if factory == nil {
		panic("meta type not found: " + string(version))
	}
	return factory.GenerateMeta(sKey, seed)
}

func MetaParse(meta interface{}) Meta {
	if meta == nil {
		return nil
	}
	var info map[string]interface{}
	value := ObjectValue(meta)
	switch value.(type) {
	case Meta:
		return value.(Meta)
	case Map:
		info = value.(Map).GetMap(false)
	case map[string]interface{}:
		info = value.(map[string]interface{})
	default:
		panic(meta)
	}
	version := MetaGetType(info)
	factory := MetaGetFactory(version)
	if factory == nil {
		factory = MetaGetFactory(0)  // unknown
	}
	return factory.ParseMeta(info)
}
