package jsons

import (
	"encoding/json"
	"testing"
)

type jsonInt interface {
	Int | Int8 | Int16 | Int32 | Int64
}

type jsonIntS interface {
	IntS | Int8S | Int16S | Int32S | Int64S
}

type jsonUint interface {
	Uint | Uint8 | Uint16 | Uint32 | Uint64
}

type jsonUintS interface {
	UintS | Uint8S | Uint16S | Uint32S | Uint64S
}

func testUnmarshalIntS[S jsonInt | jsonIntS](t *testing.T) {
	var s S

	if err := json.Unmarshal([]byte("12"), &s); err != nil {
		t.Fatalf("unmarshal %T: %v", s, err)
	} else if s != 12 {
		t.Fatalf("%T: %v", s, s)
	}

	if err := json.Unmarshal([]byte(`"34"`), &s); err != nil {
		t.Fatalf("unmarshal %T: %v", s, err)
	} else if s != 34 {
		t.Fatalf("%T: %v", s, s)
	}

	if err := json.Unmarshal([]byte("-56"), &s); err != nil {
		t.Fatalf("unmarshal %T: %v", s, err)
	} else if s != -56 {
		t.Fatalf("%T: %v", s, s)
	}

	if err := json.Unmarshal([]byte(`"-78"`), &s); err != nil {
		t.Fatalf("unmarshal %T: %v", s, err)
	} else if s != -78 {
		t.Fatalf("%T: %v", s, s)
	}

	s = 0

	if err := json.Unmarshal([]byte(`""`), &s); err != nil {
		t.Fatalf("unmarshal %T: %v", s, err)
	} else if s != 0 {
		t.Fatalf("%T: %v", s, s)
	}

	if err := json.Unmarshal([]byte(`"0"`), &s); err != nil {
		t.Fatalf("unmarshal %T: %v", s, err)
	} else if s != 0 {
		t.Fatalf("%T: %v", s, s)
	}

	if err := json.Unmarshal([]byte(`"-0"`), &s); err != nil {
		t.Fatalf("unmarshal %T: %v", s, err)
	} else if s != 0 {
		t.Fatalf("%T: %v", s, s)
	}
}

func testMarshalInt[I jsonInt, S jsonIntS](t *testing.T) {
	if p, _ := json.Marshal(I(12)); string(p) != `12` {
		t.Fatalf("%T: %s", I(12), p)
	}
	if p, _ := json.Marshal(S(12)); string(p) != `"12"` {
		t.Fatalf("%T: %s", S(12), p)
	}

	if p, _ := json.Marshal(I(-34)); string(p) != `-34` {
		t.Fatalf("%T: %s", I(-34), p)
	}
	if p, _ := json.Marshal(S(-34)); string(p) != `"-34"` {
		t.Fatalf("%T: %s", S(-34), p)
	}
}

func TestInts(t *testing.T) {
	testUnmarshalIntS[Int](t)
	testUnmarshalIntS[Int8](t)
	testUnmarshalIntS[Int16](t)
	testUnmarshalIntS[Int32](t)
	testUnmarshalIntS[Int64](t)

	testUnmarshalIntS[IntS](t)
	testUnmarshalIntS[Int8S](t)
	testUnmarshalIntS[Int16S](t)
	testUnmarshalIntS[Int32S](t)
	testUnmarshalIntS[Int64S](t)

	testMarshalInt[Int, IntS](t)
	testMarshalInt[Int8, Int8S](t)
	testMarshalInt[Int16, Int16S](t)
	testMarshalInt[Int32, Int32S](t)
	testMarshalInt[Int64, Int64S](t)
}

