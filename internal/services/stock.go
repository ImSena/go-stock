package services

import (
	"fmt"
	"go-stock/internal/models"
	"strconv"
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

func (s *Stock) AddItem(item models.Item) error {
	if item.Quantity <= 0 {
		return fmt.Errorf("Erro ao adicionar. Quantidade inválida")
	}

	existingItem, exists := s.items[strconv.Itoa(item.ID)]

	if exists {
		item.Quantity += existingItem.Quantity
	}

	s.items[strconv.Itoa(item.ID)] = item

	return nil
}
