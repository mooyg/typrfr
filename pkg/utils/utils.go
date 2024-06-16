package utils

import (
	"encoding/json"
	"log/slog"
	"math/rand"
	"os"
)

func GenCode() int {
	return 100 + rand.Intn(1000-100)
}

type Data struct {
	Id   int
	Para string
}

func GenText() string {
	content, err := os.ReadFile("para.json")

	if err != nil {
		slog.Error("error while opening the file", "err", err)
	}
	var payload []Data

	err = json.Unmarshal(content, &payload)

	if err != nil {
		slog.Error("error during unmarshal()", "err", err)
	}

	text := payload[rand.Intn(len(payload))].Para
	return text
}

func Unmarshal[T any](data []byte) T {
	var val T
	err := json.Unmarshal(data, &val)

	if err != nil {
		slog.Error("some error occured while unmarshalling")
		os.Exit(2)
	}
	return val

}
