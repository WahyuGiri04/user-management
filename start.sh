#!/bin/bash

APP_NAME="user-management"
LOG_DIR="logs/"
LOG_FILE="$LOG_DIR/app-user-management.log"
EXEC="./$APP_NAME"

# Buat folder logs kalau belum ada
mkdir -p "$LOG_DIR"

# Jalankan aplikasinya di background dan arahkan log ke file
nohup "$EXEC" >> "$LOG_FILE" 2>&1 &

echo "[$APP_NAME] berhasil dijalankan. Log: $LOG_FILE"
