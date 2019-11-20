package internal

import (
	"bufio"
	"container/list"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/pflag"
)

// Sweep ...
func Sweep(args []string, flags *pflag.FlagSet) {

	// Get path, day from args
	path, error := filepath.Abs(filepath.ToSlash(args[0]))
	if error != nil {
		log.Fatal(error)
	}
	day, error := strconv.Atoi(args[1])
	if error != nil {
		log.Fatal(error)
	}

	// Get name, assumeyes from flags
	name, error := flags.GetString("name")
	if error != nil {
		log.Fatal(error)
	}
	assumeyes, error := flags.GetBool("assumeyes")
	if error != nil {
		log.Fatal(error)
	}

	// Get files from path
	files, error := ioutil.ReadDir(path)
	if error != nil {
		log.Fatal(error)
	}

	// Calculate modLimit using day
	now := time.Now()
	modLimit := now.AddDate(0, 0, -1*day)

	// Make nameRegexp using name
	nameRegexp, error := regexp.Compile(name)
	if error != nil {
		log.Fatal(error)
	}

	// Filter files using modLimit, nameRegexp
	deleteFileList := list.New()
	for _, file := range files {
		if !file.IsDir() && file.ModTime().Before(modLimit) && nameRegexp.MatchString(file.Name()) {
			deleteFileList.PushBack(file)
		}
	}

	// Check deleteFileList exist
	if deleteFileList.Len() == 0 {
		fmt.Println("No file to delete.")
		return
	}

	// Print deleteFileList
	for element := deleteFileList.Front(); element != nil; element = element.Next() {
		var deleteFile = element.Value.(os.FileInfo)
		fmt.Println(deleteFile.Name())
	}

	// Confirm deletion
	if !assumeyes {
		scanner := bufio.NewScanner(os.Stdin)
		var confirm bool
		var answered = false
		for !answered {
			fmt.Printf("Delete these %d file(s)? [y/n] ", deleteFileList.Len())
			scanner.Scan()
			answer := scanner.Text()
			answer = strings.ToLower(answer)
			switch answer {
			case "y":
				fallthrough
			case "ye":
				fallthrough
			case "yes":
				answered = true
				confirm = true
			case "n":
				fallthrough
			case "no":
				answered = true
				confirm = false
			default:
				fmt.Println("Error: Invalid answer. Please type 'y' or 'n'")
			}
		}
		if !confirm {
			return
		}
	}

	// Delete deleteFileList
	for element := deleteFileList.Front(); element != nil; element = element.Next() {
		var deleteFile = element.Value.(os.FileInfo)
		error = os.Remove(path + "/" + deleteFile.Name())
		if error != nil {
			log.Fatal(error)
		}
	}

	// Print the results
	fmt.Printf("%d file(s) deleted.\n", deleteFileList.Len())

}
