package main

import (
	"bufio"
	"fmt"
	"go-stock/internal/models"
	"go-stock/internal/services"
	"os"
)

func main() {
	fmt.Println("Sistema de Estoque")
	var user string

	for {
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("Digite o nome do user: ")
		scanner.Scan()

		input := scanner.Text()

		if err := scanner.Err(); err != nil {
			fmt.Println("erro ao ler entrada")
			continue
		}

		user = input
		break

	}

	stock := services.NewStock()

	itens := []models.Item{
		{
			ID:       1,
			Name:     "FONE",
			Quantity: 3,
			Price:    10,
		},
		{
			ID:       2,
			Name:     "CAMISETA",
			Quantity: 3,
			Price:    10,
		},
		{
			ID:       3,
			Name:     "MACBOOK",
			Quantity: 3,
			Price:    5000,
		},
	}

	for _, item := range itens {
		err := stock.AddItem(item, user)

		if err != nil {
			fmt.Println(err)
			continue
		}
	}

	for _, item := range stock.ListItems() {
		fmt.Printf("%s", item.Info())
	}

	stock.RemoveItemQuantity(3, 3, "vendido", user)

	for _, item := range stock.ListItems() {
		fmt.Printf("%s", item.Info())
	}

	fmt.Printf("\n Valor total do estoque: %.2f \n", stock.CalculateTotalCost())

	logs := stock.ViewAuditLog()

	for _, log := range logs {
		fmt.Printf("[%s] Ação: %s - User: %s - Item ID: %d - Quantidade: %d - Motivo: %s \n",
			log.Timestamp.Format("01/02 15:04:05"),
			log.Action,
			log.User,
			log.ItemId,
			log.Quantity,
			log.Reason,
		)
	}

	itemSearch, err := services.FindBy(itens, func(item models.Item) bool {
		return item.Name == "CAMISETA"
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(itemSearch)

	corecode := services.Fornecedor{
		CNPJ:    "1234",
		Contato: "13424",
		Cidade:  "SP",
	}

	fmt.Println(corecode.GetInfo())
	if corecode.VerificarDisponibilidade(10, 15) {
		fmt.Println("Possui Disponibilidade")
	} else {
		fmt.Println("Não Possui Disponibilidade")
	}

}
