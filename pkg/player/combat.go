package player

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"github.com/pushpakjha/go-text-rpg/pkg/world"
)


func enter_fight_mode(game_world *world.World) *world.World {
	fmt.Println("You entered combat, type help for combat instructions")
	display_player_info(game_world)

	player_x_pos := game_world.Player_info.X_position
	player_y_pos := game_world.Player_info.Y_position

	current_monster := game_world.World_matrix[player_y_pos][player_x_pos].Monster_info
	print_monster_info(&current_monster)

	fight_ended := false
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		game_world, fight_ended = evaulate_fight_action(game_world, text)
		if fight_ended == true {
			break
		}
	}

	return game_world
}


func evaulate_fight_action(game_world *world.World, action string) (*world.World, bool) {
	fight_ended := false

	switch action {
	case "help":
		print_combat_help()
		break
	case "attack":
		game_world, fight_ended = attack_monster(game_world)
		break
	case "defend":
		fmt.Println("Not implemented")
		break
	case "run":
		fmt.Println("Not implemented")
		break
	case "info":
		display_player_info(game_world)
		break
	}

	return game_world, fight_ended
}


func print_combat_help() {
	available_commands :=
		"\nIn combat the available commands are:\n" +
		"help - Displays this help text\n" +
		"attack - Attack the monster, uses all equipped weapons\n" +
		"defend - Defend against the monster, gives extra defense for 1 turn" +
		"run - Try to run away, may not work\n" +
		"info - List your current player info and stats\n"
	fmt.Println(available_commands)
}


func attack_monster(game_world *world.World) (*world.World, bool) {
	fight_ended := false
	player_x_pos := game_world.Player_info.X_position
	player_y_pos := game_world.Player_info.Y_position

	current_monster := game_world.World_matrix[player_y_pos][player_x_pos].Monster_info

	var player_damage int
	if game_world.Player_info.Attack - current_monster.Attributes.Defense > 0 {
		player_damage = game_world.Player_info.Attack - current_monster.Attributes.Defense
	}
	current_monster.Attributes.Health = current_monster.Attributes.Health - player_damage
	game_world.World_matrix[player_y_pos][player_x_pos].Monster_info = current_monster

	var monster_damage int
	if current_monster.Attributes.Attack - game_world.Player_info.Defense > 0 {
		monster_damage = current_monster.Attributes.Attack - game_world.Player_info.Defense
	}
	game_world.Player_info.Health = game_world.Player_info.Health - monster_damage

	if game_world.Player_info.Health <= 0 {
		end_game()
	}
	if current_monster.Attributes.Health <= 0 {
		fight_ended = true
	}

	print_monster_info(&current_monster)

	return game_world, fight_ended
}


func print_monster_info(current_monster *world.Monster) {
	fmt.Println("\nMonster name:", current_monster.Monster_type)
	fmt.Println("Level -", current_monster.Attributes.Level)
	fmt.Println("Health -", current_monster.Attributes.Health)
	fmt.Println("Attack -", current_monster.Attributes.Attack)
	fmt.Println("Defense -", current_monster.Attributes.Defense)
	fmt.Println("Special abilities -", current_monster.Attributes.Special_abilities)
}


func end_game() {
	fmt.Println("You died! Game over!")
	os.Exit(0)
}