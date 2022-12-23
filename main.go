package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	width = 60
	height = 20
	lifespan = 1000
	clear = "\033[2J"
	head = "\033[1;1H"
	
	alive = 1
	dead = 0
)

type cells [width + 2][height + 2]int
var cs, ncs cells

func main() {
	initialize()
	for i := 0; i < lifespan; i++ {
		render()
		update()
		time.Sleep(100 * time.Millisecond)
	}

	end()
}

func initialize() {
	fmt.Print(clear)
	rand.Seed(time.Now().UnixNano())
	for y := 1; y <= height; y++ {
		for x := 1; x <= width; x++ {
			cs[x][y] = rand.Intn(2)
		}
	}
}

func render() {
	var screen string
	for y := 0; y < height+2; y++ {
		for x := 0; x < width+2; x++ {
			c := " "
			if cs[x][y] == alive {
				c = "▉"
			}
			screen += c
		}
		screen += "\n"
	}

	fmt.Print(head)
	fmt.Print(screen)
}

func update() {
	for y := 1; y <= height; y++ {
		for x := 1; x <= width; x++ {
			ncs[x][y] = dead
			cnt := cs[x-1][y-1] + cs[x][y-1] + cs[x+1][y-1] + cs[x-1][y] + cs[x+1][y] + cs[x-1][y+1] + cs[x][y+1] + cs[x+1][y+1]

			if cs[x][y] == dead && cnt == 3 {
				ncs[x][y] = alive
			} else if cs[x][y] == alive && (cnt == 2 || cnt == 3) {
				ncs[x][y] = alive
			}
		}
	}

	cs = ncs
}

func end() {
	fmt.Print("Press any key to end. ")
	bufio.NewScanner(os.Stdin).Scan()
	fmt.Print(clear)
	fmt.Print(head)
}
