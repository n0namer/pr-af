#!/bin/sh
set -u
BIN="/afhome/opencode-runtime/v1.17.15/opencode"
[ -x "$BIN" ] || { echo "OpenCode runtime missing" >&2; exit 126; }
export OPENCODE_CONFIG="/src/swe-af/opencode.json"
export OPENCODE_DISABLE_AUTOUPDATE=1
if [ "$#" -eq 8 ] && [ "$6" = "-m" ]; then
  if [ -n "${LLM_BROKER_BASE_URL:-}" ]; then
    set -- "$1" "$2" "$3" "$4" "$5" "$6" "fcm/${LLM_BROKER_MODEL:-fcm}" "$8"
  else
    model="$7"
    case "$model" in openai/*) model="compat/${model#openai/}" ;; esac
    set -- "$1" "$2" "$3" "$4" "$5" "$6" "$model" "$8"
  fi
fi
exec "$BIN" --print-logs --log-level DEBUG "$@"
