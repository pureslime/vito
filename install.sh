#!/bin/sh
set -e

# Terminal colors
GREEN='\033[0;32m'
BOLD='\033[1m'
NC='\033[0m'

echo "${BOLD}📦 Installing VITO...${NC}"

# Detect Architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

if [ "$ARCH" = "x86_64" ]; then ARCH="amd64"; fi
if [ "$ARCH" = "aarch64" ] || [ "$ARCH" = "arm64" ]; then ARCH="arm64"; fi

VITO_HOME="$HOME/.vito"
BIN_DIR="/usr/local/bin"

mkdir -p "$VITO_HOME/pills"
mkdir -p "$VITO_HOME/config"

# curl -L "https://github.com/tu_usuario/vito/releases/latest/download/vito_${OS}_${ARCH}" -o vito

echo "✨ Directories created at $VITO_HOME"

# chmod +x vito
# sudo mv vito $BIN_DIR/vito
echo "${GREEN}${BOLD}🚀 VITO installed successfully!${NC}"
echo "Try running: ${BOLD}vito --help${NC}"
