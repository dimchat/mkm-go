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
 *  Base Symmetric Key
 *  ~~~~~~~~~~~~~~~~~~
 *
 *  Abstract methods:
 *      Data() []byte
 *      Encrypt(plaintext []byte) []byte
 *      Decrypt(ciphertext []byte) []byte
 */
type BaseSymmetricKey struct {
	Dictionary
	SymmetricKey
}

func (key *BaseSymmetricKey) Init(dict map[string]interface{}) *BaseSymmetricKey {
	if key.Dictionary.Init(dict) != nil {
	}
	return key
}

func (key *BaseSymmetricKey) Equal(other interface{}) bool {
	return key.Dictionary.Equal(other)
}

func (key *BaseSymmetricKey) Get(name string) interface{} {
	return key.Dictionary.Get(name)
}

func (key *BaseSymmetricKey) Set(name string, value interface{}) {
	key.Dictionary.Set(name, value)
}

func (key *BaseSymmetricKey) Keys() []string {
	return key.Dictionary.Keys()
}

func (key *BaseSymmetricKey) GetMap(clone bool) map[string]interface{} {
	return key.Dictionary.GetMap(clone)
}

func (key *BaseSymmetricKey) Algorithm() string {
	return CryptographyKeyGetAlgorithm(key.GetMap(false))
}

func (key *BaseSymmetricKey) Match(pKey EncryptKey) bool {
	return CryptographyKeysMatch(pKey, key)
}

/**
 *  Base Public Key
 *  ~~~~~~~~~~~~~~~
 *
 *  Abstract methods:
 *      Data() []byte
 *      Verify(data []byte, signature []byte) bool
 */
type BasePublicKey struct {
	Dictionary
	PublicKey
}

func (key *BasePublicKey) Init(dict map[string]interface{}) *BasePublicKey {
	if key.Dictionary.Init(dict) != nil {
	}
	return key
}

func (key *BasePublicKey) Equal(other interface{}) bool {
	return key.Dictionary.Equal(other)
}

func (key *BasePublicKey) Get(name string) interface{} {
	return key.Dictionary.Get(name)
}

func (key *BasePublicKey) Set(name string, value interface{}) {
	key.Dictionary.Set(name, value)
}

func (key *BasePublicKey) Keys() []string {
	return key.Dictionary.Keys()
}

func (key *BasePublicKey) GetMap(clone bool) map[string]interface{} {
	return key.Dictionary.GetMap(clone)
}

func (key *BasePublicKey) Algorithm() string {
	return CryptographyKeyGetAlgorithm(key.GetMap(false))
}

func (key *BasePublicKey) Match(sKey SignKey) bool {
	return AsymmetricKeysMatch(sKey, key)
}

/**
 *  Base Private Key
 *  ~~~~~~~~~~~~~~~~
 *
 *  Abstract methods:
 *      Data() []byte
 *      Sign(data []byte) []byte
 *      PublicKey() PublicKey
 */
type BasePrivateKey struct {
	Dictionary
	PrivateKey
}

func (key *BasePrivateKey) Init(dict map[string]interface{}) *BasePrivateKey {
	if key.Dictionary.Init(dict) != nil {
	}
	return key
}

func (key *BasePrivateKey) Equal(other interface{}) bool {
	return key.Dictionary.Equal(other)
}

func (key *BasePrivateKey) Get(name string) interface{} {
	return key.Dictionary.Get(name)
}

func (key *BasePrivateKey) Set(name string, value interface{}) {
	key.Dictionary.Set(name, value)
}

func (key *BasePrivateKey) Keys() []string {
	return key.Dictionary.Keys()
}

func (key *BasePrivateKey) GetMap(clone bool) map[string]interface{} {
	return key.Dictionary.GetMap(clone)
}

func (key *BasePrivateKey) Algorithm() string {
	return CryptographyKeyGetAlgorithm(key.GetMap(false))
}
