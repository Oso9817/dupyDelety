package main

import (
	//"bufio"
	//"bytes"

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

func hasDupes(m map[string]*goimagehash.ImageHash) []string {
	x := make(map[*goimagehash.ImageHash]struct{})
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
	var hashes = make(map[string]*goimagehash.ImageHash)
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
			hashes[foo] = hash
			distance, _ := hash.Distance(hashes["ship1D.jpg"])
			fmt.Printf("DIF HASH Distance between images: %v\n", distance)
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
			hashes[foo] = hash
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
