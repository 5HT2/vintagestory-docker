#!/bin/bash

# shellcheck disable=SC1091
source "$HOME/.env"

if [[ -z "$VINTAGESTORY_PATH" ]]; then
  echo "VINTAGESTORY_PATH not set!"
  exit 1
fi

docker stop vintagestory || echo "Could not stop missing container vintagestory"
docker rm vintagestory || echo "Could not remove missing container vintagestory"

docker run --name vintagestory \
  --mount type=bind,source="$VINTAGESTORY_PATH",target=/root/.config/VintagestoryData \
  --network host -d vintagestory
