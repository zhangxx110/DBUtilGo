package main

import (
	"fmt"
	"dbutil"
	"log"
	"time"
)
var format string = "2006-01-02 15:04:05"
func testInsert() {
	var property = []string{"name", "ts"}
	ts := time.Now().Format(format)
	values := []interface{}{"zhangxx", ts}
	dbutil.Insert("t", property, values)
}
func testUpdate() {
	var property = []string{"name", "ts"}
	ts := time.Now()
	values := []interface{}{"testName", ts}
	dbutil.Update("t", property, values, "id", ">=", 16)
}
func testExcuteSql() {
	ts := time.Now()
	values := []interface{}{"xxx", ts}
	dbutil.ExcuteSql("update t set name = ? ,ts = ? where id >= 16", values)
}
func testDelete() {
	dbutil.Delete("t", "id<10")
}
func testQuery() {
	rows, _ := dbutil.Query("t", nil, "id > 10", "id desc", "")
	if rows == nil {
		return
	}
	for rows.Next() {
		var id int
		var name string
		var ts time.Time
		err := rows.Scan(&id, &name, &ts)
		if err != nil {
       	  fmt.Println(err)
    	}
		fmt.Println(id, name, ts.Format(format))
	}
}
func main() {
	err := dbutil.Init("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true")
	if err != nil {
		log.Println(err)
	}
	defer dbutil.Relase()
	testInsert()
	//testUpdate()
	//testExcuteSql()
	//testDelete()
	testQuery()
}
