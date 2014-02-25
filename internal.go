package main

import (
	. "./driver"
	. "fmt"
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
	Speed(150)

	// Stop at nearest floor
	last_floor := -1

	send_to_floor(2, last_floor)

}
