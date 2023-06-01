package linq

import "errors"

var (
	ErrorMoreThanOneItemFound = errors.New("more than one item found")
	ErrorNoItemsFound         = errors.New("no items found")
)
