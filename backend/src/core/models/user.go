package models

/*
* 	Module for creating a user model
 */

type User struct {
	Name           string `json:"name" binding:"required"`
	Id             string `json:"id" binding:"required"`
	ProfilePicture string `json:"profilePicture" binding:"required"`
}

func NewUser(name string, id string, profilePicture string) (user User) {
	user.Name = name
	user.Id = id
	user.ProfilePicture = profilePicture
	return user
}
