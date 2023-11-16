package generations

import (
	"testing"
	"randy/data"
)

func TestGetMaleName(t *testing.T){
	if len(data.MaleNames) != data.NumberOfMaleNames {
		t.Fatalf("Actual number of male names does not match with given number of male names");
	}
	for i := 0; i < data.NumberOfMaleNames; i++ {
		_, err := GetMaleName();
		if err != nil {
			t.Fatalf(err.Error());
		}
	}
}

func TestGetFemaleName(t *testing.T){
	if len(data.FemaleNames) != data.NumberOfFemaleNames {
		t.Fatalf("Actual number of female names does not match with given number of female names");
	}
	name, err := GetFemaleName();
	if err != nil {
		t.Fatalf(err.Error());
	}
	bool := true;

	for i := 0; i < data.NumberOfFemaleNames; i++ {
		if(data.FemaleNames[i] == name){
			bool = false;
		}
	}

	if bool {
		t.Fail();
	}
}