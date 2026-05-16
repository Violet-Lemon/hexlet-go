package main

import (
	"fmt"
	"github.com/Violet-Lemon/hexlet-go/v2/greeting"
	"github.com/fatih/color"
	"strings"
	//em "github.com/Violet-Lemon/hexlet-go/v2/employees"
	//pr "github.com/Violet-Lemon/hexlet-go/v2/payroll"
	"github.com/Violet-Lemon/hexlet-go/v2/orders"
	"github.com/Violet-Lemon/hexlet-go/v2/accounts"
)

/*type Movie struct {
	Title string
	Year int
	Genres []string
	Rating float64
}

type Account struct {
	Id int
	Owner string
	Balance float64
}

func Deposit(acc *Account, amount float64) {
	if amount <= 0 {
		return
	}

	fmt.Println("Баланс до: ", acc.Balance)

	acc.Balance += amount

	fmt.Println("Баланс после: ", acc.Balance)
	}*/

func main() {
	color.Cyan(greeting.Hello())
	color.Blue("I'm coloured text.")
	color.Magenta("Good luck, dear friend!")
	fmt.Println(greeting.Hello())

	/*movie1 := Movie{
		Title: "Harry Potter",
	 	Year: 2000,
		Genres: []string{"fantastic", "magic"},
	 	Rating: 100.0,
	}

	fmt.Print(movie1.Title, "(released in ", movie1.Year, ") - ", movie1.Genres, ". Rate ", movie1.Rating, "\n")

	acc1 := Account{
		Id: 1,
		Owner: "Natasha",
		Balance: 0.0,
	}

	Deposit(&acc1, 1000000.0)*/

	/*team := []em.Employee{
		{
			Name: "Bob",
			Position: "boss",
			Skills: []string{"has fun", "eats bananas", "has the little bear"},
		},
		{
			Name: "Kevin",
			Position: "musician",
			Skills: []string{"has fun", "eats bananas", "plays a guitar"},
		},
		{
			Name: "Stuart",
			Position: "the most clever",
			Skills: []string{"reads books", "eats bananas", "looks after friends"},
		},
	}

	if !team[0].SetBaseSalary(100) {
		fmt.Println("некорректная базовая ставка для %s", team[0])
	}

	if !team[1].SetBaseSalary(20) {
		fmt.Println("некорректная базовая ставка для %s", team[1])
	}

	if !team[2].SetBaseSalary(50) {
		fmt.Println("некорректная базовая ставка для %s", team[2])
	}

	fmt.Println("все нормально")

	bonusPct := 10.0

	if bonusPct < 0 {
		fmt.Println("предупреждение: отрицательный бонус, устанавливаем 0%")
		bonusPct = 0
	}

	taxPct := 13.0

	if taxPct < 0 || taxPct > 100 {
		fmt.Println("предупреждение: некорректный налог, устанавливаем 13%")
		taxPct = 13
	}

	for _, e := range team {
		gross := pr.CalcGross(e, bonusPct)
		fmt.Printf("Грязная зарплата %s (%s): %.2f\n", e.Name, e.Position, gross)
		fmt.Printf("Чистая зарплата %s (%s): %.2f\n", e.Name, e.Position, pr.CalcNet(gross, taxPct))
		}*/

	order1 := orders.Order{
		ID: 1,
		Customer: "Plushka",
		Status: "new",
	}

	fmt.Println("новый заказ", order1.Status)

	orders.SetStatusCopy(order1, "paid")

	fmt.Println("применить для копии", order1.Status)

	orders.SetStatusPtr(&order1, "paid")

	fmt.Println("применить для заказа", order1.Status)

	a := 1
	fmt.Println(Inc(&a), a)
	fmt.Println(Inc(&a), a)
	fmt.Println(Inc(nil))

	account1 := accounts.Account{
		ID: 1,
		Owner: "Buba",
	}

	fmt.Println("первоначальный баланс", account1.Balance())

	account1.Deposit(-100)

	fmt.Println("отрицательный депозит", account1.Balance())

	account1.Deposit(100)
	account1.Deposit(20)
	account1.Deposit(3)

	fmt.Println("пополнение", account1.Balance())

	books := []Book{
		{Title: "Евгений Онегин", Author: "Пушкин А.С."},
		{Title: "Мастер и Маргарита", Author: "Булгаков М.А."},
		{Title: "Сказка о Царе Солтане", Author: "Пушкин А.С."},
	}

	titles := make([]string, 0, len(books))
	for _, b := range books {
		titles = append(titles, b.Title)
	}
	fmt.Println(strings.Join(titles, ", "))

	author := "Пушкин А.С."
	count := 0
	for _, b := range books {
		if b.Author == author {
			count++
		}
	}
	fmt.Printf("Книг автора %s: %d\n", author, count)
}

func Inc(p *int) bool {
	if p == nil {
		return false
	}

	*p++

	return true
}

type Book struct {
	Title  string
	Author string
}
