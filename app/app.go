package app

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var (
	width  int
	height int
)

// Run runs server
func Run() error {
	flag.IntVar(&width, "width", 25, "-width=25")
	flag.IntVar(&height, "height", 25, "-height=25")
	flag.Parse()

	fmt.Println(ClearScreenSequence) // also we can use this constant MoveCursorAndClearScreenSequence
	rand.Seed(time.Now().UTC().UnixNano())

	newWorld, err := MakeWorld(width, height)
	if err != nil {
		return err
	}
	nextWorld, err := MakeWorld(width, height)
	if err != nil {
		return err
	}

	newWorld.Seed()
	for {
		newWorld.Display()
		Step(newWorld, nextWorld)
		// When new universe contains the next generation, you can swap universes and repeat the process
		newWorld, nextWorld = nextWorld, newWorld
		time.Sleep(SleepIteration * time.Millisecond)

		fmt.Println(ClearScreenSequence) // also we can use this constant app.MoveCursorAndClearScreenSequence
	}

}
