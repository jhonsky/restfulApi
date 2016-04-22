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

//func (r *Relationship) ConvertState() {
//	switch r.State {
//	case "1":
//		r.State = "disliked"
//	case "2":
//		r.State = "liked"
//	case "3":
//		r.State = "matched"
//	default:
//		r.State = "no relation"
//	}
//}
