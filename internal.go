package main

import (
	. "./driver"
	. "fmt"
	"time"
)

func send_to_floor(floor, last_floor int) {
	for {
		if last_floor < floor {
			for {
				Speed(150)
				sensor_floor := Get_floor_sensor()
				Println(sensor_floor)
				if sensor_floor == floor {
					Println("I am now at floor: " + string(sensor_floor))
					Speed(0)
					break
				}
			}
		}
	}
}

func test() {
	for {
		time.Sleep(100 * time.Millisecond)
		floor := Get_floor_sensor()
		Println(floor)

		if Get_floor_sensor() == 1 {
			Println("I am at: ")
			Println(floor)
			Speed(0)
			time.Sleep(1000 * time.Millisecond)
		}
	}
}

func main() {

	// Initialize
	Init()
	Speed(150)

	go send_to_floor(2, -1)

	neverQuit := make(chan string)
	<-neverQuit

}
