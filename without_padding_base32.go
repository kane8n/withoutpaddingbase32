package withoutpaddingbase32

import (
	"math"
	"strings"
)

const (
	InByteSize     = 8
	OutByteSize    = 5
	Base32Alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"
)

func EncodeToBase32String(src string) string {
	if src == "" {
		return ""
	}
	bytes := []byte(src)

	sbuf := make([]string, len(bytes)*InByteSize/OutByteSize)
	var (
		bytesPosition            uint
		bytesSubPosition         uint
		outputBase32Byte         uint
		outputBase32BytePosition uint
	)
	bytesPosition = 0
	bytesSubPosition = 0
	outputBase32Byte = 0
	outputBase32BytePosition = 0
	i := 0
	for bytesPosition < uint(len(bytes)) {
		i++
		bitsAvailableInByte := uint(math.Min(float64(InByteSize-bytesSubPosition), float64(OutByteSize-outputBase32BytePosition)))
		outputBase32Byte = outputBase32Byte << bitsAvailableInByte
		outputBase32Byte |= uint(bytes[bytesPosition] >> (InByteSize - (bytesSubPosition + bitsAvailableInByte)))
		bytesSubPosition += bitsAvailableInByte
		if bytesSubPosition >= InByteSize {
			bytesPosition++
			bytesSubPosition = 0
		}
		outputBase32BytePosition += bitsAvailableInByte
		if outputBase32BytePosition >= OutByteSize {
			outputBase32Byte &= 0x1F // 0x1F = 00011111 in binary
			sbuf = append(sbuf, string(Base32Alphabet[outputBase32Byte]))
			outputBase32BytePosition = 0
		}
	}

	if outputBase32BytePosition > 0 {
		outputBase32Byte = outputBase32Byte << (OutByteSize - outputBase32BytePosition)
		outputBase32Byte &= 0x1F // 0x1F = 00011111 in binary
		sbuf = append(sbuf, string(Base32Alphabet[outputBase32Byte]))
	}
	return strings.Join(sbuf, "")
}

func DecodeFromBase32String(base32Str string) string {
	if base32Str == "" {
		return ""
	}

	base32StringUpperCase := strings.ToUpper(base32Str)

	outputBytes := make([]byte, len(base32StringUpperCase)*OutByteSize/InByteSize)

	if len(outputBytes) == 0 {
		return ""
	}

	var (
		base32Position        uint
		base32SubPosition     uint
		outputBytePosition    uint
		outputByteSubPosition uint
	)
	base32Position = 0
	base32SubPosition = 0
	outputBytePosition = 0
	outputByteSubPosition = 0

	for outputBytePosition < uint(len(outputBytes)) {
		currentBase32Byte := strings.Index(Base32Alphabet, string(base32StringUpperCase[base32Position]))
		if currentBase32Byte < 0 {
			return ""
		}
		bitsAvailableInByte := uint(math.Min(float64(OutByteSize-base32SubPosition), float64(InByteSize-outputByteSubPosition)))
		outputBytes[outputBytePosition] = outputBytes[outputBytePosition] << bitsAvailableInByte
		outputBytes[outputBytePosition] |= byte(currentBase32Byte >> (OutByteSize - (base32SubPosition + bitsAvailableInByte)))
		outputByteSubPosition += bitsAvailableInByte
		if outputByteSubPosition >= InByteSize {
			outputBytePosition++
			outputByteSubPosition = 0
		}
		base32SubPosition += bitsAvailableInByte
		if base32SubPosition >= OutByteSize {
			base32Position++
			base32SubPosition = 0
		}
	}
	return string(outputBytes)
}
