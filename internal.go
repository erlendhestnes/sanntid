package main

import (
	. "./driver"
	. "fmt"
	"time"
)

func send_to_floor(floor, last_floor int) {

	if last_floor < floor {
		for {
			Speed(150)
			sensor_floor := Get_floor_sensor()
			Println(string(sensor_floor))
			if sensor_floor == floor {
				Println("I am now at floor: " + string(sensor_floor))
				break
			}
		}
	}
}

func main() {

	// Initialize
	Init()
	for {
		Speed(150)
		time.Sleep(25 * time.Millisecond)

		if Get_floor_sensor() == 2 {
			break
		}
	}
}
