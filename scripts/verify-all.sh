#!/usr/bin/env bash
set -euo pipefail

project_dir="$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd)"
cd "$project_dir"

unformatted="$(gofmt -l entity event match)"
if [[ -n "$unformatted" ]]; then
	printf 'gofmt required:\n%s\n' "$unformatted" >&2
	exit 1
fi

GOCACHE="${GOCACHE:-/tmp/lumi-spec-go-cache}" go vet ./...
GOCACHE="${GOCACHE:-/tmp/lumi-spec-go-cache}" go test ./...
ruby -rjson -e 'Dir["schema/*.json"].each { |path| JSON.parse(File.read(path)); puts "valid #{path}" }'

printf 'lumi-spec verification passed\n'
