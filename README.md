# Simple IT Website (Go)

A minimal, production-ready starter for a small IT services website built with Go's standard library.

## Features
- Standard library only (no external deps)
- Templating via `html/template` with a base layout and partials
- Static file serving (CSS/JS/images)
- Routes: Home, Services, About, Contact (with simple POST handler)
- Graceful shutdown
- Basic security headers
- Dockerfile and Makefile

## Project Structure
```text
it-website/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── handlers/
│   │   └── handlers.go
│   └── render/
│       └── render.go
├── web/
│   ├── static/
│   │   ├── css/
│   │   │   └── styles.css
│   │   ├── js/
│   │   │   └── main.js
│   │   └── img/
│   └── templates/
│       ├── base.tmpl
│       ├── home.tmpl
│       ├── services.tmpl
│       ├── about.tmpl
│       ├── contact.tmpl
│       └── partials/
│           ├── header.tmpl
│           └── footer.tmpl
├── Dockerfile
├── Makefile
└── .gitignore
```

## Quick Start
```bash
# 1) Build & run
go run ./cmd/server

# or
make run

# Visit http://localhost:8080
```

## Build
```bash
make build      # builds ./bin/server
make run        # runs the server
make tidy       # go mod tidy
```

## Docker
```bash
docker build -t it-website:latest .
docker run -p 8080:8080 it-website:latest
```

## Configuration
- Port defaults to `8080`. Override via `PORT` env var.
- Templates auto-load on startup; server logs template load errors and exits.

## Notes
- The contact form posts to `/contact` and just logs/echoes the data server-side.
- Add email or ticketing integration later as needed.
