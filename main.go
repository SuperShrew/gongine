package main

import "fmt"

type sprite struct {
	x,y int
	coords [][]int
	structure [][]tile
	i1 func()
	i2 func()
	i3 func()
	update func()
	anims []string
}

type tile struct {
	//x, y int
	icon string
	solid bool
	parent *sprite
}

func createEnvironment(x, y int, background string) [][]tile {
	env := make([][]tile, y)
	for i:=0; i<y; i++ {
		env[i] = make([]tile, x)
		for j:=0; j<x; j++ {
			env[i][j] = tile{background, false, nil}
		}
	}
	return env
}

func renderEnvironment(env [][]tile) {
	for _, row := range env {
		for _, t := range row {
			fmt.Print(t.icon)
		}
		fmt.Println()
	}
}

func placeSprite(s *sprite, env [][]tile) ([][]tile, error){
	if s.x < 0 || s.y < 0 || s.y >= len(env) || s.x >= len(env[0]) {
		return nil, fmt.Errorf("Sprite position out of bounds")
	}
	for _, coord := range s.coords {
		tileX := s.x + coord[0]
		tileY := s.y + coord[1]
		if tileX < 0 || tileY < 0 || tileY >= len(env) || tileX >= len(env[0]) {
			return nil, fmt.Errorf("Sprite coordinate %v , %v out of bounds", coord[1], coord[0])
		}
		if env[tileY][tileX].solid {
			return nil, fmt.Errorf("Collision detected at %v , %v", tileX, tileY)
		}
		env[tileY][tileX] = s.structure[coord[1]][coord[0]]
	}
	return env, nil
}

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
