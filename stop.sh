#!/bin/bash

APP_NAME="user-management"

# Cari PID dari proses aplikasi
PID=$(pgrep -f "$APP_NAME")

if [ -z "$PID" ]; then
  echo "[$APP_NAME] tidak sedang berjalan."
else
  echo "Menghentikan [$APP_NAME] dengan PID: $PID"
  kill "$PID"

  # Tunggu sebentar untuk memastikan proses berhenti
  sleep 1

  # Cek apakah masih berjalan
  if ps -p "$PID" > /dev/null; then
    echo "Proses masih berjalan, paksa kill..."
    kill -9 "$PID"
  else
    echo "Berhasil dihentikan."
  fi
fi
