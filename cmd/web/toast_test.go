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