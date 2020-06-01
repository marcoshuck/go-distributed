package docs

import "strings"

type Word struct {
	Reference string
	Value string
	Count int64
	WordFrequency float64
}

type Words []Word

func Inspect(b Book) Words {
	var words Words
	plainWords := strings.Split(Sanitize(b.Content), " ")
	count := len(plainWords)
	if count == 0 {
		return nil
	}
	for _, plainWord := range plainWords {
		idx := Exist(words, plainWord, b.Title)
		if idx == -1 {
			words = append(words, Word{
				Reference:     b.Title,
				Value:         plainWord,
				Count:         1,
				WordFrequency: float64(1 / count),
			})
			continue
		}
		words[idx].Count++
		words[idx].WordFrequency = float64(words[idx].Count / int64(count))
	}
}

func Sanitize(content string) string {
	content = strings.Replace(content, ".", "", -1)
	content = strings.Replace(content, ",", "", -1)
	content = strings.Replace(content, ";", "", -1)
	return content
}

func Exist(words Words, w string, title string) int {
	for i, word := range words {
		if w == word.Value && title == word.Reference {
			return i
		}
	}
	return -1
}