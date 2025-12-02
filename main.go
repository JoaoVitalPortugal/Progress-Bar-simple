package main

import (
	"fmt"
	"time"

	"github.com/schollz/progressbar/v3"
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

		bar := progressbar.New(total)
		for i := 0; i <= total; i++ {
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)

		}
		bar.Finish()
	}
}

