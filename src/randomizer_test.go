package generations

import (
	"randy/data"
	"testing"
)

func TestNegativeNumberInput(t *testing.T){
	val, err := GetRandomInteger(-10);
	if err == nil {
		t.Fatalf("Upper Limit less than 0");
	}

	if val != 0 {
		t.Fatalf("Value != 0");
	}
}

func TestLimits(t *testing.T){
	for i := 0; i < data.NumberOfMaleNames; i++ {
		val, err := GetRandomInteger(data.NumberOfMaleNames);
		if err != nil {
			t.Fatalf(err.Error());
		}

		if val >= uint(data.NumberOfMaleNames) {
			t.Fatalf("Value exceeds upper limit");
		}

		if val < 0 {
			t.Fatalf("Value falls below lower limit");
		}
	}
}
