package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Jogo da Adivinhação")
	fmt.Println("Um número aleatório será sorteado. Tente acertar. O número é um inteiro entre 0 e 100")

	x := rand.Int64N(101)
	scanner := bufio.NewScanner(os.Stdin)
	attempts := [10]int64{}

	for i := range attempts {
		fmt.Println("Qual é o seu chute?")
		scanner.Scan()
		attempt := scanner.Text()
		attempt = strings.TrimSpace(attempt)
		attemptInt, err := strconv.ParseInt(attempt, 10, 64)
		if err != nil {
			fmt.Println("O valor deve ser um inteiro!")
			return
		}

		switch {
		case attemptInt < x:
			fmt.Println("Você errou. O número sorteado é maior que", attemptInt)
		case attemptInt > x:
			fmt.Println("Você errou. O número sorteado é menor que", attemptInt)
		case attemptInt == x:
			fmt.Printf(
				"Parabéns! Você acertou! O número era: %d\n"+
					"Você acertou em %d tentativas: \n"+
					"Essas foram suas tentativas: %v\n",
				x, i+1, attempts[:i])
			return
		}

		attempts[i] = attemptInt
	}

	fmt.Printf(
		"Infelizmente, você não acertou o número, que era: %d\n"+
			"Você teve 10 tentativas.\n"+
			"Essas foram suas tentativas: %v\n",
		x, attempts)
}
