package main

// models.go используетсся для предаставления структур данных товаров и заказов в их строковом виде.
import "fmt"

func (p Product) String() string {
	return "Товар: " + p.Name + "\nОписание товара: " + p.Description + "\nЦена на товар: " + fmt.Sprint("%v", p.BasePrice)
}
func (o Order) String() string {
	result := "\n--- Чек заказа ---\n"
	for i, item := range o.Items {
		result += fmt.Sprintf("Товар № %d\nНазвание товара: %s\nЦена на товар: %v\n", i+1, item.Name, item.BasePrice)
	}
	result += fmt.Sprintf("\n---------------\nИтоговая сумма заказа: %v", o.FinalPrice)
	return result
}
