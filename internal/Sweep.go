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

	// args에서 path, day 가져오기
	path, error := filepath.Abs(filepath.ToSlash(args[0]))
	if error != nil {
		log.Fatal(error)
	}
	day, error := strconv.Atoi(args[1])
	if error != nil {
		log.Fatal(error)
	}

	// flags에서 name, assumeyes 가져오기
	name, error := flags.GetString("name")
	if error != nil {
		log.Fatal(error)
	}
	assumeyes, error := flags.GetBool("assumeyes")
	if error != nil {
		log.Fatal(error)
	}

	// path 에서 파일 목록 가져오기
	files, error := ioutil.ReadDir(path)
	if error != nil {
		log.Fatal(error)
	}

	// 현재 시간 기준으로 day일 이전을 limit으로 설정
	now := time.Now()
	limit := now.AddDate(0, 0, -1*day)

	// name으로 파일 이름 정규식 생성
	nameRegexp, error := regexp.Compile(name)
	if error != nil {
		log.Fatal(error)
	}

	// 삭제 대상 파일 필터링
	deleteFileList := list.New()
	for _, file := range files {
		if !file.IsDir() && file.ModTime().Before(limit) && nameRegexp.MatchString(file.Name()) {
			deleteFileList.PushBack(file)
		}
	}

	// 삭제 대상 파일 존재 여부 확인
	if deleteFileList.Len() == 0 {
		fmt.Println("No file to delete.")
		return
	}

	// 삭제 대상 파일 출력
	for element := deleteFileList.Front(); element != nil; element = element.Next() {
		var deleteFile = element.Value.(os.FileInfo)
		fmt.Println(deleteFile.Name())
	}

	// 삭제 여부 확인
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

	// 삭제 대상 파일 삭제
	for element := deleteFileList.Front(); element != nil; element = element.Next() {
		var deleteFile = element.Value.(os.FileInfo)
		error = os.Remove(path + "/" + deleteFile.Name())
		if error != nil {
			log.Fatal(error)
		}
	}

	// 삭제 결과 출력
	fmt.Printf("%d file(s) deleted.\n", deleteFileList.Len())

}
