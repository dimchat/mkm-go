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
	 *  Check if signature matched
	 *
	 * @return False on signature not matched
	 */
	IsValid() bool

	//-------- signature

	/**
	 *  Verify 'data' and 'signature' with public key
	 *
	 * @param publicKey - public key as meta.key
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
	 *  Get all properties
	 *
	 * @return properties
	 */
	Properties() map[string]interface{}

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
}

const (
	// Document types
	VISA = "visa"          // for login/communication
	PROFILE = "profile"    // for user info
	BULLETIN = "bulletin"  // for group info
)

/**
 *  User/Group Profile
 *  ~~~~~~~~~~~~~~~~~~
 *  This class is used to generate entity profile
 *
 *      data format: {
 *          ID: "EntityID",   // entity ID
 *          data: "{JSON}",   // data = json_encode(info)
 *          signature: "..."  // signature = sign(data, SK);
 *      }
 */
type Document interface {
	Map
	TAI

	/**
	 *  Get document type
	 *
	 * @return document type
	 */
	Type() string

	/**
	 *  Get entity ID
	 *
	 * @return entity ID
	 */
	ID() ID

	//---- properties getter/setter

	/**
	 *  Get entity name
	 *
	 * @return name string
	 */
	Name() string

	/**
	 *  Set entity name
	 *
	 * @param name - nickname of user; title of group
	 */
	SetName(name string)
}

func DocumentGetType(doc map[string]interface{}) string {
	str := doc["type"]
	if str == nil {
		return ""
	}
	return str.(string)
}

func DocumentGetID(doc map[string]interface{}) ID {
	return IDParse(doc["ID"])
}

/**
 *  Document Factory
 *  ~~~~~~~~~~~~~~~~
 */
type DocumentFactory interface {

	/**
	 *  Create document with data & signature loaded from local storage
	 *  (If data & signature empty, create a new empty document with entity ID)
	 *
	 * @param identifier - entity ID
	 * @param data       - document data
	 * @param signature  - document signature
	 * @return Document
	 */
	CreateDocument(identifier ID, data []byte, signature []byte) Document

	/**
	 *  Parse map object to entity document
	 *
	 * @param doc - info
	 * @return Document
	 */
	ParseDocument(doc map[string]interface{}) Document
}

var documentFactories = make(map[string]DocumentFactory)

func DocumentRegister(docType string, factory DocumentFactory) {
	documentFactories[docType] = factory
}

func DocumentGetFactory(docType string) DocumentFactory {
	return documentFactories[docType]
}

//
//  Factory methods
//
func DocumentCreate(docType string, identifier ID, data []byte, signature []byte) Document {
	factory := DocumentGetFactory(docType)
	if factory == nil {
		panic("document type not found: " + docType)
	}
	return factory.CreateDocument(identifier, data, signature)
}

func DocumentParse(doc interface{}) Document {
	if doc == nil {
		return nil
	}
	var info map[string]interface{}
	value := ObjectValue(doc)
	switch value.(type) {
	case Document:
		return value.(Document)
	case Map:
		info = value.(Map).GetMap(false)
	case map[string]interface{}:
		info = value.(map[string]interface{})
	default:
		panic(doc)
	}
	docType := DocumentGetType(info)
	factory := DocumentGetFactory(docType)
	if factory == nil {
		factory = DocumentGetFactory("*")  // unknown
	}
	return factory.ParseDocument(info)
}
