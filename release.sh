#!/usr/bin/env bash

for goos in darwin windows linux; do
  GOOS=$goos GOARCH=amd64 ./build.sh build_backend
  GOOS=$goos GOARCH=amd64 ./build.sh build_agent
  GOOS=$goos GOARCH=amd64 ./build.sh build_cli

  if [[ "$goos" == "windows" ]]; then
    mv bin/sensu-agent bin/sensu-agent.exe
    mv bin/sensu-backend bin/sensu-backend.exe
    mv bin/sensuctl bin/sensuctl.exe
  fi

  rm -rf release/$goos
  mkdir -p release/$goos
  mv bin/sensu* release/$goos
done

tar czvf release-$(git rev-parse --abbrev-ref HEAD)-$(git rev-parse --short HEAD).tar.gz release/*
