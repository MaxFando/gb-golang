package main

import (
	"fmt"
)

func example1() {
	var a uint32 = 1
	var b *uint32 = &a // взятие указателя на а. Тип указателя на тип T - *T
	fmt.Println("b = ", b)
	c := *b // взятие значение по указателю
	fmt.Println("c =", c)
}

func example2() {
	var a uint32 // заготавливаем переменную для считывания

	// fmt.Scan позволяет считать данные из стандартного потока ввода в переменную.
	// На вход требует указатель на эту переменную.
	_, err := fmt.Scan(&a) // считываем ввод в переменную
	if err != nil {
		fmt.Println(err)
		//os.Exit(1)
	}

	fmt.Println("your input is", a)
}

type ExampleStruct struct {
	field1 uint64
	field2 uint64
	field3 uint64
	field4 uint64
	field5 uint64
}

func exampleFunc(s *ExampleStruct) {

}

func example4() {
	var a *uint32 = nil
	fmt.Println("value of *a:", *a)
}

type User struct {
	Name    string
	Surname string
	Phone   string
}

// привязываем метод к типу *User
func (u *User) SetPhone(phone string) error {
	if err := validatePhone(phone); err != nil {
		return err
	}

	u.Phone = phone
	return nil
}

func (u User) GetFullName() string {
	return u.Name + " " + u.Surname
}

func validatePhone(phone string) error {
	return nil
}

func example5() {
	user := User{Name: "Иван", Surname: "Иванов"}
	if err := user.SetPhone("9999"); err != nil {
		panic(err)
	}

	fmt.Printf("user: %+v\n", user)
}

func example6() {
	user := User{Name: "test", Surname: "testing"}
	userFullName := user.GetFullName()

	fmt.Printf("userFullName: %s\n", userFullName)
}

//  все переменные, захватываемые замыканием, захватываются по указателю
func example7() {
	for i := 0; i < 5; i++ {
		defer func() {
			println(i)
		}()
	}
}

type Node struct {
	value    int64
	prevNode *Node
	nextNode *Node
}

func main() {
	example1()
	example2()
	example4()
	example5()
	example6()
}
