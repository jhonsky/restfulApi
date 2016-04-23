package util

import (
	_ "sync"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

type User struct {
	Id   string
	Name string
}

type UserType struct {
	User
	Type string
}

type Relationship struct {
	Id    string
	State string
}

type RelationshipType struct {
	Relationship
	Type string
}

type IsValid interface {
	Valid() bool
}

func (u *User) Valid() bool {
	return u.Id != "" && u.Name != ""
}

func (ut *Relationship) Valid() bool {
	return ut.Id != "" && ut.State != ""
}
