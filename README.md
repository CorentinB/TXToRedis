# TXToRedis
Push each line from a text file, to a Redis set

# Build

Simply:

`git clone https://github.com/CorentinB/TXToRedis.git && cd TXToRedis/`

`go build .`

# Usage

`./TXToRedis [FILE] [REDIS-HOST] [REDIS-PASSWORD] [REDIS-DB] [REDIS-KEY] [CONCURRENCY]`

For example to push each line from the file `data.txt` to a set with the `ids` value, in database `0` of a Redis instance running at `127.0.0.1:6379`, secured with the password `not-secured`, with a concurrency of `512` parallel push:

`./TXToRedis data.txt "127.0.0.1:6379" "not-secured" "0" "ids" "512"`

---

# Isn't that beautiful?

![Example](https://media.giphy.com/media/7zuMe0bNOHgRAd9Tw4/giphy.gif)