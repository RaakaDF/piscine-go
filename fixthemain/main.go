package main

import "github.com/01-edu/z01"

type Door struct {
	state bool
}

func PrintStr(s string) {
	tab := []rune(s)
	for _, r := range tab {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
func OpenDoor(Door *Door) bool {
	PrintStr("Door Opening...")
	Door.state = true
	return true
}
func CloseDoor(Door *Door) bool {
	PrintStr("Door Closing...")
	Door.state = false
	return true
}
func IsDoorOpen(Door *Door) bool {
	PrintStr("is the Door opened ?")
	return Door.state == true
}
func IsDoorClose(Door *Door) bool {
	PrintStr("is the Door closed ?")
	return Door.state == false
}
func main() {
	door := &Door{}
	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
	if door.state == true {
		CloseDoor(door)
	}
}
