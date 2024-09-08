package interfaces

import (
	"net/http"

	"work/usecase"

	"github.com/labstack/echo/v4"
)

type ReceiptController struct {
	receiptUsecase usecase.ReceiptUsecase
}

func NewReceiptController(ru usecase.ReceiptUsecase) *ReceiptController{
	return &ReceiptController{receiptUsecase: ru}
}

func (c *ReceiptController) Mount(g *echo.Group) {
	g.POST("", c.Create)
}


// IssueBill
//
// @Summery		Issuance of bill
// @Param 		body body Bill true "request body"
// @Success		201 {object} Bill
// @Failure		400 {string} string
// @Router		/ [post]
func (c *ReceiptController) Create(e echo.Context) error{
	receipt := new(usecase.CreateReceiptInput)

	if err := e.Bind(receipt); err != nil{
		return e.String(http.StatusBadRequest, "入力値が不正確です")
	}

	err := c.receiptUsecase.CreateReceipt(receipt)
	
	if err != nil{
		return e.String(http.StatusInternalServerError, "データベースへの登録に失敗しました")
	}

	return e.String(http.StatusCreated, "登録完了しました")

}
