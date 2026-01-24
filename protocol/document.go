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
package protocol

import (
	. "github.com/dimchat/mkm-go/crypto"
	. "github.com/dimchat/mkm-go/ext"
	. "github.com/dimchat/mkm-go/types"
)

/**
 *  The Additional Information (Profile)
 *
 *      'Meta' is the information for entity which never changed,
 *          which contains the key for verify signature;
 *      'TAI' is the variable part (signed by meta.key's private key),
 *          which could contain a public key for asymmetric encryption.
 */
type TAI interface {

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
	 * @param metaKey - public key in meta.key
	 * @return true on signature matched
	 */
	Verify(metaKey VerifyKey) bool

	/**
	 *  Encode properties to 'data' and sign it to 'signature'
	 *
	 * @param sKey - private key match meta.key
	 * @return signature
	 */
	Sign(sKey SignKey) []byte

	//-------- properties

	/**
	 *  Get all properties
	 *
	 * @return properties
	 */
	Properties() StringKeyMap

	/**
	 *  Get property data with key
	 *
	 * @param name - property name
	 * @return property data
	 */
	GetProperty(name string) interface{}

	/**
	 *  Update property with key and data
	 *  (this will reset 'data' and 'signature')
	 *
	 * @param name - property name
	 * @param value - property data
	 */
	SetProperty(name string, value interface{})
}

/**
 *  User/Group Profile
 *  <p>
 *      This class is used to generate entity profile
 *  </p>
 *
 *  <blockquote><pre>
 *  data format: {
 *      "did"       : "{EntityID}",      // entity ID
 *      "type"      : "visa",            // "bulletin", ...
 *      "data"      : "{JSON}",          // data = json_encode(info)
 *      "signature" : "{BASE64_ENCODE}"  // signature = sign(data, SK);
 *  }
 *  </pre></blockquote>
 */
type Document interface {
	Mapper
	TAI

	/**
	 *  Get entity ID
	 *
	 * @return entity ID
	 */
	//ID() ID

	//---- properties getter/setter

	/**
	 * Get sign time
	 */
	Time() Time

	/**
	 *  Get entity name
	 *
	 * @return name string
	 */
	//Name() string
	//SetName(name string)
}

/**
 *  Document Factory
 *  ~~~~~~~~~~~~~~~~
 */
type DocumentFactory interface {

	/**
	 *  Create document
	 *  <p>
	 *      1. Create document with data &amp; signature loaded from local storage
	 *  </p>
	 *  <p>
	 *      2. Create a new empty document with type
	 *  </p>
	 *
	 * @param data      - document data (JsON)
	 * @param signature - document signature (Base64)
	 * @return Document
	 */
	CreateDocument(data string, signature TransportableData) Document

	/**
	 *  Parse map object to entity document
	 *
	 * @param doc - info
	 * @return Document
	 */
	ParseDocument(doc StringKeyMap) Document
}

//
//  Factory methods
//

func CreateDocument(docType string, data string, signature TransportableData) Document {
	helper := GetDocumentHelper()
	return helper.CreateDocument(docType, data, signature)
}

func ParseDocument(doc interface{}) Document {
	helper := GetDocumentHelper()
	return helper.ParseDocument(doc)
}

func GetDocumentFactory(docType string) DocumentFactory {
	helper := GetDocumentHelper()
	return helper.GetDocumentFactory(docType)
}

func SetDocumentFactory(docType string, factory DocumentFactory) {
	helper := GetDocumentHelper()
	helper.SetDocumentFactory(docType, factory)
}

//
//  Conveniences
//

func DocumentConvert(array interface{}) []Document {
	profiles := FetchList(array)
	documents := make([]Document, 0, len(profiles))
	var doc Document
	for _, item := range profiles {
		doc = ParseDocument(item)
		if doc == nil {
			continue
		}
		documents = append(documents, doc)
	}
	return documents
}

func DocumentRevert(documents []Document) []StringKeyMap {
	array := make([]StringKeyMap, len(documents))
	for idx, doc := range documents {
		array[idx] = doc.Map()
	}
	return array
}
