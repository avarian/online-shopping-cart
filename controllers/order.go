package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/avarian/online-shopping-cart/model"
	"github.com/avarian/online-shopping-cart/service/repository"
	"github.com/avarian/online-shopping-cart/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PostCreateOrderRequest struct {
	VoucherCode string `json:"voucher_code"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}

type OrderController struct {
	db        *gorm.DB
	validator *util.Validator
}

func NewOrderController(db *gorm.DB, validator *util.Validator) *OrderController {
	return &OrderController{
		db:        db,
		validator: validator,
	}
}

// GetAllOrder	goDocs
// @Summary      get all own order
// @Description  get all own order need credentials
// @Tags         Order
// @Param				 Authorization	header		string	true	"Bearer {token}" default(Bearer {token})
// @Produce      application/json
// @Router       /order/all [get]
func (s *OrderController) GetOrders(c *gin.Context) {
	// log
	logCtx := log.WithFields(log.Fields{
		"api": "GetOrders",
	})

	username := c.GetString("username")
	accountRepo := repository.NewAccountRepository(s.db)
	account, result := accountRepo.OneByEmail(username)
	if result.RowsAffected == 0 || result.Error != nil {
		err := errors.New("error find account")
		if result.Error != nil {
			err = result.Error
		}
		logCtx.WithField("reason", err).Error("error find account")
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "account not found"})
		return
	}

	orderRepo := repository.NewOrderRepository(s.db)
	order, result := orderRepo.AllByAccountId(int(account.ID), "OrderItem", "OrderVoucher")
	if result.Error != nil {
		logCtx.WithField("reason", result.Error).Error("error find order")
		c.AbortWithStatusJSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Sucess!",
		"data":    order,
	})
}

// GetOneOrderDetail	goDocs
// @Summary      get one own order detail
// @Description  get one own order detail, need credentials
// @Tags         Order
// @Param				 id path int true "get detail by id"
// @Param				 Authorization	header		string	true	"Bearer {token}" default(Bearer {token})
// @Produce      application/json
// @Router       /order/{id} [get]
func (s *OrderController) GetOrderDetail(c *gin.Context) {
	// log
	logCtx := log.WithFields(log.Fields{
		"api": "GetOrder",
	})

	username := c.GetString("username")
	accountRepo := repository.NewAccountRepository(s.db)
	account, result := accountRepo.OneByEmail(username)
	if result.RowsAffected == 0 || result.Error != nil {
		err := errors.New("error find account")
		if result.Error != nil {
			err = result.Error
		}
		logCtx.WithField("reason", err).Error("error find account")
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "account not found"})
		return
	}

	idS := c.Param("id")
	id, err := strconv.Atoi(idS)
	if err != nil {
		logCtx.WithField("reason", err).Error("error parse id")
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid id"})
		return
	}

	orderRepo := repository.NewOrderRepository(s.db)
	order, result := orderRepo.OneByIdAndAccountId(id, int(account.ID), "OrderItem", "OrderVoucher")
	if result.RowsAffected == 0 || result.Error != nil {
		err := errors.New("error find order")
		if result.Error != nil {
			err = result.Error
		}
		logCtx.WithField("reason", err).Error("error find order")
		c.AbortWithStatusJSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    order,
		"message": "Sucess!",
	})
}

// CheckoutOrder	goDocs
// @Summary      create order from cart
// @Description  create order from cart, need credential
// @Tags         Order
// @Param				 Authorization	header		string	true	"Bearer {token}" default(Bearer {token})
// @Param        tags body PostCreateOrderRequest true "Body Request"
// @Produce      application/json
// @Router       /order [post]
func (s *OrderController) PostCreateOrder(c *gin.Context) {
	// bind data
	var req PostCreateOrderRequest
	if err := c.ShouldBind(&req); err != nil {
		log.WithField("reason", err).Error("error Binding")
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	// validate
	if err := s.validator.Validate.Struct(&req); err != nil {
		log.WithField("reason", err).Error("invalid Request")
		errs := err.(validator.ValidationErrors)
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": errs.Translate(s.validator.Trans)})
		return
	}

	// log
	logCtx := log.WithFields(log.Fields{
		"api": "PostCreateOrder",
	})

	username := c.GetString("username")
	accountRepo := repository.NewAccountRepository(s.db)
	account, result := accountRepo.OneByEmail(username)
	if result.RowsAffected == 0 || result.Error != nil {
		err := errors.New("error find account")
		if result.Error != nil {
			err = result.Error
		}
		logCtx.WithField("reason", err).Error("error find account")
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "account not found"})
		return
	}

	if req.Address == "" {
		req.Address = account.Address
	}

	if req.PhoneNumber == "" {
		req.PhoneNumber = account.PhoneNumber
	}

	if req.VoucherCode != "" {
		voucherRepo := repository.NewVoucherRepository(s.db)
		if _, result := voucherRepo.OneByCode(req.VoucherCode); result.Error != nil || result.RowsAffected == 0 {
			err := errors.New("voucher not found")
			if result.Error != nil {
				err = result.Error
			}
			logCtx.WithField("reason", err).Error("error find account")
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "voucher not found"})
			return
		}
	}

	if err := s.db.Transaction(func(tx *gorm.DB) error {
		var total float64

		cartRepo := repository.NewCartRepository(tx)
		cart, result := cartRepo.AllByAccountId(int(account.ID), "Item")
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return errors.New("cart empty")
		}

		orderRepo := repository.NewOrderRepository(tx)
		order, result := orderRepo.Create(model.Order{
			AccountID:   account.ID,
			Address:     req.Address,
			PhoneNumber: req.PhoneNumber,
			CreatedBy:   username,
		})
		if result.Error != nil {
			return result.Error
		}

		orderItemRepo := repository.NewOrderItemRepository(tx)
		itemRepo := repository.NewItemRepository(tx)
		for _, v := range cart {
			remaining := *v.Item.Qty - v.Qty
			if remaining < 0 {
				return errors.New("qty > item qty")
			}

			orderItem, result := orderItemRepo.Create(model.OrderItem{
				OrderID:     order.ID,
				ItemID:      v.Item.ID,
				Name:        v.Item.Name,
				Description: v.Item.Description,
				Price:       v.Item.Price,
				Qty:         v.Qty,
				Total:       v.Item.Price * float64(v.Qty),
				CreatedBy:   username,
			})
			if result.Error != nil {
				return result.Error
			}

			total += orderItem.Total

			if result := cartRepo.Delete(int(v.ID), true); result.Error != nil {
				return result.Error
			}

			if _, result := itemRepo.Update(int(v.Item.ID), model.Item{Qty: &remaining}); result.Error != nil {
				return result.Error
			}
		}

		var appliedVoucher float64
		voucherRepo := repository.NewVoucherRepository(tx)
		orderVoucherRepo := repository.NewOrderVoucherRepository(tx)
		if req.VoucherCode != "" {
			voucher, result := voucherRepo.OneByCode(req.VoucherCode)
			if result.Error != nil {
				return result.Error
			}
			if result.RowsAffected == 0 {
				return errors.New("voucher not found")
			}

			appliedVoucher = total * voucher.Percentage / 100
			if appliedVoucher > voucher.Max {
				appliedVoucher = voucher.Max
			}

			if _, result := orderVoucherRepo.Create(model.OrderVoucher{
				OrderID:     order.ID,
				VoucherID:   voucher.ID,
				Code:        voucher.Code,
				Name:        voucher.Name,
				Description: voucher.Description,
				Percentage:  voucher.Percentage,
				Max:         voucher.Max,
				Total:       total,
				Applied:     appliedVoucher,
				CreatedBy:   username,
			}); result.Error != nil {
				return result.Error
			}
		}

		totalPrice := total - appliedVoucher
		if totalPrice < 0 {
			totalPrice = 0
		}
		order, result = orderRepo.Update(int(order.ID), model.Order{Total: totalPrice, UpdatedBy: username})
		if result.Error != nil {
			return result.Error
		}

		return nil
	}); err != nil {
		logCtx.WithField("reason", err).Error("error create order")
		c.AbortWithStatusJSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Sucess!",
	})
}
