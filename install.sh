#!/usr/bin/env bash
set -e

echo Installing compose-cli...

TARGET_DIR="$HOME/.compose-cli/bin"
TARGET_DIR_BIN="$TARGET_DIR/compose-cli"

SERVER="https://github.com/Vanclief/compose-cli/raw/master/bin"

# Detect the platform (similar to $OSTYPE)
OS="$(uname)"
if [[ "$OS" == "Linux" ]]; then
	# Linux
	FILENAME="compose-cli-linux"
elif [[ "$OS" == "Darwin" ]]; then
	# MacOS, should validate if Intel or ARM
	UNAMEM="$(uname -m)"
	if [[ "$UNAMEM" == "x86_64" ]]; then
		FILENAME="compose-cli-mac-intel"
	else
		FILENAME="compose-cli-mac-arm64"
	fi
else
	echo "unrecognized OS: $OS"
	echo "Exiting..."
	exit 1
fi

# Check if ~/.compose-cli/bin exists, if not create it
if [[ ! -e "${TARGET_DIR}" ]]; then
	mkdir -p "${TARGET_DIR}"
fi

# Download the appropriate binary
echo "Downloading $SERVER/$FILENAME..."
curl -# -L "${SERVER}/${FILENAME}" -o "${TARGET_DIR_BIN}"
chmod +x "${TARGET_DIR_BIN}"
echo "Installed under ${TARGET_DIR_BIN}"

# Store the correct profile file (i.e. .profile for bash or .zshenv for ZSH).
case $SHELL in
*/zsh)
	PROFILE=${ZDOTDIR-"$HOME"}/.zshenv
	PREF_SHELL=zsh
	;;
*/bash)
	PROFILE=$HOME/.bashrc
	PREF_SHELL=bash
	;;
*/fish)
	PROFILE=$HOME/.config/fish/config.fish
	PREF_SHELL=fish
	;;
*/ash)
	PROFILE=$HOME/.profile
	PREF_SHELL=ash
	;;
*)
	echo "could not detect shell, manually add ${TARGET_DIR_BIN} to your PATH."
	exit 1
	;;
esac

# Only add if it isn't already in PATH.
if [[ ":$PATH:" != *":${TARGET_DIR_BIN}:"* ]]; then
	# Add the directory to the path and ensure the old PATH variables remain.
	echo >>$PROFILE && echo "export PATH=\"\$PATH:$TARGET_DIR\"" >>$PROFILE
fi

echo && echo "Detected your preferred shell is ${PREF_SHELL} and added compose-cli to PATH. Run 'source ${PROFILE}' or start a new terminal session to use compose-cli."

# Confirmation
echo "compose-cli successfully installed!"
