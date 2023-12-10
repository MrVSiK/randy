package generations

import (
	"randy/data"
	"strings"
)

func GetRandomColorName() (string, error) {
	val, err := GetRandomInteger(data.NumberOfColors);
	if err != nil {
		return "", err;
	}

	return data.ColorNames[val], nil;
}

func GetRandomColorHex() (string, error) {
	hex, err := GetRandomHex(16777215, 3);
	if err != nil {
		return "", err;
	}

	return strings.ToUpper("#" + hex), nil;
}