package repository

import (
	"context"

	"gorm.io/gorm"
	domainpayment "io.lees.cloud-between/core/core-domain/payment"
	"io.lees.cloud-between/storage/db-core/entity"
)

// PaymentGormRepository implements PaymentRepository.
type PaymentGormRepository struct {
	db *gorm.DB
}

func NewPaymentCoreRepository(db *gorm.DB) domainpayment.PaymentRepository {
	return &PaymentGormRepository{db: db}
}

func (r *PaymentGormRepository) Save(ctx context.Context, p *domainpayment.Payment) error {
	e := &entity.PaymentEntity{
		BaseEntity: entity.BaseEntity{
			ID:        p.ID,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		},
		UserID:   p.UserID,
		OrderID:  p.OrderID,
		Amount:   p.Amount,
		Currency: p.Currency,
		Status:   string(p.Status),
	}
	return r.db.WithContext(ctx).Create(e).Error
}

func (r *PaymentGormRepository) FindByOrderID(ctx context.Context, orderID string) (*domainpayment.Payment, error) {
	var e entity.PaymentEntity
	err := r.db.WithContext(ctx).Where("order_id = ?", orderID).First(&e).Error
	if err != nil {
		return nil, err
	}
	return toPaymentDomain(e), nil
}

func (r *PaymentGormRepository) UpdateStatus(ctx context.Context, orderID string, status domainpayment.PaymentStatus) error {
	return r.db.WithContext(ctx).Model(&entity.PaymentEntity{}).
		Where("order_id = ?", orderID).
		Updates(map[string]interface{}{
			"status":     string(status),
			"updated_at": gorm.Expr("NOW()"),
		}).Error
}

func toPaymentDomain(e entity.PaymentEntity) *domainpayment.Payment {
	return &domainpayment.Payment{
		ID:        e.ID,
		UserID:    e.UserID,
		OrderID:   e.OrderID,
		Amount:    e.Amount,
		Currency:  e.Currency,
		Status:    domainpayment.PaymentStatus(e.Status),
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

// PaymentCancelGormRepository implements PaymentCancelRepository.
type PaymentCancelGormRepository struct {
	db *gorm.DB
}

func NewPaymentCancelCoreRepository(db *gorm.DB) domainpayment.PaymentCancelRepository {
	return &PaymentCancelGormRepository{db: db}
}

func (r *PaymentCancelGormRepository) Save(ctx context.Context, cancel *domainpayment.PaymentCancel) error {
	e := &entity.PaymentCancelEntity{
		BaseEntity: entity.BaseEntity{
			ID:        cancel.ID,
			CreatedAt: cancel.CreatedAt,
		},
		PaymentID: cancel.PaymentID,
		OrderID:   cancel.OrderID,
		Reason:    cancel.Reason,
	}
	return r.db.WithContext(ctx).Create(e).Error
}
