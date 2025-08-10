# SCAN-ON-GO

A simple, fast concurrent port scanner written in Go.

## Features

- Concurrent TCP port scanning
- Customizable port range and timeout
- Save results to a file
- CLI ASCII art banner

## Usage

### 1. Build or Run

**To run directly:**
```sh
go run main.go -target <target_ip_or_domain> [options]
```

**To build:**
```sh
go build -o scan-on-go main.go
./scan-on-go -target <target_ip_or_domain> [options]
```

### 2. Options

| Flag      | Description                  | Default   |
|-----------|------------------------------|-----------|
| -target   | Target IP or domain (required) |           |
| -start    | Start port                   | 1         |
| -end      | End port                     | 1024      |
| -timeout  | Timeout in ms                | 1000      |
| -save     | Save output to file          | (none)    |

### 3. Example

```sh
go run main.go -target 127.0.0.1 -start 1 -end 100 -timeout 500 -save results.txt
```

### 4. Output

- Results are printed in the terminal.
- If `-save` is specified, results are also saved to the given file.

---

## ASCII Art Banner

The scanner displays a hash-style ASCII art banner on startup.

---

## Project Structure

```
Scan-on-GO/
├── main.go
├── scanner/
│   └── scanner.go
└── go.mod
```

---

## License

MIT