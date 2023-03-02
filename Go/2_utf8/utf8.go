package main

import (
	"fmt"
)

func encode(utf32 []rune) []byte {
	var result []byte
	for _, r := range utf32 {
		if r < 128 {
			result = append(result, byte(r))
			continue
		}

		if r < 2048 {
			result = append(result, byte(r>>6)|192)
			result = append(result, byte(r&63)|128)
			continue
		}

		if r < 65536 {
			result = append(result, byte(r>>12)|224)
			result = append(result, byte(r>>6&63)|128)
			result = append(result, byte(r&63)|128)
			continue
		}

		result = append(result, byte(r>>18)|240)
		result = append(result, byte(r>>12&63)|128)
		result = append(result, byte(r>>6&63)|128)
		result = append(result, byte(r&63)|128)
	}
	return result
}

func decode(utf8 []byte) []rune {
	var result []rune
	var r, c int
	for i := 0; i < len(utf8); i++ {
		c = int(utf8[i])
		if c < 128 {
			result = append(result, rune(c))
			continue
		}

		if c < 224 {
			r = c & 31
			r = r << 6
			i++
			c = int(utf8[i])
			r = r | (c & 63)
			result = append(result, rune(r))
			continue
		}

		if c < 240 {
			r = c & 15
			r = r << 12
			c = int(utf8[i])
			r = r | (c&63)<<6
			i++
			c = int(utf8[i])
			r = r | (c & 63)
			result = append(result, rune(r))
			continue
		}

		r = c & 7
		r = r << 18
		i++
		c = int(utf8[i])
		r = r | (c&63)<<12
		i++
		c = int(utf8[i])
		r = r | (c&63)<<6
		i++
		c = int(utf8[i])
		r = r | (c & 63)
		result = append(result, rune(r))
	}
	return result
}

func main() {
    // Test encoding
    utf32 := []rune{'H', 'e', 'l', 'l', 'o', ',', ' ', 'w', 'o', 'r', 'l', 'd', '!'}
    utf8 := encode(utf32)
    fmt.Printf("Encoded string: %s\n", string(utf8))
    fmt.Printf("UTF-8: %v\n", utf8)

    // Test decoding
    decoded := decode(utf8)
    fmt.Printf("Decoded string: %s\n", string(decoded))
    fmt.Printf("UTF-32: %v\n", decoded)
}
