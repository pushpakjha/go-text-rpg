package player

import (
	"fmt"
	"time"
	"strings"
	"strconv"
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

	action_info := strings.Fields(action)
	if strings.HasPrefix(action, "move") {
		action = "move"
	} else if strings.HasPrefix(action, "equip") {
		action = "equip"
	} else if strings.HasPrefix(action, "unequip") {
		action = "unequip"
	}

	switch action {
	case "help":
		print_help()
		break
	case "move":
		game_world = move_player(game_world, action_info)
		check_collision(game_world)
		break
	case "info":
		display_player_info(game_world)
		break
	case "scout":
		fmt.Println("Not implemented")
		break
	case "talk":
		fmt.Println("Not implemented")
		break
	case "fight":
		game_world = fight_monster(game_world)
		break
	case "run":
		fmt.Println("Not implemented")
		break
	case "equip":
		game_world = equip_item(game_world, action_info)
		break
	case "unequip":
		game_world = unequip_item(game_world, action_info)
		break
	case "interact":
		game_world = interact_with_world(game_world)
		break
	case "save":
		world.Save_world(game_world)
		break
	case "load":
		fmt.Println("Not implemented")
		break
	}

	return *game_world
}


func move_player(game_world *world.World, move_info []string) *world.World {
	var move_direction string
	var move_number_str string
	var move_number_int int

	if len(move_info) > 2 {
		move_direction = move_info[1]
		move_number_str = move_info[2]
	} else if len(move_info) > 1 {
		move_direction = move_info[1]
		move_number_str = "1"
	}

	if i, err := strconv.Atoi(move_number_str); err == nil {
		move_number_int = i
	} else {
		panic(err)
	}

	switch move_direction {
	case "up":
		if (game_world.Player_info.Y_position - move_number_int) < 0 {
			fmt.Println("\nYou can't move that far up, you hit the edge of the map.")
		} else {
			fmt.Println("\nYou move up", move_number_int, "tile(s).")
			game_world.Player_info.Y_position -= move_number_int
		}
		break
	case "down":
		if (game_world.Player_info.Y_position + move_number_int) > game_world.Max_y_size - 1 {
			fmt.Println("\nYou can't move that far down, you hit the edge of the map.")
		} else {
			fmt.Println("\nYou move down", move_number_int, "tile(s).")
			game_world.Player_info.Y_position += move_number_int
		}
		break
	case "left":
		if (game_world.Player_info.X_position - move_number_int) < 0 {
			fmt.Println("\nYou can't move that far left, you hit the edge of the map.")
		} else {
			fmt.Println("\nYou move left", move_number_int, "tile(s).")
			game_world.Player_info.X_position -= move_number_int
		}
		break
	case "right":
		if (game_world.Player_info.X_position + move_number_int) > game_world.Max_x_size - 1 {
			fmt.Println("\nYou can't move that far right, you hit the edge of the map.")
		} else {
			fmt.Println("\nYou move right", move_number_int, "tile(s).")
			game_world.Player_info.X_position += move_number_int
		}
		break
	default:
		fmt.Println("\nInvalid direction, valid options are up, down, left, right")
	}

	return game_world
}


func print_help() {
	available_commands :=
		"\nAvailable commands are:\n" +
		"help - Displays this help text\n" +
		"move - Use to move your character, provide a direction (up, down, left, right)" +
				" plus the number of tiles to move (as an integer), defaults to 1\n" +
		"info - List your current player info and stats\n" +
		"scout - Displays info about where you are and what may be near you\n" +
		"talk - Use to talk with NPCs, they have may quests or items for sale\n" +
		"fight - Use to fight a monster you run across\n" +
		"run - Use to run away from a monster, may not always work\n" +
		"equip - Use to equip an item in your inventory, type equip then the item name\n" +
		"unequip - Use to unequip an item\n" +
		"interact - Use as a general action to interact with the world\n" +
		"save - Use to save the current game with a timestamped file\n" +
		"load - Use to load a saved game\n\n"
	fmt.Println(available_commands)
}


