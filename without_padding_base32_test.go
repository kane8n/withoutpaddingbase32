package withoutpaddingbase32

import "testing"

func TestBase32String(t *testing.T) {
	testStrings := []string{"name", "foo", "foob", "fooba", "foobar", ""}
	for _, s := range testStrings {
		encStr := EncodeToBase32String(s)
		decStr := DecodeFromBase32String(encStr)
		t.Logf("test string:%s encode string:%s decode string:%s", s, encStr, decStr)
		if decStr != s {
			t.Errorf("The test character string and the character string after the decoding differ. test string:%s decode string:%s", s, decStr)
		}
	}
}
