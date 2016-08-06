package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/crypto/sha3"
)

var gpm_version = "0.1.0"

var searchDirectory = flag.String("dir", ".", "Directory to search for gifs")
var targetDirectory = flag.String("tar", "./0gif", "Directory in which to place gifs")
var dryRun = flag.Bool("dry", true, "Don't move anything")
var keepName = flag.Bool("keep-name", false, "Keep the Original Filename")
var printFiles = flag.Bool("print-files", false, "Print all file names")
var version = flag.Bool("version", false, "Display Version")

func visit(path string, f os.FileInfo, err error) error {
	path = filepath.Clean(path) // Fixes some weird bug on Windows
	if strings.HasSuffix(f.Name(), ".gif") {
		absPath, err := filepath.Abs(path)
		if err != nil {
			fmt.Printf("Error on File: %s\n", path)
			return err
		}
		if strings.Contains(absPath, "/0gif/") {
			return nil
		}
		var newName string

		if !*keepName {
			newName, err = getHashName(absPath)
			if err != nil {
				fmt.Printf("HashName Err: %s\n", err)
				return err
			}
		} else {
			newName = f.Name()
		}

		newPath := filepath.Join(*targetDirectory, newName)

		if _, err = os.Stat(newPath); os.IsNotExist(err) {
			if *printFiles {
				fmt.Printf("Moving %s to %s\n", path, newName)
			}
			if !*dryRun {
				err = os.Rename(absPath, newPath)
				if err != nil {
					fmt.Printf("Error on Move: %s\n", err)
					return err
				}
			}
		} else {
			fmt.Printf("Collision: %s with %s\n", absPath, newPath)
		}
	}
	return nil
}

func main() {
	flag.Parse()
	
	if *version {
		fmt.Printf("GifPicMover v%s\n", gpm_version)
		return
	}
	
	if s, e := filepath.Abs(*searchDirectory); e != nil {
		fmt.Printf("Can't get absolute path for searchDirectory!\n")
		return
	} else {
		searchDirectory = &s
	}

	if s, e := filepath.Abs(*targetDirectory); e != nil {
		fmt.Printf("Can't get absolute path for targetDirectory!\n")
		return
	} else {
		targetDirectory = &s
	}

	fmt.Printf("--- BEGIN GIFPICMOV ---\nSearching in: %s\nTarget: %s\n", *searchDirectory, *targetDirectory)
	if *dryRun {
		fmt.Println("Dry Run!")
	}
	os.MkdirAll(*targetDirectory, os.ModePerm)
	_ = filepath.Walk(*searchDirectory, visit)
	fmt.Printf("--- END GIFPICMOV ---\n")
}

func getHashName(fileName string) (string, error) {
	var cont []byte

	cont, err := ioutil.ReadFile(fileName)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x.gif", sha3.Sum256(cont)), nil
}
