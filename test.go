package main

import "fmt"

func main() {
	var test_env = createEnvironment(6, 6, ".")
	renderEnvironment(test_env)
	fmt.Println()
	test_sprite := sprite{
		x: 2,
		y: 3,
		coords: [][]int{{0,0}, {1,0}, {0,1}, {1,1}},
		structure: [][]tile{
			{tile{"#", false, nil}, tile{"#", false, nil}},
			{tile{"#", false, nil}, tile{"#", false, nil}},
		},
	}
	updated_env, err := placeSprite(&test_sprite, test_env)
	if err != nil {
		fmt.Println("Error placing sprite:", err)
	} else {
		renderEnvironment(updated_env)
	}
}
