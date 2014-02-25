package main

import (
	. "./driver"
	. "fmt"
	. "strconv"
	"time"
)

var last_floor int

func send_to_floor_2(ch1 chan int) {

	for {
		select {
		case floor := <-ch1:
			last_floor = Get_floor_sensor()
			Println("Going to floor : " + Itoa(floor))
			Println("From previous floor : " + Itoa(last_floor))
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
			if Get_floor_sensor() == floor {
				Println("I am now at floor: " + Itoa(Get_floor_sensor()))
				//time.Sleep(50 * time.Millisecond)
				Speed(0)
				break
			}
		}
	} else {
		Println("Going down")
		for {
			Speed(-150)
			if Get_floor_sensor() == floor {
				Println("I am now at floor: " + Itoa(Get_floor_sensor()))
				//time.Sleep(50 * time.Millisecond)
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
	Speed(0)

	go UserInput(ch1)
	go send_to_floor_2(ch1)

	neverQuit := make(chan string)
	<-neverQuit

}
