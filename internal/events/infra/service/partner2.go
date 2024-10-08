package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Partner2ReservationRequest struct {
	Spot       []string `json:"spots"`
	TicketKind string   `json:"ticket_kind"`
	Email      string   `json:"email"`
}

type Partner2ReservationResponse struct {
	ID         string `json:"id"`
	Email      string `json:"email"`
	Spot       string `json:"spot"`
	TicketKind string `json:"ticket_kind"`
	Status     string `json:"status"`
	EventID    string `json:"event_id"`
}

type Partner2 struct {
	BaseURL string
}

func (p *Partner2) MakeReservation(req *ReservationRequest) ([]ReservationResponse, error) {
	partnerRequest := Partner2ReservationRequest{
		Spot:       req.Spot,
		TicketKind: req.TicketType,
		Email:      req.Email,
	}

	body, err := json.Marshal(partnerRequest)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/events/%s/reserves", p.BaseURL, req.EventID)
	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	httpResp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}

	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("unexpected status code: %d", httpResp.StatusCode)
	}

	var partnerResp []Partner2ReservationResponse
	if err = json.NewDecoder(httpResp.Body).Decode(&partnerResp); err != nil {
		return nil, err
	}

	responses := make([]ReservationResponse, len(partnerResp))
	for i, r := range partnerResp {
		responses[i] = ReservationResponse{
			ID:     r.ID,
			Spot:   r.Spot,
			Status: r.Status,
		}
	}

	return responses, nil
}
