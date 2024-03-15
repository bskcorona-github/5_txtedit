package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filename := "editor.txt" // テキストファイルの名前

	// テキストファイルの読み込み
	content, err := readFile(filename)
	if err != nil {
		fmt.Println("ファイルを読み込めませんでした:", err)
	} else {
		fmt.Println("ファイルの内容:")
		fmt.Println(content)
	}

	// ユーザーからの入力を受け付けてテキストファイルに書き込み
	fmt.Println("ファイルに追加するテキストを入力してください。空行で終了します。")
	newContent := getUserInput()
	if newContent != "" {
		err := appendToFile(filename, newContent)
		if err != nil {
			fmt.Println("ファイルに書き込めませんでした:", err)
		} else {
			fmt.Println("ファイルに書き込みました。")
		}
	} else {
		fmt.Println("何も書き込まれませんでした。")
	}
}

// ファイルを読み込む
func readFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	content := ""
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}
	return content, scanner.Err()
}

// ユーザーからの入力を受け付ける
func getUserInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	userInput := ""
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		userInput += text + "\n"
	}
	return userInput
}

// ファイルに追記する
func appendToFile(filename, content string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}
