package usecase

import (
	"context"
	"testing"
	"work/infrastructure"

	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func Test_CreateReceipt_Success(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("insert one", func(mt *mtest.T) {
		cri := &CreateReceiptInput{
			Date: "2029/09/01",
			Category: "meal",
			Content: "A portion of dinner",
			Money: "500",
		}
		ctx := context.Background()

		// モックのmongo.Clientを直接渡す
		mydb := &infrastructure.MyMongoDB{Client: mt.Client}
		mt.AddMockResponses(mtest.CreateSuccessResponse())
		// addMockResponsesは位置関係が大切

		receiptRepository := infrastructure.NewReceiptRepositoryInfrastructure(mydb, ctx)
		receiptUsecase := NewReceiptUsecase(receiptRepository)
		if err := receiptUsecase.CreateReceipt(cri); err != nil{
			t.Fatal("The create process failed.")
		}

	})
}

func Test_CreateReceipt_Failure(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	mt.Run("insert one", func(mt *mtest.T) {
		cri := &CreateReceiptInput{
			Date: "2029/09/01",
			Category: "meal",
			Content: "A portion of dinner",
		}
		ctx := context.Background()

		// モックのmongo.Clientを直接渡す
		mydb := &infrastructure.MyMongoDB{Client: mt.Client}
		mt.AddMockResponses(mtest.CreateSuccessResponse())
		// addMockResponsesは位置関係が大切

		receiptRepository := infrastructure.NewReceiptRepositoryInfrastructure(mydb, ctx)
		receiptUsecase := NewReceiptUsecase(receiptRepository)
		if err := receiptUsecase.CreateReceipt(cri); err == nil{
			t.Fatal("The create process failed.")
		}

	})
}
