package lzw

import "bytes"

func Encode(s []byte) []byte {
	dict := make(map[string]int, 256)
	for i := 0; i < 256; i++ {
		dict[string([]byte{byte(i)})] = i
	}
	var p []byte
	var buf bytes.Buffer
	for i := 0; i < len(s); i++ {
		c := s[i]
		pc := string(append(p, c))
		if _, ok := dict[pc]; ok {
			p = append(p, c)
		} else {
			buf.Write([]byte{byte(dict[string(p)])})
			dict[pc] = len(dict)
			p = []byte{c}
		}
	}
	if len(p) > 0 {
		buf.Write([]byte{byte(dict[string(p)])})
	}
	return buf.Bytes()
}

func Decode(c []byte) []byte {
	var buf bytes.Buffer
	return buf.Bytes()
}
