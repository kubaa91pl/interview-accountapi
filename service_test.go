package service

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestService_HAPPY_PATH(t *testing.T) {
	baseURL, ok := os.LookupEnv("FORM3_API")
	if !ok {
		baseURL = "http://localhost:8080"
	}

	client := NewClient(baseURL)
	f, err := os.Open("payloads/account-payload.json")
	require.NoError(t, err)
	defer f.Close()

	var data ResponseNotification
	err = json.NewDecoder(f).Decode(&data)
	if err != nil {
		panic(err)
	}

	err = client.Create(data.Data)
	if err != nil {
		panic(err)
	}
	acc, err := client.Fetch("a9e3b971-a241-4930-a09f-a7c04bf394fe")
	if err != nil {
		fmt.Println(err)
	}

	err = client.Delete("a9e3b971-a241-4930-a09f-a7c04bf394fe", "0")
	if err != nil {
		fmt.Println(err)
	}

	err = client.Create(data.Data)
	if err != nil {
		panic(err)
	}

	fmt.Println(acc)
	require.NoError(t, err)
}