func testUnmarshalUintS[I ints, S jsonUint | jsonUintS](t *testing.T) {
	var s S

	if err := json.Unmarshal([]byte("12"), &s); err != nil {
		t.Fatalf("unmarshal %T: %v", s, err)
	} else if s != 12 {
		t.Fatalf("%T: %v", s, s)
	}

	if err := json.Unmarshal([]byte(`"34"`), &s); err != nil {
		t.Fatalf("unmarshal %T: %v", s, err)
	} else if s != 34 {
		t.Fatalf("%T: %v", s, s)
	}

	if err := json.Unmarshal([]byte("-56"), &s); err != nil {
		t.Fatalf("unmarshal %T: %v", s, err)
	} else if I(s) != -56 {
		t.Fatalf("%T: %v", s, s)
	}

	if err := json.Unmarshal([]byte(`"-78"`), &s); err != nil {
		t.Fatalf("unmarshal %T: %v", s, err)
	} else if I(s) != -78 {
		t.Fatalf("%T: %v", s, s)
	}

	s = 0

	if err := json.Unmarshal([]byte(`""`), &s); err != nil {
		t.Fatalf("unmarshal %T: %v", s, err)
	} else if I(s) != 0 {
		t.Fatalf("%T: %v", s, s)
	}

	if err := json.Unmarshal([]byte(`"0"`), &s); err != nil {
		t.Fatalf("unmarshal %T: %v", s, err)
	} else if I(s) != 0 {
		t.Fatalf("%T: %v", s, s)
	}

	if err := json.Unmarshal([]byte(`"-0"`), &s); err != nil {
		t.Fatalf("unmarshal %T: %v", s, err)
	} else if I(s) != 0 {
		t.Fatalf("%T: %v", s, s)
	}
}

func testMarshalUint[U jsonUint, S jsonUintS](t *testing.T) {
	if p, _ := json.Marshal(U(12)); string(p) != `12` {
		t.Fatalf("%T: %s", U(12), p)
	}
	if p, _ := json.Marshal(S(12)); string(p) != `"12"` {
		t.Fatalf("%T: %s", S(12), p)
	}
}

func TestUints(t *testing.T) {
	testUnmarshalUintS[int, Uint](t)
	testUnmarshalUintS[int8, Uint8](t)
	testUnmarshalUintS[int16, Uint16](t)
	testUnmarshalUintS[int32, Uint32](t)
	testUnmarshalUintS[int64, Uint64](t)

	testUnmarshalUintS[int, UintS](t)
	testUnmarshalUintS[int8, Uint8S](t)
	testUnmarshalUintS[int16, Uint16S](t)
	testUnmarshalUintS[int32, Uint32S](t)
	testUnmarshalUintS[int64, Uint64S](t)

	testMarshalUint[Uint, UintS](t)
	testMarshalUint[Uint8, Uint8S](t)
	testMarshalUint[Uint16, Uint16S](t)
	testMarshalUint[Uint32, Uint32S](t)
	testMarshalUint[Uint64, Uint64S](t)
}

func testUnmarshalFloat[F Float32 | Float32S | Float64 | Float64S](t *testing.T) {
	var f F

	if err := json.Unmarshal([]byte("1.23"), &f); err != nil {
		t.Fatalf("unmarshal %T: %v", f, err)
	} else if f != 1.23 {
		t.Fatalf("%T: %v", f, f)
	}

	if err := json.Unmarshal([]byte(`"4.56"`), &f); err != nil {
		t.Fatalf("unmarshal %T: %v", f, err)
	} else if f != 4.56 {
		t.Fatalf("%T: %v", f, f)
	}

	if err := json.Unmarshal([]byte("-7.89"), &f); err != nil {
		t.Fatalf("unmarshal %T: %v", f, err)
	} else if f != -7.89 {
		t.Fatalf("%T: %v", f, f)
	}

	if err := json.Unmarshal([]byte(`"-1.23"`), &f); err != nil {
		t.Fatalf("unmarshal %T: %v", f, err)
	} else if f != -1.23 {
		t.Fatalf("%T: %v", f, f)
	}

	if err := json.Unmarshal([]byte("1.2345e2"), &f); err != nil {
		t.Fatalf("unmarshal %T: %v", f, err)
	} else if f != 123.45 {
		t.Fatalf("%T: %v", f, f)
	}

	if err := json.Unmarshal([]byte(`"4.56789E3"`), &f); err != nil {
		t.Fatalf("unmarshal %T: %v", f, err)
	} else if f != 4567.89 {
		t.Fatalf("%T: %v", f, f)
	}

	if err := json.Unmarshal([]byte("-1.2345e1"), &f); err != nil {
		t.Fatalf("unmarshal %T: %v", f, err)
	} else if f != -12.345 {
		t.Fatalf("%T: %v", f, f)
	}

	if err := json.Unmarshal([]byte(`"-3.456789E4"`), &f); err != nil {
		t.Fatalf("unmarshal %T: %v", f, err)
	} else if f != -34567.89 {
		t.Fatalf("%T: %v", f, f)
	}

	f = 0

	if err := json.Unmarshal([]byte(`""`), &f); err != nil {
		t.Fatalf("unmarshal %T: %v", f, err)
	} else if f != 0 {
		t.Fatalf("%T: %v", f, f)
	}

	if err := json.Unmarshal([]byte(`"0"`), &f); err != nil {
		t.Fatalf("unmarshal %T: %v", f, err)
	} else if f != 0 {
		t.Fatalf("%T: %v", f, f)
	}

	if err := json.Unmarshal([]byte(`"-0"`), &f); err != nil {
		t.Fatalf("unmarshal %T: %v", f, err)
	} else if f != 0 {
		t.Fatalf("%T: %v", f, f)
	}
}

