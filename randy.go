/*
Package randy provides a simple way to generate random data.
*/
package randy

import generations "randy/src"

// Generates a name. Can be a `female` or `male` name.
func Name() (string, error) {
	return generations.GetName();
}

// Generates a male name.
func MaleName() (string, error) {
	return generations.GetMaleName();
}

// Generates a female name.
func FemaleName() (string, error) {
	return generations.GetFemaleName();
}

// Generates an email. Generated emails can contain numbers and periods. The IDs use the name database as reference.
func Email() (string, error) {
	return generations.GetEmail();
}

// Generates an email with the given provider. `provider` refers to a custom domain.
// Example: `gmail.com`
func EmailWithProvider(provider string) (string, error) {
	return generations.GetEmailWithCustomProvider(provider);
}

// Generates a color name.
// Example: `red`, `green`, etc
func Color() (string, error) {
	return generations.GetRandomColorName();
}

// Generates a color in hexadecimal format.
// Example: `#220022`
func ColorHex() (string, error) {
	return generations.GetRandomColorHex();
}

// Generates an IPv4 address.
func Ipv4() (string, error) {
	return generations.GetIpv4();
}

// Generates an IPv4 address with custom octets set. `fomartString` is of form `#.#.#.#` where each octet is denoted by a single `#`. To ensure that a particular octet is always in the IP, one can replace `#` with the value.
// Example: `255.#.13.#` is valid. `255.##.13.#`, `255.13` or `255:13:#.#` are invalid variations.
func Ipv4WithCustomOctets(formatString string) (string, error) {
	return generations.GetIpv4WithCustomOctet(formatString);
}

// Generates an IPv6 address.
func Ipv6() (string, error) {
	return generations.GetIpv6();
}

// Generates an IPv6 address with custom segments set. `fomartString` is of form `#:#:#:#:#:#:#:#` where each segment is denoted by a single `#`. To ensure that a particular segment is always in the IP, one can replace `#` with the value.
// Example: `ff21:#:255:#:#:#:#:13` is valid.
func Ipv6WithCustomSegments(formatString string) (string, error) {
	return generations.GetIpv6WithCustomSegment(formatString);
}