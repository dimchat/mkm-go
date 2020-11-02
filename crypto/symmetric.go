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
	"unsafe"
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
	EncryptKey
	DecryptKey
}

func SymmetricKeysEqual(key, other *SymmetricKey) bool {
	if key == nil {
		return other == nil
	} else if other == nil {
		return false
	}
	ptr1 := (*CryptographyKey)(unsafe.Pointer(key))
	ptr2 := (*CryptographyKey)(unsafe.Pointer(other))
	if CryptographyKeysEqual(ptr1, ptr2) {
		return true
	}
	// check by encryption
	ciphertext := (*other).Encrypt(promise)
	plaintext := (*key).Decrypt(ciphertext)
	return bytes.Equal(promise, plaintext)
}
