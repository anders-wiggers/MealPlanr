package models

/*
* 	Module for creating a user model
 */

type User struct {
	name           string
	id             string
	profilePicture string
}

func NewUser(name string, id string, profilePicture string) (user User) {
	user.name = name
	user.id = id
	user.profilePicture = profilePicture
	return user
}
