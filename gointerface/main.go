package main

import (
	"fmt"
)

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)

}

func mian() {

	u := user{"Bill", "bill@163.com"}
	sendNotification(&u)
}

func sendNotification(n notifier) {
	n.notify()

}
