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

// PublicKey defines the interface for asymmetric cryptographic public keys
//
// Used for signature verification
//
//	Key data structure: {
//	    "algorithm" : "RSA", // "ECC", ...
//	    "data"      : "{BASE64_ENCODE}",
//	    ...
//	}
type PublicKey interface {
	AsymmetricKey
	IVerifyKey
}

// PublicKeyFactory defines the factory interface for PublicKey
type PublicKeyFactory interface {

	// ParsePublicKey parses a StringKeyMap into a PublicKey instance
	//
	// Parameters
	//   - key: Key info in StringKeyMap format (matches PublicKey data structure)
	// Returns: Parsed PublicKey instance
	ParsePublicKey(key StringKeyMap) PublicKey
}

//
//  Factory method
//

func ParsePublicKey(key any) PublicKey {
	helper := GetPublicKeyHelper()
	return helper.ParsePublicKey(key)
}

func GetPublicKeyFactory(algorithm string) PublicKeyFactory {
	helper := GetPublicKeyHelper()
	return helper.GetPublicKeyFactory(algorithm)
}

func SetPublicKeyFactory(algorithm string, factory PublicKeyFactory) {
	helper := GetPublicKeyHelper()
	helper.SetPublicKeyFactory(algorithm, factory)
}
