package data

import "github.com/google/uuid"

type User struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

var data []*User = []*User{}

func CreateUser(n string) *User {

	user := &User{
		Name: n,
		ID:   uuid.New().String(),
	}
	data = append(data, user)
	return user
}

func DeleteUser(id string) bool {
	index := -1
	for i, u := range data {
		if u.ID == id {
			index = i
			break
		}
	}

	if index < 0 {
		return false
	}

	data[index] = data[len(data)-1]
	data = data[:len(data)-1]

	return true
}

func ListUsers() []*User {
	return data
}
