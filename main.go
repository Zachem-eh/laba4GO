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
	ID          uint
	Name        string
	Description string
	BasePrice   float64
}

// Добавлена структура Order, которая содержит список товаров и их итоговую цену
type Order struct {
	Items      []Product
	FinalPrice float64
}

func printProducts(products []Product) {
	for _, pr := range products {
		fmt.Println(pr)
		fmt.Println("---------------")
	}
}

var NEXT_ID uint = 4

func addProduct(reader *bufio.Reader, products *[]Product) {
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

		if err1 != nil || err2 != nil || err3 != nil {
			fmt.Println("Ошибка ввода данных!")
			return
		}
		if err4 != nil || price < 0 {
			fmt.Println("Некорректная цена!")
			continue
		}
		*products = append(*products, Product{
			ID:          NEXT_ID,
			Name:        name,
			Description: description,
			BasePrice:   price,
		})
		NEXT_ID++
		break
	}
}

func findProductByID(products []Product, id int) (int, error) {
	if id >= int(NEXT_ID) || id < 1 {
		return -1, fmt.Errorf("товар с ID %d не найден", id)
	}
	for i, p := range products {
		if int(p.ID) == id {
			return i, nil // успех — nil
		}
	}
	return -1, fmt.Errorf("товар с ID %d не найден", id)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	all_products := []Product{
		{ID: 1, Name: "Колбаса", Description: "Очень-очень вкусная", BasePrice: 167},
		{ID: 2, Name: "Батарейки АА", Description: "Максимально долговечные", BasePrice: 349},
		{ID: 3, Name: "Вода 0,5 литра", Description: "Черноголовка, негазированная", BasePrice: 19},
	}
	for {
		fmt.Println("Меню команд:")
		fmt.Println("1 - Вывести весь ассортимент")
		fmt.Println("2 - Добавить товар в ассортимент")
		fmt.Println("3 - Добавить модификаторы цены для товара")
		fmt.Println("4 - Оформить заказ")
		fmt.Println("0 - Закончить работу")
		fmt.Println("Выберите пункт из меню управления: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			return
		}
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
			fmt.Println("\n--- Ассортимент ---")
			printProducts(all_products)
		case 2:
			addProduct(reader, &all_products)
		case 3:
			clearScreen()
			fmt.Println("Спискок товаров:")
			printProducts(all_products)
			fmt.Println("Выбирете товар для модификации цены (введите ID): ")
			input, err1 := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			productId, err2 := strconv.Atoi(input)
			if err1 != nil || err2 != nil {
				fmt.Println("Некорректный ввод!")
				break
			}
			productIndex, err_ind := findProductByID(all_products, productId)
			if err_ind != nil {
				fmt.Println(err_ind)
				break
			}
			clearScreen()
			fmt.Println(all_products[productIndex].String() + "\n")
			fmt.Println("Укажите процент скидки для товара: ")
			discount, err1 := reader.ReadString('\n')
			discount = strings.TrimSpace(discount)

			fmt.Println("Укажите до скольки знаков округлить: ")
			Rounding, err3 := reader.ReadString('\n')
			Rounding = strings.TrimSpace(Rounding)

			fmt.Println("Укажите минимальную цену для товара: ")
			MinimumPrice, err4 := reader.ReadString('\n')
			MinimumPrice = strings.TrimSpace(MinimumPrice)
			if err1 != nil || err3 != nil || err4 != nil {
				fmt.Println("Некорректный ввод!")
				break
			}
			discountValue, discountErr := strconv.ParseFloat(discount, 64)
			roundingValue, roundingErr := strconv.Atoi(Rounding)
			minimumPriceValue, minimumPriceErr := strconv.ParseFloat(MinimumPrice, 64)
			if discountErr != nil || roundingErr != nil || minimumPriceErr != nil ||
				discountValue < 0 || discountValue > 100 || roundingValue < 0 {
				fmt.Println("Некорректный ввод!")
				break
			}
			// Налог здесь не применяется: цены в ассортименте хранятся без НДС,
			// НДС (vatPercent) добавляется при оформлении заказа в order.go.
			chain := CombineModifiers(
				WithDiscount(discountValue),
				WithRounding(roundingValue),
				WithMinimumPrice(minimumPriceValue),
			)
			fmt.Println()
			all_products[productIndex].BasePrice = chain(all_products[productIndex].BasePrice)
			fmt.Printf("%s -> Новая цена: %.2f руб.\n", all_products[productIndex].Name, all_products[productIndex].BasePrice)
		case 4:
			clearScreen()
			makeOrder(reader, all_products)
		default:
			fmt.Println("Нет такого пункта меню!")
		}
		fmt.Println("Нажмите Enter для продолжения работы......")
		reader.ReadString('\n')
		clearScreen()
	}
}
