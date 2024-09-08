package domain

import (
	"fmt"
)

type Receipt struct {
	Date string
	Category string 
	Content string 
	Money int 
	Remarks string 
}

func (r Receipt) Validate() error {
	if r.Money <= 0 {
		return fmt.Errorf("金額を正しく入力してください")
	}

	if r.Date == "" || r.Category == "" || r.Content == "" {
		return fmt.Errorf("必要なデータを入力してください")

	}
	return nil
}

type ReceiptRepository interface{
	Create(receipt *Receipt) error
}