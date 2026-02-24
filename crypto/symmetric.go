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

import . "github.com/dimchat/mkm-go/types"

// SymmetricKey defines the interface for symmetric cryptographic keys
//
// Used for both encryption and decryption of message data (same key for both operations)
//
//	Key data structure: {
//	    "algorithm" : "AES", // "DES", ...
//	    "data"      : "{BASE64_ENCODE}",
//	    ...
//	}
type SymmetricKey interface {
	CryptographyKey
	IEncryptKey
	IDecryptKey
}

//const (
//	AES = "AES"  //-- "AES/CBC/PKCS7Padding"
//	DES = "DES"
//)

// SymmetricKeyFactory defines the factory interface for SymmetricKey
type SymmetricKeyFactory interface {

	// GenerateSymmetricKey creates a new random SymmetricKey using the default algorithm
	//
	// Returns: Newly generated SymmetricKey instance
	GenerateSymmetricKey() SymmetricKey

	// ParseSymmetricKey parses a StringKeyMap into a SymmetricKey instance
	//
	// Parameters:
	//   - key: Key info in StringKeyMap format (matches SymmetricKey data structure)
	// Returns: Parsed SymmetricKey instance
	ParseSymmetricKey(key StringKeyMap) SymmetricKey
}

//
//  Factory methods
//

func GenerateSymmetricKey(algorithm string) SymmetricKey {
	helper := GetSymmetricKeyHelper()
	return helper.GenerateSymmetricKey(algorithm)
}

func ParseSymmetricKey(key any) SymmetricKey {
	helper := GetSymmetricKeyHelper()
	return helper.ParseSymmetricKey(key)
}

func GetSymmetricKeyFactory(algorithm string) SymmetricKeyFactory {
	helper := GetSymmetricKeyHelper()
	return helper.GetSymmetricKeyFactory(algorithm)
}

func SetSymmetricKeyFactory(algorithm string, factory SymmetricKeyFactory) {
	helper := GetSymmetricKeyHelper()
	helper.SetSymmetricKeyFactory(algorithm, factory)
}
