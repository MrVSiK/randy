package generations

import (
	"randy/data"
	"strconv"
)

func GetEmail() string {
	num := GetRandomInteger(data.LengthOfEmailProviders);
	provider := data.EmailProviders[num];

	var email = GetName();

	num = GetRandomInteger(2);

	if num == 0 {
		email += "." + GetName();
	}

	num = GetRandomInteger(2);

	if num == 0 {
		reps := GetRandomInteger(6);
		for i := 0; i < reps; i++ {
			email += strconv.Itoa(GetRandomInteger(10));
		}
	}

	return email + "@" + provider + ".com";
}