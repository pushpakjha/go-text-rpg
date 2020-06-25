package world

import (
	"fmt"
	// "time"
	// "math/rand"
)


type World struct {
	Max_x_size int
	Max_y_size int
	World_matrix [][]Tile
}


type Tile struct {
	Player bool
	Terrain string
	Objects string
	Npcs string
	Monster_info Monster
	Treasure_info Treasure
}


func Generate_world() World {
	fmt.Println("Generating world")

	// seed := rand.NewSource(time.Now().UnixNano())
	// random_seed := rand.New(seed)

	// max_x_size := random_seed.Intn(20) + 10
	// max_y_size := random_seed.Intn(20) + 10
	max_x_size := 25
	max_y_size := 25
	arrayofarrays := make([][]Tile, max_y_size)

	world_matrix := create_terrain(arrayofarrays, max_x_size, max_y_size)

	return World {
		Max_x_size: max_x_size,
		Max_y_size: max_y_size,
		World_matrix: world_matrix,
	}
}


func create_terrain(world_matrix [][]Tile, max_x int, max_y int) [][]Tile {
	fmt.Println("Populating terrain for world")

	for y := 0; y < max_y; y++ {
		world_matrix[y] = make([]Tile, max_x)
		for x := 0; x < max_x; x++ {
			var current_tile Tile
			if x < max_x/2 {
				current_tile.Terrain = "grass"
			} else {
				current_tile.Terrain = "dirt"
			}
			world_matrix[y][x] = current_tile
		}
	}

	return world_matrix
}
