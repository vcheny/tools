package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/vcheny/golib/rest"
)

func main() {
	// Command-line flags
	methodFlag := flag.String("method", "post", "HTTP method (post/get)")
	url := flag.String("url", "", "ServiceNow API URL")
	username := flag.String("username", "yan.chen", "Username")
	password := flag.String("password", "", "Password")
	raw := flag.Bool("raw", false, "Raw respone. No pretty JSON output for GET requests")

	flag.Parse()

	// Validate required flags
	if *url == "" || *username == "" || *password == "" {
		flag.Usage()
		os.Exit(1)
	}

	// Create a new Rest instance
	restClient := &rest.Rest{
		URL:      *url,
		Username: *username,
		Password: *password,
	}

	switch *methodFlag {
	case "post":
		// Read JSON data from stdin
		jsonData := readStdin()
		// Make the POST request
		response, err := restClient.Post(jsonData)
		if err != nil {
			fmt.Println("Failed to send POST request:", err)
			os.Exit(1)
		}

		fmt.Println("Post was sent successfully:", response)

	case "get":
		// Make the GET request
		responseData, err := restClient.Get()
		if err != nil {
			fmt.Println("Failed to send GET request:", err)
			os.Exit(1)
		}
		// Pretty print JSON output if the flag is set
		if *raw {
			fmt.Println("Response Body:\n", string(responseData))
		} else {
			var prettyJSON bytes.Buffer
			err = json.Indent(&prettyJSON, responseData, "", "  ")
			if err != nil {
				fmt.Println("Failed to pretty print JSON:", err)
				os.Exit(1)
			}
			fmt.Println("Response Body:\n", prettyJSON.String())
		}

	default:
		fmt.Println("Invalid HTTP method specified.")
		flag.Usage()
		os.Exit(1)
	}
}

// Read JSON data from stdin
func readStdin() []byte {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		bytes, _ := ioutil.ReadAll(os.Stdin)
		return bytes
	}
	return []byte{}
}
