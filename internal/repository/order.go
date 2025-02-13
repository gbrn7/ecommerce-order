package repository

import (
	"context"
	"ecommerce-order/internal/models"

	"gorm.io/gorm"
)

type OrderRepo struct {
	DB *gorm.DB
}

func (r *OrderRepo) InsertNewOrder(tx context.Context, order *models.Order) error {
	return r.DB.Transaction(
		func(tx *gorm.DB) error {
			err := tx.Create(order).Error
			if err != nil {
				return err
			}

			return nil
		},
	)

}

func (r *OrderRepo) UpdateStatusOrder(tx context.Context, orderID int, status string) error {
	return r.DB.Exec("UPDATE orders SET status = ? WHERE id = ?", status, orderID).Error
}

func (r *OrderRepo) GetOrderDetail(tx context.Context, orderID int) (models.Order, error) {
	var (
		resp models.Order
		err  error
	)

	err = r.DB.Preload("OrderItem").Where("id = ?", orderID).First(&resp).Error
	return resp, err
}

func (r *OrderRepo) GetOrder(tx context.Context) ([]models.Order, error) {
	var (
		resp []models.Order
		err  error
	)

	err = r.DB.Preload("OrderItem").Order("id DESC").Find(&resp).Error
	return resp, err
}
