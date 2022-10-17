package main

import "github.com/stewie1520/golog/cmd/foo"

func main() {
	f := foo.Foo{}
	f.ChangeName("Hieu")
	changeNameToHoang(f)
	f.PrintName()
}

func changeNameToHoang(f foo.Foo) {
	f.ChangeName("Hoang")
}