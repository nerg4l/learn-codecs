package lzw

import (
	"bytes"
	"compress/lzw"
	"io"
	"reflect"
	"testing"
)

func TestEncode(t *testing.T) {
	s := "TOBEORNOTTOBEORTOBEORNOT"
	got := Encode([]byte(s))
	var buf bytes.Buffer
	w := lzw.NewWriter(&buf, lzw.LSB, 8)
	if _, err := io.WriteString(w, s); err != nil {
		t.Fatal(err)
	}
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}
	want := buf.Bytes()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Encode() = %08b,\n                  want %08b", got, want)
	}
}

func TestDecode(t *testing.T) {
	want := "TOBEORNOTTOBEORTOBEORNOT"
	c := Encode([]byte(want))
	got := Decode(c)
	if string(got) != want {
		t.Errorf("Decode() = %s, want %s", got, want)
	}
}
