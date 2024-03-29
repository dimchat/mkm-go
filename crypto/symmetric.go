/* license: https://mit-license.org
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
package crypto

import (
	"bytes"
	. "github.com/dimchat/mkm-go/types"
)

const (
	AES = "AES"  //-- "AES/CBC/PKCS7Padding"
	DES = "DES"
)

/**
 *  Symmetric Cryptography Key
 *  ~~~~~~~~~~~~~~~~~~~~~~~~~~
 *  This class is used to encrypt or decrypt message data
 *
 *  key data format: {
 *      algorithm : "AES", // "DES", ...
 *      data      : "{BASE64_ENCODE}",
 *      ...
 *  }
 */
type SymmetricKey interface {
	CryptographyKey
	IEncryptKey
	IDecryptKey
}

/**
 *  Check symmetric keys
 *
 * @param pKey - symmetric key1
 * @param sKey - symmetric key2
 * @return true on keys equal
 */
func SymmetricKeysMatch(pKey EncryptKey, sKey DecryptKey) bool {
	// check by encryption
	ciphertext := pKey.Encrypt(promise)
	plaintext := sKey.Decrypt(ciphertext)
	return bytes.Equal(plaintext, promise)
}

/**
 *  Symmetric Key Factory
 *  ~~~~~~~~~~~~~~~~~~~~~
 */
type SymmetricKeyFactory interface {

	/**
	 *  Generate key
	 *
	 * @return SymmetricKey
	 */
	GenerateSymmetricKey() SymmetricKey

	/**
	 *  Parse map object to key
	 *
	 * @param key - key info
	 * @return SymmetricKey
	 */
	ParseSymmetricKey(key map[string]interface{}) SymmetricKey
}

//
//  Instances of SymmetricKeyFactory
//
var symmetricFactories = make(map[string]SymmetricKeyFactory)

func SymmetricKeySetFactory(algorithm string, factory SymmetricKeyFactory) {
	symmetricFactories[algorithm] = factory
}

func SymmetricKeyGetFactory(algorithm string) SymmetricKeyFactory {
	return symmetricFactories[algorithm]
}

//
//  Factory methods
//
func SymmetricKeyGenerate(algorithm string) SymmetricKey {
	factory := SymmetricKeyGetFactory(algorithm)
	if factory == nil {
		panic("key algorithm not support: " + algorithm)
	}
	return factory.GenerateSymmetricKey()
}

func SymmetricKeyParse(key interface{}) SymmetricKey {
	if ValueIsNil(key) {
		return nil
	}
	value, ok := key.(SymmetricKey)
	if ok {
		return value
	}
	info := FetchMap(key)
	// get key factory by algorithm
	algorithm := CryptographyKeyGetAlgorithm(info)
	factory := SymmetricKeyGetFactory(algorithm)
	if factory == nil {
		factory = SymmetricKeyGetFactory("*")  // unknown
	}
	return factory.ParseSymmetricKey(info)
}
