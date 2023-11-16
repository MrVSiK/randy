package generations

import (
	"randy/data"
)

func GetMaleName() (string, error) {
	num, err := GetRandomInteger(data.NumberOfMaleNames);
	return data.MaleNames[num], err;
}

func GetFemaleName() (string, error) {
	num, err := GetRandomInteger(data.NumberOfFemaleNames);
	return data.FemaleNames[num], err;
}

func GetName() (string, error) {
	num, err := GetRandomInteger(2);
	if err != nil {
		return "", err;
	}

	if num == 0 {
		return GetMaleName();
	} else {
		return GetFemaleName();
	}
}