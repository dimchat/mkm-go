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

	_identifier ID

	_data string       // JSONEncode(properties)
	_signature []byte  // LocalUser(id).Sign(data)

	_properties map[string]interface{}
	_status int8       // 1 for valid, -1 for invalid
}

func (doc *BaseDocument) Init(dict map[string]interface{}) Document {
	if doc.Dictionary.Init(dict) != nil {
		// lazy load
		doc._identifier = nil
		doc._data = ""
		doc._signature = nil
		doc._properties = nil
		doc._status = 0
	}
	return doc
}

// load document info from local database
func (doc *BaseDocument) InitWithData(identifier ID, data string, signature string) Document {
	dict := make(map[string]interface{})
	dict["ID"] = identifier.String()  // ID string
	dict["data"] = data               // JsON string
	dict["signature"] = signature     // Base64 string
	// create
	if doc.Dictionary.Init(dict) != nil {
		doc._identifier = identifier
		doc._data = data
		doc._signature = nil  // lazy
		doc._properties = nil
		// all documents must be verified before saving into local storage
		doc._status = 1
	}
	return doc
}

// create empty document with ID & type
func (doc *BaseDocument) InitWithID(identifier ID, docType string) Document {
	dict := make(map[string]interface{})
	dict["ID"] = identifier.String()  // ID string
	// create
	if doc.Dictionary.Init(dict) != nil {
		doc._identifier = identifier
		doc._data = ""
		doc._signature = nil
		if docType == "" {
			doc._properties = nil
		} else {
			doc._properties = make(map[string]interface{})
			doc._properties["type"] = docType
		}
		// all documents must be verified before saving into local storage
		doc._status = 0
	}
	return doc
}

func (doc *BaseDocument) data() string {
	if doc._data == "" {
		json := doc.Get("data")
		if json != nil {
			doc._data, _ = json.(string)
		}
	}
	return doc._data
}
func (doc *BaseDocument) signature() []byte {
	if doc._signature == nil {
		base64 := doc.Get("signature")
		if base64 != nil {
			doc._signature = Base64Decode(base64.(string))
		}
	}
	return doc._signature
}

//-------- TAI

func (doc *BaseDocument) IsValid() bool {
	return doc._status > 0
}

func (doc *BaseDocument) Verify(publicKey VerifyKey) bool {
	if doc._status > 0 {
		// already verify OK
		return true
	}
	data := doc.data()
	signature := doc.signature()
	if data == "" {
		// NOTICE: if data is empty, signature should be empty at the same time
		//         this happen while entity document not found
		if signature == nil || len(signature) == 0 {
			doc._status = 0
		} else {
			// data signature error
			doc._status = -1
		}
	} else if signature == nil || len(signature) == 0 {
		// signature error
		doc._status = -1
	} else if publicKey.Verify(UTF8Encode(data), signature) {
		// signature matched
		doc._status = 1
	}
	// NOTICE: if status is 0, it doesn't mean the document is invalid,
	//         try another key
	return doc._status > 0
}

func (doc *BaseDocument) Sign(privateKey SignKey) []byte {
	if doc._status > 0 {
		// already signed/verified
		return doc.signature()
	}
	// update sign time
	doc.SetProperty("time", TimeToFloat64(TimeNow()))
	// sign
	doc._data = JSONEncodeMap(doc.Properties())
	doc._signature = privateKey.Sign(UTF8Encode(doc._data))
	doc.Set("data", doc._data)
	doc.Set("signature", Base64Encode(doc._signature))
	// update status
	doc._status = 1
	return doc._signature
}

func (doc *BaseDocument) Properties() map[string]interface{} {
	if doc._status < 0 {
		// invalid
		return nil
	}
	if doc._properties == nil {
		data := doc.data()
		if data == "" {
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
	// 1. reset status
	doc._status = 0
	// 2. update property value with name
	properties := doc.Properties()
	if ValueIsNil(value) {
		delete(properties, name)
	} else {
		properties[name] = value
	}
	// 3. clear data signature after properties changed
	doc.Remove("data")
	doc.Remove("signature")
	doc._data = ""
	doc._signature = nil
}

//-------- IDocument

func (doc *BaseDocument) Type() string {
	text, ok := doc.GetProperty("type").(string)
	if ok && text != "" {
		return text
	} else {
		return DocumentGetType(doc.Map())
	}
}

func (doc *BaseDocument) ID() ID {
	if doc._identifier == nil {
		doc._identifier = DocumentGetID(doc.Map())
	}
	return doc._identifier
}

func (doc *BaseDocument) Time() Time {
	timestamp := doc.GetProperty("time")
	return TimeParse(timestamp)
}

func (doc *BaseDocument) Name() string {
	text, ok := doc.GetProperty("name").(string)
	if ok {
		return text
	} else {
		return ""
	}
}

func (doc *BaseDocument) SetName(name string) {
	doc.SetProperty("name", name)
}
