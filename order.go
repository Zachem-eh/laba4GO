package main

// order.go содержит работу с заказами: модификатор цены с ограниченным числом
// применений, промокоды, создание заказа из товаров ассортимента, расчёт
// итоговой цены и диалог оформления заказа.

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

const vatPercent = 20

func promoCodes() map[string]PriceModifier {
	return map[string]PriceModifier{
		"SALE50":  WithLimitedDiscount(50, 2), // −50 % на первые два товара заказа
		"FIRST30": WithLimitedDiscount(30, 1), // −30 % на первый товар заказа
		"MINUS10": WithDiscount(10),           // −10 % на все товары
	}
}

func WithLimitedDiscount(percent float64, limit int) PriceModifier {
	used := 0
	discount := WithDiscount(percent)
	return func(price float64) float64 {
		if used < limit {
			used++
			return discount(price)
		}
		return price
	}
}

func NewOrder(products []Product, ids []int) (Order, error) {
	var order Order
	for _, id := range ids {
		index, err := findProductByID(products, id)
		if err != nil {
			return Order{}, err
		}
		order.Items = append(order.Items, products[index])
	}
	return order, nil
}

func (o *Order) Calculate(modifier PriceModifier) {
	o.FinalPrice = 0
	for _, item := range o.Items {
		o.FinalPrice += modifier(item.BasePrice)
	}
}

func readLine(reader *bufio.Reader, prompt string) (string, error) {
	fmt.Println(prompt)
	line, err := reader.ReadString('\n')
	return strings.TrimSpace(line), err
}

func makeOrder(reader *bufio.Reader, products []Product) {
	// parseIDs разбирает строку вида "1 3 3" в срез ID.
	parseIDs := func(s string) ([]int, error) {
		fields := strings.Fields(s)
		if len(fields) == 0 {
			return nil, fmt.Errorf("не введено ни одного ID")
		}
		ids := make([]int, 0, len(fields))
		for _, field := range fields {
			id, err := strconv.Atoi(field)
			if err != nil {
				return nil, fmt.Errorf("%q не является числом", field)
			}
			ids = append(ids, id)
		}
		return ids, nil
	}

	printProducts(products)
	input, err := readLine(reader, "Введите ID товаров через пробел (например: 1 3 3):")
	if err != nil {
		fmt.Println("Ошибка ввода!")
		return
	}
	ids, err := parseIDs(input)
	if err != nil {
		fmt.Println("Некорректный ввод:", err)
		return
	}
	order, err := NewOrder(products, ids)
	if err != nil {
		fmt.Println(err)
		return
	}

	code, err := readLine(reader, "Введите промокод (или просто нажмите Enter, если его нет):")
	if err != nil {
		fmt.Println("Ошибка ввода!")
		return
	}

	var modifiers []PriceModifier
	if code != "" {
		promo, ok := promoCodes()[strings.ToUpper(code)]
		if ok {
			modifiers = append(modifiers, promo)
		} else {
			fmt.Println("Промокод не найден, заказ будет оформлен без скидки.")
		}
	}
	modifiers = append(modifiers, WithTax(vatPercent), WithRounding(2))

	chain := CombineModifiers(modifiers...)
	order.Calculate(chain)
	fmt.Println(order)
	fmt.Printf("(цены указаны без НДС, НДС %d %% включён в итоговую сумму)\n", vatPercent)
}
