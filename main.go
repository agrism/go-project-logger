package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {

	cmd := flag.String("action", "", "")
	flag.Parse()
	fmt.Printf("my action: \"%v\"\n", string(*cmd))

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	projectPathes := []string{
		"/Users/agris/Pay Later Group/micro-services/accounts/public/index.php",
		"/Users/agris/Pay Later Group/micro-services/mobile/public/index.php",
		"/Users/agris/Pay Later Group/micro-services/merchant/public/index.php",
		"/Users/agris/Pay Later Group/micro-services/customer-portal/public/index.php",
		"/Users/agris/Pay Later Group/micro-services/external-service/public/index.php",
		"/Users/agris/Pay Later Group/micro-services/documents/public/index.php",
		"/Users/agris/Pay Later Group/micro-services/address/public/index.php",
		"/Users/agris/Pay Later Group/micro-services/merchant-portal/public/index.php",
		"/Users/agris/git.paylatergroup.com/code/Micro-Service-Consumer-Level-Lending-API/public/index.php",
		"/Users/agris/Pay Later Group/legacy/LMP/public_html/index.php",
	}

	for _, projectPath := range projectPathes {
		content, err := readLines(projectPath)
		if err != nil {
			panic(err)
		}

		var newContent []string

		phpCodeSlice, err := readLines(pwd + "/phpcode.php")
		phpCodeSlice = phpCodeSlice[1:] // remove first line "<?php"

		if err != nil {
			panic(err)
		}

		if string(*cmd) == "remove" {
			fmt.Println("REMOVING from project " + projectPath)

			startDelete := false
			hasAnyLineToDelete :=false

			for _, s := range content {

				if phpCodeSlice[0] == s {
					startDelete = true
					hasAnyLineToDelete = true;
				}

				if startDelete == false {
					newContent = append(newContent, s)
				}

				if phpCodeSlice[len(phpCodeSlice)-1] == s {
					startDelete = false
				}
			}

			if(hasAnyLineToDelete != true){
				fmt.Println("No code to delete in this project")
			} else {
				writeToFile(projectPath, newContent)
				fmt.Println("Removing done....")
			}


		} else {
			fmt.Println("ADDING to project" + projectPath)

			codeAlreadyInsterted := false
			for _, s := range content {
				if s == phpCodeSlice[0]{
					codeAlreadyInsterted = true
				}
			}

			if codeAlreadyInsterted == true {
				fmt.Println("code already inserted, stop adding...")
			} else {
				for i, s := range content {

					if i == 1 && s != phpCodeSlice[0] {
						newContent = append(newContent, phpCodeSlice...)
					}

					newContent = append(newContent, s)
				}

				writeToFile(projectPath, newContent)
				fmt.Println("Adding done....")
			}
		}
	}
}

func readLines(path string) ([]string, error) {
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

func writeToFile(path string, content []string) error {

	f2, err := os.OpenFile(path, os.O_TRUNC|os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f2.Close()

	if err != nil {
		fmt.Println(fmt.Errorf("could not open file %q for truncation: %v", path, err))
	}

	if err != nil {
		panic(err)
		return err
	}

	for _, s := range content {

		if _, err := f2.WriteString(s + "\n"); err != nil {
			log.Println(err)
		}
	}

	return nil
}
