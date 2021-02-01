/* license: https://mit-license.org
 *
 *  Ming-Ke-Ming : Decentralized User Identity Authentication
 *
 *                                Written in 2021 by Moky <albert.moky@gmail.com>
 *
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
package mkm

import (
	. "github.com/dimchat/mkm-go/crypto"
	. "github.com/dimchat/mkm-go/format"
	. "github.com/dimchat/mkm-go/protocol"
	. "github.com/dimchat/mkm-go/types"
)

type BaseDocument struct {
	Dictionary
	Document

	_identifier ID
	_type string       // "visa", "bulletin", ...

	_data []byte       // JsON.encode(properties)
	_signature []byte  // LocalUser(identifier).sign(data)

	_properties map[string]interface{}
	_status int8       // 1 for valid, -1 for invalid
}

func (doc *BaseDocument) Init(dict map[string]interface{}) *BaseDocument {
	if doc.Dictionary.Init(dict) != nil {
		// lazy load
		doc._identifier = nil
		doc._type = ""
		doc._data = nil
		doc._signature = nil
		doc._properties = nil
		doc._status = 0
	}
	return doc
}

func (doc *BaseDocument) InitWithType(docType string, identifier ID, data []byte, signature []byte) *BaseDocument {
	if docType == "*" {
		docType = ""
	}
	var status int8
	dict := make(map[string]interface{})
	// ID
	dict["ID"] = identifier.String()
	// document type
	if docType != "" {
		dict["type"] = docType
	}
	// data & signature
	if data == nil || signature == nil {
		// create a new empty document with ID and doc type
		status = 0
	} else {
		// create entity document with ID, data and signature loaded from local storage
		dict["data"] = UTF8Decode(data)
		dict["signature"] = Base64Encode(signature)
		status = 1  // all documents must be verified before saving into local storage
	}
	if doc.Dictionary.Init(dict) != nil {
		doc._identifier = identifier
		doc._type = docType
		doc._data = data
		doc._signature = signature
		doc._status = status
		// set type for empty document
		if status == 1 || docType == "" {
			doc._properties = nil
		} else {
			doc._properties = make(map[string]interface{})
			doc._properties["type"] = docType
		}
	}
	return doc
}

func (doc *BaseDocument) Type() string {
	if doc._type == "" {
		docType := doc.GetProperty("type")
		if docType != nil {
			doc._type = docType.(string)
		} else {
			doc._type = DocumentGetType(doc.GetMap(false))
		}
	}
	return doc._type
}

func (doc *BaseDocument) ID() ID {
	if doc._identifier == nil {
		doc._identifier = DocumentGetID(doc.GetMap(false))
	}
	return doc._identifier
}

/**
 *  Get serialized properties
 *
 * @return JsON string
 */
func (doc *BaseDocument) Data() []byte {
	if doc._data == nil {
		utf8 := doc.Get("data")
		if utf8 != nil {
			doc._data = UTF8Encode(utf8.(string))
		}
	}
	return doc._data
}

/**
 *  Get signature for serialized properties
 *
 * @return signature data
 */
func (doc *BaseDocument) Signature() []byte {
	if doc._signature == nil {
		base64 := doc.Get("signature")
		if base64 != nil {
			doc._signature = Base64Decode(base64.(string))
		}
	}
	return doc._signature
}

func (doc *BaseDocument) IsValid() bool {
	return doc._status > 0
}

//-------- signature

func (doc *BaseDocument) Verify(publicKey VerifyKey) bool {
	if doc._status > 0 {
		// already verify OK
		return true
	}
	data := doc.Data()
	signature := doc.Signature()
	if data == nil {
		// NOTICE: if data is empty, signature should be empty at the same time
		//         this happen while profile not found
		if signature == nil {
			doc._status = 0
		} else {
			// data signature error
			doc._status = -1
		}
	} else if signature == nil {
		// signature error
		doc._status = -1
	} else if publicKey.Verify(data, signature) {
		// signature matched
		doc._status = 1
	}
	// NOTICE: if status is 0, it doesn't mean the profile is invalid,
	//         try another key
	return doc._status == 1
}

func (doc *BaseDocument) Sign(privateKey SignKey) []byte {
	if doc._status > 0 {
		// already signed/verified
		return doc._signature
	}
	doc._status = 1
	doc._data = JSONEncode(doc.Properties())
	doc._signature = privateKey.Sign(doc._data)
	doc.Set("data", UTF8Decode(doc._data))
	doc.Set("signature", Base64Encode(doc._signature))
	return doc._signature
}

//-------- properties

func (doc *BaseDocument) Properties() map[string]interface{} {
	if doc._status < 0 {
		// invalid
		return nil
	}
	if doc._properties == nil {
		data := doc.Get("data")
		if data == nil {
			// create new properties
			doc._properties = make(map[string]interface{})
		} else {
			// get properties from data
			doc._properties = JSONDecode(UTF8Encode(data.(string)))
		}
	}
	return doc._properties
}

func (doc *BaseDocument) GetProperty(name string) interface{} {
	dict := doc.Properties()
	if dict == nil {
		return nil
	}
	return dict[name]
}

func (doc *BaseDocument) SetProperty(name string, value interface{}) {
	// 1. reset status
	doc._status = 0
	// 2. update property value with name
	dict := doc.Properties()
	if value == nil {
		delete(dict, name)
	} else {
		dict[name] = value
	}
	// 3. clear data signature after properties changed
	doc.Set("data", nil)
	doc.Set("signature", nil)
	doc._data = nil
	doc._signature = nil
}

//---- properties getter/setter

/**
 *  Get entity name
 *
 * @return name string
 */
func (doc *BaseDocument) Name() string {
	name := doc.GetProperty("name")
	if name == nil {
		return ""
	}
	return name.(string)
}

func (doc *BaseDocument) SetName(name string) {
	doc.SetProperty("name", name)
}
