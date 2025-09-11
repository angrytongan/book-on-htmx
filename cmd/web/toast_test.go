package main

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestToastRandomNumber(t *testing.T) {
	// Create a minimal template for testing
	tmplContent := `{{block "toast-random-number" .}}<span>Random number: {{.Number}}</span>{{end}}`
	tpl, err := template.New("test").Parse(tmplContent)
	if err != nil {
		t.Fatalf("failed to parse template: %v", err)
	}

	app := &Application{
		tpl: tpl,
	}

	req := httptest.NewRequest(http.MethodGet, "/toast/random-number", nil)
	req.Header.Set("HX-Request", "true")

	rr := httptest.NewRecorder()

	app.toastRandomNumber(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	body := rr.Body.String()
	if !strings.Contains(body, "Random number:") {
		t.Errorf("response body should contain 'Random number:', got %v", body)
	}

	// Extract the number from the response and verify it's within range
	parts := strings.Split(body, "Random number: ")
	if len(parts) < 2 {
		t.Errorf("could not find random number in response: %v", body)
		return
	}

	// Get the part after "Random number: " and before the next tag
	numberPart := strings.Split(parts[1], "<")[0]
	number, err := strconv.Atoi(numberPart)
	if err != nil {
		t.Errorf("could not parse number from response: %v", err)
		return
	}

	if number < 0 || number > 100 {
		t.Errorf("random number should be between 0 and 100 inclusive, got %d", number)
	}
}

func TestToastRandomNumberMultipleRequests(t *testing.T) {
	// Create a minimal template for testing
	tmplContent := `{{block "toast-random-number" .}}<span>Random number: {{.Number}}</span>{{end}}`
	tpl, err := template.New("test").Parse(tmplContent)
	if err != nil {
		t.Fatalf("failed to parse template: %v", err)
	}

	app := &Application{
		tpl: tpl,
	}

	numbers := make(map[int]bool)
	
	// Make multiple requests to ensure randomness (though not a guarantee)
	for i := 0; i < 20; i++ {
		req := httptest.NewRequest(http.MethodGet, "/toast/random-number", nil)
		req.Header.Set("HX-Request", "true")

		rr := httptest.NewRecorder()
		app.toastRandomNumber(rr, req)

		body := rr.Body.String()
		parts := strings.Split(body, "Random number: ")
		if len(parts) < 2 {
			continue
		}

		numberPart := strings.Split(parts[1], "<")[0]
		number, err := strconv.Atoi(numberPart)
		if err != nil {
			continue
		}

		numbers[number] = true

		// Ensure each number is within range
		if number < 0 || number > 100 {
			t.Errorf("random number should be between 0 and 100 inclusive, got %d", number)
		}
	}

	// We should get at least some variation in 20 requests (though this could theoretically fail)
	if len(numbers) < 2 {
		t.Logf("Warning: Got %d unique numbers in 20 requests, expected more variation", len(numbers))
	}
}

func TestToastRandomLetter(t *testing.T) {
	// Create a minimal template for testing
	tmplContent := `{{block "toast-random-letter" .}}<span>Random letter: {{.Letter}}</span>{{end}}`
	tpl, err := template.New("test").Parse(tmplContent)
	if err != nil {
		t.Fatalf("failed to parse template: %v", err)
	}

	app := &Application{
		tpl: tpl,
	}

	req := httptest.NewRequest(http.MethodGet, "/toast/random-letter", nil)
	req.Header.Set("HX-Request", "true")

	rr := httptest.NewRecorder()

	app.toastRandomLetter(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	body := rr.Body.String()
	if !strings.Contains(body, "Random letter:") {
		t.Errorf("response body should contain 'Random letter:', got %v", body)
	}

	// Extract the letter from the response and verify it's a valid letter
	parts := strings.Split(body, "Random letter: ")
	if len(parts) < 2 {
		t.Errorf("could not find random letter in response: %v", body)
		return
	}

	// Get the part after "Random letter: " and before the next tag
	letterPart := strings.Split(parts[1], "<")[0]
	if len(letterPart) != 1 {
		t.Errorf("expected exactly one character, got %s", letterPart)
		return
	}

	letter := letterPart[0]
	if !((letter >= 'A' && letter <= 'Z') || (letter >= 'a' && letter <= 'z')) {
		t.Errorf("letter should be in range [A-Za-z], got %c", letter)
	}
}

func TestToastRandomLetterMultipleRequests(t *testing.T) {
	// Create a minimal template for testing
	tmplContent := `{{block "toast-random-letter" .}}<span>Random letter: {{.Letter}}</span>{{end}}`
	tpl, err := template.New("test").Parse(tmplContent)
	if err != nil {
		t.Fatalf("failed to parse template: %v", err)
	}

	app := &Application{
		tpl: tpl,
	}

	letters := make(map[byte]bool)
	
	// Make multiple requests to ensure randomness and verify range
	for i := 0; i < 20; i++ {
		req := httptest.NewRequest(http.MethodGet, "/toast/random-letter", nil)
		req.Header.Set("HX-Request", "true")

		rr := httptest.NewRecorder()
		app.toastRandomLetter(rr, req)

		body := rr.Body.String()
		parts := strings.Split(body, "Random letter: ")
		if len(parts) < 2 {
			continue
		}

		letterPart := strings.Split(parts[1], "<")[0]
		if len(letterPart) != 1 {
			continue
		}

		letter := letterPart[0]
		letters[letter] = true

		// Ensure each letter is within range [A-Za-z]
		if !((letter >= 'A' && letter <= 'Z') || (letter >= 'a' && letter <= 'z')) {
			t.Errorf("letter should be in range [A-Za-z], got %c", letter)
		}
	}

	// We should get at least some variation in 20 requests
	if len(letters) < 2 {
		t.Logf("Warning: Got %d unique letters in 20 requests, expected more variation", len(letters))
	}
}

func TestToastRandomWord(t *testing.T) {
	// Create templates for testing
	tmplContent := `{{block "toast-random-word" .}}<span>Random word: {{.Word}}</span>{{end}}`
	tpl, err := template.New("test").Parse(tmplContent)
	if err != nil {
		t.Fatalf("failed to parse template: %v", err)
	}

	// Create application with some test words
	testWords := []string{"apple", "banana", "cherry", "date", "elderberry"}
	app := &Application{
		tpl:   tpl,
		words: testWords,
	}

	req := httptest.NewRequest(http.MethodGet, "/toast/random-word", nil)
	req.Header.Set("HX-Request", "true")

	rr := httptest.NewRecorder()

	app.toastRandomWord(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	body := rr.Body.String()
	if !strings.Contains(body, "Random word:") {
		t.Errorf("response body should contain 'Random word:', got %v", body)
	}

	// Extract the word from the response and verify it's from our test set
	parts := strings.Split(body, "Random word: ")
	if len(parts) < 2 {
		t.Errorf("could not find random word in response: %v", body)
		return
	}

	wordPart := strings.Split(parts[1], "<")[0]
	found := false
	for _, word := range testWords {
		if wordPart == word {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("returned word '%s' not in expected test words %v", wordPart, testWords)
	}
}

func TestToastRandomWordError(t *testing.T) {
	// Create templates for testing
	tmplContent := `{{block "toast-error" .}}<span>Error: {{.Message}}</span>{{end}}`
	tpl, err := template.New("test").Parse(tmplContent)
	if err != nil {
		t.Fatalf("failed to parse template: %v", err)
	}

	// Create application with no words (simulate error condition)
	app := &Application{
		tpl:   tpl,
		words: []string{}, // empty word list
	}

	req := httptest.NewRequest(http.MethodGet, "/toast/random-word", nil)
	req.Header.Set("HX-Request", "true")

	rr := httptest.NewRecorder()

	app.toastRandomWord(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}

	body := rr.Body.String()
	if !strings.Contains(body, "Internal Server Error") {
		t.Errorf("response body should contain 'Internal Server Error', got %v", body)
	}
}

func TestToastRandomWordMultipleRequests(t *testing.T) {
	// Create templates for testing
	tmplContent := `{{block "toast-random-word" .}}<span>Random word: {{.Word}}</span>{{end}}`
	tpl, err := template.New("test").Parse(tmplContent)
	if err != nil {
		t.Fatalf("failed to parse template: %v", err)
	}

	// Create application with test words
	testWords := []string{"apple", "banana", "cherry", "date", "elderberry", "fig", "grape", "honeydew"}
	app := &Application{
		tpl:   tpl,
		words: testWords,
	}

	words := make(map[string]bool)

	// Make multiple requests to ensure randomness
	for i := 0; i < 20; i++ {
		req := httptest.NewRequest(http.MethodGet, "/toast/random-word", nil)
		req.Header.Set("HX-Request", "true")

		rr := httptest.NewRecorder()
		app.toastRandomWord(rr, req)

		body := rr.Body.String()
		parts := strings.Split(body, "Random word: ")
		if len(parts) < 2 {
			continue
		}

		wordPart := strings.Split(parts[1], "<")[0]
		words[wordPart] = true

		// Ensure each word is from our test set
		found := false
		for _, word := range testWords {
			if wordPart == word {
				found = true
				break
			}
		}

		if !found {
			t.Errorf("returned word '%s' not in expected test words %v", wordPart, testWords)
		}
	}

	// We should get at least some variation in 20 requests
	if len(words) < 2 {
		t.Logf("Warning: Got %d unique words in 20 requests, expected more variation", len(words))
	}
}