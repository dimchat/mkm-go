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

import (
	. "github.com/dimchat/mkm-go/types"
)

/**
 *  Serializable Data or File Content
 *
 *  <blockquote><pre>
 *  0. "{BASE64_ENCODE}"
 *  1. "data:image/png;base64,{BASE64_ENCODE}"
 *  2. "https://..."
 *  3. {
 *        "data"     : "...",        // base64_encode(fileContent)
 *        "filename" : "avatar.png",
 *
 *        "URL"      : "http://...", // download from CDN
 *         // before fileContent uploaded to a public CDN,
 *         // it can be encrypted by a symmetric key
 *        "key"      : {             // symmetric key to decrypt file data
 *            "algorithm" : "AES",   // "DES", ...
 *            "data"      : "{BASE64_ENCODE}"
 *        }
 *    }
 *  </pre></blockquote>
 */
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

	/**
	 *  Encode data
	 *
	 * @return String or Map
	 */
	Serialize() interface{}
}

/**
 *  Transportable Data
 *  <p>
 *      TED - Transportable Encoded Data
 *  </p>
 *
 *  <blockquote><pre>
 *      0. "{BASE64_ENCODE}"
 *      1. "data:image/png;base64,{BASE64_ENCODE}"
 *  </pre></blockquote>
 */
type TransportableData interface {
	Stringer
	TransportableResource

	/*
	   //
	   //  encoding algorithm
	   //
	   String DEFAULT = "base64";
	   String BASE_64 = "base64";
	   String BASE_58 = "base58";
	   String HEX     = "hex";
	*/

	/**
	 *  Get data encoding algorithm
	 *
	 * @return "base64"
	 */
	Encoding() string

	/**
	 *  Get original data
	 *
	 * @return plaintext
	 */
	Bytes() []byte

	/**
	 *  Get length of bytes
	 *
	 * @return len(plaintext)
	 */
	Size() int

	/**
	 *  Get encoded string
	 *
	 * @return "{BASE64_ENCODE}", or
	 *         "data:image/png;base64,{BASE64_ENCODE}"
	 */
	//String() string

	/**
	 *  String()
	 */
	//Serialize() interface{}
}

/**
 *  TED Factory
 */
type TransportableDataFactory interface {

	/**
	 *  Parse string object to TED
	 *
	 * @param ted - TED string
	 * @return TED object
	 */
	ParseTransportableData(ted string) TransportableData
}

//
//  Factory method
//

func ParseTransportableData(ted interface{}) TransportableData {
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
