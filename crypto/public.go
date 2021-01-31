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

/**
 *  Asymmetric Cryptography Public Key
 *  ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 *
 *  key data format: {
 *      algorithm : "RSA", // "ECC", ...
 *      data      : "{BASE64_ENCODE}",
 *      ...
 *  }
 */
type PublicKey interface {
	AsymmetricKey
	VerifyKey
}

/**
 *  Public Key Factory
 *  ~~~~~~~~~~~~~~~~~~
 */
type PublicKeyFactory interface {

	/**
	 *  Parse map object to key
	 *
	 * @param key - key info
	 * @return PublicKey
	 */
	ParsePublicKey(key map[string]interface{}) PublicKey
}

var publicFactories = make(map[string]PublicKeyFactory)

func PublicKeyRegister(algorithm string, factory PublicKeyFactory) {
	publicFactories[algorithm] = factory
}

func PublicKeyGetFactory(algorithm string) PublicKeyFactory {
	return publicFactories[algorithm]
}

//
//  Factory method
//
func PublicKeyParse(key map[string]interface{}) PublicKey {
	if key == nil {
		return nil
	}
	algorithm := CryptographyKeyAlgorithm(key)
	factory := PublicKeyGetFactory(algorithm)
	if factory == nil {
		factory = PublicKeyGetFactory("*")  // unknown
	}
	return factory.ParsePublicKey(key)
}
