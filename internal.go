package main

import (
	. "./driver"
	. "fmt"
	"time"
)

func send_to_floor(floor, last_floor int) {

	for {

		floor := Get_floor_sensor()
		Println("Going to floor nr: " + string(floor))

		switch floor {
		case 0:
			Speed(150)
			time.Sleep(25 * time.Millisecond)
			Speed(0)
			time.Sleep(1000 * time.Millisecond)
			Speed(150)
			time.Sleep(1000 * time.Millisecond)
			last_floor = floor

		case 1:
			if last_floor > floor {
				Speed(150)
				time.Sleep(25 * time.Millisecond)
				Speed(0)
				time.Sleep(1000 * time.Millisecond)
				Speed(-150)
				time.Sleep(1000 * time.Millisecond)
				last_floor = floor
			} else {
				Speed(-150)
				time.Sleep(25 * time.Millisecond)
				Speed(0)
				time.Sleep(1000 * time.Millisecond)
				Speed(150)
				time.Sleep(1000 * time.Millisecond)
				last_floor = floor
			}
		case 2:
			if last_floor > floor {
				Speed(150)
				time.Sleep(25 * time.Millisecond)
				Speed(0)
				time.Sleep(1000 * time.Millisecond)
				Speed(-150)
				time.Sleep(1000 * time.Millisecond)
				last_floor = floor
			} else {
				Speed(-150)
				time.Sleep(25 * time.Millisecond)
				Speed(0)
				time.Sleep(1000 * time.Millisecond)
				Speed(150)
				time.Sleep(1000 * time.Millisecond)
				last_floor = floor
			}
		case 3:
			Speed(-150)
			time.Sleep(25 * time.Millisecond)
			Speed(0)
			time.Sleep(1000 * time.Millisecond)
			Speed(-150)
			time.Sleep(1000 * time.Millisecond)
			last_floor = floor
		}
	}
}

func main() {

	// Initialize
	Init()
	Speed(150)

	// Stop at nearest floor
	last_floor := -1

	go send_to_floor(2, last_floor)

}
