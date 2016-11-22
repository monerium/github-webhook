package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type Commit struct {
	Id        string `json:"id"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

type Event struct {
	Ref    string `json:"ref"`
	Commit Commit `json:"head_commit"`
}

func CorrectSignature(signature, message, key []byte) bool {
	mac := hmac.New(sha1.New, key)
	mac.Write(message)
	expected := mac.Sum(nil)

	return hmac.Equal(signature, expected)
}

func main() {
	secret := os.Getenv("SECRET_KEY")
	if secret == "" {
		log.Fatal("The SECRET_KEY environment variable is not set")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("The PORT environment variable is not set")
	}

	fmt.Fprintln(os.Stderr, "Hello, 世界")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}

		signature := strings.TrimLeft(r.Header.Get("X-Hub-Signature"), "sha1=")
		s, err := hex.DecodeString(signature)
		if err != nil {
			log.Fatal(err)
		}

		if CorrectSignature(s, b, []byte(secret)) {
			_, err = io.Copy(os.Stdout, bytes.NewReader(b))
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println()

		} else {
			http.Error(w, "Unauthorized", 401)
		}
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
