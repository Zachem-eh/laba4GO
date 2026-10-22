package main

// models.go используетсся для предаставления структур данных товаров и заказов в их строковом виде.
import "fmt"

func (p Product) String() string {
	return fmt.Sprintf("ID товара: %d\nНазвание товара: %s\nОписание товара: %s\nЦена на товар: %v",
		p.ID, p.Name, p.Description, p.BasePrice)
}
func (o Order) String() string {
	result := "\n--- Чек заказа ---\n"
	for i, item := range o.Items {
		result += fmt.Sprintf("Товар № %d\nНазвание товара: %s\nЦена на товар: %.2f\n", i+1, item.Name, item.BasePrice)
	}
	result += fmt.Sprintf("\n---------------\nИтоговая сумма заказа: %.2f", o.FinalPrice)
	return result
}
