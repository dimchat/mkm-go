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

// CryptographyKey defines the base interface for cryptographic keys with designated algorithms
//
//	Key data structure: {
//	    "algorithm" : "RSA", // ECC, AES, ...
//	    "data"      : "{BASE64_ENCODE}",
//	    ...
//	}
type CryptographyKey interface {
	Mapper

	// Algorithm returns the name of the cryptographic algorithm associated with the key
	//
	// Typical return values: "RSA", "ECC", "AES", ...
	Algorithm() string

	// Data returns the key material as TransportableData (encoded binary key data)
	Data() TransportableData
}

// EncryptKey represents a key capable of encryption (symmetric or asymmetric public key)
type EncryptKey interface {
	IEncryptKey
	CryptographyKey
}
type IEncryptKey interface {

	/**
	 *  1. Symmetric Key:
	 *         ciphertext = encrypt(plaintext, PW)
	 *
	 *  2. Asymmetric Public Key:
	 *         ciphertext = encrypt(plaintext, PK);
	 */

	// Encrypt encrypts plaintext binary data using the key
	//
	// Parameters:
	//   - plaintext: Raw binary data to be encrypted (plaintext)
	//   - extra: Algorithm-specific extra parameters (e.g., "IV" for AES encryption)
	// Returns: Encrypted binary data (ciphertext)
	Encrypt(plaintext []byte, extra StringKeyMap) []byte
}

// DecryptKey represents a key capable of decryption (symmetric or asymmetric private key)
type DecryptKey interface {
	IDecryptKey
	CryptographyKey
}
type IDecryptKey interface {

	/**
	 *  1. Symmetric Key:
	 *         plaintext = decrypt(ciphertext, PW);
	 *
	 *  2. Asymmetric Private Key:
	 *         plaintext = decrypt(ciphertext, SK);
	 */

	// Decrypt decrypts ciphertext binary data using the key
	//
	// Parameters:
	//   - ciphertext: Encrypted binary data to be decrypted
	//   - params: Algorithm-specific parameters (e.g., "IV" for AES decryption)
	// Returns: Decrypted binary data (plaintext)
	Decrypt(ciphertext []byte, params StringKeyMap) []byte

	/**
	 *  Check symmetric keys by encryption.
	 *      CT = encrypt(data, PK);
	 *      OK = decrypt(CT, SK) == data;
	 */

	// MatchEncryptKey verifies if this DecryptKey matches the given EncryptKey (key pair validation)
	//
	// Validation logic:
	//   1. Encrypt test data with the provided EncryptKey
	//   2. Decrypt the ciphertext with this DecryptKey
	//   3. Check if decrypted data matches original test data
	//
	// Parameters:
	//   - pKey: EncryptKey (public/secret key) to validate against
	// Returns: true if key pair is valid, false otherwise
	MatchEncryptKey(pKey EncryptKey) bool
}

// SignKey represents a key capable of generating digital signatures (asymmetric private key)
type SignKey interface {
	ISignKey
	CryptographyKey
}
type ISignKey interface {

	/**
	 *  Get signature for data
	 *      signature = sign(data, SK);
	 */

	// Sign generates a digital signature for the given binary data
	//
	// Parameters:
	//   - data: Binary data to be signed
	// Returns: Binary digital signature of the input data
	Sign(data []byte) []byte
}

// VerifyKey represents a key capable of verifying digital signatures (asymmetric public key)
type VerifyKey interface {
	IVerifyKey
	CryptographyKey
}
type IVerifyKey interface {

	/**
	 *  Verify signature with data
	 *      OK = verify(data, signature, PK);
	 */

	// Verify checks if the given signature is valid for the provided data
	//
	// Parameters:
	//   - data: Original binary data that was signed
	//   - signature: Digital signature to verify against the data
	// Returns: true if signature is valid, false otherwise
	Verify(data []byte, signature []byte) bool

	/**
	 *  Check asymmetric keys by signature.
	 *      signature = sign(data, SK);
	 *      OK = verify(data, signature, PK);
	 */

	// MatchSignKey verifies if this VerifyKey matches the given SignKey (signature key pair validation)
	//
	// Validation logic:
	//   1. Generate signature for test data with the provided SignKey
	//   2. Verify the signature with this VerifyKey
	//   3. Check if verification succeeds
	//
	// Parameters:
	//   - sKey: SignKey (private key) to validate against
	// Returns: true if key pair is valid, false otherwise
	MatchSignKey(sKey SignKey) bool
}

// AsymmetricKey defines the interface for asymmetric cryptographic keys
type AsymmetricKey interface {
	CryptographyKey
}

//const (
//	RSA = "RSA"  //-- "RSA/ECB/PKCS1Padding", "SHA256withRSA"
//	ECC = "ECC"
//)
