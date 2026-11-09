# cgo-bpf-observer

Télémétrie eBPF côté chargeur (CO-RE, cilium/ebpf) pour le banc CGO :
réémissions TCP, RTT socket, drops qdisc — épinglés `v0.1.0` stub-first.

- Module : `github.com/Realms4239/cgo-bpf-observer` (`go 1.25`)
- État : **stub** — loader Linux absent (`internal/probe/bpf/session_stub.go`),
  la session rend `UnavailableError` honnêtement. `bpf_probe.o` (build
  Linux 6.8 `clang-18`, `make bpf`+`make bpf-check` sur l'hôte CGO-DEV)
  n'est pas encore livré : la première capture live reste BLOCKED
  (gate `docs/runbooks/bpf-observer-gates.md` dans le monolithe CGO,
  monolith doc `docs/runbooks/bpf-observer.md`).
- Armure : stdlib seule dans le stub ; le loader Linux exigera
  `github.com/cilium/ebpf v0.22.0` (approuvé humain) — pas avant l'objet.

## Verif

```bash
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go vet ./...
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build ./...
```
