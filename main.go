package main

import (
	"fmt"
	"github.com/Violet-Lemon/hexlet-go/v2/greeting"
	"github.com/fatih/color"
	em "github.com/Violet-Lemon/hexlet-go/v2/employees"
	pr "github.com/Violet-Lemon/hexlet-go/v2/payroll"
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

	team := []em.Employee{
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
	}
}
