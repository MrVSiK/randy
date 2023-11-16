package generations

import (
	"crypto/rand"
	"encoding/binary"
	"errors"
)

func GetRandomInteger(upperLimit int) (uint, error) {
	if upperLimit < 0 {
		return 0, errors.New("upper limit less than 0");
	}

	randomBytes := make([]byte, 8);

	_, err := rand.Read(randomBytes);

	if err != nil {
		return 0, err;
	}
	return uint(binary.BigEndian.Uint64(randomBytes) % uint64(upperLimit)), nil;
}