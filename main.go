package main

import (
	"fmt"
	example "github.com/yangyang5214/protoc-gen-gorm-serializer-example/example"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Name   string
	Status example.TaskStatus `gorm:"type:int"`
}

func (Task) TableName() string {
	return "task"
}

func main() {
	db, err := gorm.Open(sqlite.Open("/tmp/test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&Task{})
	if err != nil {
		panic(err)
	}

	t := &Task{
		Name:   "task_name_test",
		Status: example.TaskStatus_Exiting,
	}

	// Create
	tx := db.Create(t)
	if tx.Error != nil {
		panic(tx.Error)
	}

	// Read
	var task Task
	db.Where(Task{Name: "task_name_test"})
	db.First(&task)

	fmt.Println(fmt.Sprintf("task status is: %s", task.Status))
}
