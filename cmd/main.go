package main

import (
	"fmt"
	"math/rand"
	"time"

	"GameOfLife/app"
)

func main() {
	//fmt.Println(app.MoveCursorAndClearScreenSequence)
	fmt.Println(app.ClearScreenSequence)
	rand.Seed(time.Now().UTC().UnixNano())

	newWorld := app.MakeWorld()
	nextWorld := app.MakeWorld()
	newWorld.Seed()
	for {
		newWorld.Display()
		app.Step(newWorld, nextWorld)
		// When new universe contains the next generation, you can swap universes and repeat the process
		newWorld, nextWorld = nextWorld, newWorld
		time.Sleep(app.SleepIteration * time.Millisecond)
		//fmt.Println(app.MoveCursorAndClearScreenSequence)
		fmt.Println(app.ClearScreenSequence)
	}
}
