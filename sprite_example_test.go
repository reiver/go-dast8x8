package dast8x8_test

import (
	"github.com/reiver/go-dast8x8"

	"github.com/reiver/go-imgstr"

	"fmt"
	"image"
)

func ExampleSprite() {

	// For the sake of this example, will use the following sprite ID.
	const spriteID = 1

	// Get the sprite.
	var img image.Image = dast8x8.Sprite(spriteID)

	// For the a
	serialized := imgstr.ImageString(img)

	fmt.Print(serialized)

	// Output:
	//
	// IMAGE:iVBORw0KGgoAAAANSUhEUgAAAAgAAAAICAYAAADED76LAAAAYElEQVR4nGJhgIJdDLz/YWw3hs+MMDYTuiQ6nwVJF8P///8ZGBkZkdUyMCLrYFvcCpdwiM0DyzHCJEEmIFkBpg3+f4ZYAXIUVkci2fYfitHZDCxoOuHGgxSBTAIEAAD//5urIvgu0on0AAAAAElFTkSuQmCC
}
