package services

import (
	"fmt"
	"go-stock/internal/models"
	"strconv"
	"time"
)

type Stock struct {
	items map[string]models.Item
	logs  []models.Log
}

func NewStock() *Stock {
	return &Stock{
		items: make(map[string]models.Item),
		logs:  []models.Log{},
	}
}

func (s *Stock) AddItem(item models.Item, user string) error {
	if item.Quantity <= 0 {
		return fmt.Errorf("Erro ao adicionar. Quantidade inválida")
	}

	existingItem, exists := s.items[strconv.Itoa(item.ID)]

	if exists {
		item.Quantity += existingItem.Quantity
	}

	s.items[strconv.Itoa(item.ID)] = item

	s.logs = append(s.logs, models.Log{
		Timestamp: time.Now(),
		Action:    "Entrada de Estoque",
		User:      user,
		ItemId:    item.ID,
		Quantity:  item.Quantity,
		Reason:    "Adicionando novos itens no Estoque",
	})

	return nil
}

func (s *Stock) RemoveItemQuantity(id int, quantity int, reason string, user string) error {
	existingItem, exists := s.items[strconv.Itoa(id)]

	if !exists {
		return fmt.Errorf("Erro ao remover item: [ID:%d]", id)
	}

	if quantity <= 0 {
		return fmt.Errorf("Erro ao remover item: quantidade menor que zero")
	}

	if existingItem.Quantity < quantity {
		return fmt.Errorf("Erro ao remover item: quantidade sobressae")
	}

	existingItem.Quantity -= quantity

	if existingItem.Quantity == 0 {
		delete(s.items, strconv.Itoa(id))
	} else {
		s.items[strconv.Itoa(id)] = existingItem
	}

	s.logs = append(s.logs, models.Log{
		Timestamp: time.Now(),
		Action:    "Saída de estoque",
		User:      user,
		ItemId:    id,
		Quantity:  quantity,
		Reason:    reason,
	})

	return nil

}

func (s *Stock) ListItems() []models.Item {
	var itemList []models.Item

	for _, item := range s.items {
		itemList = append(itemList, item)
	}

	return itemList
}

func (s *Stock) ViewAuditLog() []models.Log {
	return s.logs
}

func (s *Stock) CalculateTotalCost() float64 {
	var totalCost float64

	for _, item := range s.items {
		totalCost += float64(item.Quantity) * item.Price
	}

	return totalCost
}

func FindBy[T any](data []T, comparator func(T) bool) ([]T, error) {
	var result []T

	for _, v := range data {
		if comparator(v) {
			result = append(result, v)
		}
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("Nenhum item encontrado")
	}

	return result, nil
}
