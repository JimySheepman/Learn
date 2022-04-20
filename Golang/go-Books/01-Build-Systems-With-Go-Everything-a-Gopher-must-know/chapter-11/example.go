package main

import (
	"fmt"
	"strings"
)

type User struct {
	UserId  string
	Frineds []User
}

func (u *User) GetUserId() string {
	return strings.ToUpper(u.UserId)
}

func (u *User) CountFriends() int {
	return len(u.Frineds)
}

func CommonFriend(a *User, b *User) *User {
	for _, af := range a.Frineds {
		for _, bf := range b.Frineds {
			if af.UserId == bf.UserId {
				return &af
			}
		}
	}
	return nil
}

func ExampleUser() {
	j := User{"John", nil}
	m := User{"Mary", []User{j}}
	fmt.Println(m)
}

func ExampleCommonFriend() {
	a := User{"a", nil}
	b := User{"b", []User{a}}
	c := User{"c", []User{a, b}}
	fmt.Println(CommonFriend(&b, &c))
}
func ExampleUser_GetUserId() {
	u := User{"John", nil}
	fmt.Println(u.GetUserId())
}

func ExampleUser_CountFriends() {
	u := User{"John", nil}
	fmt.Println(u.CountFriends())
}
