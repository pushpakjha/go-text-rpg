package engine

import (
	"fmt"
	"github.com/pushpakjha/go-text-rpg/pkg/world"
)

func Run_game() {
	initial_world := world.Generate_world()
	fmt.Println(initial_world.Max_X_size)
	fmt.Println(initial_world.Max_Y_size)
}
