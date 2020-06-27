package world


import (
	"fmt"
	"time"
	"math/rand"
)


type Item_attributes struct {
	Item_type string
	Attack int
	Defense int
	Special_abilities string
}


type Item struct {
	Item_name string
	Attributes Item_attributes
	Equipped bool
}


type Treasure struct {
	Treasure_text string
	Treasure_item Item
}


func Spawn_treasure(game_world *World, treasure_number int) World {
	fmt.Println("Spawning treasure")

	for i := 0;i < treasure_number; i++ {
		seed := rand.NewSource(time.Now().UnixNano())
		random_seed := rand.New(seed)

		x_position := random_seed.Intn(game_world.Max_x_size)
		y_position := random_seed.Intn(game_world.Max_y_size)

		current_treasure := Treasure{
			Treasure_text: "A chest with a small weapons cache!",
			Treasure_item: Item{
				Item_name: "Bronze short sword",
				Equipped: false,
				Attributes: Item_attributes{
					Item_type: "Sword",
					Attack: 5,
					Defense: 0,
					Special_abilities: "None",
				},
			},
		}
		game_world.World_matrix[y_position][x_position].Treasure_info = current_treasure
	}

	return *game_world

}