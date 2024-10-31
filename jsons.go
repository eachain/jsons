package jsons

import (
	"encoding/json"
	"strings"
)

type (
	Int     int
	IntS    int
	Int8    int8
	Int8S   int8
	Int16   int16
	Int16S  int16
	Int32   int32
	Int32S  int32
	Int64   int64
	Int64S  int64
	Uint    uint
	UintS   uint
	Uint8   uint8
	Uint8S  uint8
	Uint16  uint16
	Uint16S uint16
	Uint32  uint32
	Uint32S uint32
	Uint64  uint64
	Uint64S uint64
)

type (
	Float32  float32
	Float32S float32
	Float64  float64
	Float64S float64
)

type (
	Bool  bool
	BoolS bool
)

func (i *Int) UnmarshalJSON(p []byte) error {
	return unmarshal(p, (*int)(i))
}

func (i Int) MarshalJSON() ([]byte, error) {
	return json.Marshal(int(i))
}

func (i *IntS) UnmarshalJSON(p []byte) error {
	return unmarshal(p, (*int)(i))
}

func (i IntS) MarshalJSON() ([]byte, error) {
	return marshal(int(i))
}

func (i *Int8) UnmarshalJSON(p []byte) error {
	return unmarshal(p, (*int8)(i))
}

func (i Int8) MarshalJSON() ([]byte, error) {
	return json.Marshal(int8(i))
}

func (i *Int8S) UnmarshalJSON(p []byte) error {
	return unmarshal(p, (*int8)(i))
}

func (i Int8S) MarshalJSON() ([]byte, error) {
	return marshal(int8(i))
}

func (i *Int16) UnmarshalJSON(p []byte) error {
	return unmarshal(p, (*int16)(i))
}

func (i Int16) MarshalJSON() ([]byte, error) {
	return json.Marshal(int16(i))
}

func (i *Int16S) UnmarshalJSON(p []byte) error {
	return unmarshal(p, (*int16)(i))
}

func (i Int16S) MarshalJSON() ([]byte, error) {
	return marshal(int16(i))
}

func (i *Int32) UnmarshalJSON(p []byte) error {
	return unmarshal(p, (*int32)(i))
}

func (i Int32) MarshalJSON() ([]byte, error) {
	return json.Marshal(int32(i))
}

func (i *Int32S) UnmarshalJSON(p []byte) error {
	return unmarshal(p, (*int32)(i))
}

func (i Int32S) MarshalJSON() ([]byte, error) {
	return marshal(int32(i))
}

func (i *Int64) UnmarshalJSON(p []byte) error {
	return unmarshal(p, (*int64)(i))
}

func (i Int64) MarshalJSON() ([]byte, error) {
	return json.Marshal(int64(i))
}

func (i *Int64S) UnmarshalJSON(p []byte) error {
	return unmarshal(p, (*int64)(i))
}

func (i Int64S) MarshalJSON() ([]byte, error) {
	return marshal(int64(i))
}

func (u *Uint) UnmarshalJSON(p []byte) error {
	return unmarshalUint[int](p, (*uint)(u))
}

func (u Uint) MarshalJSON() ([]byte, error) {
	return json.Marshal(uint(u))
}

func (u *UintS) UnmarshalJSON(p []byte) error {
	return unmarshalUint[int](p, (*uint)(u))
}

func (u UintS) MarshalJSON() ([]byte, error) {
	return marshal(uint(u))
}

func (u *Uint8) UnmarshalJSON(p []byte) error {
	return unmarshalUint[int8](p, (*uint8)(u))
}

func (u Uint8) MarshalJSON() ([]byte, error) {
	return json.Marshal(uint8(u))
}

func (u *Uint8S) UnmarshalJSON(p []byte) error {
	return unmarshalUint[int8](p, (*uint8)(u))
}

func (u Uint8S) MarshalJSON() ([]byte, error) {
	return marshal(uint8(u))
}

func (u *Uint16) UnmarshalJSON(p []byte) error {
	return unmarshalUint[int16](p, (*uint16)(u))
}

