# dockup

A CLI tool for managing Docker containers and images.

## Docker Host Configuration

The tool requires a Docker host connection, which is configured through the `DOCKER_HOST` environment variable in a `.env` file. Here's how to set it up:

1. Create a `.env` file in the project root:

```env
# For local Docker socket (Linux/macOS/Windows with WSL2):
DOCKER_HOST=unix:///var/run/docker.sock

# For Docker Desktop over TCP:
DOCKER_HOST=tcp://host.docker.internal:2375  # macOS/Windows
# or
DOCKER_HOST=tcp://docker-desktop:2375        # some setups
```

Replace `host.docker.internal` or `docker-desktop` with your Docker Desktop host's actual hostname.

> **Note**: Using TCP without TLS (port 2375) is insecure and should only be used for local development.

## Features

- List containers
- List images
- Interactively delete containers (single or all)
- Interactively delete images (single or all)

## Prerequisites

- Go 1.20+ (or whatever your `go.mod` requires)
- Docker (local Docker daemon or Docker Desktop)

## Build / Install

To build the binary locally:

```bash
go build -o dockup
```

Or install to your `$GOBIN`:

```bash
go install ./...
```

After building you can run `./dockup` from the repository root.

## Configuration (DOCKER_HOST)

`dockup` reads the Docker host address from a `.env` file using the `DOCKER_HOST` environment variable. Create a file named `.env` in the project root with a line like:

```env
DOCKER_HOST=unix:///var/run/docker.sock
```

Common `DOCKER_HOST` values:

- Local Linux (default Docker socket):
	- `unix:///var/run/docker.sock`
- Docker Desktop (macOS / Windows) using the local socket (default):
	- `unix:///var/run/docker.sock` (on systems where Docker provides a socket)
- Docker daemon exposed over TCP (insecure; enable only for local testing):
	- `tcp://<DOCKER_HOSTNAME_OR_IP>:2375`

Important: if you use Docker Desktop and want to connect via TCP, set `DOCKER_HOST` to the hostname or address of *your Docker Desktop host*. For example, on some setups you can use:

```env
DOCKER_HOST=tcp://host.docker.internal:2375
# or
DOCKER_HOST=tcp://docker-desktop:2375
```

Replace `host.docker.internal` or `docker-desktop` with the actual hostname or IP for your environment. The key point: use the host name/address for your Docker Desktop host when connecting over TCP.

Security note: exposing the Docker daemon over an unencrypted TCP port (2375) is insecure. Prefer the local Unix socket or configure TLS for the Docker API in production.

## Usage

General form:

```bash
./dockup <command> [flags]
```

Available commands (from the codebase):

- `container` — list containers
	- `container delete` — interactive delete of a container
		- `-a` or `--all` to delete all containers (interactive confirmation)
- `image` — list images
	- `image rmi` — interactive delete of an image
		- `a` or `all` to delete all images(interactive cofirmation)

Examples:

```bash
# list containers
./dockup container

# list images
./dockup image

# delete a container (interactive)
./dockup container delete

# delete all containers (interactive confirmation)
./dockup container delete -a

# remove an image (interactive)
./dockup image rmi

# delete all images (interactive confirmation)
./dockup image rmi -a
```

## Troubleshooting

- If you get connection errors, double-check your `.env` and `DOCKER_HOST` value.
- For macOS/Windows Docker Desktop, either use the local socket (if available) or configure Docker Desktop to expose the daemon on TCP and use the correct host name (see the Configuration section above).
- If you see the same output for different arguments to `image`, it's because the command currently ignores args — see the note above on how to add filtering.

## Contributing

Contributions welcome. Please open issues or pull requests. If you implement argument parsing or filtering for the `image` command, add unit tests and update this README with examples.

---

If you'd like, I can also:

- add argument parsing to `image` so it can accept a filter (e.g. `./dockup image <name>`),
- or implement `rename` to rename containers interactively.

Tell me which you'd prefer and I will implement it.
