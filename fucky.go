package main

import (
	"bufio"
	"bytes"
	//"reflect"
	"strings"

	//"encoding/base64"
	"fmt"
	"image/jpeg"
	"image/png"

	"github.com/corona10/goimagehash"

	//"image/png"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func hasDupes(m map[string]string) []string {
	x := make(map[string]struct{})
	//finds duplicates
	var dupeFiles []string
	for i, v := range m {
		if _, has := x[v]; has {
			log.Println(i, v)
			dupeFiles = append(dupeFiles, i)
		}
		x[v] = struct{}{}
	}
	return dupeFiles

}
func main() {
	//_, bar := diffList(iterate("C:/Users/Alonzo/Programming/DisArchived/DisArchived/images"))
	//	for _, i := range bar {
	//		newPlace := strings.SplitAfter(i, "images/")
	//		err := os.Rename(i, "C:/Users/Alonzo/Programming/DisArchived/DisArchived/images/failed/"+newPlace[1])
	//if err != nil {
	//log.Println(err)
	//}
	//	}

	imageDir := "C:/Users/Alonzo/Programming/DisArchived/DisArchived/images/"
	images := iterate(imageDir)
	var hashes = make(map[string]string)
	for _, foo := range images {
		fileDir := imageDir + foo

		//file1, err := os.Open(fileDir)
		fileExt := filepath.Ext(fileDir)
		switch fileExt {
		case ".png":
			file1, err := os.Open(fileDir)
			if err != nil {
				log.Println(err)
			}
			defer file1.Close()
			img1, err := png.Decode(file1)
			if err != nil {
				log.Println(img1, err)

			}
			hash, _ := goimagehash.DifferenceHash(img1)
			hashes[foo] = hash.ToString()
			if err != nil {
				log.Println(img1, err)

			}

		case ".jpeg", ".jpg":
			file1, err := os.Open(fileDir)
			if err != nil {
				log.Println(err)
			}
			defer file1.Close()
			img1, err := jpeg.Decode(file1)
			//maps name, hash as k,v to map, next is to find duplicates and move them to a separate folder, and fix png not working

			hash, _ := goimagehash.DifferenceHash(img1)
			hashes[foo] = hash.ToString()
			if err != nil {
				log.Println(img1, err)

			}

			if err != nil {
				log.Println(err)
			}

		}

	}
	log.Print(len(hashes))
	hasDupes(hashes)
	imagePng := "ship1C.png"
	imageJpg := "ship1.jpg"

	fileExt := findExt(imagePng)
	log.Println(fileExt)

	file1, err := os.Open(imagePng)
	if err != nil {
		log.Println(err)
	}
	defer file1.Close()
	//convert image to string to base64 then hash it? lol
	img1, err := png.Decode(file1)
	if err != nil {
		log.Println(img1, err)

	}

	file2, _ := os.Open(imageJpg)
	defer file1.Close()
	defer file2.Close()

	img2, _ := jpeg.Decode(file2)

	hash1, _ := goimagehash.AverageHash(img1)
	hash2, _ := goimagehash.AverageHash(img2)
	distance, _ := hash1.Distance(hash2)
	fmt.Printf("AVG HASH Distance between images: %v\n", distance)

	hash1, _ = goimagehash.DifferenceHash(img1)
	hash2, _ = goimagehash.DifferenceHash(img2)
	distance, _ = hash1.Distance(hash2)
	fmt.Printf("DIF HASH Distance between images: %v\n", distance)
	width, height := 8, 8
	hash3, _ := goimagehash.ExtAverageHash(img1, width, height)
	hash4, _ := goimagehash.ExtAverageHash(img2, width, height)
	distance, _ = hash3.Distance(hash4)
	fmt.Printf("EXTAVG HASH Distance between images: %v\n", distance)
	fmt.Printf("hash3 bit size: %v\n", hash3.Bits())
	fmt.Printf("hash4 bit size: %v\n", hash4.Bits())

	var b bytes.Buffer
	foo := bufio.NewWriter(&b)
	_ = hash4.Dump(foo)
	foo.Flush()

}

func findExt(file string) string {
	fullFile := strings.Split(file, ".")
	return fullFile[1]
}
func diffList(fileNames []string) (map[string]*goimagehash.ImageHash, []string) {
	// slice will be in [filename, diffList]
	var brokePic []string
	var hashes = make(map[string]*goimagehash.ImageHash)

	//convert failed to base64??
	for _, v := range fileNames {
		ogPath := "C:/Users/Alonzo/Programming/DisArchived/DisArchived/images/" + v
		file1, err := os.Open(ogPath)
		if err != nil {
			log.Println(err)
		}
		defer file1.Close()
		img1, err := jpeg.Decode(file1)
		if err != nil {
			log.Println(err)
			brokePic = append(brokePic, ogPath)
			if err != nil {
				log.Println(err)
			}

		}
		hash1, err := goimagehash.DifferenceHash(img1)
		if err != nil {
			log.Println(err)

		}
		hashes[v] = hash1
	}

	return hashes, brokePic
}

func iterate(dir string) []string {
	//returns directory in []string to be used to loop hashes
	var titles []string
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {

		titles = append(titles, file.Name())

	}

	return titles
}
