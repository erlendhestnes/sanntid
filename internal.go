package main

import (
	. "./driver"
	. "./network"
	. "fmt"
	. "strconv"
	"time"
)

var current_floor int

func wait_for_input(ch1 chan int, ch2 chan string) {

	for {
		select {
		case floor := <-ch1:
			current_floor = Get_floor_sensor()
			ch2 <- Itoa(floor) + ":" + GetMyIP()
		default:
			time.Sleep(25 * time.Millisecond)
		}
	}

}

func send_to_floor(floor, current_floor int) {
	if current_floor < floor {
		Println("Going up")
		for {
			Speed(150)
			if Get_floor_sensor() == floor {
				Println("I am now at floor: " + Itoa(Get_floor_sensor()))
				Set_floor_indicator(Get_floor_sensor())
				Set_stop_lamp(1)
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
				Set_floor_indicator(Get_floor_sensor())
				Set_stop_lamp(1)
				Speed(0)
				break
			}
		}
	}
}

func KeyboardInput(ch chan int) {
	var a int

	for {
		Scan(&a)
		ch <- a
	}
}

func order(ch1 chan int) {

	i := 0

	for {

		if i < 3 {
			if Get_button_signal(BUTTON_COMMAND, i) == 1 {
				//Println("Button nr: " + Itoa(i) + " has been pressed!")
				ch1 <- i
				time.Sleep(300 * time.Millisecond)
			}
		}
		if i > 0 {
			if Get_button_signal(BUTTON_COMMAND, i) == 1 {
				//Println("Button nr: " + Itoa(i) + " has been pressed!")
				ch1 <- i
				time.Sleep(300 * time.Millisecond)
			}
		}

		i++
		i = i % 4

	}
}

func main() {

	//channels
	ch1 := make(chan int)
	ch2 := make(chan string)

	// Initialize
	Init()
	Speed(0)
	Set_stop_lamp(1)

	go order(ch1)
	go wait_for_input(ch1, ch2)

	neverQuit := make(chan string)
	<-neverQuit

}
