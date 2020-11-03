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

const _promise = "Moky loves May Lee forever!"
var promise = UTF8BytesFromString(_promise)

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
}

func CryptographyKeysEqual(key, other *CryptographyKey) bool {
	if *key == *other {
		return true
	}
	// check inner maps
	map1 := (*key).GetMap(false)
	map2 := (*other).GetMap(false)
	return MapsEqual(map1, map2)
}

type EncryptKey interface {

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

	/**
	 *  plaintext = decrypt(ciphertext, PW);
	 *  plaintext = decrypt(ciphertext, SK);
	 *
	 * @param ciphertext - encrypted data
	 * @return plaintext
	 */
	Decrypt(ciphertext []byte) []byte
}

type SignKey interface {

	/**
	 *  signature = sign(data, SK);
	 *
	 * @param data - data to be signed
	 * @return signature
	 */
	Sign(data []byte) []byte
}

type VerifyKey interface {

	/**
	 *  OK = verify(data, signature, PK)
	 *
	 * @param data - data
	 * @param signature - signature of data
	 * @return true on signature matched
	 */
	Verify(data []byte, signature []byte) bool
}
