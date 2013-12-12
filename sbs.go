// sbs can encode a struct to a byte slice and back
// This is useful if you need to save data in a key value database such as
// a leveldb, because a leveldb only takes bytes as values.
// You can use sbs also as the backend to satify the interfaces
// BinaryMarshaler and BinaryUnmarshaler by simply creating two functions for
// your struct called MarshalBinary() and UnmarshalBinary
//
// Example:
// type MyStruct struct {
//     A, B, C string
// }
//
// func (o *MyStruct) MarshalBinary() (data []byte, err error) {
//     data, err := sbs.Enc(o)
//     return
// }
//
// func (o *MyStruct) UnmarshalBinary(data []byte) error {
//     return sbs.Dec(o, data)
// }
package sbs

// Copyright © 2013 Dennis T Kaplan <http://www.robotamer.com>
// The standard MIT License can be found at robotamer.com
import (
	"bytes"
	"encoding/gob"
)

// Enc takes a populated struct and returns a []byte
func Enc(o interface{}) (b []byte, err error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err = enc.Encode(o)
	if err != nil {
		return
	}
	b = buf.Bytes()
	return
}

// Dec takes an empty struct as well as the []byte returned by Enc(),
// and populates the given empty struct from the []byte.
func Dec(o interface{}, b []byte) (err error) {
	var buf bytes.Buffer
	_, err = buf.Write(b)
	if err != nil {
		return
	}
	dec := gob.NewDecoder(&buf)
	err = dec.Decode(o)
	return
}
