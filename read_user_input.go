package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

func trim_text(txt string) string {
	txt = strings.TrimSuffix(txt, "\n")
	txt = strings.TrimSuffix(txt, "\r")
	return txt
}

func read_user_input_demo() {
	fmt.Println("Read User Input")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please write some text : ")
	phrase, _ := reader.ReadString('\n')
	flat_phrase := strings.Trim(phrase, " ")
	fmt.Print("Please enter file name : ")
	file_name, _ := reader.ReadString('\n')
	flat_file_name := strings.Split(file_name, " ")[0]
	flat_file_name = trim_text(flat_file_name)
	file_path := fmt.Sprintf("%s/%s", path.Base(""), flat_file_name)
	fmt.Printf("Writing \"%s\" to %s", flat_phrase, file_path)
	f, err := os.OpenFile(file_path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("Can not open file: %s", file_path)
		os.Exit(1)
	}
	// close file after all
	defer f.Close()
	_, err = f.Write([]byte(flat_phrase))
	if err != nil {
		fmt.Printf("Write file error: %s", err)
		os.Exit(1)
	}

}
