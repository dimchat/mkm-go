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
package protocol

import (
	. "github.com/dimchat/mkm-go/crypto"
	. "github.com/dimchat/mkm-go/format"
	. "github.com/dimchat/mkm-go/types"
)

type MetaType = string

// Meta defines the interface for User/Group entity metadata
//
// Contains core information used to generate and validate entity identities (ID/Address)
//
//	Data structure: {
//	    "type"        : i2s(1),         // Algorithm version (MetaType)
//	    "key"         : "{public key}", // Public key (PK = secp256k1(SK))
//	    "seed"        : "moKy",         // User/group name (seed for fingerprint)
//	    "fingerprint" : "..."           // Signature of seed (CT = sign(seed, SK))
//	}
//
//	 Core algorithm for fingerprint generation:
//	     fingerprint = sign(seed, SK)
type Meta interface {
	Mapper

	// Type returns the meta algorithm version (MetaType)
	//
	// Supported versions:
	//  1 = MKM : username@address (default)
	//  2 = BTC : btc_address
	//  4 = ETH : eth_address
	//  ... (additional algorithm versions)
	Type() MetaType

	// PublicKey returns the public verification key associated with the meta
	//
	// Used for signature verification (supports RSA/ECC algorithms)
	PublicKey() VerifyKey

	// Seed returns the seed string used to generate the fingerprint
	//
	// Typically the username or group name (e.g., "moKy", "Group-X")
	Seed() string

	// Fingerprint returns the cryptographic fingerprint of the seed
	//   - Generation: fingerprint = sign(seed, privateKey)
	//   - Validation: verify(seed, fingerprint, publicKey)
	// Used to verify the integrity of ID and public key
	Fingerprint() TransportableData

	//
	//  Validation
	//

	// IsValid validates the meta information for integrity and correctness
	// Must be called when receiving new meta info from the network to ensure validity
	// Returns: true if meta is valid (all fields match cryptographic verification), false otherwise
	IsValid() bool

	// GenerateAddress creates an Address instance for the specified network type
	//
	// Uses the meta's fingerprint/seed to derive the entity address
	//
	// Parameters:
	//   - network: EntityType specifying the target network (ID type)
	// Returns: Derived Address instance for the entity
	GenerateAddress(network EntityType) Address
}

// MetaFactory defines the factory interface for Meta
type MetaFactory interface {

	// CreateMeta constructs a Meta instance from explicit components
	//
	// Parameters:
	//   - pKey: Public verification key (VerifyKey)
	//   - seed: Seed string (entity name, ID.name)
	//   - fingerprint: Pre-computed fingerprint (signature of seed from sKey.Sign(seed))
	// Returns: Newly constructed Meta instance
	CreateMeta(pKey VerifyKey, seed string, fingerprint TransportableData) Meta

	// GenerateMeta creates a Meta instance by generating the fingerprint from a private key
	//
	// Automatically computes the fingerprint via sKey.Sign(seed) and derives the public key
	//
	// Parameters:
	//   - sKey: Private signing key (SignKey) to generate the fingerprint
	//   - seed: Seed string (entity name, ID.name)
	// Returns: Newly generated Meta instance with computed fingerprint
	GenerateMeta(sKey SignKey, seed string) Meta

	// ParseMeta converts a StringKeyMap into a Meta instance
	//
	// Expects the map to follow the Meta info structure format (type/key/seed/fingerprint)
	//
	// Parameters:
	//   - meta: Meta information in StringKeyMap format
	// Returns: Parsed Meta instance (nil if parsing/validation fails)
	ParseMeta(meta StringKeyMap) Meta
}

//
//  Factory methods
//

func CreateMeta(version MetaType, pKey VerifyKey, seed string, fingerprint TransportableData) Meta {
	helper := GetMetaHelper()
	return helper.CreateMeta(version, pKey, seed, fingerprint)
}

func GenerateMeta(version MetaType, sKey SignKey, seed string) Meta {
	helper := GetMetaHelper()
	return helper.GenerateMeta(version, sKey, seed)
}

func ParseMeta(meta any) Meta {
	helper := GetMetaHelper()
	return helper.ParseMeta(meta)
}

func GetMetaFactory(version MetaType) MetaFactory {
	helper := GetMetaHelper()
	return helper.GetMetaFactory(version)
}

func SetMetaFactory(version MetaType, factory MetaFactory) {
	helper := GetMetaHelper()
	helper.SetMetaFactory(version, factory)
}
