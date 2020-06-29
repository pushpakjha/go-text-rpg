package engine

import (
	"fmt"
	"bufio"
	"strings"
	"os"
	"github.com/pushpakjha/go-text-rpg/pkg/world"
	"github.com/pushpakjha/go-text-rpg/pkg/player"
)

func Run_game() {
	// Generate world, player, monsters, NPCs etc.
	game_world := world.Generate_world()
	game_world = world.Spawn_monsters(&game_world, 5)
	game_world = world.Spawn_treasure(&game_world, 3)
	game_world = player.Spawn_player(&game_world)

	display_instructions()
	world.Write_game_file(&game_world, "game_state.json")

	// Infinite game loop
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		game_world = player.Evaulate_action(&game_world, text)

		// Write current game state to file
		world.Write_game_file(&game_world, "game_state.json")
	}
}


func display_instructions() {
	instructions := "\nWelcome to the Golang RPG!\n" +
	"You will spawn in a randomly generated world with NPCs, monsters and treasure.\n" +
	"Explore the world and stay alive as long as possible, death is permanent!\n" +
	"If you don't know where to start, type the help command.\n" +
	"Everything in the game is done by typing out your actions.\n" +
	"Good luck!\n\n"
	fmt.Println(instructions)
}