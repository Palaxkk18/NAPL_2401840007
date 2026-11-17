package main

import (
	"errors"
	"fmt"
)

type Item struct {
	ID       int
	Name     string
	Quantity int
	Price    float64
}

var inventory []Item

func AddItem(item Item) error {
	if item.Quantity < 0 {
		return errors.New("quantity cannot be negative")
	}
	if item.Price <= 0 {
		return errors.New("price must be a positive number")
	}
	inventory = append(inventory, item)
	return nil
}
func UpdateItem(id int, updatedItem Item) error {
	for i, item := range inventory {
		if item.ID == id {
			if updatedItem.Quantity < 0 {
				return errors.New("quantity cannot be negative")
			}
			if updatedItem.Price <= 0 {
				return errors.New("price must be a positive number")
			}
			inventory[i] = updatedItem
			return nil
		}
	}
	return errors.New("item not found")
}
func RemoveItem(id int) error {
	for i, item := range inventory {
		if item.ID == id {

			inventory = append(inventory[:i], inventory[i+1:]...)
			return nil
		}
	}
	return errors.New("item not found")
}
func summmary(ch chan float64) {
	total := 0.0

	for _, item := range inventory {
		total += float64(item.Quantity) * item.Price
	}

	ch <- total
}
func TestAddItem() {
	err := AddItem(Item{ID: 1, Name: "Laptop", Quantity: 5, Price: 50000})
	if err != nil {
		fmt.Println("TestAddItem failed:", err)
	}
}
func TestUpdateItem() {
	err := UpdateItem(1, Item{ID: 1, Name: "Laptop", Quantity: 10, Price: 45000})
	if err != nil {
		fmt.Println("TestUpdateItem failed:", err)
	}
}
func TestRemoveItem() {
	err := RemoveItem(1)
	if err != nil {
		fmt.Println("TestRemoveItem failed:", err)
	}
}
func main() {
	displayInventory := func() {
		fmt.Println("Current Inventory:")
		for _, item := range inventory {
			fmt.Printf("ID: %d, Name: %s, Quantity: %d, Price: %.2f\n", item.ID, item.Name, item.Quantity, item.Price)
		}
	}
	TestAddItem()
	err := AddItem(Item{
		ID: 2, Name: "Smartphone",
		Quantity: 15, Price: 20000,
	})
	if err != nil {
		fmt.Println("Error adding item:", err)
	}
	displayInventory()
	updateInventory := func() {
		err := UpdateItem(1, Item{
			ID: 1, Name: "Laptop",
			Quantity: 10, Price: 45000,
		})
		if err != nil {
			fmt.Println("Error updating item:", err)
		}
		err = UpdateItem(2, Item{
			ID: 2, Name: "Smartphone",
			Quantity: 15, Price: 20000,
		})
		if err != nil {
			fmt.Println("Error updating item:", err)
		}
	}
	updateInventory()
	displayInventory()

	ch := make(chan float64)
	go summmary(ch)
	total := <-ch
	fmt.Printf("Total Inventory Value: %.2f\n", total)
	removeInventory := func() {
		err := RemoveItem(1)

		if err != nil {
			fmt.Println("Error removing item:", err)
		}
	}
	removeInventory()
	displayInventory()
}
