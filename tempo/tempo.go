package tempo

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

func CallTempoAPI(from string, to string) (string, error) {
	// Load environment variables from a .env file
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Error finding user's home directory:", err)
	}
	environmentFile := homeDir + string(os.PathSeparator) + ".tempoworklog"
	err = godotenv.Load(environmentFile)
	if err != nil {
		log.Fatal("Error loading .env file: ", environmentFile)
	}
	// Set your Tempo API token, the base URL for Tempo, and the issue key you're interested in.
	apiToken := os.Getenv("TEMPO_API_TOKEN")
	tempoBaseURL := os.Getenv("TEMPO_BASE_URL") // Something like "https://api.tempo.io/core/3"

	baseURL := fmt.Sprintf("%s/worklogs", tempoBaseURL)
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		panic(err)
	}
	// Prepare query parameters
	params := url.Values{}

	params.Add("from", from) //
	params.Add("to", to)     //

	// Encode and assign the parameters to the URL
	parsedURL.RawQuery = params.Encode()
	urlFinite := parsedURL.String()
	req, err := http.NewRequest("GET", urlFinite, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return "", err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	req.Header.Add("Content-Type", "application/json")

	// Make the HTTP request.
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making the tempo request:", err)
		fmt.Println("Url:", urlFinite)
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body.
	body, err := io.ReadAll(io.Reader(resp.Body))
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return "", err
	}

	return string(body), nil
}

func CallTempoNext(url string) (string, error) {
	// Load environment variables from a .env file
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Error finding user's home directory:", err)
	}
	environmentFile := homeDir + string(os.PathSeparator) + ".tempoworklog"
	err = godotenv.Load(environmentFile)
	if err != nil {
		log.Fatal("Error loading .env file: ", environmentFile)
	}
	// Set your Tempo API token, the base URL for Tempo, and the issue key you're interested in.
	apiToken := os.Getenv("TEMPO_API_TOKEN")

	urlFinite := url
	req, err := http.NewRequest("GET", urlFinite, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return "", err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	req.Header.Add("Content-Type", "application/json")

	// Make the HTTP request.
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making the tempo next request:", err)
		fmt.Println("Url:", url)
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body.
	body, err := io.ReadAll(io.Reader(resp.Body))
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return "", err
	}

	return string(body), nil
}
