package main

import "fmt"

type Backpack struct {
	owner string
	items map[string]int
}

func (backpack *Backpack) addItems(itemName string, qty int) {
	_, found := backpack.items[itemName]
	if !found {
		backpack.items[itemName] = qty
	} else {
		backpack.items[itemName] += qty
	}

}

func main() {
	fmt.Println("*** receiver function demo ***")

	jack := Backpack{
		"Jack",
		make(map[string]int),
	}

	jack.addItems("pencil", 5)
	jack.addItems("laptop", 1)
	jack.addItems("pencil", 5)

	for key, val := range jack.items {
		fmt.Printf("item %s, qty %d\n", key, val)
	}
}
