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
		config.FilePath = args[0]

		// Redis host
		config.Host = args[1]

		// Redis Password
		config.Password = args[2]

		// Redis DB
		if _, err := strconv.ParseInt(args[3], 10, 64); err == nil {
			config.DB, _ = strconv.ParseInt(args[3], 10, 64)
		} else {
			color.Red("Bad argument for DB: " + args[3])
			color.Red("Usage: ./TXToRedis [FILE] [REDIS-HOST] [REDIS-PASSWORD] [REDIS-DB] [REDIS-KEY] [CONCURRENCY]")
			os.Exit(1)
		}

		// Redis Key
		config.Key = args[4]

		if _, err := strconv.ParseInt(args[5], 10, 64); err == nil {
			config.Concurrency, _ = strconv.ParseInt(args[5], 10, 64)
		} else {
			color.Red("Bad argument for concurrency: " + args[5])
			color.Red("Usage: ./TXToRedis [FILE] [REDIS-HOST] [REDIS-PASSWORD] [REDIS-DB] [REDIS-KEY] [CONCURRENCY]")
			os.Exit(1)
		}
	} else {
		color.Red("Only " + string(len(args)) + " specified.")
		color.Red("Usage: ./TXToRedis [FILE] [REDIS-HOST] [REDIS-PASSWORD] [REDIS-DB] [REDIS-KEY] [CONCURRENCY]")
		os.Exit(1)
	}
}
