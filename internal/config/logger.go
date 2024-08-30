package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
)

type betterStackWritter struct{}

func (b *betterStackWritter) Write(p []byte) (n int, err error) {
	log, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Erro ao converter para json", err)
		return 0, err
	}

	go func(log []byte) {

		url := "https://in.logs.betterstack.com"
		req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(log))
		if err != nil {
			fmt.Println("Erro ao criar request", err)
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer PWGMe7XafNwkyhs3GX5ziB4g")

		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Erro ao enviar log", err)
		}
		defer resp.Body.Close()

		//fmt.Printf("Log enviado, status de resposta: %d\n", resp.StatusCode)
	}(log)
	return len(p), nil
}

func Logger() {
	options := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}

	writers := io.MultiWriter(os.Stdout, &betterStackWritter{})
	logger := slog.New(slog.NewTextHandler(writers, options))
	slog.SetDefault(logger)
}
