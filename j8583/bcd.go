package j8583

import (
	"encoding/hex"
)

func lbcd(data []byte) []byte {
	if len(data) % 2 != 0 {
		return bcd(append(data, "0"...))
	}
	return bcd(data)
}

func rbcd(data []byte) []byte {
	if len(data) % 2 != 0 {
		return bcd(append([]byte("0"), data...))
	}
	return bcd(data)
}

// Encode numeric in ascii into bsd (be sure len(data) % 2 == 0)
func bcd(data []byte) []byte {
	out := make([]byte, len(data) / 2 + 1)
	n, err := hex.Decode(out, data)
	if err != nil {
		panic(err.Error())
	}
	return out[:n]
}

func bcdl2Ascii(data []byte, length int) []byte {
	return bcd2Ascii(data)[:length]
}

func bcdr2Ascii(data []byte, length int) []byte {
	out := bcd2Ascii(data)
	return out[len(out) - length:]
}

func bcd2Ascii(data []byte) []byte {
	out := make([]byte, len(data) * 2)
	n := hex.Encode(out, data)
	return out[:n]
}

func RBCD2Byte(value string) []byte {
	val := []byte(value)
	return rbcd(val)
}

func BCD2Byte(value string) []byte {
	val := []byte(value)
	return lbcd(val)
}



