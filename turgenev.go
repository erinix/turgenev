package main

import (
	"fmt"
	"time"
)

var (
	SessionStart int64 = time.Now().Unix()
	TTableFile string = "ttable.txt"
	verbose bool = false
)

type Action byte

const (
	MakeMove Action = iota
	SetCompWhite
)

func main() {
	LoadTTable(TTableFile)
	s := InitialState()

	for {
		if Mode == TUI { PrintState(s, Orientation) }
		c, a := Prompt(s)

		if a == MakeMove {
			s = c
			if Mode == TUI { PrintState(s, Orientation) }
			if s.TrueLegalSuccessors().Len() == 0 {
				verbose = true
				PrintLog("Break point A\n")
				PrintLog(fmt.Sprintf(
		"TrueLegalSuccessors() has %d fewer element(s) than LegalSuccessors()\n",
		s.LegalSuccessors().Len() - s.TrueLegalSuccessors().Len()))
				break
			}
		}

		if Mode == TUI { fmt.Printf("Thinking... ") }
		c = s.SingleSearchWithTTable(4)

		if c == nil {
			PrintLog("Break point B\n")
			break
		}

		DumpTTable(TTableFile)

		if Mode == TUI {
			fmt.Printf("My move: ")
		} else {
			fmt.Printf("move ")
		}
		fmt.Println(MoveString(s, c, Coordinate))
		PrintLog("\t\t\tOUTPUT: move " + MoveString(s, c, Coordinate) + "\n")

		s = c
		if s.TrueLegalSuccessors().Len() == 0 {
			PrintLog("Break point C\n")
			break
		}
	}

	PrintResults(s)

	for {}
}