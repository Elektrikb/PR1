package main

import (
	"fmt"
	"os"
	"strings"
)

// FileOwner представляет владельца файла и его содержимого
type FileOwner struct {
	path    string
	content string
}

// newFileOwner создает новый экземпляр FileOwner (владелец файла)
func newFileOwner(path string) (*FileOwner, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("ошибка при чтении файла: %v", err)
	}
	return &FileOwner{
		path:    path,
		content: string(content),
	}, nil
}

// getContent возвращает неизменяемую ссылку на содержимое файла (заимствование)
func (f *FileOwner) getContent() *string {
	return &f.content
}

// countWords подсчитывает общее количество слов в тексте
func countWords(content *string) int {
	words := strings.Fields(*content)
	return len(words)
}

// countWordOccurrences подсчитывает количество вхождений слова (без учета регистра)
func countWordOccurrences(content *string, target string) int {
	words := strings.Fields(*content)
	count := 0
	for _, word := range words {
		if strings.EqualFold(word, target) {
			count++
		}
	}
	return count
}

// analyzeFile анализирует файл и выводит результаты
func analyzeFile(filePath string, targetWord string) error {
	// Владение файлом и его содержимым передается FileOwner
	fileOwner, err := newFileOwner(filePath)
	if err != nil {
		return err
	}

	// Заимствуем содержимое для анализа
	content := fileOwner.getContent()

	totalWords := countWords(content)
	occurrences := countWordOccurrences(content, targetWord)

	fmt.Printf("Результаты анализа файла %s:\n", fileOwner.path)
	fmt.Printf("Общее количество слов: %d\n", totalWords)
	fmt.Printf("Количество вхождений слова \"%s\": %d\n", targetWord, occurrences)

	return nil
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Использование: program <путь_к_файлу> <слово_для_поиска>")
		os.Exit(1)
	}

	filePath := os.Args[1]
	targetWord := os.Args[2]

	if err := analyzeFile(filePath, targetWord); err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		os.Exit(1)
	}
}
