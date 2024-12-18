package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// Server'ı başlat
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 8080...")

	// Bağlantıları kabul et
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

	// Soruları ve doğru cevapları tanımla
	questions := []string{
		"What is the capital of France?\n",
		"Who is the president of the United States?\n",
		"What is the largest ocean on Earth?\n",
		"What is the square root of 64?\n",
		"Who developed the theory of relativity?\n",
	}

	answers := []string{
		"paris",
		"biden",
		"pacific",
		"8",
		"einstein",
	}

	// Puanları takip et
	score1 := 0
	score2 := 0

	// Her bir kullanıcıya aynı soruları sırayla sor
	for i := 0; i < len(questions); i++ {
		question := questions[i]
		correctAnswer := answers[i]

		// Kullanıcı 1'e soru sor ve cevabını al
		conn1.Write([]byte(question))
		var answer1 string
		fmt.Fscan(conn1, &answer1)
		answer1 = strings.ToLower(strings.TrimSpace(answer1))

		// Kullanıcı 1 cevabını kontrol et
		if answer1 == correctAnswer {
			score1++
			conn1.Write([]byte("Correct!\n"))
		} else {
			conn1.Write([]byte("Wrong!\n"))
		}

		// Kullanıcı 2'ye soru sor ve cevabını al
		conn2.Write([]byte(question))
		var answer2 string
		fmt.Fscan(conn2, &answer2)
		answer2 = strings.ToLower(strings.TrimSpace(answer2))

		// Kullanıcı 2 cevabını kontrol et
		if answer2 == correctAnswer {
			score2++
			conn2.Write([]byte("Correct!\n"))
		} else {
			conn2.Write([]byte("Wrong!\n"))
		}
	}

	// Sonuçları bildir
	conn1.Write([]byte(fmt.Sprintf("User 1's score: %d\n", score1)))
	conn2.Write([]byte(fmt.Sprintf("User 2's score: %d\n", score2)))

	// Kazananı belirle
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
