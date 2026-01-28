/* license: https://mit-license.org
 * ==============================================================================
 * The MIT License (MIT)
 *
 * Copyright (c) 2026 Albert Moky
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
package ext

import (
	"bytes"

	. "github.com/dimchat/mkm-go/crypto"
	. "github.com/dimchat/mkm-go/types"
)

// Sample data for checking keys
const _promise = "Moky loves May Lee forever!"

var promise = []byte(_promise)

/**
 *  Check asymmetric keys
 *
 * @param sKey - private key
 * @param pKey - public key
 * @return true on keys matched
 */
func AsymmetricKeysMatch(sKey SignKey, pKey VerifyKey) bool {
	// try to verify with signature
	signature := sKey.Sign(promise)
	return pKey.Verify(promise, signature)
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
	params := NewMap()
	ciphertext := pKey.Encrypt(promise, params)
	plaintext := sKey.Decrypt(ciphertext, params)
	return bytes.Equal(plaintext, promise)
}

/**
 *  CryptographyKey GeneralFactory
 */
type GeneralCryptoHelper interface {
	//SymmetricKeyHelper
	//PrivateKeyHelper
	//PublicKeyHelper

	//
	//  Algorithm
	//
	GetKeyAlgorithm(key StringKeyMap, defaultValue string) string
}

var sharedGeneralCryptoHelper GeneralCryptoHelper = nil

func SetGeneralCryptoHelper(helper GeneralCryptoHelper) {
	sharedGeneralCryptoHelper = helper
}

func GetGeneralCryptoHelper() GeneralCryptoHelper {
	return sharedGeneralCryptoHelper
}
