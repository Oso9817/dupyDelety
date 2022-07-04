package main

import (

	//"encoding/base64"
	"errors"
	"github.com/corona10/goimagehash"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	//"net/http"

	//"image/png"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var (
	errInvalid = errors.New("invalid input data")
)

func HasDupes(m map[string]*goimagehash.ImageHash, imageDir string) int {
	//checkes provided map and comppares hash distance to determine similarity, removes those files
	var lastKey string
	removedQTY := 0
	var lastHash *goimagehash.ImageHash
	//var duplicates []string
	for currentKey, currentHash := range m {
		for k, h := range m {
			if currentKey == k {
				continue
			}
			if lastKey != "" && lastHash != nil {
				if k != currentKey && h != currentHash {
					distance, err := currentHash.Distance(h)
					if err != nil {
						log.Println(err)
					}

					if distance <= 7 {
						os.Remove(imageDir + currentKey)
						removedQTY += 1

					}
				}

			}
		}
		lastKey = currentKey
		lastHash = currentHash
	}
	log.Printf("Removed %v duplicate files!", removedQTY)
	return removedQTY

}

func ProcessImage(file string) (*goimagehash.ImageHash, error) {
	file1, err := os.Open(file)
	//if it cant open the file, move to the next, opening file is the most crucial step
	if err != nil {
		return nil, err

	}
	defer file1.Close()
	img1, _, err := image.Decode(file1)
	//EOF was a common error, image is malformed and needs to be repaired
	if errors.Is(err, io.EOF) || errors.Is(err, image.ErrFormat) {
		log.Printf("File: %v has reached EOF, file needs repaired, or is invalid: %v", file1.Name(), err)

	} else if err != nil {
		return nil, err
	}

	hash, _ := goimagehash.DifferenceHash(img1)

	if err != nil {
		return nil, err

	}
	return hash, err
}

//creates hash map paired with file names to check later for comparision
func HashMap(imageDir string, images []string) (map[string]*goimagehash.ImageHash, error) {
	var hashes = make(map[string]*goimagehash.ImageHash)
	if imageDir == "" || images == nil {
		return nil, errInvalid
	}
	//loop through range of image file names
	for _, foo := range images {
		file := filepath.Join(imageDir, foo)

		imageHash, err := ProcessImage(file)
		if errors.Is(err, io.EOF) || errors.Is(err, image.ErrFormat) {
			continue
		} else if err != nil {
			log.Println(err)
			continue
		}
		hashes[foo] = imageHash

	}
	return hashes, nil
}

func Iterate(dir string) ([]string, error) {
	//returns directory in []string to be used to loop hashes
	var titles []string
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		titles = append(titles, file.Name())
	}

	return titles, nil
}
