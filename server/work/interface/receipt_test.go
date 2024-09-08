package interfaces

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"work/infrastructure"
	"work/usecase"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func Test_Create_Success(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("insert one", func(mt *mtest.T) {
			ctx := context.Background()

			mydb := &infrastructure.MyMongoDB{Client: mt.Client}
			mt.AddMockResponses(mtest.CreateSuccessResponse())

			receiptRepository := infrastructure.NewReceiptRepositoryInfrastructure(mydb, ctx)
			receiptUsecase := usecase.NewReceiptUsecase(receiptRepository)
			receiptController := NewReceiptController(receiptUsecase)

			reqBody := strings.NewReader(`{
				"date": "2029/09/01",
				"category": "meal",
				"content": "A portion of dinner",
				"money": "500"
			}`)

			e := echo.New()
			e.POST("/receipt", receiptController.Create)

			req := httptest.NewRequest(http.MethodPost, "/receipt", reqBody)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			e.ServeHTTP(rec, req)


			if rec.Code != http.StatusCreated {
					t.Errorf("Expected status code 201, but got %d", rec.Code)
			}

			if body := rec.Body.String(); body != "登録完了しました" {
					t.Errorf("Expected response body '登録完了しました', but got '%s'", body)
			}
	})
}


func Test_Create_Failure(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	mt.Run("insert one", func(mt *mtest.T) {
			ctx := context.Background()

			mydb := &infrastructure.MyMongoDB{Client: mt.Client}
			mt.AddMockResponses(mtest.CreateSuccessResponse())

			receiptRepository := infrastructure.NewReceiptRepositoryInfrastructure(mydb, ctx)
			receiptUsecase := usecase.NewReceiptUsecase(receiptRepository)
			receiptController := NewReceiptController(receiptUsecase)

			reqBody := strings.NewReader(`{
				"date": "2029/09/01",
				"category": "meal",
				"content": "A portion of dinner",
				"money": "0"
			}`)

			e := echo.New()
			e.POST("/receipt", receiptController.Create)

			req := httptest.NewRequest(http.MethodPost, "/receipt", reqBody)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			e.ServeHTTP(rec, req)

			if rec.Code != http.StatusInternalServerError {
					t.Errorf("Expected status code 500, but got %d", rec.Code)
			}

			if body := rec.Body.String(); body != "データベースへの登録に失敗しました" {
					t.Errorf("Expected response body 'データベースへの登録に失敗しました', but got '%s'", body)
			}
	})
}
