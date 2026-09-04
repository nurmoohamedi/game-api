#!/usr/bin/env bash
set -e

cd "$(dirname "$0")"

if ! docker info >/dev/null 2>&1; then
  echo "Docker daemon is not running, starting Docker Desktop..."
  open -a Docker
  echo "Waiting for Docker to start..."
  until docker info >/dev/null 2>&1; do
    sleep 1
  done
fi

docker compose up -d

echo "Waiting for Postgres to become healthy..."
until [ "$(docker inspect -f '{{.State.Health.Status}}' game-api-postgres 2>/dev/null)" = "healthy" ]; do
  sleep 1
done

echo "Postgres is up on localhost:5432 (db: game_api)"
