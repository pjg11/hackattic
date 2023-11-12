package connection

import (
	"net/http"
	"fmt"
	"io"
	"os"
	"bytes"
)

var token = os.Getenv("HACKATTIC_TOKEN")

func Challenge(name string) (io.ReadCloser, error) {
	resp, err := http.Get(fmt.Sprintf("https://hackattic.com/challenges/%s/problem?access_token=%s", name, token))
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func Solve(name string, b []byte) error {
	resp, err := http.Post(fmt.Sprintf("https://hackattic.com/challenges/%s/solve?access_token=%s&playground=1", name, token), "application/json", bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
	return nil
}
