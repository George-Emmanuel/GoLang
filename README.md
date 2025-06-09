<h1 align="center">🚀 GoLang Starter Project</h1>

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.22-blue?style=for-the-badge&logo=go" />
  <img src="https://img.shields.io/badge/License-MIT-green?style=for-the-badge" />
  <img src="https://img.shields.io/badge/PRs-Welcome-brightgreen?style=for-the-badge" />
</p>

---

<h2>📁 Project Structure</h2>

<pre>
project-name/
├── cmd/            # Application entry points
│   └── app/        # Main app
├── pkg/            # Library code to be used by external apps
├── internal/       # Private application code
├── config/         # Configuration files (YAML, JSON, ENV)
├── deploy/         # Docker, CI/CD, Helm charts etc.
├── scripts/        # Developer tools and automation scripts
├── test/           # External test data and mocks
├── go.mod          # Go module definition
├── go.sum          # Go module checksums
└── README.md
</pre>

---

<h2>⚙️ Setup</h2>

```bash
git clone https://github.com/yourusername/project-name.git
cd project-name
go mod tidy

go run ./cmd/app

go fmt ./...
go vet ./...
go test ./...

go build -o bin/app ./cmd/app
```
<h2>✅ Features</h2>
Idiomatic Go directory layout
Built-in logging and error handling
Graceful shutdown via context
Modular, testable components
Config via .env, JSON, or YAML
Docker-ready and CI/CD friendly

<h2>🛠️ Tech Stack</h2>
| Tool            | Purpose           |
| --------------- | ----------------- |
| Go              | Language          |
| slog/zap        | Logging           |
| chi/mux         | HTTP Routing      |
| cobra           | CLI Apps          |
| viper/envconfig | Config Management |

<h2>📚 References</h2>
<a href="https://github.com/golang-standards/project-layout" target="_blank">Go Project Layout</a>
<a href="https://golang.org/doc/effective_go.html" target="_blank">Effective Go</a>
<a href="https://github.com/tmrts/go-patterns" target="_blank">Go Patterns</a>
