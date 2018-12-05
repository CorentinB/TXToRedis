package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/labstack/gommon/color"
)

var checkPre = color.Yellow("[") + color.Green("✓") + color.Yellow("]") + color.Yellow("[")

func pushValue(value string, wg *sync.WaitGroup) {
	defer wg.Done()
	exist, err := config.Client.SIsMember(config.Key, value).Result()
	if exist == false {
		err = config.Client.SAdd(config.Key, value, 0).Err()
		if err != nil {
			log.Fatal(err)
		} else {
			config.Count++
			fmt.Println(checkPre + color.Cyan(config.Count) +
				color.Yellow("]") + checkPre +
				color.Cyan(value) + color.Yellow("]") +
				color.Green(" Added to the DB!"))
		}
	}
}
