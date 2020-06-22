package engine

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"github.com/pushpakjha/go-text-rpg/pkg/world"
)

func Run_game() {
	// Generate world
	initial_world := world.Generate_world()
	fmt.Println("Max X size", initial_world.Max_x_size)
	fmt.Println("Max Y size", initial_world.Max_y_size)

	// Infinite game loop


	// Write current game state to file
	write_game_file(initial_world)
}


func write_game_file(world_matrix world.World) {

	file, _ := json.MarshalIndent(world_matrix, "", " ")
	_ = ioutil.WriteFile("game_state.json", file, 0644)
}