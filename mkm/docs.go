/* license: https://mit-license.org
 *
 *  Ming-Ke-Ming : Decentralized User Identity Authentication
 *
 *                                Written in 2020 by Moky <albert.moky@gmail.com>
 *
 * ==============================================================================
 * The MIT License (MIT)
 *
 * Copyright (c) 2020 Albert Moky
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
	. "github.com/dimchat/mkm-go/protocol"
	. "github.com/dimchat/mkm-go/types"
)

/**
 *  User Document
 *  ~~~~~~~~~~~~~
 *
 *  This interface is defined for authorizing other apps to login,
 *  which can generate a temporary asymmetric key pair for messaging.
 */
type BaseVisa struct {
	BaseDocument

	_key EncryptKey
}

func (doc *BaseVisa) Init(dict map[string]interface{}) Visa {
	if doc.BaseDocument.Init(dict) != nil {
		// lazy load
		doc._key = nil
	}
	return doc
}

func (doc *BaseVisa) InitWithData(identifier ID, data string, signature []byte) Visa {
	if doc.BaseDocument.InitWithData(identifier, data, signature) != nil {
		// lazy load
		doc._key = nil
	}
	return doc
}

func (doc *BaseVisa) InitWithType(identifier ID, docType string) Visa {
	if doc.BaseDocument.InitWithType(identifier, docType) != nil {
		// lazy load
		doc._key = nil
	}
	return doc
}

//-------- IVisa

func (doc *BaseVisa) Key() EncryptKey {
	if doc._key == nil {
		info := doc.GetProperty("key")
		pKey := PublicKeyParse(info)
		key, ok := pKey.(EncryptKey)
		if ok {
			doc._key = key
		}
	}
	return doc._key
}

func (doc *BaseVisa) SetKey(key EncryptKey) {
	info, ok := key.(Map)
	if ok {
		doc.SetProperty("key", info.GetMap(false))
	} else {
		doc.SetProperty("key", key)
	}
	doc._key = key
}

func (doc *BaseVisa) Avatar() string {
	url, ok := doc.GetProperty("avatar").(string)
	if ok {
		return url
	} else {
		return ""
	}
}

func (doc *BaseVisa) SetAvatar(url string) {
	doc.SetProperty("avatar", url)
}

/**
 *  Group Document
 *  ~~~~~~~~~~~~~~
 */
type BaseBulletin struct {
	BaseDocument

	_assistants []ID
}

func (doc *BaseBulletin) Init(dict map[string]interface{}) Bulletin {
	if doc.BaseDocument.Init(dict) != nil {
		// lazy load
		doc._assistants = nil
	}
	return doc
}

func (doc *BaseBulletin) InitWithData(identifier ID, data string, signature []byte) Bulletin {
	if doc.BaseDocument.InitWithData(identifier, data, signature) != nil {
		// lazy load
		doc._assistants = nil
	}
	return doc
}

func (doc *BaseBulletin) InitWithType(identifier ID, docType string) Bulletin {
	if doc.BaseDocument.InitWithType(identifier, docType) != nil {
		// lazy load
		doc._assistants = nil
	}
	return doc
}

//-------- IBulletin

func (doc *BaseBulletin) Assistants() []ID {
	if doc._assistants == nil {
		assistants := doc.GetProperty("assistants")
		if assistants != nil {
			doc._assistants = IDConvert(assistants)
		}
	}
	return doc._assistants
}

func (doc *BaseBulletin) SetAssistants(bots []ID) {
	if ValueIsNil(bots) {
		doc.SetProperty("assistants", nil)
	} else {
		doc.SetProperty("assistants", IDRevert(bots))
	}
	doc._assistants = bots
}
