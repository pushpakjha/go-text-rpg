package world


import (
	"fmt"
	"time"
	"math/rand"
)


type Monster_attributes struct {
	Level int
	Health int
	Attack int
	Defense int
	Special_abilities string
}


type Monster_info struct {
	Monster_type string
	Attributes Monster_attributes

}


func Spawn_monsters(game_world *World, monster_number int) World {
	fmt.Println("Spawning monsters")

	for i := 0;i < monster_number; i++ {
		seed := rand.NewSource(time.Now().UnixNano())
		random_seed := rand.New(seed)

		x_position := random_seed.Intn(game_world.Max_x_size)
		y_position := random_seed.Intn(game_world.Max_y_size)

		monster_level := random_seed.Intn(2) + 1

		current_monster := Monster_info{
			Monster_type: "Wolf",
			Attributes: Monster_attributes{
				Level: monster_level,
				Health: monster_level * 10,
				Attack: monster_level * 3,
				Defense: monster_level * 2,
				Special_abilities: "None",
			},
		}
		game_world.World_matrix[y_position][x_position].Monster = current_monster
	}

	return *game_world

}