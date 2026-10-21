package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

type Product struct {
	Name        string
	Description string
	BasePrice   float64
}

// Добавлена структура Order, которая содержит список товаров и их итоговую цену
type Order struct {
	Items      []Product
	FinalPrice float64
}

func (pr Product) toString() string {
	return fmt.Sprintf("Название товара: %s\nОписание товара: %s\nЦена на товар: %v",
		pr.Name, pr.Description, pr.BasePrice)
}

func printProducts(products []Product) {
	fmt.Println("\n--- Ассортимент ---") // Использвование интерфейса Stringer
	for _, pr := range products {
		fmt.Println(pr.toString())
	}
	fmt.Println("---------------")
}

func addProduct(products *[]Product) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Введите название товара:")
		name, err1 := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		fmt.Println("Введите описание товара:")
		description, err2 := reader.ReadString('\n')
		description = strings.TrimSpace(description)

		fmt.Println("Введите цену товара:")
		input, err3 := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		price, err4 := strconv.ParseFloat(input, 64)

		if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
			fmt.Println("Ошибка ввода данных!")
			continue
		}
		*products = append(*products, Product{name, description, price})
		break
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	all_products := []Product{{"Колбаса", "Очень-очень вкусная", 167},
		{"Батарейки АА", "Максимально долгoвечные", 349},
		{"Вода 0,5 литра", "Черноголовка, негазированная", 19}}
	for {
		fmt.Println("Меню команд:")
		fmt.Println("1 - Вывести весь ассортимент")
		fmt.Println("2 - Добавить товар в ассортимент")
		fmt.Println("3 - Добавить модификаторы цены для товара")
		fmt.Println("0 - Закончить работу")
		fmt.Println("Выберите пункт из меню управления: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		a, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Некорректный ввод! Пожалуйста, введите число.")
			continue
		}

		switch a {
		case 0:
			return
		case 1:
			printProducts(all_products)
		case 2:
			addProduct(&all_products)
		case 3:
			clearScreen()
			fmt.Println("Спискок товаров:")
			for i := 0; i < len(all_products); i++ {
				fmt.Println("Товар №", i+1, "\nНазвание товара: ", all_products[i].Name, "\nЦена на товар: ", all_products[i].BasePrice)
			}
			fmt.Println("Выбирете товар для модификации цены: ")
			reader := bufio.NewReader(os.Stdin)
			input, err1 := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			productIndex, err2 := strconv.Atoi(input)
			if err1 != nil || err2 != nil {
				fmt.Println("Некорректный ввод!")
			}
			if productIndex < 1 || productIndex > len(all_products) {
				fmt.Println("Некорректный номер товара!")
				break
			}
			clearScreen()
			fmt.Println("Товар №", productIndex, "\nНазвание товара: ", all_products[productIndex-1].Name, "\nЦена на товар: ", all_products[productIndex-1].BasePrice)
			fmt.Println("Укажите процент скидки для товара: ")
			discount, err1 := reader.ReadString('\n')
			discount = strings.TrimSpace(discount)

			fmt.Println("Укажите процент налога для товара: ")
			Tax, err2 := reader.ReadString('\n')
			Tax = strings.TrimSpace(Tax)

			fmt.Println("Укажите до скольки знаков округлить: ")
			Rounding, err3 := reader.ReadString('\n')
			Rounding = strings.TrimSpace(Rounding)

			fmt.Println("Укажите минимальную цену для товара: ")
			MinimumPrice, err4 := reader.ReadString('\n')
			MinimumPrice = strings.TrimSpace(MinimumPrice)
			if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
				fmt.Println("Некорректный ввод!")
				break
			}
			discountValue, discountErr := strconv.ParseFloat(discount, 64)
			taxValue, taxErr := strconv.ParseFloat(Tax, 64)
			roundingValue, roundingErr := strconv.Atoi(Rounding)
			minimumPriceValue, minimumPriceErr := strconv.ParseFloat(MinimumPrice, 64)
			if discountErr != nil || taxErr != nil || roundingErr != nil || minimumPriceErr != nil || roundingValue <= 0 {
				fmt.Println("Некорректный ввод!")
				break
			}
			chain := CombineModifiers(
				WithDiscount(discountValue),
				WithTax(taxValue),
				WithRounding(roundingValue),
				WithMinimumPrice(minimumPriceValue),
			)
			fmt.Println("\n")
			all_products[productIndex-1].BasePrice = chain(all_products[productIndex-1].BasePrice)
			fmt.Printf("%s -> Новая цена: %.2f руб.\n", all_products[productIndex-1].Name, chain(all_products[productIndex-1].BasePrice))
		default:
			fmt.Println("Нет такого пункта меню!")
		}
		fmt.Println("Нажмите Enter для продолжения работы......")
		reader.ReadString('\n')
		clearScreen()
	}
}
