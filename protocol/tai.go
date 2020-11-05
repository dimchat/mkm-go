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
package protocol

import (
	. "github.com/dimchat/mkm-go/crypto"
)

/**
 *  The Additional Information
 *
 *      'Meta' is the information for entity which never changed,
 *          which contains the key for verify signature;
 *      'TAI' is the variable part,
 *          which could contain a public key for asymmetric encryption.
 */
type TAI interface {

	/**
	 *  Get entity ID
	 *
	 * @return entity ID
	 */
	ID() ID

	/**
	 *  Check if signature matched
	 *
	 * @return False on signature not matched
	 */
	IsValid() bool

	//-------- signature

	/**
	 *  Verify 'data' and 'signature' with public key
	 *
	 * @param publicKey - public key in meta.key
	 * @return true on signature matched
	 */
	Verify(publicKey VerifyKey) bool

	/**
	 *  Encode properties to 'data' and sign it to 'signature'
	 *
	 * @param privateKey - private key match meta.key
	 * @return signature
	 */
	Sign(privateKey SignKey) []byte

	//-------- properties

	/**
	 *  Get all names for properties
	 *
	 * @return profile properties key set
	 */
	PropertyNames() []string

	/**
	 *  Get profile property data with key
	 *
	 * @param name - property name
	 * @return property data
	 */
	GetProperty(name string) interface{}

	/**
	 *  Update profile property with key and data
	 *  (this will reset 'data' and 'signature')
	 *
	 * @param name - property name
	 * @param value - property data
	 */
	SetProperty(name string, value interface{})

	//---- properties getter/setter

	/**
	 *  Get public key to encrypt message for user
	 *
	 * @return public key
	 */
	GetKey() EncryptKey

	/**
	 *  Set public key for other user to encrypt message
	 *
	 * @param publicKey - public key in profile.key
	 */
	SetKey(publicKey EncryptKey)
}
