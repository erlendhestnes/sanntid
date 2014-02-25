package main

import (
	. "./driver"
	. "fmt"
	//. "strconv"
	"time"
)

var last_floor int = -1

func send_to_floor(chan floor int, last_floor int) {
	if last_floor < <-floor {
		Println("Going up")
		for {
			Speed(150)
			Println(Get_floor_sensor())
			if Get_floor_sensor() == <-floor {
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
			if Get_floor_sensor() == <-floor {
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
	Speed(150)
	time.Sleep(25 * time.Millisecond)

	go UserInput(ch1)
	go send_to_floor(ch1, last_floor)

	neverQuit := make(chan string)
	<-neverQuit

}
