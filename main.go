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

func main() {
	reader := bufio.NewReader(os.Stdin)
	//all_products := []Product{}
	for {
		fmt.Println("Меню команд:")
		fmt.Println("1 - Добавить продукт в ассортимент")
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
		default:
			fmt.Println("Нет такого пункта меню!")
		}
		fmt.Println("Нажмите Enter для продолжения работы......")
		reader.ReadString('\n')
		clearScreen()
	}
}
