package world

import (
	"fmt"
	"os"
	"strings"
	"time"
	"encoding/json"
	"io/ioutil"
)


type World struct {
	Max_x_size int
	Max_y_size int
	Player_info Player
	World_matrix [][]Tile
}


type Tile struct {
	Terrain string
	Objects string
	Npcs string
	Monster_info Monster
	Treasure_info Treasure
}

type Player struct {
	X_position int
	Y_position int
	Level int
	Health int
	Attack int
	Defense int
	Special_abilities string
	Inventory []Item
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

func Save_world(world_matrix *World) {

	currentTime := time.Now()
	save_file_name := "game_state_" + currentTime.Format("2006.01.02 15:04:05") + ".json"
	save_file_name = strings.Replace(save_file_name, " ", "_", -1)
	save_file_name = strings.Replace(save_file_name, ":", ".", -1)

	Write_game_file(world_matrix, save_file_name)
	fmt.Println("Game saved to", save_file_name)
	os.Exit(0)
}


func Write_game_file(world_matrix *World, file_name string) {
	file, _ := json.MarshalIndent(world_matrix, "", " ")
	_ = ioutil.WriteFile(file_name, file, 0644)
}