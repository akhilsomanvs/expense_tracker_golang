package handlers

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
	ConnPool *pgxpool.Pool
}

func (h *Handler) DBHealthCheck(w http.ResponseWriter, r *http.Request) {
	var version string
	err := h.ConnPool.QueryRow(
		r.Context(),
		"SELECT version()",
	).Scan(&version)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(version))
}
