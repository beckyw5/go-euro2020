package main

import (
	"io"
	"log"
	"os"
	"time"

	"euro2020/euro2020"
)

func main() {
    now := time.Now()

    file, err := os.OpenFile("matches.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
    if err != nil {
        log.Fatalf("error opening the file: %v", err)
    }
    defer file.Close()

    wrt := io.MultiWriter(os.Stdout, file

    log.SetOutput(wrt)

    matches, err := euro2020.GetAllMatches()
    if err != nil {
        log.Fatalf("error while getting all matches: %v", err)
    }

    for _, match := range matches {
        log.Println("--------------------")
        log.Println("Home Team %s", match.homeTeam)
        log.Println("Away Team %s", match.awayTeam)
        log.Println("Winning Team %s", match.score.winner)
        log.Println("--------------------")
    }

    log.Printf("took %v", time.Now().Sub(now).String())
}