package utils

import (
	"bufio"
	log "github.com/sirupsen/logrus"
	"os"
)

type Questions struct {
	Question []string `json:"question"`
}

//TODO: продумать логику добавления вопросов
func QuestionParser(questions Questions) error {
	fl, err := os.Open("questions.json")
	if err != nil {
		log.Fatalf("error while opening file with questions, description: %v", err)
	}
	defer fl.Close()
	scann := bufio.NewScanner(fl)
	scann.Buffer([]byte{}, 1024*1024*1024*1024)
	for scann.Scan() {
		scann.Text()
	}

	return nil
}
