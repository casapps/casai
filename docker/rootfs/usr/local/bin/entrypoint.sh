#!/usr/bin/env bash
# shellcheck shell=bash
# entrypoint.sh — container startup script
# Called by: tini → entrypoint.sh → app
# Prepares cache/config/data dirs; the binary itself creates its own
# user/group and drops privileges (AI.md PART 5 → "Container Runtime Rules").
mkdir -p /home/casai/.cache/casapps/casai /home/casai/.config/casapps/casai /home/casai/.local/share/casapps/casai

exec "$@"
