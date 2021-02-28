package main

import (
	"mekar/domain"

	logs "github.com/MaulIbra/logs_module")


func main() {
	logs.InitLog("log")
	domain.Init()
}