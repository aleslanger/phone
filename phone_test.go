package phone

import (
	"testing"
)

func TestFormatPhone(t *testing.T) {
	phone := "+420123456789"
	expected := "+420 123 456 789"
	result := FormatPhone(phone)

	if result != expected {
		t.Errorf("FormatPhone(%s) = %s; expected %s", phone, result, expected)
	}
}

func TestValidatePhoneNo_ValidPhone(t *testing.T) {
	phone := "+376123456789"
	expected := true
	result := ValidatePhoneNo(phone)

	if result != expected {
		t.Errorf("ValidatePhoneNo(%s) = %t; expected %t", phone, result, expected)
	}
}

func TestValidatePhoneNo_InvalidPhone(t *testing.T) {
	phone := "12345678"
	expected := false
	result := ValidatePhoneNo(phone)

	if result != expected {
		t.Errorf("ValidatePhoneNo(%s) = %t; expected %t", phone, result, expected)
	}
}

func TestValidatePhonePrefix_ValidPrefix(t *testing.T) {
	prefix := "+420"
	expected := true
	result := ValidatePhonePrefix(prefix)

	if result != expected {
		t.Errorf("ValidatePhonePrefix(%s) = %t; expected %t", prefix, result, expected)
	}
}

func TestValidatePhonePrefix_InvalidPrefix(t *testing.T) {
	prefix := "+123"
	expected := false
	result := ValidatePhonePrefix(prefix)

	if result != expected {
		t.Errorf("ValidatePhonePrefix(%s) = %t; expected %t", prefix, result, expected)
	}

}
