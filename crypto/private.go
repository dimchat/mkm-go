/* license: https://mit-license.org
 * ==============================================================================
 * The MIT License (MIT)
 *
 * Copyright (c) 2021 Albert Moky
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

import . "github.com/dimchat/mkm-go/types"

/**
 *  Asymmetric Cryptography Private Key
 *  ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 *  This class is used to decrypt symmetric key or sign message data
 *
 *  key data format: {
 *      algorithm : "RSA", // "ECC", ...
 *      data      : "{BASE64_ENCODE}",
 *      ...
 *  }
 */
type PrivateKey interface {
	AsymmetricKey
	SignKey

	/**
	 *  Get public key from private key
	 *
	 * @return public key paired to this private key
	 */
	PublicKey() PublicKey
}

/**
 *  Private Key Factory
 *  ~~~~~~~~~~~~~~~~~~~
 */
type PrivateKeyFactory interface {

	/**
	 *  Generate key
	 *
	 * @return PrivateKey
	 */
	GeneratePrivateKey() PrivateKey

	/**
	 *  Parse map object to key
	 *
	 * @param key - key info
	 * @return PrivateKey
	 */
	ParsePrivateKey(key map[string]interface{}) PrivateKey
}

var privateFactories = make(map[string]PrivateKeyFactory)

func PrivateKeyRegister(algorithm string, factory PrivateKeyFactory) {
	privateFactories[algorithm] = factory
}

func PrivateKeyGetFactory(algorithm string) PrivateKeyFactory {
	return privateFactories[algorithm]
}

//
//  Factory methods
//
func PrivateKeyGenerate(algorithm string) PrivateKey {
	factory := PrivateKeyGetFactory(algorithm)
	if factory == nil {
		panic("key algorithm not support: " + algorithm)
	}
	return factory.GeneratePrivateKey()
}

func PrivateKeyParse(key interface{}) PrivateKey {
	if key == nil {
		return nil
	}
	value, ok := key.(PrivateKey)
	if ok {
		return value
	}
	// get key info
	var info map[string]interface{}
	wrapper, ok := key.(Map)
	if ok {
		info = wrapper.GetMap(false)
	} else {
		info, ok = key.(map[string]interface{})
		if !ok {
			panic(key)
			return nil
		}
	}
	// get key factory by algorithm
	algorithm := CryptographyKeyGetAlgorithm(info)
	factory := PrivateKeyGetFactory(algorithm)
	if factory == nil {
		factory = PrivateKeyGetFactory("*")  // unknown
	}
	return factory.ParsePrivateKey(info)
}
