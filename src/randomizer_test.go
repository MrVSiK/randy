package generations

import (
	"encoding/binary"
	"encoding/hex"
	"github.com/MrVSiK/randy/data"
	"testing"
)


func TestIntegerLimits(t *testing.T){
	var i uint;
	for i = 0; i < data.NumberOfMaleNames; i++ {
		val, err := GetRandomInteger(data.NumberOfMaleNames);
		if err != nil {
			t.Fatalf(err.Error());
		}

		if val >= uint(data.NumberOfMaleNames) {
			t.Fatalf("value exceeds/equals upper limit");
		}
	}
}

func TestHex(t *testing.T){
	buffer := make([]byte, 8);
	var i uint;
	for i = 0; i < data.NumberOfMaleNames; i++ {
		val, err := GetRandomHex(data.NumberOfMaleNames, 2);
		if err != nil {
			t.Fatalf(err.Error());
		}
		binary.BigEndian.PutUint64(buffer, uint64(data.NumberOfMaleNames));
		upperLimit := hex.EncodeToString(buffer)[12:];

		if val >= upperLimit {
			t.Fatalf("value exceeds/equals upper limit");
		}
	}
}