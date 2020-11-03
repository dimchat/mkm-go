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
	. "github.com/dimchat/mkm-go/format"
	. "github.com/dimchat/mkm-go/types"
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
	ID() *String

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
	Verify(publicKey *VerifyKey) bool

	/**
	 *  Encode properties to 'data' and sign it to 'signature'
	 *
	 * @param privateKey - private key match meta.key
	 * @return signature
	 */
	Sign(privateKey *SignKey) []byte

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
	GetKey() *EncryptKey

	/**
	 *  Set public key for other user to encrypt message
	 *
	 * @param publicKey - public key in profile.key
	 */
	SetKey(publicKey *EncryptKey)
}

type Profile struct {
	Dictionary
	TAI

	_identifier *String  // ID or String
	_data []byte               // JsON.encode(properties)
	_signature []byte          // LocalUser(identifier).sign(data)

	_properties map[string]interface{}
	_status int8              // 1 for valid, -1 for invalid
}

func (tai *Profile) Init(dictionary map[string]interface{}) *Profile {
	if tai.Dictionary.Init(dictionary) != nil {
		// lazy load
		tai._identifier = nil
		tai._data = nil
		tai._signature = nil
		tai._properties = nil
		tai._status = 0
	}
	return tai
}

func (tai *Profile) ID() *String {
	if tai._identifier == nil {
		id := tai.Get("ID")
		if tai != nil {
			str, ok := id.(string)
			if ok {
				tai._identifier = CreateString(str)
			}
		}
	}
	return tai._identifier
}

/**
 *  Get serialized properties
 *
 * @return JsON string
 */
func (tai *Profile) Data() []byte {
	if tai._data == nil {
		json := tai.Get("data")
		if json != nil {
			str, ok := json.(string)
			if ok {
				tai._data = UTF8BytesFromString(str)
			}
		}
	}
	return tai._data
}

/**
 *  Get signature for serialized properties
 *
 * @return signature data
 */
func (tai *Profile) Signature() []byte {
	if tai._signature == nil {
		b64 := tai.Get("signature")
		if b64 != nil {
			str, ok := b64.(string)
			if ok {
				tai._signature = Base64Decode(str)
			}
		}
	}
	return tai._signature
}

func (tai *Profile) IsValid() bool {
	return tai._status >= 0
}

//-------- signature

func (tai *Profile) Verify(publicKey *VerifyKey) bool {
	if tai._status > 0 {
		// already verify OK
		return true
	}
	data := tai.Data()
	signature := tai.Signature()
	if data == nil {
		// NOTICE: if data is empty, signature should be empty at the same time
		//         this happen while profile not found
		if signature == nil {
			tai._status = 0
		} else {
			// data signature error
			tai._status = -1
		}
	} else if tai._signature == nil {
		// signature error
		tai._status = -1
	} else if (*publicKey).Verify(data, signature) {
		// signature matched
		tai._status = 1
	}
	// NOTICE: if status is 0, it doesn't mean the profile is invalid,
	//         try another key
	return tai._status == 1
}

func (tai *Profile) Sign(privateKey *SignKey) []byte {
	if tai._status > 0 {
		// already signed/verified
		return tai._signature
	}
	tai._status = 1
	tai._data = JSONBytesFromMap(tai.getProperties())
	tai._signature = (*privateKey).Sign(tai._data)
	tai.Set("data", UTF8StringFromBytes(tai._data))
	tai.Set("signature", Base64Encode(tai._signature))
	return tai._signature
}

//-------- properties

func (tai *Profile) getProperties() map[string]interface{} {
	if tai._status < 0 {
		// invalid
		return nil
	}
	if tai._properties == nil {
		data := tai.Get("data")
		if data == nil {
			// create new properties
			tai._properties = make(map[string]interface{})
		} else {
			str, ok := data.(string)
			if ok {
				bytes := UTF8BytesFromString(str)
				tai._properties = JSONMapFromBytes(bytes)
			} else {
				panic("failed to convert string")
			}
		}
	}
	return tai._properties
}

func (tai *Profile) PropertyNames() []string {
	dict := tai.getProperties()
	if dict == nil {
		return nil
	}
	return MapKeys(dict)
}

func (tai *Profile) GetProperty(name string) interface{} {
	dict := tai.getProperties()
	if dict == nil {
		return nil
	}
	return dict[name]
}

func (tai *Profile) SetProperty(name string, value interface{}) {
	// 1. reset status
	tai._status = 0
	// 2. update property value with name
	dict := tai.getProperties()
	if value == nil {
		delete(dict, name)
	} else {
		dict[name] = value
	}
	// 3. clear data signature after properties changed
	tai.Set("data", nil)
	tai.Set("signature", nil)
	tai._data = nil
	tai._signature = nil
}

//---- properties getter/setter

/**
 *  Get entity name
 *
 * @return name string
 */
func (tai *Profile) GetName() string {
	name := tai.GetProperty("name")
	if name == nil {
		return ""
	}
	str, ok := name.(string)
	if ok {
		return str
	} else {
		return ""
	}
}

func (tai *Profile) SetName(name string) {
	tai.SetProperty("name", name)
}
