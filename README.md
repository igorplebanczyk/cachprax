# Cachprax

Cachprax is a simple caching proxy server. It was made as a learning project, and was not tested in a production
environment (caching proxies are pretty much obsolete anyway). Any feedback is welcome.
<br><br>
*Requires a Linux system*
<br>
*Requires Go 1.22 or later.*

## Features

* Runs a caching proxy server that caches responses from an origin server on a given port.<br>
* The cache items expire after a given time (default 10 minutes).<br>
* The cache is purged after a given time (default 30 minutes).<br>
* The user can start, stop, get the status of the proxy, clear the cache and get the number of items in it.<br>
* A simple command for testing the connection to the origin server is provided.
* Configuration values can be managed using the `config` command.

## Commands:

* **Start the caching proxy server**:
    ```bash
    cachprax start
    --origin: The URL of the origin server.
    --port: The port on which the caching proxy server will listen.
    --cache-expire: Cache expiration time in minutes.
    --cache-purge: Cache purge time in minutes.
    ```
  *If any of the flags are not provided, the default values will be used.*


* **Stop the caching proxy server**:
    ```bash
    cachprax stop
    ```

* **Get the status of the caching proxy server**:
    ```bash
    cachprax status
    ```

* **Manage the cache**:
    ```bash
    cachprax cache
    --clear: Clear the cache.
    --count: Get the number of cached items.
    ```

* **Test connection to the origin server**:
    ```bash
    cachprax conntest --origin "http://example.com"
    ```

* **Manage configuration values**:
    ```bash
    cachprax config 
    --set [key] [value]: Set a specific configuration value.
    --reset: Reset the configuration file to default values.
    ```

## Configuration

Config values are stored in a `cachprax.yaml` file in the home directory.<br>
They can be modified directly in the file or using the `config` command.<br>
Config values are only used if the corresponding flags are not provided when starting the server, with the exception of
`cache_port` which can not be changed through a flag and is thus always pulled from the file.

### Values

* `origin`: The default URL of the origin server.
* `proxy_port`: The default port on which the caching proxy server will listen.
* `cache_expire`: The default cache expiration time in minutes.
* `cache_purge`: The default cache purge time in minutes.
* `cache_port`: The port on which the server used for cache operations will listen.

### Default Values

```yaml
origin: ""
proxy_port: 3000
cache_expire: 10
cache_purge: 30
cache_port: 3001
```

## Installation

### Option 1: Install from GitHub Releases (Precompiled Binaries)

To install Cachprax using a precompiled binary from the GitHub releases:

* Go to the Cachprax Releases page and download the latest binary for your operating system.

* After downloading, make the binary executable:
    ```bash
    chmod +x cachprax
    ```

* Move the binary to a directory in your $PATH (e.g., /usr/local/bin):
    ```bash
    sudo mv cachprax /usr/local/bin/
    ```

* Confirm that Cachprax is installed by running:
    ```bash
    cachprax --help
    ```

### Option 2: Compile and Install from Source

To compile and install the application yourself:

* Ensure that you have Go installed by running:
    ```bash
    go version
    ```

* Clone the repository:
    ```bash
    git clone https://github.com/your-repo/cachprax.git
    cd cachprax
    ```

* Build the application:
    ```bash
    go build -o cachprax
    ```

* Move the compiled binary to a directory in your $PATH:
    ```bash
    sudo mv cachprax /usr/local/bin/
    ```

* Confirm that Cachprax is installed by running:
    ```bash
    cachprax --help
    ```

## Notes

* The server is always started as a separate process, whose PID is outputed to the standard output when starting the
  server.
* Cache operations (`cache --clear` and `cache --count`) are performed on a separate server running on a different port,
  which can be changed through the config file.
* A `cachprax.yaml` file is created in the home directory to store configuration values.
* A `cachprax.json` file is created in the `tmp` directory to store server metadata of the currently running server.
  This file is used for `stop` and `status` commands. Deleting this file will cause those commands to not work, in which
  case to stop the server, one must manually kill its process.
* Designed as part of this [project](https://roadmap.sh/projects/caching-server)