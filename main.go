package main

import (
	"fmt"
	"time"

	"github.com/houndgo/suuid"
	"github.com/theplant/batchputs"
)

func main() {
	db := openAndMigrate()
	rows := [][]interface{}{}
	for i := 0; i < 300; i++ {
		ti := time.Now()
		uid := suuid.New().String()
		rows = append(rows, []interface{}{
			fmt.Sprintf("sid_%d", i),
			fmt.Sprintf("name %d", i),
			i,
			ti,
			uid,
		})
	}
	fmt.Println(rows)
	columns := []string{"s_id", "name", "count", "created_at", "uid"}
	dialect := "mysql"

	start := time.Now()
	err := batchputs.Put(db.DB(), dialect, "countries", "s_id", columns, rows)
	if err != nil {
		panic(err)
	}
	duration := time.Since(start)
	fmt.Println("Inserts 30000 records using less than 3 seconds:", duration.Seconds() < 3)

}