func check_collision(game_world *world.World) string {
	player_x_pos := game_world.Player_info.X_position
	player_y_pos := game_world.Player_info.Y_position

	var collision_type string
	if game_world.World_matrix[player_y_pos][player_x_pos].Monster_info.Monster_type != "" {
		fmt.Println("Found a monster!")
		collision_type = "monster"
	} else if game_world.World_matrix[player_y_pos][player_x_pos].Treasure_info.Treasure_text != "" {
		fmt.Println("Found treasure!")
		collision_type = "treasure"
	}
	return collision_type
}


func display_player_info(game_world *world.World) {
	fmt.Println("\nPlayer Info:")
	fmt.Println("Level -", game_world.Player_info.Level)
	fmt.Println("Health -", game_world.Player_info.Health)
	fmt.Println("Attack -", game_world.Player_info.Attack)
	fmt.Println("Defense -", game_world.Player_info.Defense)
	fmt.Println("Special abilities -", game_world.Player_info.Special_abilities)

	fmt.Println("\nInventory:")
	if len(game_world.Player_info.Inventory) > 0 {
		for x := 0; x < len(game_world.Player_info.Inventory); x++ {
			fmt.Printf("%s:\n", game_world.Player_info.Inventory[x].Item_name)
			fmt.Println("Equipped -", game_world.Player_info.Inventory[x].Equipped)
			fmt.Println("Item type -", game_world.Player_info.Inventory[x].Attributes.Item_type)
			fmt.Println("Attack -", game_world.Player_info.Inventory[x].Attributes.Attack)
			fmt.Println("Defense -", game_world.Player_info.Inventory[x].Attributes.Defense)
			fmt.Println("Special ability -",
						game_world.Player_info.Inventory[x].Attributes.Special_abilities)
		}
	} else {
		fmt.Println("Empty")
	}
}


func interact_with_world(game_world *world.World) *world.World {
	if check_collision(game_world) == "treasure" {
		player_x_pos := game_world.Player_info.X_position
		player_y_pos := game_world.Player_info.Y_position
		fmt.Println(
			"\nFound a new item",
			game_world.World_matrix[player_y_pos][player_x_pos].Treasure_info.Treasure_item.Item_name)
		item := game_world.World_matrix[player_y_pos][player_x_pos].Treasure_info.Treasure_item
		game_world.Player_info.Inventory = append(game_world.Player_info.Inventory, item)
		game_world.World_matrix[player_y_pos][player_x_pos].Treasure_info = world.Treasure{}
	} else {
		fmt.Println("\nNothing to interact with")
	}
	return game_world
}


func equip_item(game_world *world.World, equip_info []string) *world.World {
	var item_to_equip string
	if len(equip_info) > 1 {
		item_to_equip = strings.Join(equip_info[1:], " ")
	}

	inventory_length := len(game_world.Player_info.Inventory)
	for x := 0;x < inventory_length; x++ {
		if strings.ToLower(game_world.Player_info.Inventory[x].Item_name) ==
		   strings.ToLower(item_to_equip) {
			game_world.Player_info.Inventory[x].Equipped = true
			fmt.Println("\nYou equipped the", game_world.Player_info.Inventory[x].Item_name)
		}
	}
	return game_world
}


func unequip_item(game_world *world.World, unequip_info []string) *world.World {
	var item_to_unequip string
	if len(unequip_info) > 1 {
		item_to_unequip = strings.Join(unequip_info[1:], " ")
	}

	inventory_length := len(game_world.Player_info.Inventory)
	for x := 0;x < inventory_length; x++ {
		if strings.ToLower(game_world.Player_info.Inventory[x].Item_name) ==
		   strings.ToLower(item_to_unequip) {
			game_world.Player_info.Inventory[x].Equipped = false
			fmt.Println("\nYou unequipped the", game_world.Player_info.Inventory[x].Item_name)
		}
	}
	return game_world
}


func fight_monster(game_world *world.World) *world.World {
	if check_collision(game_world) == "monster" {
		game_world = enter_fight_mode(game_world)
	} else {
		fmt.Println("No monster here to fight")
	}

	return game_world
}


