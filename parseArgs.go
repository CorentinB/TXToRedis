package main

import (
	"os"
	"strconv"

	"github.com/fatih/color"
)

func parseArgs(args []string) {
	if len(args) > 6 {
		color.Red("Usage: ./TXToRedis [FILE] [REDIS-HOST] [REDIS-PASSWORD] [REDIS-DB] [REDIS-KEY] [CONCURRENCY]")
		os.Exit(1)
	} else if len(args) == 6 {
		// File path
		config.FilePath = args[1]

		// Redis host
		config.Host = args[2]

		// Redis Password
		config.Password = args[3]

		// Redis DB
		if _, err := strconv.ParseInt(args[4], 10, 64); err == nil {
			config.DB, _ = strconv.ParseInt(args[4], 10, 64)
		} else {
			color.Red("Usage: ./youtube-ig [FILE] [REDIS-HOST] [REDIS-PASSWORD] [REDIS-DB] [REDIS-KEY] [CONCURRENCY]")
			os.Exit(1)
		}

		// Redis Key
		config.Key = args[5]

		if _, err := strconv.ParseInt(args[6], 10, 64); err == nil {
			config.Concurrency, _ = strconv.ParseInt(args[6], 10, 64)
		} else {
			color.Red("Usage: ./youtube-ig [FILE] [REDIS-HOST] [REDIS-PASSWORD] [REDIS-DB] [REDIS-KEY] [CONCURRENCY]")
			os.Exit(1)
		}
	}
}
