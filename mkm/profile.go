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
	. "github.com/dimchat/mkm-go/protocol"
	. "github.com/dimchat/mkm-go/types"
)

//  virtual class, need to implement methods:
//      ID() ID
//      GetKey() EncryptKey
//      SetKey(publicKey EncryptKey)
//  before using it
type BaseProfile struct {
	Dictionary
	Profile

	_data []byte       // JsON.encode(properties)
	_signature []byte  // LocalUser(identifier).sign(data)

	_status int8       // 1 for valid, -1 for invalid

	_properties map[string]interface{}
}

func (tai *BaseProfile) Init(dictionary map[string]interface{}) *BaseProfile {
	if tai.Dictionary.Init(dictionary) != nil {
		// lazy load
		tai._data = nil
		tai._signature = nil
		tai._properties = nil
		tai._status = 0
	}
	return tai
}

/**
 *  Get serialized properties
 *
 * @return JsON string
 */
func (tai *BaseProfile) Data() []byte {
	if tai._data == nil {
		str, ok := tai.Get("data").(string)
		if ok {
			tai._data = UTF8Encode(str)
		}
	}
	return tai._data
}

/**
 *  Get signature for serialized properties
 *
 * @return signature data
 */
func (tai *BaseProfile) Signature() []byte {
	if tai._signature == nil {
		str, ok := tai.Get("signature").(string)
		if ok {
			tai._signature = Base64Decode(str)
		}
	}
	return tai._signature
}

func (tai BaseProfile) IsValid() bool {
	return tai._status >= 0
}

//-------- signature

func (tai *BaseProfile) Verify(publicKey VerifyKey) bool {
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
	} else if publicKey.Verify(data, signature) {
		// signature matched
		tai._status = 1
	}
	// NOTICE: if status is 0, it doesn't mean the profile is invalid,
	//         try another key
	return tai._status == 1
}

func (tai *BaseProfile) Sign(privateKey SignKey) []byte {
	if tai._status > 0 {
		// already signed/verified
		return tai._signature
	}
	tai._status = 1
	tai._data = JSONEncode(tai.getProperties())
	tai._signature = privateKey.Sign(tai._data)
	tai.Set("data", UTF8Decode(tai._data))
	tai.Set("signature", Base64Encode(tai._signature))
	return tai._signature
}

//-------- properties

func (tai *BaseProfile) getProperties() map[string]interface{} {
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
				bytes := UTF8Encode(str)
				tai._properties = JSONDecode(bytes)
			} else {
				panic("failed to convert string")
			}
		}
	}
	return tai._properties
}

func (tai *BaseProfile) PropertyNames() []string {
	dict := tai.getProperties()
	if dict == nil {
		return nil
	}
	return MapKeys(dict)
}

func (tai *BaseProfile) GetProperty(name string) interface{} {
	dict := tai.getProperties()
	if dict == nil {
		return nil
	}
	return dict[name]
}

func (tai *BaseProfile) SetProperty(name string, value interface{}) {
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
func (tai *BaseProfile) GetName() string {
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

func (tai *BaseProfile) SetName(name string) {
	tai.SetProperty("name", name)
}
