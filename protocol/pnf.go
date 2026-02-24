/* license: https://mit-license.org
 * ==============================================================================
 * The MIT License (MIT)
 *
 * Copyright (c) 2026 Albert Moky
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

/**
 *  PNF - Portable Network File
 */

// TransportableFile defines the interface for transportable file resources
//
//	Supported serialization formats (subset of TransportableResource):
//	    2. "{URL}"
//	    3. {
//	        "data"     : "...",        // base64_encode(fileContent)
//	        "filename" : "avatar.png",
//
//	        "URL"      : "http://...", // download from CDN (file may be encrypted)
//	        "key"      : {             // symmetric key to decrypt file data
//	            "algorithm" : "AES",   // "DES", ...
//	            "data"      : "{BASE64_ENCODE}",
//	            ... }
//	    }
//
// Note: Large file content should be uploaded to CDN (set URL) instead of embedding in the "data" field
type TransportableFile interface {
	Mapper
	TransportableResource

	// Data returns the raw file content as TransportableData (encoded binary data)
	//
	// For large files, this may return nil (URL is used instead)
	Data() TransportableData
	SetData(data TransportableData)

	// Filename returns the name of the file (e.g., "avatar.png")
	Filename() string
	SetFilename(filename string)

	// URL returns the CDN download URL of the file
	URL() URL
	SetURL(url URL)

	// Password returns the decrypt key for CDN-downloaded encrypted data
	//
	// Default: Plain key (returns original data when decrypted)
	Password() DecryptKey
	SetPassword(key DecryptKey)

	// String returns the encoded string representation of the TransportableFile
	//
	// Possible return values:
	//   - "{URL}"
	//   - "{...}"
	String() string

	// Map returns the inner string-keyed map with encoded file data
	// Updates the inner map with the latest encoded data before returning
	//Map() StringKeyMap

	// Serialize implements TransportableResource interface
	// Serialization logic:
	//   - If only "URL" and "filename" exist: returns String() result
	//   - Otherwise: returns Map() result
	//Serialize() any
}

/**
 *  PNF Factory
 */

// TransportableFileFactory defines the factory interface for TransportableFile
type TransportableFileFactory interface {

	// CreateTransportableFile creates a TransportableFile (PNF) object
	//
	// Parameters:
	//   - data: Raw file content (unencrypted) as TransportableData (nil for CDN-only files)
	//   - filename: Name of the file (e.g., "avatar.png")
	//   - url: CDN download URL (optional, used for large files)
	//   - password: Decrypt key for CDN-downloaded encrypted data (nil for unencrypted files)
	// Returns: Newly created TransportableFile (PNF) object
	CreateTransportableFile(
		data TransportableData, filename string,
		url URL, password DecryptKey,
	) TransportableFile

	// ParseTransportableFile parses a map object into a TransportableFile (PNF) object
	//
	// Parameters:
	//   - pnf: Input PNF info (StringKeyMap representing format 0 or 1 of TransportableFile)
	// Returns: Parsed TransportableFile (PNF) object
	ParseTransportableFile(pnf StringKeyMap) TransportableFile
}

//
//  Factory methods
//

func CreateTransportableFile(data TransportableData, filename string,
	url URL, password DecryptKey) TransportableFile {
	helper := GetTransportableFileHelper()
	return helper.CreateTransportableFile(data, filename, url, password)
}

func ParseTransportableFile(pnf any) TransportableFile {
	helper := GetTransportableFileHelper()
	return helper.ParseTransportableFile(pnf)
}

func GetTransportableFileFactory() TransportableFileFactory {
	helper := GetTransportableFileHelper()
	return helper.GetTransportableFileFactory()
}

func SetTransportableFileFactory(factory TransportableFileFactory) {
	helper := GetTransportableFileHelper()
	helper.SetTransportableFileFactory(factory)
}
