package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type user struct {
	Id       string
	UserName string
	Email    string
}

type users []struct {
	Id       string
	UserName string
	Email    string
}

func main() {
	// var m users
	// jsonList, _ := ioutil.ReadFile("users")
	// json.Unmarshal(jsonList, &m)
	// unmarshalledJsonList := fmt.Sprint(m)
	// var fetchedUser user
	// for _, u := range unmarshalledJsonList {
	// 	fmt.Println(u)
	// }
	// fmt.Println(unmarshalledJsonList)
}

func createUserStruct(id string, userName string, email string) user {
	return user{Id: id, UserName: userName, Email: email}
}

func addUser(id string, userName string, email string) user {
	fileName := "store/users"
	// var returnedUser user
	var unmarshaledUserList users
	fileCheckResult := checkIfFileExists("users")
	fmt.Println("from addUser: fileCheckResult ", fileCheckResult)
	if fileCheckResult == false {
		newUserList := []user{{Id: id, UserName: userName, Email: email}}
		newUserJson, _ := json.Marshal(newUserList)
		ioutil.WriteFile(fileName, newUserJson, 0644)
		fmt.Println("User does not exists. Creating user.")
		return fetchUser(userName)
	} else {
		checkUserResult := checkIfUserExists(userName)
		fmt.Println("from addUser: checkUserResult ", checkUserResult)
		if checkUserResult == true {
			return user{Id: "user already exists", UserName: "", Email: ""}
		} else {
			newUser := createUserStruct(id, userName, email)
			userListFromFile, _ := ioutil.ReadFile(fileName)
			err := json.Unmarshal(userListFromFile, &unmarshaledUserList)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Print("from addUser: unmarshaledUserList ", unmarshaledUserList)
			unmarshaledUserList := append(unmarshaledUserList, newUser)
			remarshaledUserList, e := json.Marshal(unmarshaledUserList)
			if e != nil {
				fmt.Println(e)
			}
			ioutil.WriteFile(fileName, remarshaledUserList, 0664)
			return fetchUser(userName)
		}
		// if fetchUserResult.Id != "empty" {
		// 	returnedUser = fetchUserResult
		// }
	}
}

func fetchUser(userName string) user {
	var fetchedUser user
	var unmarshaledUserList users
	fileCheckResult := checkIfFileExists("users")
	fmt.Println("from fetchUser: fileCheckResult", fileCheckResult)
	if fileCheckResult == true {
		userCheckResult := checkIfUserExists(userName)
		fmt.Println("from fetchUser: userCheckResult: ", userCheckResult)
		if userCheckResult == true {
			userListFromFile, _ := ioutil.ReadFile("store/users")
			err := json.Unmarshal(userListFromFile, &unmarshaledUserList)
			if err != nil {
				fmt.Println("from fetchUser: ", err)
			}
			for _, u := range unmarshaledUserList {
				if u.UserName == userName {
					fetchedUser = u
				}
			}
			return fetchedUser
		} else {
			return user{Id: "no user fetched", UserName: "", Email: ""}
		}
	} else {
		return user{Id: "no file found", UserName: "", Email: ""}
	}
}

func updateUser(userName string) {

}

func deleteUser(userName string) {

}

func checkIfFileExists(fileName string) bool {
	fileExists := false
	files, err := ioutil.ReadDir("store")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println("from checkIfFileExists: fileName ", file.Name())
		if file.Name() == fileName {
			fileExists = true
		}
	}
	return fileExists
}

func checkIfUserExists(userName string) bool {
	foundUser := false
	var unmarshaledUserList users
	userListFromFile, err := ioutil.ReadFile("store/users")
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(userListFromFile, &unmarshaledUserList)
	for _, u := range unmarshaledUserList {
		if u.UserName == userName {
			foundUser = true
		}
	}
	return foundUser
}
