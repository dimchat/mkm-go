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
package format

import . "github.com/dimchat/mkm-go/types"

// TransportableResource defines the interface for serializable data or file content
//
//	Supported serialization formats:
//	    0. "{BASE64_ENCODE}"
//	    1. "data:image/png;base64,{BASE64_ENCODE}"
//	    2. "https://..."
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
type TransportableResource interface {

	/*  Format
	 *  ~~~~~~
	 *
	 *      TED - TransportableData
	 *          0. "{BASE64_ENCODE}"
	 *          1. "data:image/png;base64,{BASE64_ENCODE}"
	 *
	 *      PNF - TransportableFile
	 *          2. "https://..."
	 *          3. {...}
	 */

	// Serialize encodes the resource data into a transportable format
	//
	// The return type can be a string (formats 0,1,2) or a map (format 3)
	Serialize() any
}

/**
 *  TED - Transportable Encoded Data
 */

// TransportableData defines the interface for transportable encoded data
//
//	Supported serialization formats (subset of TransportableResource):
//	    0. "{BASE64_ENCODE}"
//	    1. "data:image/png;base64,{BASE64_ENCODE}"
type TransportableData interface {
	Stringer
	TransportableResource

	// Supported encoding algorithms (reference):
	//   - "base64" (DEFAULT)
	//   - "base58"
	//   - "hex"

	// Encoding returns the data encoding algorithm used
	// Typical return value: "base64"
	Encoding() string

	// Bytes returns the original raw (plaintext) binary data
	Bytes() []byte

	// Size returns the length of the original binary data: len(plaintext)
	Size() int

	// String returns the encoded string representation
	//
	// Possible return values:
	//   - "{BASE64_ENCODE}"
	//   - "data:image/png;base64,{BASE64_ENCODE}"
	//String() string

	// Serialize implements TransportableResource interface (aliases String())
	//Serialize() any
}

/**
 *  TED Factory
 */

// TransportableDataFactory defines the factory interface for TransportableData
type TransportableDataFactory interface {

	// ParseTransportableData parses a TED string into a TransportableData object
	//
	// Parameters:
	//   - ted: Input TED string (format 0 or 1 of TransportableData)
	//
	// Returns: Parsed TransportableData object
	ParseTransportableData(ted string) TransportableData
}

//
//  Factory method
//

func ParseTransportableData(ted any) TransportableData {
	helper := GetTransportableDataHelper()
	return helper.ParseTransportableData(ted)
}

func GetTransportableDataFactory() TransportableDataFactory {
	helper := GetTransportableDataHelper()
	return helper.GetTransportableDataFactory()
}

func SetTransportableDataFactory(factory TransportableDataFactory) {
	helper := GetTransportableDataHelper()
	helper.SetTransportableDataFactory(factory)
}
