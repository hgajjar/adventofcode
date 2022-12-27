package crt

import "fmt"

type pixel string

type CRT struct {
	pixels [6][40]pixel
	sprite [40]pixel
}

func New() CRT {
	return CRT{
		pixels: [6][40]pixel{},
		sprite: [40]pixel{"#", "#", "#"},
	}
}

func (c *CRT) DrawPixel(cycle, registerX int) {
	row := (cycle - 1) / 40
	column := (cycle - 1) % 40

	if c.isSpritePixel(column) {
		c.pixels[row][column] = "#"
	} else {
		c.pixels[row][column] = "."
	}

	c.updateSpritePosition(registerX)

	if cycle == (len(c.pixels) * len(c.pixels[1])) {
		c.flush()
	}
}

func (c *CRT) Print() {
	for _, row := range c.pixels {
		for _, pixel := range row {
			fmt.Print(pixel)
		}
		fmt.Print("\n")
	}
}

func (c *CRT) updateSpritePosition(registerX int) {
	for i := range c.sprite {
		if i >= registerX-1 && i < registerX+2 {
			c.sprite[i] = "#"
		} else {
			c.sprite[i] = "."
		}
	}
}

func (c *CRT) isSpritePixel(p int) bool {
	for i, sp := range c.sprite {
		if i == p && sp == "#" {
			return true
		}
	}

	return false
}

func (c *CRT) flush() {
	c.Print()
	c.pixels = [6][40]pixel{}
}