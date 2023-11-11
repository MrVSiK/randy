package generations

import (
	"randy/data"
)

func GetMaleName() string {
	num := GetRandomInteger(data.NumberOfMaleNames);
	return data.MaleNames[num];
}

func GetFemaleName() string {
	num := GetRandomInteger(data.NumberOfFemaleNames);
	return data.FemaleNames[num];
}

func GetName() string {
	num := GetRandomInteger(2);

	if num == 0 {
		return GetMaleName();
	} else {
		return GetFemaleName();
	}
}