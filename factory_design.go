package main

import "fmt"

type IProduct interface {
	setStock(int)
	getStock() int
	setName(string)
	getName() string
}

type Computer struct {
	name  string
	stock int
}

func (c *Computer) setStock(stock int) {
	c.stock = stock
}

func (c *Computer) getStock() int {
	return c.stock
}

func (c *Computer) setName(name string) {
	c.name = name
}

func (c *Computer) getName() string {
	return c.name
}

type Laptop struct {
	Computer
}

func NewLaptop() IProduct {
	return &Laptop{
		Computer{
			"Laptop",
			25,
		},
	}
}

type Desktop struct {
	Computer
}

func NewDesktop() IProduct {
	return &Desktop{
		Computer{
			"Desktop",
			35,
		},
	}
}

func GetComputerFactory(computerType string) (IProduct, error) {
	if computerType == "Laptop" {
		return NewLaptop(), nil
	} else if computerType == "Desktop" {
		return NewDesktop(), nil
	}
	return nil, fmt.Errorf("Invalid computer type passed")
}

func PrintNameAndStock(product IProduct) {
	fmt.Printf("Product Name: %s, with stock %d\n", product.getName(), product.getStock())
}

func main() {
	laptop, _ := GetComputerFactory("Laptop")
	desktop, _ := GetComputerFactory("Desktop")

	PrintNameAndStock(laptop)
	PrintNameAndStock(desktop)
}
