//go:build !linux || !amd64

// Stub identique à Linux : UnavailableError typé, jamais d'import cilium/ebpf.
// C'est ce qui rend `GOOS=linux CGO_ENABLED=0 go vet` vert cross-OS tout en
// gardant la promesse « jamais ebpf sans objet et sans capacités ».
package bpf

import (
	"errors"
)

// Session — stub: API miroir du loader (chargeur Linux ailleurs).
// Note: import du package probe droppé (contient pkg/model) — le stub
// reste honnête sans dépendance tierce.

var (
	ErrSessionUnavailable = errors.New("ebpf : session indisponible (stub hors Linux/amd64)")
)

// Open — stub: le chargeur Linux n'existe pas ici, la session le dit.
func Open() (*Session, error) {
	return nil, &UnavailableError{Reason: "hors Linux/amd64, loader absent (CGO-DEV SSH pending)"}
}

// Check — stub: nomme le fait manquant (chargeur absent).
func Check() error {
	return &UnavailableError{Reason: "chargeur absent hors Linux/amd64"}
}

// Session — stub: API miroir du loader.
type Session struct{}

// Lookup — stub: jamais trouvé.
func (s *Session) Lookup(local, peer string) (Flow, bool) { return Flow{}, false }

// Snapshot — stub: vide (absence ≠ 0).
func (s *Session) Snapshot() []Flow { return nil }

// Globals — stub: raison nommée.
func (s *Session) Globals() (Globals, error) {
	return Globals{}, &UnavailableError{Reason: "hors Linux/amd64"}
}

// Attach — stub: "", jamais "fentry".
func (s *Session) Attach() string { return "" }

// Close — stub: rien à fermer.
func (s *Session) Close() error { return nil }
