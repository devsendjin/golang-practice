package maps

import (
	"fmt"

	"github.com/devsendjin/golang-practice/helpers"
)

func concept() {
	helpers.PrintDivider(&helpers.PrintDividerOptions{Label: "Map basic ⬇", DrawEmptyLineBelow: true})

	websites := map[string]string{
		"Google": "https://google.com",
		"Aws":    "https://aws.com",
	}

	fmt.Println(websites)
	fmt.Println(websites["Aws"])

	helpers.PrintDivider(&helpers.PrintDividerOptions{Label: "Add to map ⬇", DrawEmptyLines: true})

	websites["LinkedIn"] = "https://linkedin.com"
	fmt.Println(websites)

	helpers.PrintDivider(&helpers.PrintDividerOptions{Label: "Remove from map ⬇", DrawEmptyLines: true})
	delete(websites, "Google")
	fmt.Println(websites)
}

func MapsSample() {
	// fmt.Println(rand.New(rand.NewSource(time.Now().UnixNano())))
	concept()
}
