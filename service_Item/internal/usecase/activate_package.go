package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"service-package/internal/domain"
	"time"
)

// Interface ini mendefinisikan kontrak untuk MessageUseCase. Di sini, metode ActivatePackage harus diimplementasikan oleh struct yang mengimplementasi interface ini.
type MessageUseCase interface {
	ActivatePackage(ctx context.Context, msg domain.Message) (domain.Response, error)
}

type messageUseCase struct{}

func NewMessageUseCase() MessageUseCase {
	return &messageUseCase{}
}

func (uc *messageUseCase) ActivatePackage(ctx context.Context, msg domain.Message) (domain.Response, error) {
	// Business logic to activate the package
	apiUrl := "https://packageactivate.free.beeceptor.com"
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return domain.Response{}, err
	}
	// Tambahkan query parameters atau headers jika diperlukan
	q := req.URL.Query()
	q.Add("itemId", msg.ItemId)
	req.URL.RawQuery = q.Encode()

	// Atur timeout dan buat HTTP client
	client := &http.Client{Timeout: 5 * time.Second}

	// Panggil API eksternal
	resp, err := client.Do(req)
	if err != nil {
		return domain.Response{}, err
	}
	defer resp.Body.Close()
	// Proses response dari API eksternal
	if resp.StatusCode != http.StatusOK {
		return domain.Response{}, errors.New("failed to validate item")
	}

	var apiResponse struct {
		IsValid bool   `json:"isValid"`
		Status  string `json:"status"`
		Message string `json:"message"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return domain.Response{}, err
	}

	if !apiResponse.IsValid {
		return domain.Response{
			OrderType:     msg.OrderType,
			OrderService:  "validateItem",
			TransactionId: msg.TransactionId,
			UserId:        msg.UserId,
			ItemId:        msg.ItemId,
			RespCode:      400,
			RespStatus:    "Failed",
			RespMessage:   "Item is Empty not found",
		}, nil
	}

	return domain.Response{
		OrderType:     msg.OrderType,
		OrderService:  "validateItem",
		TransactionId: msg.TransactionId,
		UserId:        msg.UserId,
		ItemId:        msg.ItemId,
		RespCode:      200,
		RespStatus:    "Success",
		RespMessage:   "Item is Available",
	}, nil

}