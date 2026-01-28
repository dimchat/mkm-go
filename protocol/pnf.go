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
	"net/url"

	. "github.com/dimchat/mkm-go/crypto"
	. "github.com/dimchat/mkm-go/format"
	. "github.com/dimchat/mkm-go/types"
)

/**
 *  Transportable File
 *  <p>
 *      PNF - Portable Network File
 *  </p>
 *
 *  <blockquote><pre>
 *  0.  "{URL}"
 *  1. {
 *         "data"     : "...",        // base64_encode(fileContent)
 *         "filename" : "avatar.png",
 *
 *         "URL"      : "http://...", // download from CDN
 *         // before fileContent uploaded to a public CDN,
 *         // it can be encrypted by a symmetric key
 *         "key"      : {             // symmetric key to decrypt file data
 *             "algorithm" : "AES",   // "DES", ...
 *             "data"      : "{BASE64_ENCODE}",
 *             ...
 *         }
 *      }
 *  </pre></blockquote>
 */
type TransportableFile interface {
	Mapper
	TransportableResource

	/** When file data is too big, don't set it in this dictionary,
	 *  but upload it to a CDN and set the download URL instead.
	 */
	Data() TransportableData
	SetData(data TransportableData)

	Filename() string
	SetFilename(filename string)

	/** Download URL
	 */
	URL() url.URL
	SetURL(url url.URL)

	/** Password for decrypting the downloaded data from CDN,
	 *  default is a plain key, which just return the same data when decrypting.
	 */
	Password() DecryptKey
	SetPassword(key DecryptKey)

	/** Get encoded string
	 *
	 * @return "URL", or
	 *         "{...}"
	 */
	String() string

	/**
	 *  Encode data and update inner map
	 *
	 * @return inner map
	 */
	//Map() StringKeyMap

	/**
	 *  if contains "URL" and "filename" only,
	 *      String();
	 *  else,
	 *      Map();
	 */
	//Serialize() interface{}
}

/**
 *  PNF Factory
 */
type TransportableFileFactory interface {

	/**
	 *  Create PNF
	 *
	 * @param data
	 *        file content (not encrypted)
	 *
	 * @param filename
	 *        file name
	 *
	 * @param url
	 *        download URL
	 *
	 * @param password
	 *        decrypt key for downloaded data
	 *
	 * @return PNF object
	 */
	CreateTransportableFile(
		data TransportableData, filename string,
		url url.URL, password DecryptKey,
	) TransportableFile

	/**
	 *  Parse string/map object to PNF
	 *
	 * @param pnf
	 *        PNF info
	 *
	 * @return PNF object
	 */
	ParseTransportableFile(pnf StringKeyMap) TransportableFile
}

//
//  Factory methods
//

func CreateTransportableFile(data TransportableData, filename string,
	url url.URL, password DecryptKey) TransportableFile {
	helper := GetTransportableFileHelper()
	return helper.CreateTransportableFile(data, filename, url, password)
}

func ParseTransportableFile(pnf StringKeyMap) TransportableFile {
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
