package main

import (
	. "./driver"
	. "fmt"
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

	go send_to_floor(1, 2)

	neverQuit := make(chan string)
	<-neverQuit

}
