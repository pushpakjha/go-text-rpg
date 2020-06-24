package player

import (
	"fmt"
	"time"
	"math/rand"
	"github.com/pushpakjha/go-text-rpg/pkg/world"
)

func Spawn_player(game_world *world.World) world.World {
	fmt.Println("Spawning player")

	seed := rand.NewSource(time.Now().UnixNano())
	random_seed := rand.New(seed)

	x_position := random_seed.Intn(game_world.Max_x_size)
	y_position := random_seed.Intn(game_world.Max_y_size)

	game_world.World_matrix[y_position][x_position].Player = true

	return *game_world

}
