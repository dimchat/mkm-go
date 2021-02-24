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

type IDocumentExt interface {
	IDocument

	// Check valid status
	checkStatus(publicKey VerifyKey) int8
	status() int8
}

/**
 *  Base Document
 *  ~~~~~~~~~~~~~
 */
type BaseDocument struct {
	Dictionary
	IDocumentExt

	_shadow IDocumentExt

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

func (doc *BaseDocument) SetShadow(shadow IDocumentExt) {
	doc._shadow = shadow
}
func (doc *BaseDocument) Shadow() IDocumentExt {
	return doc._shadow
}

//-------- TAI

func (doc *BaseDocument) IsValid() bool {
	return doc.Shadow().IsValid()
}

func (doc *BaseDocument) Verify(publicKey VerifyKey) bool {
	return doc.Shadow().Verify(publicKey)
}

func (doc *BaseDocument) Sign(privateKey SignKey) (data, signature []byte) {
	return doc.Shadow().Sign(privateKey)
}

func (doc *BaseDocument) Properties() map[string]interface{} {
	if doc._status < 0 {
		// invalid
		return nil
	}
	if doc._properties == nil {
		doc._properties = doc.Shadow().Properties()
	}
	return doc._properties
}

func (doc *BaseDocument) GetProperty(name string) interface{} {
	return doc.Shadow().GetProperty(name)
}

func (doc *BaseDocument) SetProperty(name string, value interface{}) {
	// reset status
	doc._status = 0
	// update property value with name
	doc.Shadow().SetProperty(name, value)
}

//-------- IDocument

func (doc *BaseDocument) Type() string {
	if doc._type == "" {
		doc._type = doc.Shadow().Type()
	}
	return doc._type
}

func (doc *BaseDocument) ID() ID {
	if doc._identifier == nil {
		doc._identifier = doc.Shadow().ID()
	}
	return doc._identifier
}

func (doc *BaseDocument) Name() string {
	return doc.Shadow().Name()
}

func (doc *BaseDocument) SetName(name string) {
	doc.Shadow().SetName(name)
}

//-------- IDocumentExt

func (doc *BaseDocument) checkStatus(publicKey VerifyKey) int8 {
	if doc._status < 1 {
		doc._status = doc.Shadow().checkStatus(publicKey)
	}
	return doc._status
}

func (doc *BaseDocument) status() int8 {
	return doc._status
}
