package usecase

import (
	"strconv"

	"work/domain"
)

type ReceiptUsecase interface{
	CreateReceipt(input *CreateReceiptInput) error
}

type CreateReceiptInput struct{
	Date string `json:"date" form:"date"` // required
	Category string `json:"category" form:"category"`  // required
	Content string `json:"content" form:"content"`	// required
	Money string `json:"money" form:"money"` // required
	Remarks string `json:"remarks" form:"remarks"` // option
}


type receiptUsecase struct {
	receiptRepository domain.ReceiptRepository
}

func NewReceiptUsecase(rp domain.ReceiptRepository) ReceiptUsecase{
	return &receiptUsecase{ receiptRepository: rp }
}

func (r *receiptUsecase) CreateReceipt(input *CreateReceiptInput) error {
	money, err := strconv.Atoi(input.Money)
	if err != nil{
		return err
	}

	receipt := &domain.Receipt{
		Date: input.Date,
		Category: input.Category,
		Content: input.Content,
		Money: money,
		Remarks: input.Remarks,
	}

	if err := receipt.Validate(); err != nil {
		return err
	}

	if err := r.receiptRepository.Create(receipt); err != nil {
		return err
	}

	return nil
}
