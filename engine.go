package main

import "fmt"

type sprite struct {
	x,y int
	coords [][]int
	structure [][]tile
	//i1 func()
	//i2 func()
	//i3 func()
	update func()
	anims []string
}

type tile struct {
	//x, y int
	icon string
	solid bool
	parent *sprite
}

//func (spr *sprite) scale(factor float64) {
//	spr.something *= factor
//}

func createEnvironment(x, y int, background string) [][]tile {
	env := make([][]tile, y) // allocates zeroed array and returns a slice referring to the array (for more malleable arrays)
	for i:=0; i<y; i++ {
		env[i] = make([]tile, x)
		for j:=0; j<x; j++ {
			env[i][j] = tile{background, false, nil} // enabled by make()
		}
	}
	return env
}

func renderEnvironment(env [][]tile) { //simple iterate through 2D list and output func
	for _, row := range env {
		for _, t := range row {
			fmt.Print(t.icon) 
		}
		fmt.Println()
	}
}

func placeSprite(s *sprite, env [][]tile) ([][]tile, error) { // ooooh this is gonna take some explaining
	if s.x < 0 || s.y < 0 || s.y >= len(env) || s.x >= len(env[0]) { // checks if the sprite's top left coordinates are out of bounds
		return nil, fmt.Errorf("sprite position out of bounds") // return error and nil
	}
	for _, coord := range s.coords { // iterates through stored sprite tile positions
		tileX := s.x + coord[0] // assigning x y coords
		tileY := s.y + coord[1]
		if tileX < 0 || tileY < 0 || tileY >= len(env) || tileX >= len(env[0]) { // if any of the coords contained in the sprite's hitbox are out of bounds
			return nil, fmt.Errorf("sprite coordinate %v , %v out of bounds", coord[1], coord[0])
		}
		if env[tileY][tileX].solid { // also if any tiles of the sprite intersect a solid tile
			return nil, fmt.Errorf("collision detected at %v , %v", tileX, tileY)
		}
		env[tileY][tileX] = s.structure[coord[1]][coord[0]] // after all checks passed then tile is written to LOCAL FUNCTION environment variable and NOT the public one (since checks can still fail)
	}
	return env, nil
}

// add collision func (intersection)
func collision(s1 *sprite, s2 *sprite) (bool, error) {
	for _, coord1 := range s1.coords { // iterate through sprite 1 coords
		x1 := s1.x + coord1[0]
		y1 := s1.y + coord1[1]
		for _, coord2 := range s2.coords { // iterate through sprite 2 coords for each sprite 1 coord
			x2 := s2.x + coord2[0]
			y2 := s2.y + coord2[1]
			if x1 == x2 && y1 == y2 { // if any coords intersect
				return true, nil
			}
		}
	}
	return false, nil
}

func adjacent(s1 *sprite, s2 *sprite) (bool, error) {
	directions := [][]int{
		{0, -1}, // up
		{0, 1},  // down
		{-1, 0}, // left
		{1, 0},  // right
	}
	for _, coord1 := range s1.coords {
		x1 := s1.x + coord1[0]
		y1 := s1.y + coord1[1]
		for _, coord2 := range s2.coords {
			x2 := s2.x + coord2[0]
			y2 := s2.y + coord2[1]
			for _, dir := range directions {
				
			}

		}
	}

/*func main() {
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
}*/
