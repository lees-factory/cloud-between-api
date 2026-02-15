#!/bin/bash
set -e

APP_DIR="/home/ubuntu/cloud-between-api"
APP_NAME="cloud-between-api"
LOG_FILE="$APP_DIR/app.log"

cd "$APP_DIR"

echo "=== Deploy started at $(date) ==="

# 1. Graceful stop
PID=$(pgrep -f "./$APP_NAME" || true)
if [ -n "$PID" ]; then
    echo "Stopping existing process (PID: $PID)..."
    kill -15 $PID || true
    # Wait up to 10 seconds for graceful shutdown
    for i in $(seq 1 10); do
        if ! pgrep -f "$APP_NAME" > /dev/null 2>&1; then
            echo "Process stopped."
            break
        fi
        sleep 1
    done
    # Force kill if still running
    if pgrep -f "$APP_NAME" > /dev/null 2>&1; then
        echo "Force killing process..."
        kill -9 $PID || true
        sleep 1
    fi
else
    echo "No existing process found."
fi

# 2. Replace binary
if [ -f "$APP_NAME.new" ]; then
    mv "$APP_NAME.new" "$APP_NAME"
    chmod +x "$APP_NAME"
    echo "Binary replaced."
else
    echo "ERROR: New binary not found!"
    exit 1
fi

# 3. Start with nohup
echo "Starting application..."
nohup ./"$APP_NAME" > "$LOG_FILE" 2>&1 &
NEW_PID=$!
echo "Application started (PID: $NEW_PID)"

# 4. Health check
echo "Waiting for health check..."
sleep 3

for i in $(seq 1 5); do
    if curl -sf http://localhost:8081/health > /dev/null 2>&1; then
        echo "Health check passed!"
        echo "=== Deploy completed at $(date) ==="
        exit 0
    fi
    echo "Health check attempt $i/5 failed, retrying..."
    sleep 2
done

echo "ERROR: Health check failed after 5 attempts"
echo "Last 20 lines of log:"
tail -20 "$LOG_FILE"
exit 1
