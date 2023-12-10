package generations

import (
	"strings"
	"testing"
)

func TestGetEmail(t *testing.T) {
	email, err := GetEmail();
	if err != nil {
		t.Fatalf(err.Error());
	}

	if !strings.Contains(email, "@") {
		t.Fatalf("invalid email format");
	}

	if email == "" {
		t.Fatalf("invalid email");
	}
}

func TestGetEmailWithCustomProvider(t *testing.T) {
	email, err := GetEmailWithCustomProvider("test.org");
	if err != nil {
		t.Fatalf(err.Error());
	}

	if !strings.Contains(email, "@") {
		t.Fatalf("invalid custom email format");
	}

	if !strings.HasSuffix(email, "test.org") {
		t.Fatalf("invalid custom email");
	}
}