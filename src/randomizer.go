package generations

import (
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"errors"
)

func GetRandomInteger(upperLimit uint) (uint, error) {
	if upperLimit < 0 {
		return 0, errors.New("upper limit cannot be less than 0");
	}

	randomBytes := make([]byte, 8);

	_, err := rand.Read(randomBytes);

	if err != nil {
		return 0, err;
	}
	return uint(binary.BigEndian.Uint64(randomBytes) % uint64(upperLimit)), nil;
}

func GetRandomHex(upperLimmit, numberOfBytes uint) (string, error){
	if upperLimmit < 0 {
		return "", errors.New("upper limit cannot be less than 0")
	}

	val, err := GetRandomInteger(upperLimmit);
	if err != nil {
		return "", err;
	}
	buffer := make([]byte, 8);

	binary.BigEndian.PutUint64(buffer, uint64(val));

	return hex.EncodeToString(buffer)[16 - numberOfBytes*2:], nil;
}