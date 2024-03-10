package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "time"
)

func main() {
    today := time.Now().Format("2006-01-02")
    filename := today + ".txt"
    file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println("Erreur lors de l'ouverture du fichier:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(os.Stdin)

    fmt.Println("Bienvenue dans le chatbot. Tapez 'exit' pour quitter.")
    for {
        fmt.Print("Vous: ")
        scanner.Scan()
        input := scanner.Text()

        if strings.ToLower(input) == "exit" {
            break
        }

        message := fmt.Sprintf("[%s] %s\n", time.Now().Format("2006-01-02 15:04:05"), input)
      
        _, err := file.WriteString(message)
        if err != nil {
            fmt.Println("Erreur lors de l'écriture dans le fichier:", err)
            return
        }
    }

    fmt.Println("À bientôt!")
