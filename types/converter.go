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
package types

// Data Converter
type Converter interface {

	GetString  (value any, defaultValue string) string

	GetBool    (value any, defaultValue bool) bool

	GetInt     (value any, defaultValue int) int
	GetInt8    (value any, defaultValue int8) int8
	GetInt16   (value any, defaultValue int16) int16
	GetInt32   (value any, defaultValue int32) int32
	GetInt64   (value any, defaultValue int64) int64

	GetUInt    (value any, defaultValue uint) uint
	GetUInt8   (value any, defaultValue uint8) uint8
	GetUInt16  (value any, defaultValue uint16) uint16
	GetUInt32  (value any, defaultValue uint32) uint32
	GetUInt64  (value any, defaultValue uint64) uint64

	GetFloat32 (value any, defaultValue float32) float32
	GetFloat64 (value any, defaultValue float64) float64

	GetTime    (value any, defaultValue Time) Time

}

var sharedConverter Converter = &DataConverter{}

func SetConverter(converter Converter) {
	sharedConverter = converter
}

//
//  Interfaces
//

func ConvertString(value any, defaultValue string) string {
	return sharedConverter.GetString(value, defaultValue)
}

func ConvertBool(value any, defaultValue bool) bool {
	return sharedConverter.GetBool(value, defaultValue)
}

func ConvertInt(value any, defaultValue int) int {
	return sharedConverter.GetInt(value, defaultValue)
}
func ConvertInt8(value any, defaultValue int8) int8 {
	return sharedConverter.GetInt8(value, defaultValue)
}
func ConvertInt16(value any, defaultValue int16) int16 {
	return sharedConverter.GetInt16(value, defaultValue)
}
func ConvertInt32(value any, defaultValue int32) int32 {
	return sharedConverter.GetInt32(value, defaultValue)
}
func ConvertInt64(value any, defaultValue int64) int64 {
	return sharedConverter.GetInt64(value, defaultValue)
}

func ConvertUInt(value any, defaultValue uint) uint {
	return sharedConverter.GetUInt(value, defaultValue)
}
func ConvertUInt8(value any, defaultValue uint8) uint8 {
	return sharedConverter.GetUInt8(value, defaultValue)
}
func ConvertUInt16(value any, defaultValue uint16) uint16 {
	return sharedConverter.GetUInt16(value, defaultValue)
}
func ConvertUInt32(value any, defaultValue uint32) uint32 {
	return sharedConverter.GetUInt32(value, defaultValue)
}
func ConvertUInt64(value any, defaultValue uint64) uint64 {
	return sharedConverter.GetUInt64(value, defaultValue)
}

func ConvertFloat32(value any, defaultValue float32) float32 {
	return sharedConverter.GetFloat32(value, defaultValue)
}
func ConvertFloat64(value any, defaultValue float64) float64 {
	return sharedConverter.GetFloat64(value, defaultValue)
}

func ConvertTime(value any, defaultValue Time) Time {
	return sharedConverter.GetTime(value, defaultValue)
}
