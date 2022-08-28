package lists

import (
	"fmt"

	"github.com/devsendjin/golang-practice/helpers"
)

type Product struct {
	id    string
	title string
	price float64
}

func basicSlices() {
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println(prices)

	featuredPrices := prices[1:]
	featuredPrices[0] = 199.99
	highlightedPrices := featuredPrices[:1]
	fmt.Println(prices)
	fmt.Println(highlightedPrices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

	fmt.Println()

	highlightedPrices = highlightedPrices[:3]
	fmt.Println(highlightedPrices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
}

func structWithSlices() {
	// 1)
	hobbies := [3]string{"Sports", "Cooking", "Reading"}
	fmt.Println(hobbies)
	helpers.PrintDivider(&helpers.PrintDividerOptions{Step: 1})

	// 2)
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])
	helpers.PrintDivider(&helpers.PrintDividerOptions{Step: 2})

	// 3)
	mainHobbies := hobbies[:2]
	fmt.Println(mainHobbies)
	helpers.PrintDivider(&helpers.PrintDividerOptions{Step: 3})

	// 4)
	fmt.Println("cap:", cap(mainHobbies))
	mainHobbies = mainHobbies[1:len(hobbies)]
	fmt.Println(mainHobbies)
	helpers.PrintDivider(&helpers.PrintDividerOptions{Step: 4})

	// 5)
	courseGoals := []string{"Learn Go!", "Learn all the basics"}
	fmt.Println(courseGoals)
	helpers.PrintDivider(&helpers.PrintDividerOptions{Step: 5})

	// 6)
	courseGoals[1] = "Learn all the details!"
	courseGoals = append(courseGoals, "Learn all the basics!")
	fmt.Println(courseGoals)
	helpers.PrintDivider(&helpers.PrintDividerOptions{Step: 6})

	// 6)
	products := []Product{
		{
			"first-product",
			"A First Product",
			12.99,
		},
		{
			"secord-product",
			"A Secord Product",
			129.99,
		},
	}

	newProduct := Product{
		"third-product",
		"A Third Product",
		15.99,
	}

	products = append(products, newProduct)

	fmt.Println(products)
}

func unpackingSlices() {
	prices := []float64{10.99, 9.99}
	fmt.Println(prices)

	helpers.PrintDivider(&helpers.PrintDividerOptions{Label: "unpackingSlices", Step: 1})

	prices = append(prices, 101.99, 80.99, 20.59)
	prices = prices[1:]
	fmt.Println(prices)

	helpers.PrintDivider(&helpers.PrintDividerOptions{Label: "unpackingSlices", Step: 2})

	discountPrices := []float64{5.99, 12.99, 50.59}
	prices = append(prices, discountPrices...)
	fmt.Println(prices)
}

func ListsSample() {
	basicSlices()
	structWithSlices()
	unpackingSlices()
}
