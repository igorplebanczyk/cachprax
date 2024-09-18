# Cachprax
Cachprax is a simple caching proxy server. It was made as a learning project, and was not tested in a production environment (caching proxies are pretty much obsolete anyway).
<br><br>
*Requires Go 1.22 or later.*

## Features
* Runs a caching proxy server that caches responses from an origin server on a given port.<br>
* The cache items expire after a given time (default 5 minutes).<br>
* The cache is purged after a given time (default 10 minutes).<br>
* The user can start, stop, get the status of the proxy, clear the cache and get the number of items in it.<br>
* A simple command for testing the connection to the origin server is provided.

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

## Commands:
* Start the caching proxy server:
    ```bash
    cachprax start --origin "http://httpbin.org" --port 3000
    --origin: The URL of the origin server (required).
    --port: The port on which the caching proxy server will listen (required).
    --cache-expire: Cache expiration time in minutes (optional, default is 5 minutes).
    --cache-purge: Cache purge time in minutes (optional, default is 10 minutes).
    ```

* Stop the caching proxy server:
    ```bash
    cachprax stop
    ```

* Get the status of the caching proxy server:
    ```bash
    cachprax status
    ```

* Clear the cache:
    ```bash
    cachprax cache --clear
    ```
  
* Get the number of cached items:
    ```bash
    cachprax cache --count
    ```

* Test connection to the origin server:
    ```bash
    cachprax conntest --origin "http://httpbin.org"
    ```
## Notes
* Port 3001 is reserved for internal usage.
* A `cachprax.json` file is created in the `tmp` directory to store server metadata.