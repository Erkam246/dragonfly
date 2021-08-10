package recipes

import "github.com/df-mc/dragonfly/server/item"

// Recipe is embedded by all recipe types.
type Recipe interface {
	__()
}

// ShapelessRecipe is a recipe that has no particular shape.
type ShapelessRecipe struct {
	Recipe
	// Inputs is a list of items that serve as the input of the shapeless recipe. These items are the items
	// required to craft the output. The amount of input items must be exactly equal to Width * Height.
	Inputs []Item
	// Output is an item that is created as a result of crafting the recipe.
	Output item.Stack
	// Priority ...
	Priority int32
}

// ShapedRecipe is a recipe that has a specific shape that must be used to craft the output of the recipe.
type ShapedRecipe struct {
	Recipe
	// Dimensions are the dimensions for the shaped recipe.
	Dimensions Dimensions
	// Inputs is a list of items that serve as the input of the shaped recipe. These items are the items
	// required to craft the output. The amount of input items must be exactly equal to Width * Height.
	Inputs []Item
	// Output is an item that is created as a result of crafting the recipe.
	Output item.Stack
	// Priority ...
	Priority int32
}
