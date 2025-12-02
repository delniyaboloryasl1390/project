package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

var imageLinks = []string{
	"https://diakoland.com/wp-content/uploads/2025/04/years-profile-picture-for-girls-diakoland-photo-37.webp",
	"https://recap.ir/wp-content/uploads/2024/08/%D8%B9%DA%A9%D8%B3-%DA%AF%D9%84-%D8%AF%D8%B1-%D8%A7%D9%81%D9%82-%D8%A2%D8%B3%D9%85%D8%A7%D9%86.jpg",
	"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQqxDsfSxKmSSPQhBOqDDbixGUVZ0lF8Jkpjw&s",
	"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTWoPa8ytEfb1tbaKb_74M6mEyyy6XMProyjw&s",
	"https://golishi.ir/wp-content/uploads/2024/06/photo_2024-06-15_19-49-04-462x462.webp",
}

func downloadFile(url string, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	fmt.Println("Link received")
	for i, link := range imageLinks {
		fmt.Printf("image number%d: %s\n", i+1, link)
	}
	for i, link := range imageLinks {

		extension := filepath.Ext(link)
		if extension == "" {
			extension = ".webp"
		}
		filename := fmt.Sprintf("downloaded_image_%d%s", i+1, extension)

		err := downloadFile(link, filename)
		if err != nil {
			fmt.Printf("Sorry for the photo download %d (%s): %v\n", i+1, link, err)
		} else {
			fmt.Printf("image %d Downloaded successfully: %s\n", i+1, filename)
		}
	}

}
