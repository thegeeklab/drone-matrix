---
title: drone-matrix
---

[![Build Status](https://img.shields.io/drone/build/thegeeklab/drone-matrix?logo=drone&server=https%3A%2F%2Fdrone.thegeeklab.de)](https://drone.thegeeklab.de/thegeeklab/drone-matrix)
[![Docker Hub](https://img.shields.io/badge/dockerhub-latest-blue.svg?logo=docker&logoColor=white)](https://hub.docker.com/r/thegeeklab/drone-matrix)
[![Quay.io](https://img.shields.io/badge/quay-latest-blue.svg?logo=docker&logoColor=white)](https://quay.io/repository/thegeeklab/drone-matrix)
[![GitHub contributors](https://img.shields.io/github/contributors/thegeeklab/drone-matrix)](https://github.com/thegeeklab/drone-matrix/graphs/contributors)
[![Source: GitHub](https://img.shields.io/badge/source-github-blue.svg?logo=github&logoColor=white)](https://github.com/thegeeklab/drone-matrix)
[![License: MIT](https://img.shields.io/github/license/thegeeklab/drone-matrix)](https://github.com/thegeeklab/drone-matrix/blob/main/LICENSE)

Drone plugin to to send messages to a Matrix room.

<!-- prettier-ignore-start -->
<!-- spellchecker-disable -->
{{< toc >}}
<!-- spellchecker-enable -->
<!-- prettier-ignore-end -->

## Usage

```YAML
kind: pipeline
name: default

steps:
- name: notify
  image: thegeeklab/matrix
  settings:
    homeserver: https://matrix.org
    roomid: abcdefghijklmnopqrstuvwxyz:matrix.org
    username: octocat
    password: secret
```

### Parameters

<!-- prettier-ignore-start -->
<!-- spellchecker-disable -->
{{< propertylist name=drone-matrix.data sort=name >}}
<!-- spellchecker-enable -->
<!-- prettier-ignore-end -->

## Build

Build the binary with the following command:

```Shell
export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0
export GO111MODULE=on

make build
```

Build the Docker image with the following command:

```Shell
docker build --file docker/Dockerfile.amd64 --tag thegeeklab/drone-matrix .
```

## Test

```Shell
docker run --rm \
  -e PLUGIN_ROOMID=0123456789abcdef:matrix.org \
  -e PLUGIN_USERNAME=yourbot \
  -e PLUGIN_PASSWORD=p455w0rd \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  thegeeklab/drone-matrix
```
