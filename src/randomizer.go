package generations

import (
	"crypto/rand"
	"encoding/binary"
)

func GetRandomInteger(upperLimit int) int {
	randomBytes := make([]byte, 8);

	_, err := rand.Read(randomBytes);

	if err != nil {
		LogToConsole("Error generating random number");
		LogToConsole(err.Error());
		panic(nil);
	}

	return int(binary.BigEndian.Uint64(randomBytes)) % upperLimit;
}