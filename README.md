[![Build & Test](https://github.com/rogerwesterbo/openfeature/actions/workflows/build-and-test.yml/badge.svg)](https://github.com/rogerwesterbo/openfeature/actions/workflows/build-and-test.yml)

# OpenFeature

Showing how to use OpenFeature in Golang API

Reference docs:

- https://openfeature.dev/docs/reference/intro
- https://openfeature.dev/docs/tutorials/getting-started/go

# Prerequisites

- Some where to run containers (Docker, docker desktop https://docs.docker.com/get-docker, Podman https://podman.io, etc)
- Golang 1.2x (https://go.dev/)

## Optional requisites:

- Visual Studio Code (for debugging)

# Run the demo

1. open a terminal and go to this repo root
2. `docker compose up ` (`Ctrl+c` to abort)
3. open a browser and go to http://localhost:7357/hello
4. open file `flagd/flags.flagd.json`
   - change flags.welcome-message.defaultVariant to "`on`"
5. refresh site on http://localhost:7357/hello
   - the text should be different 😊

# Run with VSCode debugger

1. open a terminal and go to this repo root
2. `code .`
   - VS Code should start 😎
3. in the previous terminal or a new terminal in VS Code, run `go get ./...`
4. new terminal, run `docker compose up flagd` (`Ctrl+x` to abort)
5. in VS Code, go to the "Run and Debug" section (windows press `Ctrl+Shift+d`), and press ▶️ the play button (or `F5`)
6. View the "Debug Console" to see output

# To test in a cluster

- https://openfeature.dev/docs/tutorials/ofo/
- https://killercoda.com/open-feature/scenario/openfeature-operator-demo
- https://artifacthub.io/packages/helm/open-feature-operator/open-feature-operator
