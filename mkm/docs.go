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
 *  Document Shadow
 *  ~~~~~~~~~~~~~~~
 */
type BaseDocumentShadow struct {
	IDocumentExt

	_doc IDocumentExt
}

func (shadow *BaseDocumentShadow) Init(doc IDocumentExt) *BaseDocumentShadow {
	shadow._doc = doc
	return shadow
}

func (shadow *BaseDocumentShadow) Document() IDocumentExt {
	return shadow._doc
}
func (shadow *BaseDocumentShadow) getMap() map[string]interface{} {
	return shadow.Document().(Map).GetMap(false)
}

//-------- TAI

func (shadow *BaseDocumentShadow) IsValid() bool {
	return shadow.Document().status() > 0
}

func (shadow *BaseDocumentShadow) Verify(publicKey VerifyKey) bool {
	// NOTICE: if status is 0, it doesn't mean the profile is invalid,
	//         try another key
	return shadow.Document().checkStatus(publicKey) == 1
}

func (shadow *BaseDocumentShadow) Sign(privateKey SignKey) (data, signature []byte) {
	data = JSONEncode(shadow.Document().Properties())
	signature = privateKey.Sign(data)
	dict := shadow.getMap()
	dict["data"] = UTF8Decode(data)
	dict["signature"] = Base64Encode(signature)
	return data, signature
}

func (shadow *BaseDocumentShadow) Properties() map[string]interface{} {
	data := DocumentGetData(shadow.getMap())
	if data == nil {
		// create new properties
		return make(map[string]interface{})
	} else {
		// get properties from data
		return JSONDecode(data)
	}
}

func (shadow *BaseDocumentShadow) GetProperty(name string) interface{} {
	dict := shadow.Document().Properties()
	if dict == nil {
		return nil
	} else {
		return dict[name]
	}
}

func (shadow *BaseDocumentShadow) SetProperty(name string, value interface{}) {
	// update property value with name
	properties := shadow.Document().Properties()
	if value == nil {
		delete(properties, name)
	} else {
		properties[name] = value
	}
	// clear data signature after properties changed
	dict := shadow.getMap()
	delete(dict, "data")
	delete(dict, "signature")
}

//-------- IDocument

func (shadow *BaseDocumentShadow) Type() string {
	doc := shadow.Document()
	docType := doc.GetProperty("type")
	if docType != nil {
		return docType.(string)
	} else {
		return DocumentGetType(shadow.getMap())
	}
}

func (shadow *BaseDocumentShadow) ID() ID {
	return DocumentGetID(shadow.getMap())
}

func (shadow *BaseDocumentShadow) Name() string {
	name := shadow.Document().GetProperty("name")
	if name == nil {
		return ""
	} else {
		return name.(string)
	}
}

func (shadow *BaseDocumentShadow) SetName(name string) {
	shadow.Document().SetProperty("name", name)
}

//-------- IDocumentExt

func (shadow *BaseDocumentShadow) checkStatus(publicKey VerifyKey) int8 {
	dict := shadow.getMap()
	data := DocumentGetData(dict)
	signature := DocumentGetSignature(dict)
	if data == nil {
		// NOTICE: if data is empty, signature should be empty at the same time
		//         this happen while profile not found
		if signature == nil {
			return 0
		} else {
			// data signature error
			return -1
		}
	} else if signature == nil {
		// signature error
		return -1
	} else if publicKey.Verify(data, signature) {
		// signature matched
		return 1
	} else {
		// return old status
		return shadow.Document().status()
	}
}

func (shadow *BaseDocumentShadow) status() int8 {
	panic("Don't call me!")
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
