package generations

import (
	"strings"
	"testing"
)

func TestGetIpv4(t *testing.T){
	ip, err := GetIpv4();
	if err != nil {
		t.Fatalf(err.Error());
	}

	if !strings.Contains(ip, ".") {
		t.Fatalf("delimiter not found")
	}

	if len(strings.Split(ip, ".")) != 4 {
		t.Fatalf("4 octets not created");
	}

}

func TestCheckIpStringForIPv4(t *testing.T){
	check, err := CheckIpString("255.#.#.#");
	if err != nil {
		t.Fatalf(err.Error());
	}

	if !check {
		t.Fatalf("invalid ip string identified as valid");
	}

	check, err = CheckIpString("###.##.#.#");
	if err == nil || check {
		t.Fatalf("multiple # in octet")
	}

	check, err = CheckIpString("#.#.#");
	if err == nil || check {
		t.Fatalf("invalid number of octets")
	}

	check, err = CheckIpString("#.#.#:#");
	if err == nil || check {
		t.Fatalf("invalid delimiter use")
	}
}

func TestCustomOctetIpV4(t *testing.T) {
	ip, err := GetIpv4WithCustomOctet("255.#.53.#");
	if err != nil {
		t.Fatalf(err.Error());
	}

	ipSplit := strings.Split(ip, ".");

	if ipSplit[0] != "255" || ipSplit[2] != "53" {
		t.Fatalf("format string ignored");
	}
}

func TestGetIpv6(t *testing.T) {
	ip, err := GetIpv6();
	if err != nil {
		t.Fatalf(err.Error());
	}

	if !strings.Contains(ip, ":") {
		t.Fatalf("delimiter not found")
	}

	if len(strings.Split(ip, ":")) != 8 {
		t.Fatalf("8 segments not created");
	}

	println(ip);
}

func TestCheckIpStringForIPv6(t *testing.T){
	check, err := CheckIpString("2001:#:#:#:#:#:#:#");
	if err != nil {
		t.Fatalf(err.Error());
	}

	if !check {
		t.Fatalf("invalid ip string identified as valid");
	}

	check, err = CheckIpString("###:##:#:#:#:#:#:#");
	if err == nil || check {
		t.Fatalf("multiple # in segment")
	}

	check, err = CheckIpString("#:#:#:#");
	if err == nil || check {
		t.Fatalf("invalid number of segments")
	}

	check, err = CheckIpString("#:#.#:#");
	if err == nil || check {
		t.Fatalf("invalid delimiter use")
	}
}

func TestCustomOctetIpV6(t *testing.T) {
	ip, err := GetIpv6WithCustomSegment("2001:0db8:#:#:#:ff00:#:#");
	if err != nil {
		t.Fatalf(err.Error());
	}

	ipSplit := strings.Split(ip, ".");

	if ipSplit[0] != "2001" || ipSplit[1] != "0db8" || ipSplit[5] != "ff00" {
		t.Fatalf("format string ignored");
	}
}