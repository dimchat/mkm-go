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

/**
 *  Base Document
 *  ~~~~~~~~~~~~~
 */
type BaseDocument struct {
	Dictionary
	IDocument

	_identifier ID
	_type string       // "visa", "bulletin", ...

	_properties map[string]interface{}
	_status int8       // 1 for valid, -1 for invalid
}

func (doc *BaseDocument) Init(dict map[string]interface{}) *BaseDocument {
	if doc.Dictionary.Init(dict) != nil {
		// lazy load
		doc._identifier = nil
		doc._type = ""
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
	if ValueIsNil(data) || ValueIsNil(signature) {
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

func (doc *BaseDocument) data() []byte {
	return DocumentGetData(doc.GetMap(false))
}
func (doc *BaseDocument) signature() []byte {
	return DocumentGetSignature(doc.GetMap(false))
}

/**
 *  Get serialized properties
 *
 * @return JsON string
 */
func DocumentGetData(dict map[string]interface{}) []byte {
	utf8 := dict["data"]
	if utf8 == nil {
		return nil
	} else {
		return UTF8Encode(utf8.(string))
	}
}

/**
 *  Get signature for serialized properties
 *
 * @return signature data
 */
func DocumentGetSignature(dict map[string]interface{}) []byte {
	base64 := dict["signature"]
	if base64 == nil {
		return nil
	} else {
		return Base64Decode(base64.(string))
	}
}

//-------- TAI

func (doc *BaseDocument) IsValid() bool {
	return doc._status == 1
}

func (doc *BaseDocument) Verify(publicKey VerifyKey) bool {
	if doc._status > 0 {
		// already verify OK
		return true
	}
	data := doc.data()
	signature := doc.signature()
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
	// NOTICE: if status is 0, it doesn't mean the document is invalid,
	//         try another key
	return doc._status == 1
}

func (doc *BaseDocument) Sign(privateKey SignKey) (data, signature []byte) {
	data = JSONEncodeMap(doc.Properties())
	signature = privateKey.Sign(data)
	doc.Set("data", UTF8Decode(data))
	doc.Set("signature", Base64Encode(signature))
	return data, signature
}

func (doc *BaseDocument) Properties() map[string]interface{} {
	if doc._status < 0 {
		// invalid
		return nil
	}
	if doc._properties == nil {
		data := doc.data()
		if data == nil {
			// create new properties
			doc._properties = make(map[string]interface{})
		} else {
			// get properties from data
			doc._properties = JSONDecodeMap(data)
		}
	}
	return doc._properties
}

func (doc *BaseDocument) GetProperty(name string) interface{} {
	dict := doc.Properties()
	if dict == nil {
		return nil
	} else {
		return dict[name]
	}
}

func (doc *BaseDocument) SetProperty(name string, value interface{}) {
	// reset status
	doc._status = 0
	// update property value with name
	properties := doc.Properties()
	if ValueIsNil(value) {
		delete(properties, name)
	} else {
		properties[name] = value
	}
	// clear data signature after properties changed
	doc.Remove("data")
	doc.Remove("signature")
}

//-------- IDocument

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

func (doc *BaseDocument) Name() string {
	name := doc.GetProperty("name")
	if name == nil {
		return ""
	} else {
		return name.(string)
	}
}

func (doc *BaseDocument) SetName(name string) {
	doc.SetProperty("name", name)
}
