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

import (
	. "github.com/dimchat/mkm-go/ext"
	. "github.com/dimchat/mkm-go/types"
)

/**
 *  Asymmetric Cryptography Private Key
 *  <p>
 *      This class is used to decrypt symmetric key or sign message data
 *  </p>
 *
 *  <blockquote><pre>
 *  key data format: {
 *      "algorithm" : "RSA", // "ECC", ...
 *      "data"      : "{BASE64_ENCODE}",
 *      ...
 *  }
 *  </pre></blockquote>
 */
type PrivateKey interface {
	AsymmetricKey
	ISignKey

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
	ParsePrivateKey(key StringKeyMap) PrivateKey
}

//
//  Factory methods
//

func GeneratePrivateKey(algorithm string) PrivateKey {
	helper := GetPrivateKeyHelper()
	return helper.GeneratePrivateKey(algorithm)
}

func ParsePrivateKey(key interface{}) PrivateKey {
	helper := GetPrivateKeyHelper()
	return helper.ParsePrivateKey(key)
}

func GetPrivateKeyFactory(algorithm string) PrivateKeyFactory {
	helper := GetPrivateKeyHelper()
	return helper.GetPrivateKeyFactory(algorithm)
}

func SetPrivateKeyFactory(algorithm string, factory PrivateKeyFactory) {
	helper := GetPrivateKeyHelper()
	helper.SetPrivateKeyFactory(algorithm, factory)
}
