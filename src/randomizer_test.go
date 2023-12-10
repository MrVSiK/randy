package generations

import (
	"encoding/binary"
	"encoding/hex"
	"randy/data"
	"testing"
)

func TestIntegerNegativeNumberInput(t *testing.T){
	val, err := GetRandomInteger(-10);
	if err == nil {
		t.Fatalf("Upper Limit less than 0");
	}

	if val != 0 {
		t.Fatalf("Value != 0");
	}
}

func TestIntegerLimits(t *testing.T){
	for i := 0; i < data.NumberOfMaleNames; i++ {
		val, err := GetRandomInteger(data.NumberOfMaleNames);
		if err != nil {
			t.Fatalf(err.Error());
		}

		if val >= uint(data.NumberOfMaleNames) {
			t.Fatalf("Value exceeds/equals upper limit");
		}
	}
}

func TestHex(t *testing.T){
	buffer := make([]byte, 8);
	for i := 0; i < data.NumberOfMaleNames; i++ {
		val, err := GetRandomHex(data.NumberOfMaleNames);
		if err != nil {
			t.Fatalf(err.Error());
		}
		binary.BigEndian.PutUint64(buffer, uint64(data.NumberOfMaleNames));
		upperLimit := hex.EncodeToString(buffer)[12:];

		if val >= upperLimit {
			t.Fatalf("Value exceeds/equals upper limit");
		}
	}
}