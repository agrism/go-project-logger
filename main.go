package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main()  {

	cmd := flag.String("action", "", "")
	flag.Parse()
	fmt.Printf("my action: \"%v\"\n", string(*cmd))

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pathes := []string{"/Users/agris/Pay Later Group/micro-services/accounts/public/index.php",
		"/Users/agris/Pay Later Group/micro-services/mobile/public/index.php"}

	for _, path := range pathes{
		content, err := readLines(path)
		if err != nil {
			panic(err)
		}

		var newContent []string

		phpCodeSlice, err := readLines(pwd + "/phpcode.php")

		if err != nil {
			panic(err)
		}

		fmt.Println(phpCodeSlice[len(phpCodeSlice)-1])

		if string(*cmd) == "remove" {
			fmt.Println("REMOVE")

			startDelete := false

			for _, s := range content {

				if phpCodeSlice[0] == s {
					startDelete = true
				}

				if startDelete == false {
					newContent = append(newContent, s)
				}

				if phpCodeSlice[len(phpCodeSlice)-1] == s {
					startDelete = false
				}
			}
		} else {
			fmt.Println("ADD")
			for i, s := range content {
				if i == 1 && s != phpCodeSlice[0] {
					newContent = append(newContent, phpCodeSlice...)
				}

				newContent = append(newContent, s);
			}
		}

		writeToFile(path, newContent)
	}
}

func readLines(path string) ([]string, error)  {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func writeToFile(path string, content []string) error  {

	f2, err := os.OpenFile(path, os.O_TRUNC|os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f2.Close()

	if err != nil {
		 fmt.Println(fmt.Errorf("could not open file %q for truncation: %v", path, err))
	}

	fmt.Println("truncate end")

	if err != nil {
		panic(err)
		return err
	}

	for _, s := range content{

		if _, err := f2.WriteString(s + "\n"); err != nil {
			log.Println(err)
		}
	}

	return nil
}