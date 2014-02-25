package main

import (
	. "./driver"
	. "fmt"
	//. "strconv"
	"time"
)

var last_floor int = -1

func send_to_floor_2(ch1 chan int, last_floor int) {

	_ = last_floor

	for {
		select {
		case floor := <-ch1:
			Println(floor)
			send_to_floor(floor, last_floor)
		default:
			time.Sleep(25 * time.Millisecond)
		}
	}

}

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
				last_floor = Get_floor_sensor()
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
				last_floor = Get_floor_sensor()
				Speed(0)
				break
			}
		}
	}
}

func UserInput(ch chan int) {
	var a int

	for {
		Scan(&a)
		ch <- a
	}
}

func main() {

	//channels
	ch1 := make(chan int)

	// Initialize
	Init()
	time.Sleep(25 * time.Millisecond)

	go UserInput(ch1)
	go send_to_floor_2(ch1, last_floor)

	neverQuit := make(chan string)
	<-neverQuit

}
