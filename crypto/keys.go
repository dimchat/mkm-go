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
	. "github.com/dimchat/mkm-go/format"
	. "github.com/dimchat/mkm-go/types"
)

/**
 *  Cryptography Key
 *  ~~~~~~~~~~~~~~~~
 *  Cryptography key with designated algorithm
 *
 *  key data format: {
 *      algorithm : "RSA", // ECC, AES, ...
 *      data      : "{BASE64_ENCODE}",
 *      ...
 *  }
 */
type CryptographyKey interface {
	Map

	/**
	 *  Get key algorithm name
	 *
	 * @return algorithm name
	 */
	Algorithm() string

	/**
	 *  Get key data
	 *
	 * @return key data
	 */
	Data() []byte
}

type EncryptKey interface {
	CryptographyKey

	/**
	 *  ciphertext = encrypt(plaintext, PW)
	 *  ciphertext = encrypt(plaintext, PK)
	 *
	 * @param plaintext - plain data
	 * @return ciphertext
	 */
	Encrypt(plaintext []byte) []byte
}

type DecryptKey interface {
	CryptographyKey

	/**
	 *  plaintext = decrypt(ciphertext, PW);
	 *  plaintext = decrypt(ciphertext, SK);
	 *
	 * @param ciphertext - encrypted data
	 * @return plaintext
	 */
	Decrypt(ciphertext []byte) []byte

	/**
	 *  OK = decrypt(encrypt(data, SK), PK) == data
	 *
	 * @param pKey - public key
	 * @return true on signature matched
	 */
	Match(pKey EncryptKey) bool
}

type SignKey interface {
	CryptographyKey

	/**
	 *  signature = sign(data, SK);
	 *
	 * @param data - data to be signed
	 * @return signature
	 */
	Sign(data []byte) []byte
}

type VerifyKey interface {
	CryptographyKey

	/**
	 *  OK = verify(data, signature, PK)
	 *
	 * @param data - data
	 * @param signature - signature of data
	 * @return true on signature matched
	 */
	Verify(data []byte, signature []byte) bool

	/**
	 *  OK = verify(data, sign(data, SK), PK)
	 *
	 * @param sKey - private key
	 * @return true on signature matched
	 */
	Match(sKey SignKey) bool
}

/**
 *  Asymmetric Cryptography Key
 *  ~~~~~~~~~~~~~~~~~~~~~~~~~~~
 */
type AsymmetricKey interface {
	CryptographyKey
}

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

const _promise = "Moky loves May Lee forever!"
var promise = UTF8Encode(_promise)

func CryptographyKeyGetAlgorithm(key map[string]interface{}) string {
	text, ok := key["algorithm"].(string)
	if ok {
		return text
	} else {
		return ""
	}
}
