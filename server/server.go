package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	port := "8080"

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Server is listening on port " + port)

	conn1, err := listener.Accept()
	if err != nil {
		fmt.Println("Error accepting connection:", err)
		return
	}
	defer conn1.Close()
	fmt.Println("User 1 connected.")

	conn2, err := listener.Accept()
	if err != nil {
		fmt.Println("Error accepting connection:", err)
		return
	}
	defer conn2.Close()
	fmt.Println("User 2 connected.")

	questions := []string{
		"Samsun'un plakası nedir?\n",
		"Karabük'ün en ünlü ilçesi nedir?\n",
		"Dünyanın en büyük okyanusu nedir?\n",
		"64'ün karekökü nedir?\n",
		"Yer çekimini kim bulmuştur?\n",
	}

	answers := []string{
		"55",
		"Safranbolu",
		"pacific",
		"8",
		"Newton",
	}

	score1 := 0
	score2 := 0

	for i := 0; i < len(questions); i++ {
		question := questions[i]
		correctAnswer := answers[i]

		conn1.Write([]byte(question))
		var answer1 string
		fmt.Fscan(conn1, &answer1)
		answer1 = strings.ToLower(strings.TrimSpace(answer1))

		if answer1 == correctAnswer {
			score1++
			conn1.Write([]byte("Correct!\n"))
		} else {
			conn1.Write([]byte("Wrong!\n"))
		}

		conn2.Write([]byte(question))
		var answer2 string
		fmt.Fscan(conn2, &answer2)
		answer2 = strings.ToLower(strings.TrimSpace(answer2))

		if answer2 == correctAnswer {
			score2++
			conn2.Write([]byte("Correct!\n"))
		} else {
			conn2.Write([]byte("Wrong!\n"))
		}

		fmt.Println("Değer", i)

	}
	conn1.Write([]byte("skip"))
	conn2.Write([]byte("skip"))

	conn1.Write([]byte(fmt.Sprintf("User 1's score: %d\n", score1)))
	conn1.Write([]byte(fmt.Sprintf("User 2's score: %d\n", score2)))

	conn2.Write([]byte(fmt.Sprintf("User 1's score: %d\n", score1)))
	conn2.Write([]byte(fmt.Sprintf("User 2's score: %d\n", score2)))

	if score1 > score2 {
		conn1.Write([]byte("User 1 wins!\n"))
		conn2.Write([]byte("User 2 loses!\n"))
	} else if score2 > score1 {
		conn2.Write([]byte("User 2 wins!\n"))
		conn1.Write([]byte("User 1 loses!\n"))
	} else {
		conn1.Write([]byte("It's a draw!\n"))
		conn2.Write([]byte("It's a draw!\n"))
	}
}
