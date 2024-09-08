package infrastructure

import (
	"context"
	"work/domain"
)


type ReceiptRepositoryInfrastructure struct {
	db *MyMongoDB
	ctx context.Context
}

func NewReceiptRepositoryInfrastructure(db *MyMongoDB, ctx context.Context) domain.ReceiptRepository {
	return &ReceiptRepositoryInfrastructure{db, ctx}
}

func (r *ReceiptRepositoryInfrastructure) Create(receipt *domain.Receipt) error {
	_, err := r.db.Client.Database("mongo").Collection("receipt").InsertOne(r.ctx, receipt)
	if err != nil{
		return err
	}
	return nil
}