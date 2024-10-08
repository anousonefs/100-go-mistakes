package main

import (
	"13_util_packages/stringset"
	"13_util_packages/util"
	"fmt"
)

func main() {
	// util package
	stringSet := util.NewStringSet("a", "c", "z", "t")
	print(util.SortStringSet(stringSet))

	// stringset package
	set := stringset.New("c", "a", "b")
	fmt.Println(stringset.Sort(set))

	// stringset custom type package
	set2 := stringset.New2("c", "a", "b")
	fmt.Println(set2.Sort2())

	/* bear in mind that naming a package after what it provides
	and not what it contains can be an efficient way
	to increase its expressiveness. */
}
