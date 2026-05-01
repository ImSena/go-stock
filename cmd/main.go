package main

import (
	"fmt"
	"go-stock/internal/services"
)

func main() {
	fmt.Println("Sistema de Estoque")

	stock := services.NewStock()
	fmt.Println(stock)
}
