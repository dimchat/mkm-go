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
 *  Base Crypto Key
 *  ~~~~~~~~~~~~~~~
 *
 * @abstract:
 *      Data() []byte
 */
type BaseKey struct {
	Dictionary
	ICryptographyKey
}

func (key *BaseKey) Init(this CryptographyKey, dict map[string]interface{}) *BaseKey {
	if key.Dictionary.Init(this, dict) != nil {
	}
	return key
}

func (key *BaseKey) Algorithm() string {
	return CryptographyKeyGetAlgorithm(key.GetMap(false))
}

/**
 *  Base Symmetric Key
 *  ~~~~~~~~~~~~~~~~~~
 *
 * @abstract:
 *      Data() []byte
 *      Encrypt(plaintext []byte) []byte
 *      Decrypt(ciphertext []byte) []byte
 */
type BaseSymmetricKey struct {
	BaseKey
	ISymmetricKey
}

func (key *BaseSymmetricKey) Init(this SymmetricKey, dict map[string]interface{}) *BaseSymmetricKey {
	if key.BaseKey.Init(this, dict) != nil {
	}
	return key
}

func (key *BaseSymmetricKey) Match(pKey EncryptKey) bool {
	return CryptographyKeysMatch(pKey, key)
}

/**
 *  Base Public Key
 *  ~~~~~~~~~~~~~~~
 *
 * @abstract:
 *      Data() []byte
 *      Verify(data []byte, signature []byte) bool
 */
type BasePublicKey struct {
	BaseKey
	IPublicKey
}

func (key *BasePublicKey) Init(this PublicKey, dict map[string]interface{}) *BasePublicKey {
	if key.BaseKey.Init(this, dict) != nil {
	}
	return key
}

func (key *BasePublicKey) Match(sKey SignKey) bool {
	return AsymmetricKeysMatch(sKey, key)
}

/**
 *  Base Private Key
 *  ~~~~~~~~~~~~~~~~
 *
 * @abstract:
 *      Data() []byte
 *      Sign(data []byte) []byte
 *      PublicKey() PublicKey
 */
type BasePrivateKey struct {
	BaseKey
	IPrivateKey
}

func (key *BasePrivateKey) Init(this PrivateKey, dict map[string]interface{}) *BasePrivateKey {
	if key.BaseKey.Init(this, dict) != nil {
	}
	return key
}
