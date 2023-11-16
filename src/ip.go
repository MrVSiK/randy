package generations

import (
	"errors"
	"strconv"
	"strings"
)

func GetIpv4() (string, error) {
	var ip string;

	for i := 0; i<4; i++ {
		val, err := GetRandomInteger(256);
		if err != nil {
			return "", err;
		}

		ip += strconv.FormatUint(uint64(val), 10);
		if i != 3 {
			ip += ".";
		}
	}

	return ip, nil;
}

func CheckIpString(formatString string) (bool, error) {
	dots := 0;
	colons := 0;

	for i := range formatString {
		if formatString[i] == '.' {
			dots++;
		} else if formatString[i] == ':' {
			colons++;
		}
	}

	if dots > 0 && colons > 0 {
		return false, errors.New("invalid separator, use either `.` or `:`");
	}

	if dots > 0 {
		ip := strings.Split(formatString, ".");
		if len(ip) != 4 {
			return false, errors.New("invalid IPv4: wrong number of octets");
		}

		for i, v := range ip {
			if v != "#" {
				var builder strings.Builder;
				builder.WriteString("invalid octet: ");
				builder.WriteString(strconv.Itoa(i+1));
				builder.WriteString("; use only `#`");
				return false, errors.New(builder.String());
			}
		}
	}

	if colons > 0 {
		ip := strings.Split(formatString, ":");
		if len(ip) != 6 {
			return false, errors.New("invalid IPv6: wrong number of octets");
		}

		for i, v := range ip {
			if v != "#" {
				var builder strings.Builder;
				builder.WriteString("invalid segment: ");
				builder.WriteString(strconv.Itoa(i+1));
				builder.WriteString("; use only `#`");
				return false, errors.New(builder.String());
			}
		}
	}

	return true, nil;
}

func ParseIpString(formatString string) ([]string, error) {
	var ipString []string;

	check, err := CheckIpString(formatString);
	if !check {
		return ipString, err;
	}

	numberOfParts := 0;
	for i := range formatString {
		if formatString[i] == '.' || formatString[i] == ':' {
			numberOfParts++;
		}
	}

	switch(numberOfParts){
		case 3: {
			ipString = strings.Split(formatString, ".");
		}
		case 5: {
			ipString = strings.Split(formatString, ":");
		}
		default: {
			return ipString, errors.New("invalid format string");
		}
	}

	return ipString, nil;
}

func GetIpv4WithCustomOctet(formatString string) (string, error) {
	var ip string;

	parts, err := ParseIpString(formatString);
	if err != nil {
		return ip, err;
	}

	if len(parts) != 4 {
		return ip, errors.New("invalid IPv4");
	}

	for i := range parts {
		if parts[i] == "#" {
			val, err := GetRandomInteger(256);
			if err != nil {
				return ip, err;
			}
			ip += strconv.FormatUint(uint64(val), 10)
		} else {
			ip += parts[i];
		}

		if i != 3 {
			ip += ".";
		}
	}

	return ip, nil;
}