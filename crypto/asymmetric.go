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
	"unsafe"
)

const (
	RSA = "RSA"  //-- "RSA/ECB/PKCS1Padding", "SHA256withRSA"
	ECC = "ECC"
)

type AsymmetricKey interface {
	CryptographyKey
}

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

	Match(privateKey *PrivateKey) bool
}

func PublicKeysEqual(key, other *PublicKey) bool {
	ptr1 := (*CryptographyKey)(unsafe.Pointer(key))
	ptr2 := (*CryptographyKey)(unsafe.Pointer(other))
	return CryptographyKeysEqual(ptr1, ptr2)
}

/**
 *  Check whether they are key pair
 *
 * @param privateKey - private key that can generate the same public key
 * @return true on keys matched
 */
func AsymmetricKeyMatch(publicKey *PublicKey, privateKey *PrivateKey) bool {
	// 1. if the SK has the same public key, return true
	pKey := (*privateKey).GetPublicKey()
	if PublicKeysEqual(pKey, publicKey) {
		return true
	}
	// 2. check by signature
	signature := (*privateKey).Sign(promise)
	return (*publicKey).Verify(promise, signature)
}

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
	GetPublicKey() *PublicKey
}

func PrivateKeysEqual(key, other *PrivateKey) bool {
	ptr1 := (*CryptographyKey)(unsafe.Pointer(key))
	ptr2 := (*CryptographyKey)(unsafe.Pointer(other))
	if CryptographyKeysEqual(ptr1, ptr2) {
		return true
	}
	// check by public
	publicKey := (*key).GetPublicKey()
	return AsymmetricKeyMatch(publicKey, other)
}
