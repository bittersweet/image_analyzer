package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

func RGBToHex(r, g, b uint8) string {
	return fmt.Sprintf("#%02X%02X%02X", r, g, b)
}

func convertColorToHex(c color.Color) string {
	r, g, b, _ := c.RGBA()
	return RGBToHex(uint8(r>>8), uint8(g>>8), uint8(b>>8))
}

func listImages() []string {
	s := make([]string, 0)
	files, _ := ioutil.ReadDir("./")
	imageRegexp, _ := regexp.Compile("\\.png|\\.jpg")

	for _, f := range files {
		if imageRegexp.MatchString(f.Name()) == true {
			s = append(s, f.Name())
		}
	}
	return s
}

func processImage(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	m, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	bounds := m.Bounds()
	total := 0
	colormap := make(map[string]int)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			total++
			color := m.At(x, y)

			hex := convertColorToHex(color)
			colormap[hex]++
		}
	}

	highest := 0
	var highest_hex string
	for hex, amount := range colormap {
		if amount > highest {
			highest_hex = hex
			highest = amount
		}
	}

	percentage := float64(highest) / float64(total) * 100

	if isDirectInput() {
		fmt.Print(percentage)
	} else {
		fmt.Printf("%s, Total: %d, highest: %d, percentage: %f\n", filename, total, highest, percentage)
		fmt.Println("hex: ", highest_hex)
	}
}

func isDirectInput() bool {
	return len(os.Args) > 1
}

func main() {
	if isDirectInput() {
		processImage(os.Args[1])
	} else {
		images := listImages()

		for _, image := range images {
			processImage(image)
		}
	}
}
