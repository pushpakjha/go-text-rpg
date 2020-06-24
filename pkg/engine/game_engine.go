package engine

import (
	"encoding/json"
	"io/ioutil"
	"github.com/pushpakjha/go-text-rpg/pkg/world"
	"github.com/pushpakjha/go-text-rpg/pkg/player"
)

func Run_game() {
	// Generate world, player, monsters, NPCs etc.
	game_world := world.Generate_world()
	game_world = player.Spawn_player(&game_world)

	// Infinite game loop


	// Write current game state to file
	write_game_file(game_world)
}


func write_game_file(world_matrix world.World) {

	file, _ := json.MarshalIndent(world_matrix, "", " ")
	_ = ioutil.WriteFile("game_state.json", file, 0644)
}