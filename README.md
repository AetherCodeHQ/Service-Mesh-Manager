# 🌐 Service Mesh Manager

![Go](https://img.shields.io/badge/Go-1.21%2B-00ADD8?style=flat-square&logo=go&logoColor=white)
![Version](https://img.shields.io/badge/Version-v2.1.0-00ADD8?style=flat-square)
![License](https://img.shields.io/badge/License-MIT-green?style=flat-square)
![PRs](https://img.shields.io/badge/PRs-Welcome-brightgreen?style=flat-square)

> Infrastructure tool by [AetherCodeHQ](https://github.com/AetherCodeHQ)

`infrastructure` `devops` `cli` `golang` `json`

---

## What is Service-Mesh-Manager?

**Service-Mesh-Manager** is an infrastructure tool for monitoring, inspecting, and managing systems and services.

## Features

- ✅ JSON data handling
- 🚀 **Zero dependencies** — only Go standard library
- 📦 **Single binary** — compile and run anywhere
- 🔄 **Offline capable** — no internet required

## Installation

```bash
# Clone
git clone https://github.com/AetherCodeHQ/Service-Mesh-Manager.git
cd Service-Mesh-Manager

# Build
go build -o service-mesh-manager .

# Run
./service-mesh-manager <registry.json> [list|register|route]
```

### Or directly with `go run`:
```bash
go run main.go <registry.json> [list|register|route]
```

## Usage

```bash
# Basic usage
./service-mesh-manager <registry.json> [list|register|route]

# With flags
./service-mesh-manager <registry.json> [list|register|route] value <registry.json> [list|register|route]
```

### Example Output

```
$ ./service-mesh-manager <registry.json> [list|register|route]
<registry.json> [list|register|route]
%-20s %s:%d %v\n
\n%d services registered\n
```

## Project Structure

```
Service-Mesh-Manager/
  main.go          # Entry point (65 lines)
  go.mod            # Go module definition
  go.sum            # Dependency checksums
  README.md         # This file
  LICENSE           # MIT License
  CHANGELOG.md      # Version history
```

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

MIT License - see [LICENSE](LICENSE) for details.

---

Built with ❤️ by [AetherCodeHQ](https://github.com/AetherCodeHQ)
