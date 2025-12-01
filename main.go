package main

import (
	"fmt"
	"time"
)

func main() {
	var resposta1 int
	fmt.Println(`Faça uma escolha:
	1- Sair
	2- Veja a o progresso
	`)
	fmt.Scan(&resposta1)
	if resposta1 == 1 {
		fmt.Println("Saindo...")
		return
	} else if resposta1 == 2 {

		total := 50

		for i := 0; i <= total; i++ {
			percent := (i * 100) / total
			bar := ""

			for j := 0; j < i; j++ {
				bar += "="
			}
			for j := i; j < total; j++ {
				bar += " "
			}

			fmt.Printf("\r[%s] %d--", bar, percent)

			time.Sleep(60 * time.Millisecond)
		}
	}
}
