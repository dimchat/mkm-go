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
 *  This interface is defined for authorizing other apps to login,
 *  which can generate a temporary asymmetric key pair for messaging.
 */
type BaseVisa struct {
	BaseDocument
	Visa

	_key EncryptKey
}

func (visa *BaseVisa) Init(dict map[string]interface{}) *BaseVisa {
	if visa.BaseDocument.Init(dict) != nil {
		// lazy load
		visa._key = nil
	}
	return visa
}

func (visa *BaseVisa) InitWithID(identifier ID, data []byte, signature []byte) *BaseVisa {
	if visa.BaseDocument.InitWithType(VISA, identifier, data, signature) != nil {
		// lazy load
		visa._key = nil
	}
	return visa
}

func (visa *BaseVisa) Key() EncryptKey {
	if visa._key == nil {
		info := visa.GetProperty("key")
		pKey := PublicKeyParse(info)
		if pKey != nil {
			vKey, ok := pKey.(EncryptKey)
			if ok {
				visa._key = vKey
			}
		}
	}
	return visa._key
}

func (visa *BaseVisa) SetKey(key EncryptKey) {
	if key == nil {
		visa.SetProperty("key", nil)
	} else {
		info, ok := key.(Map)
		if ok {
			visa.SetProperty("key", info.GetMap(false))
		}
	}
	visa._key = key
}

func (visa *BaseVisa) Avatar() string {
	url := visa.GetProperty("avatar")
	if url == nil {
		return ""
	}
	return url.(string)
}

func (visa *BaseVisa) SetAvatar(url string) {
	visa.SetProperty("avatar", url)
}

/**
 *  Group Document
 *  ~~~~~~~~~~~~~~
 */
type BaseBulletin struct {
	BaseDocument
	Bulletin

	_assistants []ID
}

func (doc *BaseBulletin) Init(dict map[string]interface{}) *BaseBulletin {
	if doc.BaseDocument.Init(dict) != nil {
		// lazy load
		doc._assistants = nil
	}
	return doc
}

func (doc *BaseBulletin) InitWithID(identifier ID, data []byte, signature []byte) *BaseBulletin {
	if doc.BaseDocument.InitWithType(BULLETIN, identifier, data, signature) != nil {
		// lazy load
		doc._assistants = nil
	}
	return doc
}

func (doc *BaseBulletin) Assistants() []ID {
	if doc._assistants == nil {
		assistants := doc.GetProperty("assistants")
		if assistants != nil {
			doc._assistants = IDConvert(assistants.([]interface{}))
		}
	}
	return doc._assistants
}

func (doc *BaseBulletin) SetAssistants(bots []ID) {
	if bots == nil {
		doc.SetProperty("assistants", nil)
	} else {
		doc.SetProperty("assistants", IDRevert(bots))
	}
	doc._assistants = bots
}
