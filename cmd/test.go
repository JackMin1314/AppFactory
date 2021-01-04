package main

import (
	excel "AppFactory/internal/app/HandleExcel"
	log "AppFactory/pkg"
	"fmt"
)

func main() {
	datasAll := [][]string{
		{"7002",	"韩翠萍",	"7班",	"119",	"138",	"101",	"79",	"82",	"519",	"1",	"1",	"1"},
		{"7003",	"梁欣欣",	"9班",	"114",	"146",	"112",	"72",	"73",	"517",	"1",	"2",	"1"},
	}
	fmt.Println(datasAll[0])
	fmt.Println(len(datasAll[0]))
	log.InitLogger()
	logger := log.GetLogInstance()
	excel.WriteDataToTable(logger,datasAll)

}
