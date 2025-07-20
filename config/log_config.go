package config

import (
	"io"
	"log"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

func SetupLog() {

	// Buat folder logs jika belum ada
	logDir := "logs"
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		os.Mkdir(logDir, 0755)
		if err != nil {
			log.Fatalf("Failed to create log directory: %v", err)
			panic("Failed to create log directory: " + err.Error())
		}
	}

	// Konfigurasi logger file dengan rotasi harian
	logger := &lumberjack.Logger{
		Filename:   logDir + "/app.log",
		MaxSize:    10, // megabytes
		MaxBackups: 10,
		MaxAge:     1, // days
		Compress:   true,
	}

	// Gabungkan output ke terminal (os.Stdout) dan file
	multiWritter := io.MultiWriter(os.Stdout, logger)

	// Atur log agar ke dua tempat
	log.SetOutput(multiWritter)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
