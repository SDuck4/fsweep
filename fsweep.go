package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func main() {

	// argument에서 path, day 가져오기
	path, error := filepath.Abs(filepath.ToSlash(os.Args[1]))
	if error != nil {
		log.Fatal(error)
	}
	day, error := strconv.Atoi(os.Args[2])
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

	// 파일 수정 시간이 limit 보다 이전인 경우, 파일 삭제
	deleteCount := 0
	for _, file := range files {
		if file.ModTime().Before(limit) {
			filePath := path + "/" + file.Name()
			error = os.Remove(filePath)
			if error != nil {
				log.Fatal(error)
			}
			deleteCount++
		}
	}

	// 결과 출력
	fmt.Printf("%d file(s) deleted: files modified before %s", deleteCount, limit)

}