func (u Uint16) MarshalJSON() ([]byte, error) {
	return json.Marshal(uint16(u))
}

func (u *Uint16S) UnmarshalJSON(p []byte) error {
	return unmarshalUint[int16](p, (*uint16)(u))
}

func (u Uint16S) MarshalJSON() ([]byte, error) {
	return marshal(uint16(u))
}

func (u *Uint32) UnmarshalJSON(p []byte) error {
	return unmarshalUint[int32](p, (*uint32)(u))
}

func (u Uint32) MarshalJSON() ([]byte, error) {
	return json.Marshal(uint32(u))
}

func (u *Uint32S) UnmarshalJSON(p []byte) error {
	return unmarshalUint[int32](p, (*uint32)(u))
}

func (u Uint32S) MarshalJSON() ([]byte, error) {
	return marshal(uint32(u))
}

func (u *Uint64) UnmarshalJSON(p []byte) error {
	return unmarshalUint[int64](p, (*uint64)(u))
}

func (u Uint64) MarshalJSON() ([]byte, error) {
	return json.Marshal(uint64(u))
}

func (u *Uint64S) UnmarshalJSON(p []byte) error {
	return unmarshalUint[int64](p, (*uint64)(u))
}

func (u Uint64S) MarshalJSON() ([]byte, error) {
	return marshal(uint64(u))
}

func (f *Float32) UnmarshalJSON(p []byte) error {
	return unmarshal(p, (*float32)(f))
}

func (f Float32) MarshalJSON() ([]byte, error) {
	return json.Marshal(float32(f))
}

func (f *Float32S) UnmarshalJSON(p []byte) error {
	return unmarshal(p, (*float32)(f))
}

func (f Float32S) MarshalJSON() ([]byte, error) {
	return marshal(float32(f))
}

func (f *Float64) UnmarshalJSON(p []byte) error {
	return unmarshal(p, (*float64)(f))
}

func (f Float64) MarshalJSON() ([]byte, error) {
	return json.Marshal(float64(f))
}

func (f *Float64S) UnmarshalJSON(p []byte) error {
	return unmarshal(p, (*float64)(f))
}

func (f Float64S) MarshalJSON() ([]byte, error) {
	return marshal(float64(f))
}

func (b *Bool) UnmarshalJSON(p []byte) error {
	return unmarshal(p, (*bool)(b))
}

func (b Bool) MarshalJSON() ([]byte, error) {
	return json.Marshal(bool(b))
}

func (b *BoolS) UnmarshalJSON(p []byte) error {
	return unmarshal(p, (*bool)(b))
}

func (b BoolS) MarshalJSON() ([]byte, error) {
	return marshal(bool(b))
}

func unmarshal(p []byte, u any) error {
	if len(p) == 0 {
		return nil
	}
	if p[0] == '"' {
		if len(p) == 2 && p[1] == '"' {
			return nil // empty string
		}

		var s string
		err := json.Unmarshal(p, &s)
		if err != nil {
			return err
		}
		s = strings.TrimLeft(s, "0")
		if s == "" {
			s = "0"
		}
		p = []byte(s)
	}
	return json.Unmarshal(p, u)
}

type uints interface {
	uint | uint8 | uint16 | uint32 | uint64
}

type ints interface {
	int | int8 | int16 | int32 | int64
}

func unmarshalUint[I ints, U uints](p []byte, u *U) error {
	if len(p) == 0 {
		return nil
	}
	isNeg := p[0] == '-' || (len(p) > 1 && p[0] == '"' && p[1] == '-')
	if isNeg {
		var i I
		err := unmarshal(p, &i)
		if err != nil {
			return err
		}
		*u = U(i)
		return nil
	}
	return unmarshal(p, u)
}

func marshal(u any) ([]byte, error) {
	p, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}
	if len(p) == 0 {
		return nil, nil
	}
	if p[0] == '"' {
		return p, nil
	}
	return json.Marshal(string(p))
}
