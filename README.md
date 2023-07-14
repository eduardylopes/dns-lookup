# DNS Lookup CLI

A command-line interface (CLI) tool to perform DNS lookups and retrieve information about public IPs and name servers.

## Installation

1. Make sure you have Go installed on your machine. If not, you can download and install it from the official Go website: [https://golang.org](https://golang.org)

2. Clone the repository or download the source code.

   ```
   git clone https://github.com/eduardylopes/dns-lookup-cli.git
   ```

3. Navigate to the project directory.

   ```
   cd dns-lookup-cli
   ```

4. Build the executable.

   ```
   go build -o dns-lookup
   ```

5. Add the executable to your system's PATH.

   - **Linux/Mac**:

     ```
     export PATH=$PATH:/path/to/dns-lookup-cli
     ```

   - **Windows**:
     Add the `dns-lookup-cli` directory to the `PATH` environment variable.

## Usage

```
dns-lookup [command] [flags]
```

Available commands:

- `ip`: Perform a search for public IPs.
- `ns`: Perform a search for name servers.

Flags:

- `--host`: Host name.

Examples:

1. Perform a search for public IPs:

   ```
   dns-lookup ip --host value1
   ```

2. Perform a search for name servers:

   ```
   dns-lookup ns --host value1
   ```
