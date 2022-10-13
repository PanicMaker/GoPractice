package main

import (
	"fmt"
	"testing"
)

func Test_sqlDb(t *testing.T) {
	db := sqlDb()
	todo := todoModel{
		Title:     "123",
		Completed: 0,
	}

	err := db.Save(&todo).Error
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", todo)
}
