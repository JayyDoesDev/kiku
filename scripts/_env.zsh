#!/bin/zsh
# Usage: source ./scripts/_env.zsh

# Remove conflicting aliases first
unalias gobuild 2>/dev/null
unalias gotest 2>/dev/null
unalias gorun 2>/dev/null
unalias gofmtall 2>/dev/null
unalias gorestart 2>/dev/null

# --- Clean Apple metadata in current directory ---
_clean_current_dir() {
  dot_clean -m . 2>/dev/null
  find . -name '._*' -type f -delete 2>/dev/null
}

# --- Go Commands Wrapped with Cleaning ---

gobuild() {
  _clean_current_dir
  go build ./...
}

gotest() {
  _clean_current_dir
  go test ./...
}

gorun() {
  _clean_current_dir
  go run main.go
}

gofmtall() {
  _clean_current_dir
  gofmt -s -w .
}

gorestart() {
  _clean_current_dir
  pkill -f main.go
  go run main.go
}

# --- Confirmation ---
echo "[env] Go project shell environment loaded with auto-cleaning."
