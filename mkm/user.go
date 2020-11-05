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
	. "github.com/dimchat/mkm-go/crypto"
	. "github.com/dimchat/mkm-go/protocol"
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

func (user *User) Init(identifier ID) *User {
	if user.Entity.Init(identifier) != nil {
		// ...
	}
	return user
}

func (user User) GetDataSource() UserDataSource {
	return user._delegate.(UserDataSource)
}

/**
 *  Get all contacts of the user
 *
 * @return contact list
 */
func (user User) GetContacts() []ID {
	delegate := user.GetDataSource()
	return delegate.GetContacts(user.ID())
}

func (user User) metaKey() PublicKey {
	meta := user.GetMeta()
	if meta == nil {
		return nil
	} else {
		return meta.Key()
	}
}

func (user User) profileKey() EncryptKey {
	profile := user.GetProfile()
	if profile == nil || !profile.IsValid() {
		// profile not found or not valid
		return nil
	} else {
		return profile.GetKey()
	}
}

// NOTICE: meta.key will never changed, so use profile.key to encrypt
//         is the better way
func (user User) encryptKey() EncryptKey {
	// 0. get key from data source
	delegate := user.GetDataSource()
	key := delegate.GetPublicKeyForEncryption(user.ID())
	if key != nil {
		return key
	}
	// 1. get key from profile
	key = user.profileKey()
	if key != nil {
		return key
	}
	// 2. get key from meta
	mKey := user.metaKey()
	if mKey != nil {
		key, ok := mKey.(EncryptKey)
		if ok {
			return key
		}
	}
	//panic("failed to get encrypt key for user")
	return nil
}

// NOTICE: I suggest using the private key paired with meta.key to sign message
//         so here should return the meta.key
func (user User) verifyKeys() []VerifyKey {
	// 0. get keys from data source
	delegate := user.GetDataSource()
	keys := delegate.GetPublicKeysForVerification(user.ID())
	if keys != nil && len(keys) > 0 {
		return keys
	}
	keys = make([]VerifyKey, 0)
	// 1. get key from profile
	pKey := user.profileKey()
	if pKey != nil {
		key, ok := pKey.(VerifyKey)
		if ok {
			keys = append(keys, key)
		}
	}
	// 2. get key from meta
	mKey := user.metaKey()
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
func (user User) Verify(data []byte, signature []byte) bool {
	keys := user.verifyKeys()
	if keys == nil {
		//panic("failed to get keys for verification")
		return false
	}
	for _, key := range keys {
		if key.Verify(data, signature) {
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
func (user User) Encrypt(plaintext []byte) []byte {
	key := user.encryptKey()
	if key == nil {
		//panic("failed to get key for encryption")
		return nil
	} else {
		return key.Encrypt(plaintext)
	}
}

//
//  Interfaces for Local User
//

// NOTICE: I suggest use the private key which paired to meta.key
//         to sign message
func (user User) signKey() SignKey {
	delegate := user.GetDataSource()
	return delegate.GetPrivateKeyForSignature(user.ID())
}

// NOTICE: if you provide a public key in profile for encryption
//         here you should return the private key paired with profile.key
func (user User) decryptKeys() []DecryptKey {
	delegate := user.GetDataSource()
	return delegate.GetPrivateKeysForDecryption(user.ID())
}

/**
 *  Sign data with user's private key
 *
 * @param data - message data
 * @return signature
 */
func (user User) Sign(data []byte) []byte {
	key := user.signKey()
	if key == nil {
		panic("failed to get key for signing")
		return nil
	}
	return key.Sign(data)
}

/**
 *  Decrypt data with user's private key(s)
 *
 * @param ciphertext - encrypted data
 * @return plain text
 */
func (user User) Decrypt(ciphertext []byte) []byte {
	keys := user.decryptKeys()
	if keys == nil {
		panic("failed to get keys for decryption")
		return nil
	}
	var plaintext []byte
	for _, key := range keys {
		plaintext = key.Decrypt(ciphertext)
		if plaintext != nil {
			// OK
			return plaintext
		}
	}
	// decryption failed
	return nil
}
