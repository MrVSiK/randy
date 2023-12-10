package generations

import (
	"testing"
	"randy/data"
)

func TestGetMaleName(t *testing.T){
	if uint(len(data.MaleNames)) != data.NumberOfMaleNames {
		t.Fatalf("number of male names does not match length of array");
	}

	var i uint;
	for i = 0; i < data.NumberOfMaleNames; i++ {
		_, err := GetMaleName();
		if err != nil {
			t.Fatalf(err.Error());
		}
	}
}

func TestGetFemaleName(t *testing.T){
	if uint(len(data.FemaleNames)) != data.NumberOfFemaleNames {
		t.Fatalf("number of female names does not match length of array");
	}
	name, err := GetFemaleName();
	if err != nil {
		t.Fatalf(err.Error());
	}
	bool := true;

	var i uint;
	for i = 0; i < data.NumberOfFemaleNames; i++ {
		if(data.FemaleNames[i] == name){
			bool = false;
		}
	}

	if bool {
		t.Fail();
	}
}