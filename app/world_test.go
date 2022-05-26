package app

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAliveCell_NextState_withZeroNeighbours(t *testing.T) {
	newWorld := MakeWorld()
	alive := newWorld.NextState(0, true)

	if alive {
		t.Error(t, "Cell should die due to loneliness")
	}
	require.False(t, alive, "Cell should die due to loneliness")
}

func TestAliveCell_NextState_withTwoNeighbours(t *testing.T) {
	newWorld := MakeWorld()
	alive := newWorld.NextState(2, true)
	fmt.Printf("al: '%v'\n", alive)
	if !alive {
		t.Error(t, "Cell should stay alive")
	}
	require.True(t, alive, "Cell should stay alive")
}

func TestAliveCell_NextState_withThreeNeighbours(t *testing.T) {
	newWorld := MakeWorld()
	alive := newWorld.NextState(3, true)

	if !alive {
		t.Error(t, "Cell should stay alive")
	}
	require.True(t, alive, "Cell should stay alive")
}

func TestAliveCell_NextState_withFourNeighbours(t *testing.T) {
	newWorld := MakeWorld()
	alive := newWorld.NextState(4, true)

	if alive {
		t.Error(t, "Cell should die due to over-population")
	}
	require.False(t, alive, "Cell should die due to over-population")
}

func TestDeadCell_NextState_withThreeNeighbours(t *testing.T) {
	newWorld := MakeWorld()
	alive := newWorld.NextState(3, false)

	if !alive {
		t.Error(t, "Cell should get alive with 3 neighbours")
	}
	require.True(t, alive, "Cell should get alive with 3 neighbours")
}
