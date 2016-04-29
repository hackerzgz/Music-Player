package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"music-player/mlib"
	"music-player/mp"
)

var lib *mlib.MusicManager
var id int = 1
var ctr1, signal chan int

// -------- Lib Command ---------+

func handleLibCommands(tokens []string) {
	switch tokens[1] {
	case "list":
		ShowMusicList()
	case "add":
		if len(tokens) == 6 {
			id++
			lib.Add(&mlib.MusicEntry{strconv.Itoa(id), tokens[2], tokens[3], tokens[4], tokens[5]})
		} else {
			fmt.Println("USAGE: lib add <name><artist><source><type>")
		}
	case "remove":
		if len(tokens) == 3 {
			index, _ := strconv.Atoi(tokens[2])
			lib.Remove(index)
		} else {
			fmt.Println("USAGE: lib remove <name>")
		}
	default:
		fmt.Println("UNrecognized lib command: ", tokens[1])

	}
}

func ShowMusicList() {
	for i := 0; i < lib.Len(); i++ {
		e, _ := lib.Get(i)
		fmt.Println(i+1, ": ", e.Name, " by ", e.Artist, " Type --> ", e.Type)
	}
}

// -------- Lib Command ---------+

// -------- Play Command ---------+

func handlePlayCommand(tokens []string) {
	if len(tokens) != 2 {
		fmt.Println("USAGE: play <name>")
		return
	}
	e := lib.Find(tokens[1])
	if e == nil {
		fmt.Println("The music", tokens[1], "does not exist.")
		return
	}
	mp.Play(e.Source, e.Type)
}

// -------- Play Command ---------+

// -------- Main ---------+

func main() {
	lib = mlib.NewMusicManager()
	fmt.Println(` 
      Enter following commands to control the player: 
      lib list -- View the existing music lib 
      lib add <name><artist><source><type> -- Add a music to the music lib 
      lib remove <name> -- Remove the specified music from the lib 
      play <name> -- Play the specified music 
       q   -- quit 
       `)
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter command-> ")
		rawLine, _, _ := r.ReadLine()
		line := string(rawLine)
		if line == "q" || line == "e" {
			break
		}
		tokens := strings.Split(line, " ")
		if tokens[0] == "lib" {
			handleLibCommands(tokens)
		} else if tokens[0] == "play" {
			handlePlayCommand(tokens)
		} else {
			fmt.Println("Unrecognized command:", tokens[0])
		}
	}
}

// -------- Main ---------+
