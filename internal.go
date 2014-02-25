package main

import (
	. "./driver"
	. "fmt"
	. "strconv"
	"time"
)

var current_floor int

func wait_for_input(ch1 chan int) {

	for {
		select {
		case floor := <-ch1:
			Set_stop_lamp(0)
			current_floor = Get_floor_sensor()
			Println("Going to floor : " + Itoa(floor))
			Println("From previous floor : " + Itoa(current_floor))
			send_to_floor(floor, current_floor)
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
				//Set_floor_indicator(Get_floor_sensor())
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
				//Set_floor_indicator(Get_floor_sensor())
				Set_stop_lamp(1)
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

func order(ch1 chan int) {

	i := 0

	for {

		if i < 3 {
			if Get_button_signal(BUTTON_COMMAND, i) == 1 {
				ch1 <- i
				time.Sleep(300 * time.Millisecond)
			}
		}
		if i > 0 {
			if Get_button_signal(BUTTON_COMMAND, i) == 1 {
				ch1 <- i
				time.Sleep(300 * time.Millisecond)
			}
		}

		i++
		i = i % 4

	}
}

func button_test() {
	for {
		if Get_button_signal(BUTTON_COMMAND, 0) {
			Println("btn pressed!")
		}
	}
}

func main() {

	//channels
	ch1 := make(chan int)

	// Initialize
	Init()
	Speed(0)
	Set_stop_lamp(1)

	go button_test()
	//go order(ch1)
	go wait_for_input(ch1)

	neverQuit := make(chan string)
	<-neverQuit

}
