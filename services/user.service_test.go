package services

import (
	"SE_School/models"
	"SE_School/tests/unit"
	"testing"
)

func TestUserService_AddUserWillNotAcceptInvalidEmail(t *testing.T) {
	service := UserService{unit.NewInMemoryUserRepository()}
	user := models.User{Email: "wrongEmail", Password: ""}

	err := service.AddUser(user)

	if err == nil || err.Error() != "email is not valid" {
		t.Fail()
	}
}

func TestUserService_AddUserWillNotAcceptUserWithRepeatingEmail(t *testing.T) {
	service := UserService{unit.NewInMemoryUserRepository()}
	user := models.User{Email: "test_email@abc.com", Password: ""}

	err := service.AddUser(user)

	if err == nil || err.Error() != "user with such email already exists" {
		t.Fail()
	}
}

func TestUserService_LoginUserWillNotLoginUserWithNewEmail(t *testing.T) {
	service := UserService{unit.NewInMemoryUserRepository()}
	user := models.User{Email: "new_email@abc.com", Password: ""}

	err := service.LoginUser(user)

	if err == nil || err.Error() != "no user with specified email" {
		t.Fail()
	}
}

func TestUserService_LoginUserWillNotLoginUserWithWrongPassword(t *testing.T) {
	service := UserService{unit.NewInMemoryUserRepository()}
	user := models.User{Email: "test_email@abc.com", Password: ""}

	err := service.LoginUser(user)

	if err == nil || err.Error() != "incorrect password specified" {
		t.Fail()
	}
}
