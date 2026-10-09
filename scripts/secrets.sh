#!/usr/bin/env bash
set -euo pipefail

name=nlm-env
ps1=$(wslpath -w "$(dirname "$0")/credman.ps1")
cm() { powershell.exe -NoProfile -ExecutionPolicy Bypass -File "$ps1" "$@" | tr -d '\r'; }

case "$1" in
push)
	base64 -w0 "$HOME/.nlm/env" | cm set "$name"
	shred -u "$HOME/.nlm/env"
	;;
exec)
	shift
	env_text=$(cm get "$name" | base64 -d)
	while IFS='=' read -r key value; do
		[ -n "$key" ] && export "$key=$(printf '%s' "$value" | sed 's/^"//;s/"$//')"
	done <<<"$env_text"
	exec "$@"
	;;
clear)
	cm del "$name"
	;;
esac
