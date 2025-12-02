package main

import (
	"fmt"
)

var imageLinks = []string{
	"https://diakoland.com/wp-content/uploads/2025/04/years-profile-picture-for-girls-diakoland-photo-37.webp",
	"https://recap.ir/wp-content/uploads/2024/08/%D8%B9%DA%A9%D8%B3-%DA%AF%D9%84-%D8%AF%D8%B1-%D8%A7%D9%81%D9%82-%D8%A2%D8%B3%D9%85%D8%A7%D9%86.jpg",
	"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQqxDsfSxKmSSPQhBOqDDbixGUVZ0lF8Jkpjw&s",
	"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTWoPa8ytEfb1tbaKb_74M6mEyyy6XMProyjw&s",
	"https://golishi.ir/wp-content/uploads/2024/06/photo_2024-06-15_19-49-04-462x462.webp",
}

func main() {
	fmt.Println("Link received")
	for i, link := range imageLinks {
		fmt.Printf("image number%d: %s\n", i+1, link)
	}
}
