package app

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMakeWorld(t *testing.T) {
	tests := []struct {
		name                     string
		width                    int
		height                   int
		expectWidth              int
		expectHeight             int
		expectLenUniverse        int
		expectLenOfUniverseValue int
		expectError              error
	}{
		{
			name:                     "expect successful creating of new world",
			width:                    20,
			height:                   11,
			expectLenUniverse:        11,
			expectLenOfUniverseValue: 20,
			expectWidth:              20,
			expectHeight:             11,
			expectError:              nil,
		},
		{
			name:        "expect failed, height less than zero",
			width:       20,
			height:      -30,
			expectError: ErrHeightNotValid,
		},
		{
			name:        "expect failed, height is zero",
			width:       20,
			height:      0,
			expectError: ErrHeightNotValid,
		},
		{
			name:        "expect failed, width less than zero",
			width:       -20,
			height:      30,
			expectError: ErrWidthNotValid,
		},
		{
			name:        "expect failed, width is zero",
			width:       0,
			height:      30,
			expectError: ErrWidthNotValid,
		},
		{
			name:        "expect failed, width and height both are equally zero value",
			width:       0,
			height:      0,
			expectError: ErrNotValidWidthAndHeight,
		},
		{
			name:        "expect failed, width and height both are less than zero",
			width:       -39,
			height:      -69,
			expectError: ErrNotValidWidthAndHeight,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			world, err := MakeWorld(tt.width, tt.height)
			require.Equal(t, tt.expectError, err)

			if err == nil {
				require.Equal(t, tt.expectLenUniverse, len(world.universe))
				require.Equal(t, tt.expectLenOfUniverseValue, len(world.universe[0]))
				require.Equal(t, tt.expectWidth, world.width)
				require.Equal(t, tt.expectHeight, world.height)
			}
		})
	}

}

func TestWorld_NextState(t *testing.T) {
	tests := []struct {
		name        string
		neighbors   int
		alive       bool
		expectAlive bool
		errMessage  string
	}{
		{
			name:        "expect Alive in the next state, cell with two neighbours",
			neighbors:   2,
			alive:       true,
			expectAlive: true,
			errMessage:  "Cell should stay alive",
		},
		{
			name:        "expect Dead in the next state, cell with zero neighbours",
			neighbors:   0,
			alive:       true,
			expectAlive: false,
			errMessage:  "Cell should die due to loneliness",
		},
		{
			name:        "expect Alive, alive cell with three neighbours",
			neighbors:   3,
			alive:       true,
			expectAlive: true,
			errMessage:  "Cell should stay alive",
		},
		{
			name:        "expect Dead, alive cell with four neighbours",
			neighbors:   4,
			alive:       true,
			expectAlive: false,
			errMessage:  "Cell should die due to over-population",
		},
		{
			name:        "expect ALive, dead cell with 3 neighbours",
			neighbors:   3,
			alive:       false,
			expectAlive: true,
			errMessage:  "Cell should be alive with 3 neighbours",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newWorld, err := MakeWorld(25, 25)
			require.NoError(t, err)

			gotState := newWorld.NextState(tt.neighbors, tt.alive)
			if !tt.expectAlive {
				require.False(t, gotState, tt.errMessage)
			} else {
				require.True(t, gotState, tt.errMessage)
			}
		})
	}
}

func TestWorld_Alive_ExpectNoNeighbour(t *testing.T) {
	tests := []struct {
		name            string
		width           int
		height          int
		xWidth          int
		yHeight         int
		expectAlive     bool
		expectNeighbour bool
		neighbourY      int
		neighbourX      int
	}{
		{
			name:        "expect no neighbour 1",
			width:       80,
			height:      15,
			xWidth:      -1,
			yHeight:     -1,
			expectAlive: false,
		},
		{
			name:        "expect no neighbour 2",
			width:       80,
			height:      15,
			xWidth:      0,
			yHeight:     0,
			expectAlive: false,
		},
		{
			name:        "expect no neighbour 3",
			width:       80,
			height:      15,
			xWidth:      0,
			yHeight:     14,
			expectAlive: false,
		},
		{
			name:        "expect no neighbour 4",
			width:       80,
			height:      15,
			xWidth:      0,
			yHeight:     -1,
			expectAlive: false,
		},
		{
			name:        "expect alive neighbour D",
			width:       80,
			height:      15,
			xWidth:      300,
			yHeight:     265,
			expectAlive: true,
		},
		{
			name:        "expect alive neighbour E",
			width:       80,
			height:      15,
			xWidth:      79,
			yHeight:     14,
			expectAlive: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wl, err := MakeWorld(tt.width, tt.height)
			require.NoError(t, err)

			wl.universe[1][1] = true   //A
			wl.universe[1][2] = true   //B
			wl.universe[10][60] = true //D
			wl.universe[14][79] = true //E

			gotAlive := wl.Alive(tt.xWidth, tt.yHeight)
			require.Equal(t, tt.expectAlive, gotAlive)
		})
	}
}

func TestWorld_Neighbors(t *testing.T) {
	wl, err := MakeWorld(10, 10)
	require.NoError(t, err)

	// Here we should create alive cells in our world
	/*
		X	_ 0 1 2 3
		Y	0       D
			1   A C
			2   B   E
	*/
	wl.universe[1][1] = true //A
	wl.universe[1][2] = true //B
	wl.universe[2][1] = true //C
	wl.universe[3][0] = true //D
	wl.universe[3][2] = true //E

	t.Run("cell have 2 neighbours", func(t *testing.T) {
		gotNeighbours := wl.Neighbors(1, 1) // 2
		require.Equal(t, 2, gotNeighbours)
	})
	t.Run("cell have 4 neighbours", func(t *testing.T) {
		gotNeighbours := wl.Neighbors(2, 2) // 4
		require.Equal(t, 4, gotNeighbours)
	})
}
