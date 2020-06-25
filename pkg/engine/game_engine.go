package engine

import (
	"fmt"
	"bufio"
	"strings"
	"os"
	"encoding/json"
	"io/ioutil"
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

	// Infinite game loop
	for {

		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("help", text) == 0 {
			available_commands := "\nAvailable commands are:\n" +
			"help - Displays this help text\n" +
			"move - Use move plus a direction to move your character (up, down, left, right)\n" +
			"info - List your current player info and stats\n" +
			"scout - Displays info about where you are and what may be near you\n" +
			"talk - Use to talk with NPCs, they have may quests or items for sale\n" +
			"fight - Use to fight a monster you run across\n" +
			"run - Use to run away from a monster, may not always work\n" +
			"interact - Used as a general action to interact with the world\n\n"
			fmt.Println(available_commands)
		}

		// Write current game state to file
		write_game_file(game_world)

	}

}


func write_game_file(world_matrix world.World) {
	file, _ := json.MarshalIndent(world_matrix, "", " ")
	_ = ioutil.WriteFile("game_state.json", file, 0644)
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