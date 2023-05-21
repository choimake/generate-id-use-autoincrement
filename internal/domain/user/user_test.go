package user_test

import (
	_ "errors"
	_ "fmt"
	"generate-id-use-autoincrement/internal/domain/user"
	"testing"
)

func TestNewUser(t *testing.T) {
	// ユーザ名を指定して新しいユーザを作成
	username := "John"
	newUser := user.NewUser(username)

	// ユーザ名が正しく設定されているか検証
	if newUser.Name() != username {
		t.Errorf("expected username: %s, got: %s", username, newUser.Name())
	}

	// IDがnilであることを検証
	if newUser.Id() != nil {
		t.Error("expected user ID to be nil, but it's not nil")
	}
}

func TestReconstructUser(t *testing.T) {
	// 正しい値の場合
	validID := int64(1)
	username := "Jane"
	reconstructUser, err := user.ReconstructUser(validID, username)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if reconstructUser.Id() == nil {
		t.Error("expected reconstructUser ID to be non-nil, but it's nil")
	}
	if reconstructUser.Name() != username {
		t.Errorf("expected username: %s, got: %s", username, reconstructUser.Name())
	}

	// 不正な値の場合
	invalidID := int64(0)
	_, err = user.ReconstructUser(invalidID, username)
	if err == nil {
		t.Error("expected error, but got nil")
	} else {
		expectedErrorMessage := "failed to reconstruct the reconstructUser: the ID must be a value greater than or equal to 1"
		if err.Error() != expectedErrorMessage {
			t.Errorf("expected error message: %s, got: %s", expectedErrorMessage, err.Error())
		}
	}
}
