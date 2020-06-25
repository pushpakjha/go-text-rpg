package player

import (
	"fmt"
	"time"
	"strings"
	"math/rand"
	"github.com/pushpakjha/go-text-rpg/pkg/world"
)


func Spawn_player(game_world *world.World) world.World {
	fmt.Println("Spawning player")

	seed := rand.NewSource(time.Now().UnixNano())
	random_seed := rand.New(seed)

	x_position := random_seed.Intn(game_world.Max_x_size)
	y_position := random_seed.Intn(game_world.Max_y_size)

	player := world.Player{
		X_position: x_position,
		Y_position: y_position,
		Level: 1,
		Health: 20,
		Attack: 5,
		Defense: 1,
		Special_abilities: "None",
	}

	game_world.Player_info = player

	return *game_world
}


func Evaulate_action(game_world *world.World, action string) world.World {

	move_info := strings.Fields(action)
	if strings.HasPrefix(action, "move") {
		action = "move"
	}

	switch action {
	case "help":
		available_commands := "\nAvailable commands are:\n" +
		"help - Displays this help text\n" +
		"move - Use move plus a direction to move your character (up, down, left, right)\n" +
		"info - List your current player info and stats\n" +
		"scout - Displays info about where you are and what may be near you\n" +
		"talk - Use to talk with NPCs, they have may quests or items for sale\n" +
		"fight - Use to fight a monster you run across\n" +
		"run - Use to run away from a monster, may not always work\n" +
		"interact - Use as a general action to interact with the world\n\n"
		fmt.Println(available_commands)
		break
	case "move":
		game_world = move_player(game_world, move_info)
		break
	case "info":
		fmt.Println("Not implemented")
		break
	case "scout":
		fmt.Println("Not implemented")
		break
	case "talk":
		fmt.Println("Not implemented")
		break
	case "fight":
		fmt.Println("Not implemented")
		break
	case "run":
		fmt.Println("Not implemented")
		break
	case "interact":
	  fmt.Println("Nothing to interact with")
	  break
	}

	return *game_world

}


func move_player(game_world *world.World, move_info []string) *world.World {

	var move_direction string
	if len(move_info) > 1 {
		move_direction = move_info[1]
	}

	switch move_direction {
	case "up":
		if (game_world.Player_info.Y_position - 1) < 0 {
			fmt.Println("You can't move down, you're at the edge of the map.\n")
		} else {
			fmt.Println("You move up one tile.\n")
			game_world.Player_info.Y_position -= 1
		}
		break
	case "down":
		if (game_world.Player_info.Y_position + 1) > game_world.Max_y_size - 1 {
			fmt.Println("You can't move up, you're at the edge of the map.\n")
		} else {
			fmt.Println("You move down one tile.\n")
			game_world.Player_info.Y_position += 1
		}
		break
	case "left":
		if (game_world.Player_info.X_position - 1) < 0 {
			fmt.Println("You can't move down, you're at the edge of the map.\n")
		} else {
			fmt.Println("You move left one tile.\n")
			game_world.Player_info.X_position -= 1
		}
		break
	case "right":
		if (game_world.Player_info.X_position + 1) > game_world.Max_x_size - 1 {
			fmt.Println("You can't move down, you're at the edge of the map.\n")
		} else {
			fmt.Println("You move right one tile.\n")
			game_world.Player_info.X_position += 1
		}
		break
	default:
		fmt.Println("\nInvalid direction, valid options are up, down, left, right")
	}

	return game_world

}