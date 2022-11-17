package hangmanweb

import (
	hc "hangmanweb/hangman-classic"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

type Hangman struct {
	Username    string
	Letter      string
	MotTab      []string
	Motstr      string
	Mot         string
	Attempts    int
	LettersUsed []string
	Win         bool
}

func Initword() ([]string, string, string) {
	data, err := ioutil.ReadFile("./hangman-classic/main/words.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	var hold string
	list := []string{}
	for _, m := range string(data) {
		if m != 10 {
			hold = hold + string(m)
		} else {
			if hold != "" {
				list = append(list, hold)
				hold = ""
			}
		}
	}

	rand.Seed(time.Now().UnixNano())

	randomword := list[rand.Intn(len(list))]

	randomwordhide := hc.CreateWord(randomword)
	randomwordhide = randomwordhide[:(len(randomwordhide) - 1)]
	randomwordhidetab := []string{}
	for i := 0; i < len(randomwordhide)-1; i++ {
		randomwordhidetab = append(randomwordhidetab, string(randomwordhide[i]))
	}

	randomword = randomword[:(len(randomword) - 1)]

	return randomwordhidetab, randomword, randomwordhide

}
func TabtoStr(tab []string) string {
	str := ""
	for _, ch := range tab {
		str += string(ch)
	}
	return str
}
