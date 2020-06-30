package world

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	save_file_name := "game_state_" + currentTime.Format("2006.01.02_15.04.05") + ".json"

	Write_game_file(world_matrix, save_file_name)
	fmt.Println("Game saved to", save_file_name)
	fmt.Println("Returning to game")
	return
}


func Write_game_file(world_matrix *World, file_name string) {
	file, _ := json.MarshalIndent(world_matrix, "", " ")
	_ = ioutil.WriteFile(file_name, file, 0644)
}


func Load_world() World {
	file_name_map := make(map[string]string)

	file_number := 1
	files, _ := ioutil.ReadDir("./")
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".json") {
			fmt.Printf("%d) %s\n", file_number, f.Name())
			file_ind := strconv.Itoa(file_number) 
			file_name_map[file_ind] = f.Name()
			file_number += 1
		}
	}

	for {
		fmt.Println("Pick a number to load that game save file, the current game is game_state.json")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		if val, ok := file_name_map[text]; ok {
			fmt.Println("Loading:", val)
			file, _ := ioutil.ReadFile(val)
			game_world := World{}
			_ = json.Unmarshal([]byte(file), &game_world)
			fmt.Println("Finished loading, returning to game")
			return game_world
		} else {
			fmt.Println("Invalid option, pick a number to load that game save file")
		}
	}
}


func Quit_game() {
	fmt.Println("Make sure you save first! Are you sure you want to quit? yes/no")
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		if text == "yes" {
			fmt.Println("Quitting game")
			os.Exit(0)
		} else if text == "y" {
			fmt.Println("Quitting game")
			os.Exit(0)
		} else if text == "no" {
			fmt.Println("Returning to game")
			return
		}  else if text == "n" {
			fmt.Println("Returning to game")
			return
		} else {
			fmt.Println("Invalid option, type yes or no")
		}
	}
}