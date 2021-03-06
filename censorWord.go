package swear_free

import (
	"encoding/json"
	"fmt"
	"github.com/OhhhThatVarun/swear-free/lists"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const defaultLocale = "En"

var libLocale = ""
var libReplaceCharacter = "*"

func SetLocale(locale string) {
	libLocale = locale
}

func SetReplaceCharacter(replaceCharacter string) {
	libReplaceCharacter = replaceCharacter
}

func CensorWord(str string) string {

	// Load the json according to the locale
	unCensored := loadJsonFile()

	var newSlice []string
	strSlice := strings.Fields(str)

	// TODO: This lookup should be improved by using map or something
	for position, word := range strSlice {

		var swearWordFound = false

		for _, forbiddenWord := range unCensored.Words {

			if test := strings.EqualFold(strings.ToLower(word), forbiddenWord); test == true {
				replacement := strings.Repeat(libReplaceCharacter, len(word))
				strSlice[position] = replacement
				newSlice = append(strSlice[:position], strSlice[position:]...)
				swearWordFound = true
			}
		}

		if !swearWordFound {
			newSlice = append(strSlice[:position], strSlice[position:]...)
		}
	}
	return strings.Join(newSlice, " ") // convert []string slice back to string
}

func loadJsonFile() lists.WordJson {
	if libLocale == "" {
		fmt.Println(fmt.Sprintf("No locale defined, using the default locale: %s", defaultLocale))
		libLocale = defaultLocale
	}

	var jsonFileName = fmt.Sprintf("%s/lists/ForbiddenWords%s.json", getWorkingDir(), libLocale)
	jsonFile, err := os.Open(jsonFileName)
	if err != nil {
		fmt.Println("currently this locale is not supported")
		panic(fmt.Sprintf("Tried loading %s", jsonFileName))
	}
	fmt.Println(fmt.Sprintf("Successfully loded %s", jsonFileName))
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var wordJson lists.WordJson
	_ = json.Unmarshal(byteValue, &wordJson)
	return wordJson
}

func getWorkingDir() string {
	wd, _ := os.Getwd()
	for !strings.HasSuffix(wd, "swear-free") {
		wd = filepath.Dir(wd)
	}
	return wd
}
