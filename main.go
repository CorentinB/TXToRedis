package main

import (
	"os"
	"time"

	"github.com/go-redis/redis"
	"github.com/labstack/gommon/color"
)

var config = struct {
	Host        string
	Password    string
	DB          int64
	Key         string
	Concurrency int64
	FilePath    string
	Client      *redis.Client
	Count       int
}{
	Count: 0}

func main() {
	start := time.Now()

	// Parse args
	parseArgs(os.Args[1:])

	// Process list
	processList()

	color.Println(color.Cyan("Done in ") + color.Yellow(time.Since(start)) + color.Cyan("!"))
}
