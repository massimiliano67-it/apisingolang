package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	translatev3 "cloud.google.com/go/translate/apiv3"
	"github.com/gorilla/mux"
	translatepb "google.golang.org/genproto/googleapis/cloud/translate/v3"
)

// translateText translates input text and returns translated text.
func translateText(w io.Writer, projectID string, sourceLang string, targetLang string, text string) error {
	// projectID := "my-project-id"
	// sourceLang := "en-US"
	// targetLang := "fr"
	// text := "Text you wish to translate"

	ctx := context.Background()
	client, err := translatev3.NewTranslationClient(ctx)
	if err != nil {
		return fmt.Errorf("NewTranslationClient: %v", err)
	}
	defer client.Close()

	req := &translatepb.TranslateTextRequest{
		Parent:             fmt.Sprintf("projects/%s/locations/global", projectID),
		SourceLanguageCode: sourceLang,
		TargetLanguageCode: targetLang,
		MimeType:           "text/plain", // Mime types: "text/plain", "text/html"
		Contents:           []string{text},
	}

	resp, err := client.TranslateText(ctx, req)
	if err != nil {
		return fmt.Errorf("TranslateText: %v", err)
	}

	// Display the translation for each input text provided
	for _, translation := range resp.GetTranslations() {
		fmt.Fprintf(w, "Translated text: %v\n", translation.GetTranslatedText())
	}

	return nil
}

// Message ...
type Message struct {
	SOURCELANG string `json:"sourcelang"`
	TARGETLANG string `json:"targetLang"`
	Body       string `json:"body"`
	TRANS      string `json:"trans,omitempty"`
}

func translate(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	w.Header().Set("Content-Type", "application/json")
	var message Message
	_ = json.NewDecoder(r.Body).Decode(&message)
	fmt.Println("message ", message)
	fmt.Println("sourcelang ", message.SOURCELANG)
	fmt.Println("targetLang ", message.TARGETLANG)
	fmt.Println("body ", message.Body)
	if err := translateText(&buf, "cloud-run-cd", message.SOURCELANG, message.TARGETLANG, message.Body); err != nil {
		message.TRANS = err.Error()
	} else {
		message.TRANS = "++" + buf.String()
	}
	json.NewEncoder(w).Encode(&message)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/translate", translate).Methods("POST")
	fmt.Println("Listen on port 8080.....")
	http.ListenAndServe(":8080", router)

}
