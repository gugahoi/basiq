# Basiq

`basiq` is a CLI tool for interacting with the Basiq API.

Open API specification: [https://api.basiq.io/openapi](https://api.basiq.io/openapi)

## Installation

### Download a release

Download the latest binary for your platform from the [releases page](../../releases/latest), extract it, and move it to a directory in your `$PATH`.

```bash
# Example for macOS (arm64)
curl -L https://github.com/gugahoi/basiq/releases/latest/download/basiq_Darwin_arm64.tar.gz | tar xz
mv basiq /usr/local/bin/
```

### Build from source

Requires Go 1.21+:

```bash
git clone https://github.com/gugahoi/basiq.git
cd basiq
make build
mv basiq /usr/local/bin/
```

## Configuration

Set your Basiq API key as an environment variable:

```bash
export BASIQ_APIKEY=<your-api-key>
```

## Supported commands

### Webhooks

*   [x] [create](https://api.basiq.io/reference/addwebhook)
*   [x] [list](https://api.basiq.io/reference/listappwebhooks)
*   [x] [get](https://api.basiq.io/reference/getwebhook)
*   [x] [delete](https://api.basiq.io/reference/deletewebhook)
*   [x] [update](https://api.basiq.io/reference/updatewebhook)

### Events

*   [x] [get](https://api.basiq.io/reference/gettypebyid)
*   [x] [list](https://api.basiq.io/reference/listeventtypes)
*   [x] [listall](https://api.basiq.io/reference/getevents)
*   [x] [types](https://api.basiq.io/reference/geteventtypebyid)
*   [x] [test](https://api.basiq.io/reference/testmessage)

## Shell Completion

Shell completion is supported out of the box:

```bash
basiq completion <bash|zsh|fish>

# Example setup for ZSH, in your ~/.zshrc add:
if type "basiq" &>/dev/null; then
    source <(eval "basiq completion zsh")
fi
```
