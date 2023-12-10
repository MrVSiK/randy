package generations

import (
	"github.com/MrVSiK/randy/data"
	"strconv"
)

func GetEmail() (string, error) {
	num, err := GetRandomInteger(data.LengthOfEmailProviders);
	if err != nil {
		return "", err;
	}
	provider := data.EmailProviders[num];

	email, err := GetName();
	if err != nil {
		return "", err;
	}

	num, err = GetRandomInteger(2);
	if err != nil {
		return "", err;
	}

	if num == 0 {
		newEmail, err := GetName();
		if err != nil {
			return "", err;
		}
		email += "." + newEmail;
	}

	num, err = GetRandomInteger(2);
	if err != nil {
		return "", err;
	}

	if num == 0 {
		reps, err := GetRandomInteger(6);
		if err != nil {
			return "", err;
		}
		for i := 0; i < int(reps); i++ {
			val, err := GetRandomInteger(10);
			if err != nil {
				return "", err;
			}

			email += strconv.FormatUint(uint64(val), 10);
		}
	}

	return email + "@" + provider + ".com", nil;
}

func GetEmailWithCustomProvider(provider string) (string, error) {
	num, err := GetRandomInteger(data.LengthOfEmailProviders);
	if err != nil {
		return "", err;
	}

	email, err := GetName();
	if err != nil {
		return "", err;
	}

	num, err = GetRandomInteger(2);
	if err != nil {
		return "", err;
	}

	if num == 0 {
		newEmail, err := GetName();
		if err != nil {
			return "", err;
		}
		email += "." + newEmail;
	}

	num, err = GetRandomInteger(2);
	if err != nil {
		return "", err;
	}

	if num == 0 {
		reps, err := GetRandomInteger(6);
		if err != nil {
			return "", err;
		}
		for i := 0; i < int(reps); i++ {
			val, err := GetRandomInteger(10);
			if err != nil {
				return "", err;
			}

			email += strconv.FormatUint(uint64(val), 10);
		}
	}

	return email + "@" + provider, nil;
}