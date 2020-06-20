package world

import (
	"fmt"
	"time"
	"math/rand"
)


type World struct {
	Max_X_size int
	Max_Y_size int
	WorldMatrix [][]string
}


func Generate_world() World {
	fmt.Println("Generating world")

	seed := rand.NewSource(time.Now().UnixNano())
	random_seed := rand.New(seed)

	max_x_size := random_seed.Intn(20) + 10
	max_y_size := random_seed.Intn(20) + 10
	var arrayofarrays [][]string

	return World {
		Max_X_size: max_x_size,
		Max_Y_size: max_y_size,
		WorldMatrix: arrayofarrays,
	}
}
