package user_test

import (
	"generate-id-use-autoincrement/internal/domain/user"
	"testing"
)

func TestNewId(t *testing.T) {
	// 正しい値の場合
	validValue := int64(10)
	validId, err := user.NewId(validValue)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if validId.Value() != validValue {
		t.Errorf("expected value: %d, got: %d", validValue, validId.Value())
	}

	// 不正な値の場合
	invalidValue := int64(0)
	_, err = user.NewId(invalidValue)
	if err == nil {
		t.Error("expected error, but got nil")
	} else {
		expectedErrorMessage := "the ID must be a value greater than or equal to 1"
		if err.Error() != expectedErrorMessage {
			t.Errorf("expected error message: %s, got: %s", expectedErrorMessage, err.Error())
		}
	}
}
