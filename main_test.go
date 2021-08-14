package main

import (
	"fmt"
	"testing"
)

func Test_fetchUser(t *testing.T) {
	fetchedUser := fetchUser("megajon")
	fmt.Println("from Test_fetchUser: user", fetchedUser)
	expectedUser := createUserStruct("1", "megajon", "jonathan.seubert@megajon.com")
	fmt.Println("from Test_fetchUser: expectedUser ", expectedUser)
	if fetchedUser.Id == "no file found" {
		fmt.Println("from Test_fetchUser: user file does not exist")
	}
	fmt.Println("from Test_fetchUser: fetched user ", fetchedUser)
}

func Test_createUserListStruct(t *testing.T) {

	// userListStruct := createUserListStruct()

}

func Test_addUser(t *testing.T) {
	id, userName, email := "3", "jon", "jon@jon.com"
	newUser := addUser(id, userName, email)
	expectedUser := user{Id: id, UserName: userName, Email: email}
	if newUser.Id == "user already exists" {
		fmt.Println("from Test_addUser: newUser already exists.")
	} else if newUser == expectedUser {
		fmt.Println("from Test_addUser: newUser successfully created.")
	} else {
		t.Fatal("from Test_addUser: something else went wrong.")
	}
}

func Test_checkIfFileExists(t *testing.T) {
	result := checkIfFileExists("users")
	if result == true {
		fmt.Println("from Test_checkIfFileExists: the file exists")
	}

	if result == false {
		fmt.Println("from Test_checkIfFileExists: the file does not exist")
	}
}

func Test_checkIfUserExists(t *testing.T) {
	userCheckResult := checkIfUserExists("megajon")
	if userCheckResult != true {
		t.Fatal("from Test_checkIfUserExists: The user does not exist")
	} else {
		fmt.Println("from Test_checkIfUserExists: The user exists")
	}
}
