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

type MetaType = string

/**
 *  User/Group Meta data
 *  <p>
 *      This class is used to generate entity meta
 *  </p>
 *
 *  <blockquote><pre>
 *  data format: {
 *      "type"        : i2s(1),         // algorithm version
 *      "key"         : "{public key}", // PK = secp256k1(SK);
 *      "seed"        : "moKy",         // user/group name
 *      "fingerprint" : "..."           // CT = sign(seed, SK);
 *  }
 *
 *  algorithm:
 *      fingerprint = sign(seed, SK);
 *  </pre></blockquote>
 */
type Meta interface {
	Mapper

	/**
	 *  Meta algorithm version
	 *
	 *  <pre>
	 *  1 = MKM : username@address (default)
	 *  2 = BTC : btc_address
	 *  4 = ETH : eth_address
	 *  ...
	 *  </pre>
	 */
	Type() MetaType

	/**
	 *  Public key (used for signature)
	 *
	 *      RSA / ECC
	 */
	PublicKey() VerifyKey

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
	Fingerprint() TransportableData

	//
	//  Validation
	//

	/**
	 *  Check meta valid
	 *  <p>(must call this when received a new meta from network)</p>
	 *
	 * @return true on valid
	 */
	IsValid() bool

	/**
	 *  Generate address with network(type)
	 *
	 * @param network - ID.type
	 * @return Address
	 */
	GenerateAddress(network EntityType) Address
}

/**
 *  Meta Factory
 *  ~~~~~~~~~~~~
 */
type MetaFactory interface {

	/**
	 *  Create meta
	 *
	 * @param pKey        - public key
	 * @param seed        - ID.name
	 * @param fingerprint - sKey.sign(seed)
	 * @return Meta
	 */
	CreateMeta(pKey VerifyKey, seed string, fingerprint TransportableData) Meta

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
	ParseMeta(meta StringKeyMap) Meta
}

//
//  Factory methods
//

func CreateMeta(version MetaType, pKey VerifyKey, seed string, fingerprint TransportableData) Meta {
	helper := GetMetaHelper()
	return helper.CreateMeta(version, pKey, seed, fingerprint)
}

func GenerateMeta(version MetaType, sKey SignKey, seed string) Meta {
	helper := GetMetaHelper()
	return helper.GenerateMeta(version, sKey, seed)
}

func ParseMeta(meta interface{}) Meta {
	helper := GetMetaHelper()
	return helper.ParseMeta(meta)
}

func GetMetaFactory(version MetaType) MetaFactory {
	helper := GetMetaHelper()
	return helper.GetMetaFactory(version)
}

func SetMetaFactory(version MetaType, factory MetaFactory) {
	helper := GetMetaHelper()
	helper.SetMetaFactory(version, factory)
}
