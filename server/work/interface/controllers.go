package interfaces

import (
	"github.com/labstack/echo/v4"
)

type Controllers struct {
	receiptController *ReceiptController
}

func NewControllers(rc *ReceiptController) *Controllers {
	return &Controllers{ receiptController: rc }
}

func (c *Controllers) Mount(e *echo.Echo) {
	c.receiptController.Mount(e.Group("/receipt"))
}
