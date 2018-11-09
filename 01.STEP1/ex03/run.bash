#!/bin/bash

# 引数無し
go run main.go

# 引数あり
go run main.go -n 7777 -s test

# 引数1つ
go run main.go -n 7777
go run main.go -s test

# 引数順番逆
go run main.go -s test -n 7777

