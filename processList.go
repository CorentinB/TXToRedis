package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/go-redis/redis"
)

func processList() {
	var count int64
	var wg sync.WaitGroup

	// Connect to Redis server
	config.Client = redis.NewClient(&redis.Options{
		Addr:     config.Host,
		Password: config.Password,
		DB:       int(config.DB),
	})

	// Check connection
	pong, err := config.Client.Ping().Result()
	if pong != "PONG" {
		fmt.Println("Unable to connect to Redis DB.")
		log.Fatal(err)
		os.Exit(1)
	}

	// Open file
	file, err := os.Open(config.FilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Scan the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		count++
		wg.Add(1)
		go pushValue(scanner.Text(), &wg)
		if count == config.Concurrency {
			wg.Wait()
			count = 0
		}
	}

	// Log if error
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	wg.Wait()
}
