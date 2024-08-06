#!/bin/bash

REPO="julianofirme/clamper"
LATEST_RELEASE=$(curl -s https://api.github.com/repos/$REPO/releases/latest | grep "tag_name" | awk '{print substr($2, 2, length($2)-3) }')
OS=$(uname -s)
ARCH=$(uname -m)

case $OS in
    Linux)
        BINARY="clamper-linux"
        ;;
    Darwin)
        BINARY="clamper-darwin"
        ;;
    CYGWIN*|MINGW32*|MSYS*|MINGW*)
        BINARY="clamper-windows.exe"
        ;;
    *)
        echo "Unsupported OS: $OS"
        exit 1
        ;;
esac

echo "Downloading the latest release ($LATEST_RELEASE) for $OS..."
curl -LO "https://github.com/$REPO/releases/download/$LATEST_RELEASE/$BINARY"

if [ $? -ne 0 ]; then
    echo "Failed to download the binary. Please check your internet connection or the repository URL."
    exit 1
fi

chmod +x $BINARY
sudo mv $BINARY /usr/local/bin/clamper

if [ $? -ne 0 ]; then
    echo "Failed to move the binary to /usr/local/bin. Please check your permissions."
    exit 1
fi

echo "clamper has been successfully installed in /usr/local/bin"

# Verify installation
if ! command -v clamper &> /dev/null; then
    echo "Installation failed. Please check the installation steps or your system configuration."
    exit 1
else
    echo "clamper is installed and ready to use!"
fi