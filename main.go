package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/sys/unix"
)

type FilesList struct {
	Filename []map[string]int64 `json:"files"`
}

func main() {

	LOG_FILE := "./execution.log"
	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		fmt.Println("Error with log file", err)
	}
	//close the file at the end of main
	defer logFile.Close()

	//set output of logs to logfile
	log.SetOutput(logFile)

	arg1 := flag.String("mount", "/", "default mount point is root /")
	//fmt.Println("Size:", arg1)

	arg2 := flag.String("sort", "mount", "default sort is mount")

	flag.Parse()
	mountPoint := *arg1
	sortBy := *arg2

	fileDirList := make(map[string]int64)

	var stat unix.Statfs_t

	unix.Statfs(mountPoint, &stat)
	/*
		// Available blocks * size per block = available space in bytes
		//var available_space = stat.Bavail * uint64(stat.Bsize)

		//fmt.Println(available_space / (1024 * 1024))

			totalSize := stat.Blocks * uint64(stat.Bsize)
			totalSizeMb := float64(totalSize) / (1024 * 1024)

			free := stat.Bfree * uint64(stat.Bsize)
			freeMb := float64(free) / (1024 * 1024)

			avail := stat.Bavail * uint64(stat.Bsize)
			availMB := float64(avail) / (1024 * 1024)



			fmt.Println("-------------------- The size of mount point --------------------")
			fmt.Println("Size:", totalSizeMb)
			fmt.Println("Free:", freeMb)
			fmt.Println("Available:", availMB) //reserved filesystem block for root
			fmt.Println("Used:", totalSizeMb-freeMb)
			fmt.Println("-------------------- End of Printing Mount Storage --------------------")
	*/
	files, err := ioutil.ReadDir(mountPoint)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		absPath := filepath.Join(mountPoint, f.Name())
		//log.Printf("Returned %s %d", absPath, f.Size())
		var size int64
		if f.IsDir() {
			//fmt.Println(absPath)
			size, err = dirSize(absPath)
			if err != nil {
				//log.("Error occured ", err)
				log.Println("Error", err)
				fileDirList[absPath] = -1

			}
			fileDirList[absPath] = size
		} else {
			fileDirList[absPath] = f.Size()
		}

	}

	var innterArray []map[string]int64
	if sortBy == "diskusage" {

		pair := sortByValue(fileDirList)
		for _, e := range pair {
			innterArray = append(innterArray, map[string]int64{e.Key: e.Value})
		}

	} else {

		keys := sortByKey(fileDirList)

		for _, k := range keys {
			innterArray = append(innterArray, map[string]int64{k: fileDirList[k]})
		}

	}

	var fileList FilesList
	fileList.Filename = innterArray

	jsonStr, err := json.Marshal(fileList)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Println(string(jsonStr))
	}

}

// go mod init main
// go get golang.org/x/sys/unix
// go install golang.org/x/sys/unix
