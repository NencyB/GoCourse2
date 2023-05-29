package main

import "fmt"

func main() {
	// onetohun()
	// uppercasealpha()
	// appending()
	// slicee()
	mapping()
}

func onetohun() {
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}
}

func uppercasealpha() {
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("%#U\n", i)
		}
	}
}

func appending() {
	x := []int{4, 5, 7, 8, 42}
	// fmt.Println(x)
	x = append(x, 77, 88, 99, 1014)
	// fmt.Println(x)
	y := []int{234, 456, 678, 987}
	x = append(x, y...)

	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}

func slicee() {
	jb := []string{"James", "Bond", "Chocolate", "martini"}
	fmt.Println(jb)
	mp := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut"}
	fmt.Println(mp)
	xp := [][]string{jb, mp}
	fmt.Println(xp)
}

func mapping() {
	mp := map[string]string{
		"hey": "hello",
		"bye": "tata",
	}
	fmt.Print(mp)
}
