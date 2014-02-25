package main

import (
	. "./driver"
	. "fmt"
	. "strconv"
	"time"
)

func send_to_floor_2(ch1 chan int) {

	ch2 := make(chan int)
	<-ch2 = Get_floor_sensor()

	for {
		select {
		case floor := <-ch1:
			Println("Going to floor : " + Itoa(floor))
			Println("From previous floor : " + Itoa(<-ch2))
			send_to_floor(floor, ch2)
		default:
			time.Sleep(25 * time.Millisecond)
		}
	}

}

func send_to_floor(floor int, last_floor chan int) {
	if <-last_floor < floor {
		Println("Going up")
		for {
			Speed(150)
			if Get_floor_sensor() == floor {
				Println("I am now at floor: " + Itoa(Get_floor_sensor()))
				time.Sleep(25 * time.Millisecond)
				<-last_floor = floor
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
				time.Sleep(25 * time.Millisecond)
				<-last_floor = floor
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
	Speed(0)
	last_floor = Get_floor_sensor()

	go UserInput(ch1)
	go send_to_floor_2(ch1)

	neverQuit := make(chan string)
	<-neverQuit

}
