package handlers

import (
    "fmt"
    "log"
    "net/http"
    "time"

    "it-website/internal/render"
)

type Handlers struct {
    tpl *render.Templates
}

func New(tpl *render.Templates) *Handlers {
    return &Handlers{tpl: tpl}
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
    data := map[string]any{
        "Title":   "Gbenga IT — Cloud & DevOps Solutions",
        "Tagline": "Reliable cloud engineering, Kubernetes, and automation for modern teams.",
    }
    if err := h.tpl.Render(w, "home.tmpl", data); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func (h *Handlers) Services(w http.ResponseWriter, r *http.Request) {
    data := map[string]any{
        "Title": "Services",
        "List": []struct {
            Name string
            Desc string
        }{
            {"Cloud Architecture", "Design secure, scalable systems on Azure, AWS, or GCP."},
            {"Kubernetes & AKS", "Cluster design, GitOps, and day-2 operations."},
            {"CI/CD & Automation", "Pipelines with GitHub Actions, ArgoCD, Terraform."},
            {"Observability", "Logging, metrics, tracing, SLOs and diagnostics."},
        },
    }
    if err := h.tpl.Render(w, "services.tmpl", data); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func (h *Handlers) About(w http.ResponseWriter, r *http.Request) {
    data := map[string]any{
        "Title":   "About",
        "Content": "We are an IT consulting team specializing in Cloud, DevOps, and Kubernetes. We help teams deliver faster and safer.",
    }
    if err := h.tpl.Render(w, "about.tmpl", data); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func (h *Handlers) Contact(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        data := map[string]any{
            "Title": "Contact",
        }
        if err := h.tpl.Render(w, "contact.tmpl", data); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    case http.MethodPost:
        if err := r.ParseForm(); err != nil {
            http.Error(w, "invalid form", http.StatusBadRequest)
            return
        }
        name := r.Form.Get("name")
        email := r.Form.Get("email")
        message := r.Form.Get("message")
        log.Printf("[contact] %s <%s> at %s: %s", name, email, time.Now().Format(time.RFC3339), message)

        // Echo back a simple confirmation
        w.Header().Set("Content-Type", "text/plain")
        fmt.Fprintf(w, "Thanks, %s! We received your message.", name)
    default:
        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
    }
}
