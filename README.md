# dockup

dockup is a small CLI to inspect and clean up Docker containers and images.

It is a thin wrapper around the Docker Engine API using the official Go docker client. The CLI reads the Docker host address from an environment variable (`DOCKER_HOST`) loaded from a `.env` file in the project root.

## Features

- List containers
- List images
- Interactively delete containers (single or all)
- (Planned) Rename containers

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
- `rename` — planned (currently not implemented)

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
```

## Why `./dockup image <arg>` prints the same digest

The `image` command implemented in `cmd/images.go` is a simple "list all images" command — it does not currently accept or use additional arguments. Passing arguments after `image` will be ignored by the current implementation and you'll always get the same listing output (the command prints image IDs and their tags). If you need `image` to filter by name/tag or accept arguments, that behavior must be implemented in the code (e.g. parse `args` inside the command's `Run` function and pass filter options to `cli.ImageList`).

## Where the code reads the Docker host

The code that reads the `.env` file and returns the client is in `cmd/list.go`:

- `loadEnv()` loads `.env` and returns `DOCKER_HOST`.
- `dockerClient()` creates the Docker client with `client.NewClientWithOpts(client.WithHost(host), client.WithAPIVersionNegotiation())`.

If you want to change how the host is discovered (for example to prefer environment variables set in the shell over `.env`), update `loadEnv()` accordingly.

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
