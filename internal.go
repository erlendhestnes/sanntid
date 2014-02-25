package main

import (
	. "./driver"
	. "fmt"
	. "strconv"
	"time"
)

func send_to_floor(floor, last_floor int) {
	if last_floor < floor {
		Println("Going up")
		for {
			Speed(150)
			Println(Get_floor_sensor())
			if Get_floor_sensor() == floor {
				Println("I am now at floor: ")
				Println(Get_floor_sensor())
				time.Sleep(25 * time.Millisecond)
				Speed(0)
				break
			}
		}
	} else {
		Println("Going down")
		for {
			Speed(-150)
			Println(Get_floor_sensor())
			if Get_floor_sensor() == floor {
				Println("I am now at floor: ")
				Println(Get_floor_sensor())
				time.Sleep(25 * time.Millisecond)
				Speed(0)
				break
			}
		}
	}
}

func main() {

	// Initialize
	Init()
	Speed(150)
	time.Sleep(25 * time.Millisecond)

	go send_to_floor(0, 2)

	neverQuit := make(chan string)
	<-neverQuit

}
