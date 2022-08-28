package structWithFiles

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type product struct {
	id          string
	title       string
	description string
	price       float64
}

func (prod *product) printData() {
	fmt.Printf("ID: %v\n", prod.id)
	fmt.Printf("Title: %v\n", prod.title)
	fmt.Printf("Description: %v\n", prod.description)
	fmt.Printf("Price: USD %.2f\n", prod.price)
}

func (prod *product) store() {
	file, _ := os.Create(fmt.Sprintf("%v-%v.txt", prod.id, prod.title))

	content := fmt.Sprintf(
		"ID: %v\nTitle: %v\nDescription: %v\nPrice: USD %.2f\n",
		prod.id, prod.title, prod.description, prod.price,
	)

	file.WriteString(content)

	file.Close()
}

func readUserInput(reader *bufio.Reader, promtText string) string {
	fmt.Print(promtText)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\n", "", -1)
	return userInput
}

func getProduct() *product {
	fmt.Println("Please enterthe product data.")
	fmt.Println("----------------------------")

	reader := bufio.NewReader(os.Stdin)

	idInput := readUserInput(reader, "Product ID: ")
	titleInput := readUserInput(reader, "Product Title: ")
	descriptionInput := readUserInput(reader, "Product Description: ")
	priceInput := readUserInput(reader, "Product Price: ")

	priceValue, _ := strconv.ParseFloat(priceInput, 64)

	product := newProduct(idInput, titleInput, descriptionInput, priceValue)

	return product
}

func newProduct(id, t, d string, p float64) *product {
	return &product{id, t, d, p}
}

func StructWithFilesSample() {
	createdProduct := getProduct()

	fmt.Println()
	createdProduct.printData()
	createdProduct.store()
	// firstProduct := Product{
	// 	"first-product",
	// 	"A book",
	// 	"A nice book",
	// 	10.99,
	// }
	// secondProduct := newProduct(
	// 	"second-product",
	// 	"A Carpet",
	// 	"A beautiful carpet",
	// 	99.99,
	// )

	// firstProduct.printData()
	// fmt.Println()
	// secondProduct.printData()
}
