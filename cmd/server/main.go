package main

import (
    "context"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"

    "it-website/internal/handlers"
    "it-website/internal/render"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    // Load templates at startup
    tpl, err := render.LoadTemplates("web/templates")
    if err != nil {
        log.Fatalf("error loading templates: %v", err)
    }

    mux := http.NewServeMux()

    // Static files
    fileServer := http.FileServer(http.Dir("web/static"))
    mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

    // Routes
    h := handlers.New(tpl)
    mux.HandleFunc("/", h.Home)
    mux.HandleFunc("/services", h.Services)
    mux.HandleFunc("/about", h.About)
    mux.HandleFunc("/contact", h.Contact)

    // Wrap with security headers
    handler := securityHeaders(mux)

    srv := &http.Server{
        Addr:              ":" + port,
        Handler:           handler,
        ReadTimeout:       10 * time.Second,
        ReadHeaderTimeout: 10 * time.Second,
        WriteTimeout:      15 * time.Second,
        IdleTimeout:       60 * time.Second,
    }

    // Graceful shutdown
    go func() {
        log.Printf("server listening on http://localhost:%s", port)
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("listen: %s\n", err)
        }
    }()

    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    log.Println("shutting down server...")

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    if err := srv.Shutdown(ctx); err != nil {
        log.Fatalf("server forced to shutdown: %v", err)
    }
    log.Println("server exiting")
}

func securityHeaders(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Basic hardening
        w.Header().Set("X-Content-Type-Options", "nosniff")
        w.Header().Set("X-Frame-Options", "DENY")
        w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
        w.Header().Set("X-XSS-Protection", "0")
        w.Header().Set("Content-Security-Policy", "default-src 'self'; img-src 'self' data:; script-src 'self'; style-src 'self' 'unsafe-inline';")
        next.ServeHTTP(w, r)
    })
}
