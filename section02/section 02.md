# Variables and Functions 

```go
go run main.go
```

There's no variable inside function without using it 

func main does not take any arguments. 

# Pointers

All a pointer is a pointer is it points to a specific location in a memory and gives you means of getting that particular location in memory 

`*` - for actual pointer 

`&` - for reference 

```go
func changeUsingPointer(s *string) {
	log.Println("s is set to", s) // memory address
	newValue := "Red"
	*s = newValue
}
```

# Types and Structs 

 An explanation about scopes. Variable shadowing.

Type can be use as a variable. 

He puts variable inside this Type with capital letters because he might need this type to be available to another package

Go is not an OOP language

If I declare func like this, it's only available within this package

```go
func whatever() {

}
```

But if I declare it with a capital letter, it's visible outside of that package 

```go
func Whatever() {

}
```

# Receiver: Structs with functions 

```go
type myStruct struct {
	FirstName string
}

// assign a function to this struct
func (m *myStruct) printFirstName() string {
	return m.FirstName
}
```

You can perform some business logic and that business logic actually become a part of this struct. 

# Maps and Slices

Map is extremely useful because it's very fast. Map is immutable. 

I might want to pass that value, but I might want to pass a pointer to that value because I want to make sure when I pass that variable around to another function or another package that I'm actually passing a pointer to it.

So when it's changed in the other package or changed in another function, I know it's going to affect the one that I created here, but you don't have to do that with maps because maps stay the same no matter where they go.

You never have to bother pointing a passing a pointer to a map.

You can just pass the map itself and that map will remain constant no matter where in the program it is accessed.

That's extremely useful.

> 1. Передача структур:
>
> - Когда вы передаете структуру как значение, создается её копия
> - Изменения в копии не влияют на оригинальную структуру
> - Чтобы изменения отражались в оригинале, нужно передавать указатель
>
> 2. Передача map:
>
> - Map в Go передается по ссылке по умолчанию
> - При передаче map в другую функцию или пакет, вы работаете с одним и тем же объектом
> - Изменения в map будут видны во всех местах, где она используется
> - Не требуется явно использовать указатели

```go
// Структура - требует указателя для изменений
type Person struct {
    Name string
}

// Map - изменяется без указателя
users := map[string]string{
    "key": "value"
}
```

Map is programmatically built into the system, not sorted. You always must look up by key. Can't be sorted. 

# Interfaces 

In order for something to implement an interface, it must implement the same function as the interface in question. 

> Интерфейсы в Go - это более гибкий и легкий механизм абстракции, отличающийся от классического концепта классов в объектно-ориентированном программировании.
>
> В Go интерфейсы являются фундаментальным механизмом абстракции и полиморфизма.
>
> Интерфейс в Go - это тип, который определяет набор методов. 

# Packages 

Create your own module

```go
go mod init github.com/myusername/packagename 
```

Conventionally you start your name things by prefixing it with the name of the git repository

That command create `go.mod` file 

Enable Go modules integration - activate checkbox in GoLand settings  

?? Go: Infer Gopath in VS Code

# Channels 

A channel - a place to send information which will be received in one or more places in my program. 

# Reading and writing JSON 

The functions we used in the json package in golang are called **`Marshal`** and **`Unmarshal`**. 

**`Unmarshal`** has 2 parameters. One is a slice of bytes 

**`MarshalIndent`** hands back a slice of bytes and potentially an error. It's more readable. Use it for development purposes. And we need to convert slice of bytes to a string. 

**`Marshal`** wouldn't have any nice tab or line breaks but it will be used in production.  

> В Go, функция `json.Unmarshal()` требует указатель на переменную, в которую будут записаны распакованные данные. Это связано с тем, что функция должна модифицировать сам объект, а не его копию.
>
> 1. Сигнатура метода `Unmarshal()` требует указатель:
> 2. Без `&` вы передаете копию среза, а не ссылку на него. Это означает, что распакованные данные не будут сохранены в исходной переменной.
>
> Использование `&` позволяет методу `Unmarshal()` напрямую изменять структуру данных, передавая ей адрес переменной, а не её копию. Это стандартный подход в Go для методов, которые должны модифицировать входные данные.

# Writing tests in GO 

main_test.go - file name 

function must start with the word "Test"

```cmd
go test
go test - v
go test -cover

go test -coverprofile=coverage
go tool cover -html=coverage.out
```

