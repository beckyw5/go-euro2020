package main

func main() {
    now := time.Now()

    log.Printf("took %v", time.Now().Sub(now))
}