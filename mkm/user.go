/* license: https://mit-license.org
 *
 *  Ming-Ke-Ming : Decentralized User Identity Authentication
 *
 *                                Written in 2020 by Moky <albert.moky@gmail.com>
 *
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
package mkm

import (
	"mkm-go/crypto"
	"unsafe"
)

/**
 *  User account for communication
 *  ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 *  This class is for creating user account
 *
 *  functions:
 *      (User)
 *      1. verify(data, signature) - verify (encrypted content) data and signature
 *      2. encrypt(data)           - encrypt (symmetric key) data
 *      (LocalUser)
 *      3. sign(data)    - calculate signature of (encrypted content) data
 *      4. decrypt(data) - decrypt (symmetric key) data
 */
type User struct {
	Entity
}

func (user *User) GetDataSource() *UserDataSource {
	return (*UserDataSource)(unsafe.Pointer(user._delegate))
}

/**
 *  Get all contacts of the user
 *
 * @return contact list
 */
func (user *User) GetContacts() []*ID {
	ds := user.GetDataSource()
	return (*ds).GetContacts(user._identifier)
}

func (user *User) metaKey() *crypto.VerifyKey {
	meta := (*user).Meta()
	if meta == nil {
		return nil
	}
	key := (*meta).Key()
	if key == nil {
		return nil
	}
	return (*crypto.VerifyKey)(unsafe.Pointer(key))
}

func (user *User) profileKey() *crypto.EncryptKey {
	profile := (*user).Profile()
	if profile == nil || !(*profile).IsValid() {
		// profile not found or not valid
		return nil
	}
	return (*profile).GetKey()
}

// NOTICE: meta.key will never changed, so use profile.key to encrypt
//         is the better way
func (user *User) encryptKey() *crypto.EncryptKey {
	// 0. get key from data source
	ds := (*user).GetDataSource()
	key := (*ds).GetPublicKeyForEncryption((*user)._identifier)
	if key != nil {
		return key
	}
	// 1. get key from profile
	key = (*user).profileKey()
	if key != nil {
		return key
	}
	// 2. get key from meta
	mKey := (*user).metaKey()
	if mKey != nil {
		k, ok := (*mKey).(crypto.EncryptKey)
		if ok {
			return &k
		}
	}
	//panic("failed to get encrypt key for user")
	return nil
}

// NOTICE: I suggest using the private key paired with meta.key to sign message
//         so here should return the meta.key
func (user *User) verifyKeys() []*crypto.VerifyKey {
	// 0. get keys from data source
	ds := (*user).GetDataSource()
	keys := (*ds).GetPublicKeysForVerification((*user)._identifier)
	if keys != nil && len(keys) > 0 {
		return keys
	}
	keys = make([]*crypto.VerifyKey, 0)
	// 1. get key from profile
	pKey := (*user).profileKey()
	if pKey != nil {
		k, ok := (*pKey).(crypto.VerifyKey)
		if ok {
			keys = append(keys, &k)
		}
	}
	// 2. get key from meta
	mKey := (*user).metaKey()
	if mKey != nil {
		keys = append(keys, mKey)
	}
	return keys
}

/**
 *  Verify data and signature with user's public keys
 *
 * @param data - message data
 * @param signature - message signature
 * @return true on correct
 */
func (user *User) Verify(data []byte, signature []byte) bool {
	keys := (*user).verifyKeys()
	if keys == nil {
		//panic("failed to get keys for verification")
		return false
	}
	for _, key := range keys {
		if (*key).Verify(data, signature) {
			// matched!
			return true
		}
	}
	return false
}

/**
 *  Encrypt data, try profile.key first, if not found, use meta.key
 *
 * @param plaintext - message data
 * @return encrypted data
 */
func (user *User) Encrypt(plaintext []byte) []byte {
	key := (*user).encryptKey()
	if key == nil {
		panic("failed to get key for encryption")
		return nil
	}
	return (*key).Encrypt(plaintext)
}

//
//  Interfaces for Local User
//

// NOTICE: I suggest use the private key which paired to meta.key
//         to sign message
func (user *User) signKey() *crypto.SignKey {
	ds := (*user).GetDataSource()
	return (*ds).GetPrivateKeyForSignature((*user)._identifier)
}

// NOTICE: if you provide a public key in profile for encryption
//         here you should return the private key paired with profile.key
func (user *User) decryptKeys() []*crypto.DecryptKey {
	ds := (*user).GetDataSource()
	return (*ds).GetPrivateKeysForDecryption((*user)._identifier)
}

/**
 *  Sign data with user's private key
 *
 * @param data - message data
 * @return signature
 */
func (user *User) Sign(data []byte) []byte {
	key := (*user).signKey()
	if key == nil {
		panic("failed to get key for signing")
		return nil
	}
	return (*key).Sign(data)
}

/**
 *  Decrypt data with user's private key(s)
 *
 * @param ciphertext - encrypted data
 * @return plain text
 */
func (user *User) Decrypt(ciphertext []byte) []byte {
	keys := (*user).decryptKeys()
	if keys == nil {
		panic("failed to get keys for decryption")
		return nil
	}
	var plaintext []byte
	for _, key := range keys {
		plaintext = (*key).Decrypt(ciphertext)
		if plaintext != nil {
			// OK
			return plaintext
		}
	}
	// decryption failed
	return nil
}
