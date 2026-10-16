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

type PriceModifier func(float64) float64

type Product struct {
	Name        string
	Description string
	BasePrice   float64
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
		price, err3 := strconv.ParseFloat(input, 64)

		if err1 != nil || err2 != nil || err3 != nil {
			fmt.Println("Ошибка ввода данных!")
			continue
		}
		*products = append(*products, Product{name, description, price})
		break
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	all_products := []Product{}
	for {
		fmt.Println("Меню команд:")
		fmt.Println("1 - Вывести весь ассортимент")
		fmt.Println("2 - Добавить товар в ассортимент")
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
			//
		case 2:
			addProduct(&all_products)
		default:
			fmt.Println("Нет такого пункта меню!")
		}
		fmt.Println("Нажмите Enter для продолжения работы......")
		reader.ReadString('\n')
		clearScreen()
	}
}