func testMarshalFloat[F Float32 | Float64, S Float32S | Float64S](t *testing.T) {
	if p, _ := json.Marshal(F(1.23)); string(p) != `1.23` {
		t.Fatalf("%T: %s", F(1.23), p)
	}
	if p, _ := json.Marshal(S(4.56)); string(p) != `"4.56"` {
		t.Fatalf("%T: %s", S(4.56), p)
	}

	if p, _ := json.Marshal(F(1.2345e2)); string(p) != `123.45` {
		t.Fatalf("%T: %s", F(1.23345e2), p)
	}
	if p, _ := json.Marshal(S(4.5678e3)); string(p) != `"4567.8"` {
		t.Fatalf("%T: %s", S(4.5678e3), p)
	}

	if p, _ := json.Marshal(F(-1.23)); string(p) != `-1.23` {
		t.Fatalf("%T: %s", F(-1.23), p)
	}
	if p, _ := json.Marshal(S(-4.56)); string(p) != `"-4.56"` {
		t.Fatalf("%T: %s", S(-4.56), p)
	}

	if p, _ := json.Marshal(F(-1.2345e2)); string(p) != `-123.45` {
		t.Fatalf("%T: %s", F(-1.23345e2), p)
	}
	if p, _ := json.Marshal(S(-4.5678e3)); string(p) != `"-4567.8"` {
		t.Fatalf("%T: %s", S(-4.5678e3), p)
	}
}

func TestFloats(t *testing.T) {
	testUnmarshalFloat[Float32](t)
	testUnmarshalFloat[Float32S](t)
	testUnmarshalFloat[Float64](t)
	testUnmarshalFloat[Float64S](t)

	testMarshalFloat[Float32, Float32S](t)
	testMarshalFloat[Float64, Float64S](t)
}

func testUnmarshalBool[B Bool | BoolS](t *testing.T) {
	var b B

	b = false
	if err := json.Unmarshal([]byte("true"), &b); err != nil {
		t.Fatalf("unmarshal %T: %v", b, err)
	} else if !b {
		t.Fatalf("%T: %v", b, b)
	}

	b = false
	if err := json.Unmarshal([]byte(`"true"`), &b); err != nil {
		t.Fatalf("unmarshal %T: %v", b, err)
	} else if !b {
		t.Fatalf("%T: %v", b, b)
	}

	b = true
	if err := json.Unmarshal([]byte("false"), &b); err != nil {
		t.Fatalf("unmarshal %T: %v", b, err)
	} else if b {
		t.Fatalf("%T: %v", b, b)
	}

	b = true
	if err := json.Unmarshal([]byte(`"false"`), &b); err != nil {
		t.Fatalf("unmarshal %T: %v", b, err)
	} else if b {
		t.Fatalf("%T: %v", b, b)
	}

	b = false
	if err := json.Unmarshal([]byte(`""`), &b); err != nil {
		t.Fatalf("unmarshal %T: %v", b, err)
	} else if b {
		t.Fatalf("%T: %v", b, b)
	}
}

func testMarshalBool(t *testing.T) {
	if p, _ := json.Marshal(Bool(true)); string(p) != `true` {
		t.Fatalf("%T: %s", Bool(true), p)
	}
	if p, _ := json.Marshal(BoolS(true)); string(p) != `"true"` {
		t.Fatalf("%T: %s", BoolS(true), p)
	}

	if p, _ := json.Marshal(Bool(false)); string(p) != `false` {
		t.Fatalf("%T: %s", Bool(false), p)
	}
	if p, _ := json.Marshal(BoolS(false)); string(p) != `"false"` {
		t.Fatalf("%T: %s", BoolS(false), p)
	}
}

func TestBool(t *testing.T) {
	testUnmarshalBool[Bool](t)
	testUnmarshalBool[BoolS](t)
	testMarshalBool(t)
}
