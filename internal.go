package main

import (
	. "./driver"
	. "fmt"
	. "strconv"
	"time"
)

var current_floor int

func Wait_for_input(ch1 chan int, ch2 chan string) {

	_ = ch2

	for {
		select {
		case floor := <-ch1:
			Send_to_floor(floor)
		default:
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func Send_to_floor(floor int) {
	current_floor = Get_floor_sensor()
	if current_floor < floor {
		Println("Going up")
		for {
			Speed(150)
			if Get_floor_sensor() == floor {
				Println("I am now at floor: " + Itoa(Get_floor_sensor()))
				Set_stop_lamp(1)
				time.Sleep(25 * time.Millisecond)
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
				Set_stop_lamp(1)
				time.Sleep(25 * time.Millisecond)
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

func Order(ch1 chan int) {

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
		time.Sleep(25 * time.Millisecond)

	}
}

func Floor_indicator(last_floor chan int) {
	Println("executing floor indicator!")
	for {
		/*
			if Get_floor_sensor() != -1 {
				Set_floor_indicator(Get_floor_sensor())
				Println(Get_floor_sensor())
				last_floor <- Get_floor_sensor()
				time.Sleep(25 * time.Millisecond)
			}
		*/
		time.Sleep(100 * time.Millisecond)
		Println(Get_floor_sensor())
	}
}

func main() {

	//channels
	ch1 := make(chan int)
	ch2 := make(chan string)
	ch3 := make(chan int)

	// Initialize
	Init()
	Speed(0)
	Set_stop_lamp(1)

	go Floor_indicator(ch3)
	go Order(ch1)
	go Wait_for_input(ch1, ch2)

	neverQuit := make(chan string)
	<-neverQuit

}
