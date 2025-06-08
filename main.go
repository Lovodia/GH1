package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	gracefulShutdown()
}

func gracefulShutdown() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM) // SIGINT для остановки через ctr+c, SIGTERM для остановки в других, непредвиденных случаях

	timer := time.After(10 * time.Second)
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop() // требуется явная остановка time.NewTicker для того что бы не было утечки ресурсов но тут можно не использовать

	i := 1

	for {
		select {
		case <-timer:
			fmt.Println("time`s up")
			return
		case sig := <-sigChan:
			fmt.Println("Shutting down...:", sig)
			return
		case <-ticker.C:
			fmt.Println(i)
			i++
		}
	}
}
