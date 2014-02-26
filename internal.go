package main

import (
	. "./driver"
	. "fmt"
	. "strconv"
	"time"
)

var current_floor int

func Wait_for_input(int_button, ext_button chan int, int_order chan string) {

	_ = int_order
	var floor int

	for {
		select {
		case floor = <-int_button:
			Send_to_floor(floor)
			Set_button_lamp(BUTTON_COMMAND, floor, 0)
		case floor = <-ext_button:
			Send_to_floor(floor)
			Set_button_lamp(BUTTON_CALL_UP, floor, 0)

		default:
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func Send_to_floor(floor int) {
	current_floor = Get_floor_sensor()
	//Elev_set_door_open_lamp(0)
	Set_stop_lamp(0)

	if current_floor < floor {
		Println("Going up")
		for {
			Speed(150)
			if Get_floor_sensor() == floor {
				Println("I am now at floor: " + Itoa(Get_floor_sensor()))
				Set_stop_lamp(1)
				//Elev_set_door_open_lamp(1)
				time.Sleep(25 * time.Millisecond)
				Speed(0)
				break
			}
			time.Sleep(25 * time.Millisecond)
		}
	} else {
		Println("Going down")
		for {
			Speed(-150)
			if Get_floor_sensor() == floor {
				Println("I am now at floor: " + Itoa(Get_floor_sensor()))
				Set_stop_lamp(1)
				//Elev_set_door_open_lamp(1)
				time.Sleep(25 * time.Millisecond)
				Speed(0)
				break
			}
			time.Sleep(25 * time.Millisecond)
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

func Ext_order(int_button chan int) {

	i := 0

	for {

		if i < 3 {
			if Get_button_signal(BUTTON_CALL_UP, i) == 1 {
				//Println("Button nr: " + Itoa(i) + " has been pressed!")
				int_button <- i
				time.Sleep(300 * time.Millisecond)
			}
		}
		if i > 0 {
			if Get_button_signal(BUTTON_CALL_DOWN, i) == 1 {
				//Println("Button nr: " + Itoa(i) + " has been pressed!")
				int_button <- i
				time.Sleep(300 * time.Millisecond)
			}
		}

		i++
		i = i % 4
		time.Sleep(25 * time.Millisecond)

	}
}

func Int_order(int_button chan int) {

	i := 0

	for {

		if Get_button_signal(BUTTON_COMMAND, i) == 1 {
			//Println("Button nr: " + Itoa(i) + " has been pressed!")
			int_button <- i
			Set_button_lamp(BUTTON_COMMAND, i, 1)
			time.Sleep(300 * time.Millisecond)
		}

		i++
		i = i % 4
		time.Sleep(25 * time.Millisecond)

	}
}

func Floor_indicator(last_floor chan string) {
	Println("executing floor indicator!")
	_ = last_floor
	var floor int
	for {
		floor = Get_floor_sensor()
		if floor != -1 {
			Set_floor_indicator(floor)
			//last_floor <- floor
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func main() {

	//channels
	int_button := make(chan int)
	ext_button := make(chan int)

	int_order := make(chan string)
	ext_order := make(chan string)
	last_order := make(chan string)

	// Initialize
	Init()
	Speed(0)
	Set_stop_lamp(1)

	go Floor_indicator(last_order)
	go Int_order(int_button)
	go Ext_order(ext_button)
	go Wait_for_input(int_button, int_order)

	neverQuit := make(chan string)
	<-neverQuit

}
