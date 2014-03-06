package main

import (
	//"bufio"
	. "fmt"
	//"io"
	"os"
)

type Dict struct {
	Ip_order string
	Floor    int
	Dir      string
}

type Jobs struct {
	Ip   string
	Dest []Dict
}

type Order struct {
	Ip    string
	Pos   int
	Floor int
}

type Queues struct {
	Int_queue  []Jobs
	Ext_queue  []Dict
	Last_queue []Dict
}

var Fo *os.File

func main() {

	Fo, _ = os.Create("output.txt")

	int_queue := []Jobs{}
	floor := []Dict{}
	int_queue = append(int_queue, Jobs{"147", append(floor, Dict{"int", 2, "standby"})})
	int_queue[0].Dest = append(int_queue[0].Dest, Dict{"int", 1, "standby"})
	int_queue[0].Dest = append(int_queue[0].Dest, Dict{"int", 3, "standby"})

	int_queue = append(int_queue, Jobs{"152", append(floor, Dict{"int", 2, "standby"})})
	int_queue[1].Dest = append(int_queue[1].Dest, Dict{"int", 1, "standby"})
	int_queue[1].Dest = append(int_queue[1].Dest, Dict{"int", 3, "standby"})

	last_queue := []Dict{}
	ext_queue := []Dict{}
	ext_queue = append(ext_queue, Dict{"ext", 3, "up"}, Dict{"ext", 1, "down"}, Dict{"ext", 1, "up"})
	last_queue = append(last_queue, Dict{"134", 3, "up"}, Dict{"154", 1, "down"}, Dict{"183", 2, "up"})

	the_queue := Queues{int_queue, ext_queue, last_queue}
	_ = the_queue
	Print_graphics()
	//Format_queues(the_queue)
	//Format_queues(the_queue)

	/*
		Fprintf(fo, "Internal que for:\n")
		Fprintln(fo, int_queue[0].Ip)
		Fprintf(fo, "---\n")
		Fprintln(fo, int_queue[0].Dest[0].Floor)
		Fprintln(fo, int_queue[0].Dest[1].Floor)
		Fprintln(fo, int_queue[0].Dest[2].Floor)
	*/

	Fo.Close()
}

func Format_int_queue(int_queue []Jobs) {
	Fprintf(Fo, "Int queues:\n")
	if len(int_queue) != 0 {
		for _, job := range int_queue {
			Fprint(Fo, job.Ip)
			Fprintf(Fo, ":")
			for j := 0; j < len(job.Dest); j++ {
				Fprint(Fo, job.Dest[j].Floor)
				Fprintf(Fo, " ")
			}
			Fprintf(Fo, "\n")
		}
	} else {
		Fprintf(Fo, "<empty> \n")
	}
}

func Format_ext_queue(ext_queue []Dict) {
	Fprintf(Fo, "Ext queue:\n")
	if len(ext_queue) != 0 {
		for j := 0; j < len(ext_queue); j++ {
			Fprint(Fo, ext_queue[j].Floor)
			Fprintf(Fo, "->")
			Fprint(Fo, ext_queue[j].Dir)
			Fprintf(Fo, "\n")
		}
	} else {
		Fprintf(Fo, "<empty> \n")
	}
}

func Format_last_queue(last_queue []Dict) {
	Fprintf(Fo, "Last queue:\n")
	if len(last_queue) != 0 {
		for j := 0; j < len(last_queue); j++ {
			Fprint(Fo, last_queue[j].Ip_order)
			Fprintf(Fo, ":")
			Fprint(Fo, last_queue[j].Floor)
			Fprintf(Fo, "\n")
		}
	} else {
		Fprintf(Fo, "<empty> \n")
	}
}

func Format_queues(queues Queues) {
	Fprintf(Fo, "---------Queues--------\n")
	Format_int_queue(queues.Int_queue)
	Fprintf(Fo, "\n")
	Format_ext_queue(queues.Ext_queue)
	Fprintf(Fo, "\n")
	Format_last_queue(queues.Last_queue)
	Fprintf(Fo, "-----------------------\n")
	Fprintf(Fo, "	   V\n")
}
