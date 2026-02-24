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

import . "github.com/dimchat/mkm-go/types"

// ID defines the interface for entity identifiers (User/Group/Bot/...)
//
// Data format: "name@address[/terminal]"
//
// Field definitions:
//   - name: Entity name (seed for fingerprint used to build the address)
//   - address: Unique string identifier for the entity
//   - terminal: Entity login resource/device (OPTIONAL field)
type ID interface {
	Stringer

	Name() string
	Address() Address
	Terminal() string

	// Type returns the entity type of the ID (same as Address.Network())
	//
	// Returns: EntityType (User, Group, Bot, ...)
	Type() EntityType

	IsUser() bool
	IsGroup() bool
	IsBroadcast() bool
}

// IDFactory defines the factory interface for ID
type IDFactory interface {

	// GenerateID creates a new ID instance using meta info, network type and terminal
	//
	// Parameters:
	//   - meta: Meta information used to generate the underlying address
	//   - network: EntityType specifying the ID type (User/Group/Bot/...)
	//   - terminal: Optional terminal identifier (empty string for no terminal)
	// Returns: Newly generated ID instance
	GenerateID(meta Meta, network EntityType, terminal string) ID

	// CreateID constructs an ID instance from explicit name, address and terminal components
	//
	// Parameters:
	//   - name: Entity name component
	//   - address: Pre-constructed Address instance
	//   - terminal: Optional terminal identifier (empty string for no terminal)
	// Returns: Newly created ID instance
	CreateID(name string, address Address, terminal string) ID

	// ParseID converts an ID string (in "name@address[/terminal]" format) into an ID instance
	//
	// Parameters:
	//   - did: String representation of the ID to parse
	// Returns: Parsed ID instance (nil if parsing fails)
	ParseID(did string) ID
}

//
//  Factory methods
//

func GenerateID(meta Meta, network EntityType, terminal string) ID {
	helper := GetIDHelper()
	return helper.GenerateID(meta, network, terminal)
}

func CreateID(name string, address Address, terminal string) ID {
	helper := GetIDHelper()
	return helper.CreateID(name, address, terminal)
}

func ParseID(did any) ID {
	helper := GetIDHelper()
	return helper.ParseID(did)
}

func GetIDFactory() IDFactory {
	helper := GetIDHelper()
	return helper.GetIDFactory()
}

func SetIDFactory(factory IDFactory) {
	helper := GetIDHelper()
	helper.SetIDFactory(factory)
}

//
//  Conveniences
//

// IDConvert converts a generic array (any type) into a slice of ID instances
//
// # Extracts string elements from the input array, parses each to an ID, and filters out nil values
//
// Parameters:
//   - array: Input array (typically []string or wrapped list) containing ID strings
//
// Returns: Slice of valid ID instances (empty slice if no valid IDs)
func IDConvert(array any) []ID {
	members := FetchList(array)
	identifiers := make([]ID, 0, len(members))
	var did ID
	for _, item := range members {
		did = ParseID(item)
		if did == nil {
			continue
		}
		identifiers = append(identifiers, did)
	}
	return identifiers
}

// IDRevert converts a slice of ID instances back to a slice of their string representations
//
// # Uses the String() method (implements Stringer) for each ID to get its string format
//
// Parameters:
//   - identifiers: Slice of ID instances to convert
//
// Returns: Slice of ID strings in "name@address[/terminal]" format
func IDRevert(identifiers []ID) []string {
	array := make([]string, len(identifiers))
	for idx, did := range identifiers {
		array[idx] = did.String()
	}
	return array
}
