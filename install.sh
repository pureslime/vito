#!/bin/sh
set -e

# Configuración
VITO_HOME="$HOME/.vito"
BIN_DIR="$VITO_HOME/bin"

echo "🚀 Installing VITO from GitHub..."

# Detección de OS y ARCH
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)
[ "$ARCH" = "x86_64" ] && ARCH="amd64"
[ "$ARCH" = "aarch64" ] || [ "$ARCH" = "arm64" ] && ARCH="arm64"

# Construir URL del asset
# Ejemplo: https://github.com/pureslime/vito/releases/latest/download/vito-linux-amd64
URL="https://github.com/pureslime/VITO/releases/download/latest/vito"

# Crear carpetas
mkdir -p "$BIN_DIR" "$VITO_HOME/pills" "$VITO_HOME/config"

# Descargar
echo "📥 Downloading $URL..."
if ! curl -sSL "$URL" -o "$BIN_DIR/vito"; then
    echo "❌ Error: Could not download binary. Check if the release asset exists."
    exit 1
fi

chmod +x "$BIN_DIR/vito"

echo "✅ Installed at $BIN_DIR/vito"
echo "👉 Run 'export PATH=\"\$HOME/.vito/bin:\$PATH\"' to use it."
