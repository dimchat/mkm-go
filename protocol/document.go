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
	. "github.com/dimchat/mkm-go/format"
	. "github.com/dimchat/mkm-go/types"
)

// TAI (The Additional Information / Profile) defines the interface for mutable entity metadata
//
// Key distinctions from Meta:
//   - Meta: Immutable entity information (contains verification keys for signatures)
//   - TAI: Mutable entity profile data (signed with the private key matching meta.key)
//   - TAI may contain public keys for asymmetric encryption
type TAI interface {

	// IsValid checks if the TAI's signature is valid and matches the contained data
	//
	// Returns: false if signature verification fails, true if all checks pass
	IsValid() bool

	//-------- signature

	// Verify validates the TAI's data and signature using the provided meta public key
	//
	// Parameters
	//   - metaKey: Public verification key from meta.key (matches the signing private key)
	// Returns: true if signature matches the data and public key, false otherwise
	Verify(metaKey VerifyKey) bool

	// Sign encodes TAI properties to data and generates a signature for the data
	//
	// Resets the internal 'data' and 'signature' fields with the new values
	//
	// Parameters
	//   - sKey: Private signing key that matches meta.key (used to sign the data)
	// Returns: Raw binary signature of the encoded properties data
	Sign(sKey SignKey) []byte

	//-------- properties

	// Properties returns all key-value properties stored in the TAI
	//
	// Returns: StringKeyMap containing all TAI properties
	Properties() StringKeyMap

	// GetProperty retrieves the value of a specific property by name
	//
	// Parameters
	//   - name: Name of the property to retrieve
	// Returns: Value of the property (any type), nil if the property does not exist
	GetProperty(name string) any

	// SetProperty updates a property with the given name and value
	//
	// Note: This operation resets the internal 'data' and 'signature' fields (requires re-signing)
	//
	// Parameters:
	//   - name: Name of the property to update
	//   - value: New value for the property (any type)
	SetProperty(name string, value any)
}

// DocumentType defines the type of a Document (string alias for type safety)
// Example values: "visa", "bulletin", etc.
type DocumentType = string

// Document defines the interface for User/Group profile documents
//
// Extends Mapper and TAI interfaces, representing a signed entity profile document
//
//	Data structure: {
//	    "did"       : "{EntityID}",      // entity ID
//	    "type"      : "visa",            // "bulletin", ...
//	    "data"      : "{JSON}",          // data = json_encode(info)
//	    "signature" : "{BASE64_ENCODE}"  // signature = sign(data, SK);
//	}
type Document interface {
	Mapper
	TAI

	// ID returns the entity ID associated with the document
	//
	// Returns: Entity ID (ID interface) of the document owner
	//ID() ID

	//---- properties getter/setter

	// Time returns the timestamp when the document was signed
	//
	// Represents the creation/last update time of the document
	Time() Time

	// Name returns the entity name associated with the document
	//
	// Returns: String name of the document owner (entity)
	//Name() string
	//SetName(name string)
}

// DocumentFactory defines the factory interface for Document
type DocumentFactory interface {

	// CreateDocument creates a Document instance from raw data and signature
	//
	// Two primary use cases:
	//   1. Load existing document from local storage (with pre-existing data and signature)
	//   2. Create a new empty document (with empty data/signature) by specifying DocumentType
	//
	// Parameters:
	//   - data: Encoded document data (JSON string)
	//   - signature: Signature of the data (BASE64-encoded TransportableData)
	// Returns: Newly created Document instance
	CreateDocument(data string, signature TransportableData) Document

	// ParseDocument converts a StringKeyMap into a Document instance
	//
	// Expects the map to follow the Document data structure format (did/type/data/signature)
	//
	// Parameters
	//   - doc: Document information in StringKeyMap format
	// Returns: Parsed Document instance (nil if parsing/validation fails)
	ParseDocument(doc StringKeyMap) Document
}

//
//  Factory methods
//

func CreateDocument(docType DocumentType, data string, signature TransportableData) Document {
	helper := GetDocumentHelper()
	return helper.CreateDocument(docType, data, signature)
}

func ParseDocument(doc any) Document {
	helper := GetDocumentHelper()
	return helper.ParseDocument(doc)
}

func GetDocumentFactory(docType DocumentType) DocumentFactory {
	helper := GetDocumentHelper()
	return helper.GetDocumentFactory(docType)
}

func SetDocumentFactory(docType DocumentType, factory DocumentFactory) {
	helper := GetDocumentHelper()
	helper.SetDocumentFactory(docType, factory)
}

//
//  Conveniences
//

func DocumentConvert(array any) []Document {
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
