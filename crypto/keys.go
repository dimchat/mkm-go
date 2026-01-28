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
 *  <p>
 *      Cryptography key with designated algorithm
 *  </p>
 *
 *  <blockquote><pre>
 *  key data format: {
 *      "algorithm" : "RSA", // ECC, AES, ...
 *      "data"      : "{BASE64_ENCODE}",
 *      ...
 *  }
 *  </pre></blockquote>
 */
type CryptographyKey interface {
	Mapper

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
	Data() TransportableData
}

type EncryptKey interface {
	IEncryptKey
	CryptographyKey
}
type IEncryptKey interface {

	/**
	 *  1. Symmetric Key:
	 *     <blockquote><pre>
	 *         ciphertext = encrypt(plaintext, PW)
	 *     </pre></blockquote>
	 *
	 *  2. Asymmetric Public Key:
	 *     <blockquote><pre>
	 *         ciphertext = encrypt(plaintext, PK);
	 *     </pre></blockquote>
	 *
	 * @param plaintext
	 *        plain data
	 *
	 * @param extra
	 *        store extra variables ('IV' for 'AES')
	 *
	 * @return ciphertext
	 */
	Encrypt(plaintext []byte, extra StringKeyMap) []byte
}

type DecryptKey interface {
	IDecryptKey
	CryptographyKey
}
type IDecryptKey interface {

	/**
	 *  1. Symmetric Key:
	 *     <blockquote><pre>
	 *         plaintext = decrypt(ciphertext, PW);
	 *     </pre></blockquote>
	 *
	 *  2. Asymmetric Private Key:
	 *     <blockquote><pre>
	 *         plaintext = decrypt(ciphertext, SK);
	 *     </pre></blockquote>
	 *
	 * @param ciphertext
	 *        encrypted data
	 *
	 * @param params
	 *        extra params ('IV' for 'AES')
	 *
	 * @return plaintext
	 */
	Decrypt(ciphertext []byte, params StringKeyMap) []byte

	/**
	 *  Check symmetric keys by encryption.
	 *  <blockquote><pre>
	 *      CT = encrypt(data, PK);
	 *      OK = decrypt(CT, SK) == data;
	 *  </pre></blockquote>
	 *
	 * @param pKey
	 *        encrypt (public) key
	 *
	 * @return true on signature matched
	 */
	MatchEncryptKey(pKey EncryptKey) bool
}

type SignKey interface {
	ISignKey
	CryptographyKey
}
type ISignKey interface {

	/**
	 *  Get signature for data
	 *  <blockquote><pre>
	 *      signature = sign(data, SK);
	 *  </pre></blockquote>
	 *
	 * @param data
	 *        data to be signed
	 *
	 * @return signature
	 */
	Sign(data []byte) []byte
}

type VerifyKey interface {
	IVerifyKey
	CryptographyKey
}
type IVerifyKey interface {

	/**
	 *  Verify signature with data
	 *  <blockquote><pre>
	 *      OK = verify(data, signature, PK);
	 *  </pre></blockquote>
	 *
	 * @param data
	 *        data
	 *
	 * @param signature
	 *        signature of data
	 *
	 * @return true on signature matched
	 */
	Verify(data []byte, signature []byte) bool

	/**
	 *  Check asymmetric keys by signature.
	 *  <blockquote><pre>
	 *      signature = sign(data, SK);
	 *      OK = verify(data, signature, PK);
	 *  </pre></blockquote>
	 *
	 * @param sKey
	 *        private key
	 *
	 * @return true on signature matched
	 */
	MatchSignKey(sKey SignKey) bool
}

/**
 *  Asymmetric Cryptography Key
 *  ~~~~~~~~~~~~~~~~~~~~~~~~~~~
 */
type AsymmetricKey interface {
	CryptographyKey
}

//const (
//	RSA = "RSA"  //-- "RSA/ECB/PKCS1Padding", "SHA256withRSA"
//	ECC = "ECC"
//)
