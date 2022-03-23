package main

import "fmt"

type IProduct interface {
	setStock(stock int)
	getStock() int
	setName(name string)
	getName() string
}

type Computer struct {
	name  string
	stock int
	iType string
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

type Desktop struct {
	Computer
}

func newLaptop() IProduct {
	return &Laptop{
		Computer: Computer{
			name:  "Laptop Lenovo H100",
			stock: 10,
			iType: "Laptop",
		},
	}
}

func newDesktop() IProduct {
	return &Desktop{
		Computer: Computer{
			name:  "Desktop Mac",
			stock: 5,
			iType: "Desktop",
		},
	}
}

func GetComputerFactory(compputerType string) (IProduct, error) {
	if compputerType == "laptop" {
		return newLaptop(), nil
	}
	if compputerType == "desktop" {
		return newDesktop(), nil
	}
	return nil, fmt.Errorf("Invalid product type")
}

func printIPproductValues(ip IProduct) {
	fmt.Printf("Product name: %s, stock: %d\n", ip.getName(), ip.getStock())
}

func main() {
	laptop, _ := GetComputerFactory("laptop")
	printIPproductValues(laptop)

	desktop, _ := GetComputerFactory("desktop")
	printIPproductValues(desktop)

	_, err := GetComputerFactory("termo")

	if err != nil {
		fmt.Println(err)
	}
}
