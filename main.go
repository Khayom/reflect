package main

import (
    "fmt"
    "reflect"
)

type Example struct {
    Name string
    Age  int
    Status bool
}

func main() {
    var one = Example{Name: "Cristiano", Age: 40, Status: true}
    Reflect(one)
}

func Reflect(i Example) {
    value := reflect.ValueOf(i)
    reflectType := value.Type()

    for i := 0; i < reflectType.NumField(); i++ {
        field := reflectType.Field(i)
        value := value.Field(i)
        fmt.Printf("Field Name: %s\nField Type: %s\nField Value: %v\n", field.Name, field.Type, value)
        fmt.Println("---------------------------")
    }
}