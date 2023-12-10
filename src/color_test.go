package generations

import (
	"randy/data"
	"strings"
	"testing"
	"unicode"
)

func TestGetColorName(t *testing.T) {
	if uint(len(data.ColorNames)) != data.NumberOfColors {
		t.Fatalf("number of colors does not match length of array")
	}

	var i uint;
	for i = 0; i < data.NumberOfColors; i++ {
		_, err := GetRandomColorName();
		if err != nil {
			t.Fatalf(err.Error());
		}
	}
}

func TestGetColorHex(t *testing.T) {
	color, err := GetRandomColorHex();
	if err != nil {
		t.Fatalf(err.Error());
	}

	if !strings.HasPrefix(color, "#") {
		t.Fatalf("does not begin with `#`");
	}

	_, aft, _ := strings.Cut(color, "#");

	if len(aft) != 6 {
		t.Fatalf("invalid hex");
	}

	for _, char := range aft {
		if !unicode.IsDigit(char) && !unicode.IsUpper(char) {
			t.Fatalf("hex not in uppercase");
		}
	}
}