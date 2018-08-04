package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func convertExp(src string) {
	// おそらく、fileのパスが結構間違っているのでここでimageを読み込めていない可能性がある
	// srcで渡しているのがfile名のみとなっている。コマンドを実行したところからの相対パスを取りたい

	file, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	newFileName := replaceExt(src, ".jpeg", ".png")
	fmt.Println("this is a new file name")
	fmt.Println(newFileName)

	saveFile, err := os.Create(newFileName)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer saveFile.Close()
	png.Encode(saveFile, img)
}

func replaceExt(filePath, from, to string) string {
	ext := filepath.Ext(filePath)
	if len(from) > 0 && ext != from {
		return filePath
	}
	return filePath[:len(filePath)-len(ext)] + to
}

func main() {
	var (
		path = flag.String("path", ".", "specify path to start")
	)
	flag.Parse()

	err := filepath.Walk(*path, func(path string, info os.FileInfo, err error) error {
		fmt.Println(info.Name())
		if err != nil {
			log.Fatal(err)
		}

		if info.IsDir() {
			return nil
		}

		pos := strings.LastIndex(path, ".")
		if path[pos:] == ".jpeg" {
			convertExp(path)
			return nil
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
