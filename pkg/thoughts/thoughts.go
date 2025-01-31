package thoughts

import (
	"bufio"
	"embed"
	"fmt"
	"log"

	"math/rand"
)

//go:embed assets/*.txt
var quotesDir embed.FS

func scanFile(fileName string) []string {
	file, err := quotesDir.Open(fmt.Sprintf("assets/%s", fileName))
	if err != nil {
		log.Panic("files not found")
	}

	defer file.Close()
	if err != nil {
		log.Fatal("could not read embedded data file")
	}

	sc := bufio.NewScanner(file)

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	lines := make([]string, 0)

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

func GetRandomQuote() string {
	fileNames := []string{"40k-darktide.txt", "40k-mechanicus.txt", "40k-general.txt"}

	quotes := make([]string, 0)

	for i := range fileNames {
		lines := scanFile(fileNames[i])
		quotes = append(quotes, lines...)
	}

	quote := quotes[rand.Intn(len(quotes))]
	return quote
}
