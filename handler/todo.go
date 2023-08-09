package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/TechBowl-japan/go-stations/model"
	"github.com/TechBowl-japan/go-stations/service"
)

// A TODOHandler implements handling REST endpoints.
type TODOHandler struct {
	svc *service.TODOService
}

// NewTODOHandler returns TODOHandler based http.Handler.
func NewTODOHandler(svc *service.TODOService) *TODOHandler {
	return &TODOHandler{
		svc: svc,
	}
}

func (h *TODOHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var req model.CreateTODORequest

		// JSONデコード
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&req)
		if err != nil {
			http.Error(w, "Failed to decode request", http.StatusBadRequest)
			return
		}

		// subjectの値を確認
		if req.Subject == "" {
			http.Error(w, "subject cannot be empty", http.StatusBadRequest)
			return
		} else {
			resp, err := h.Create(r.Context(), &req)
			if err != nil {
				http.Error(w, "Creation failed", http.StatusInternalServerError)
			}
			encoder := json.NewEncoder(w)
			if err := encoder.Encode(resp); err != nil {
				http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			}
		}
	} else {
		fmt.Fprintln(w, "This is NOT a POST request")
	}
}

// Create handles the endpoint that creates the TODO.
func (h *TODOHandler) Create(ctx context.Context, req *model.CreateTODORequest) (*model.CreateTODOResponse, error) {
	if h.svc == nil {
		fmt.Println(&h.svc)
		return nil, errors.New("TODOService is not initialized")
	}
	todo, err := h.svc.CreateTODO(ctx, req.Subject, req.Description)
	if err != nil {
		return nil, err
	}

	// 保存したTODOをレスポンスとして返す
	resp := model.CreateTODOResponse{
		TODO: *todo,
	}
	fmt.Println(resp)

	return &resp, err
}

// Read handles the endpoint that reads the TODOs.
func (h *TODOHandler) Read(ctx context.Context, req *model.ReadTODORequest) (*model.ReadTODOResponse, error) {
	_, _ = h.svc.ReadTODO(ctx, 0, 0)
	return &model.ReadTODOResponse{}, nil
}

// Update handles the endpoint that updates the TODO.
func (h *TODOHandler) Update(ctx context.Context, req *model.UpdateTODORequest) (*model.UpdateTODOResponse, error) {
	_, _ = h.svc.UpdateTODO(ctx, 0, "", "")
	return &model.UpdateTODOResponse{}, nil
}

// Delete handles the endpoint that deletes the TODOs.
func (h *TODOHandler) Delete(ctx context.Context, req *model.DeleteTODORequest) (*model.DeleteTODOResponse, error) {
	_ = h.svc.DeleteTODO(ctx, nil)
	return &model.DeleteTODOResponse{}, nil
}
