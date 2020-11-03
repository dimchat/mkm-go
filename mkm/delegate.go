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
)

type EntityDataSource interface {

	/**
	 *  Get meta for entity ID
	 *
	 * @param identifier - entity ID
	 * @return meta object
	 */
	GetMeta(identifier *ID) *Meta

	/**
	 *  Get profile for entity ID
	 *
	 * @param identifier - entity ID
	 * @return profile object
	 */
	GetProfile(identifier *ID) *Profile
}

/**
 *  User Data Source
 *  ~~~~~~~~~~~~~~~~
 *
 *  (Encryption/decryption)
 *  1. public key for encryption
 *     if profile.key not exists, means it is the same key with meta.key
 *  2. private keys for decryption
 *     the private keys paired with [profile.key, meta.key]
 *
 *  (Signature/Verification)
 *  3. private key for signature
 *     the private key paired with meta.key
 *  4. public keys for verification
 *     [meta.key]
 */
type UserDataSource interface {
	EntityDataSource

	/**
	 *  Get contacts list
	 *
	 * @param user - user ID
	 * @return contacts list (ID)
	 */
	GetContacts(user *ID) []*ID

	/**
	 *  Get user's public key for encryption
	 *  (profile.key or meta.key)
	 *
	 * @param user - user ID
	 * @return public key
	 */
	GetPublicKeyForEncryption(user *ID) *EncryptKey

	/**
	 *  Get user's private keys for decryption
	 *  (which paired with [profile.key, meta.key])
	 *
	 * @param user - user ID
	 * @return private keys
	 */
	GetPrivateKeysForDecryption(user *ID) []*DecryptKey

	/**
	 *  Get user's private key for signature
	 *  (which paired with profile.key or meta.key)
	 *
	 * @param user - user ID
	 * @return private key
	 */
	GetPrivateKeyForSignature(user *ID) *SignKey

	/**
	 *  Get user's public keys for verification
	 *  [profile.key, meta.key]
	 *
	 * @param user - user ID
	 * @return public keys
	 */
	GetPublicKeysForVerification(user *ID) []*VerifyKey
}

type GroupDataSource interface {
	EntityDataSource

	/**
	 *  Get group founder
	 *
	 * @param group - group ID
	 * @return fonder ID
	 */
	GetFounder(group *ID) *ID

	/**
	 *  Get group owner
	 *
	 * @param group - group ID
	 * @return owner ID
	 */
	GetOwner(group *ID) *ID

	/**
	 *  Get group members list
	 *
	 * @param group - group ID
	 * @return members list (ID)
	 */
	GetMembers(group *ID) []*ID
}
