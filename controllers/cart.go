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

type PostCreateCartFromItemRequest struct {
	ItemId int `json:"item_id"  validate:"required"`
	Qty    int `json:"qty"  validate:"required,gt=0"`
}

type PutEditCartRequest struct {
	Qty *int `json:"qty"  validate:"required"`
}

type CartController struct {
	db        *gorm.DB
	validator *util.Validator
}

func NewCartController(db *gorm.DB, validator *util.Validator) *CartController {
	return &CartController{
		db:        db,
		validator: validator,
	}
}

// GetAllCarts	goDocs
// @Summary      get all own carts
// @Description  get all own carts. need credentials.
// @Tags         Cart
// @Param				 Authorization	header		string	true	"Bearer {token}" default(Bearer {token})
// @Produce      application/json
// @Router       /cart/all [get]
func (s *CartController) GetCarts(c *gin.Context) {
	// log
	logCtx := log.WithFields(log.Fields{
		"api": "GetCarts",
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

	cartRepo := repository.NewCartRepository(s.db)
	cart, result := cartRepo.AllByAccountId(int(account.ID), "Item")
	if result.Error != nil {
		logCtx.WithField("reason", result.Error).Error("error find cart")
		c.AbortWithStatusJSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Sucess!",
		"data":    cart,
	})
}

// GetOneCartDetail	goDocs
// @Summary      get one own carts detail
// @Description  get one own carts detail, need credentials
// @Tags         Cart
// @Param				 id path int true "get detail by id"
// @Param				 Authorization	header		string	true	"Bearer {token}" default(Bearer {token})
// @Produce      application/json
// @Router       /cart/{id} [get]
func (s *CartController) GetCartDetail(c *gin.Context) {
	// log
	logCtx := log.WithFields(log.Fields{
		"api": "GetCart",
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

	cartRepo := repository.NewCartRepository(s.db)
	cart, result := cartRepo.OneByIdAndAccountId(id, int(account.ID), "Item")
	if result.RowsAffected == 0 || result.Error != nil {
		err := errors.New("error find cart")
		if result.Error != nil {
			err = result.Error
		}
		logCtx.WithField("reason", err).Error("error find cart")
		c.AbortWithStatusJSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    cart,
		"message": "Sucess!",
	})
}

// AddToCart	goDocs
// @Summary      add to own cart from item
// @Description  add to own cart from item, need credential
// @Tags         Cart
// @Param				 Authorization	header		string	true	"Bearer {token}" default(Bearer {token})
// @Param        tags body PostCreateCartFromItemRequest true "Body Request"
// @Produce      application/json
// @Router       /cart [post]
func (s *CartController) PostCreateCartFromItem(c *gin.Context) {
	// bind data
	var req PostCreateCartFromItemRequest
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
		"api": "PostCreateCart",
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

	itemRepo := repository.NewItemRepository(s.db)
	item, result := itemRepo.OneById(req.ItemId)
	if result.RowsAffected == 0 || result.Error != nil {
		err := errors.New("error find item")
		if result.Error != nil {
			err = result.Error
		}
		logCtx.WithField("reason", err).Error("error find item")
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "item not found"})
		return
	}

	if req.Qty > *item.Qty {
		logCtx.WithField("reason", "request qty > item qty").Error("error add qty")
		c.AbortWithStatusJSON(http.StatusPreconditionFailed, gin.H{"error": "qty > item qty"})
		return
	}

	cartRepo := repository.NewCartRepository(s.db)
	cart, result := cartRepo.Create(model.Cart{
		AccountID: account.ID,
		ItemID:    item.ID,
		Qty:       req.Qty,
		CreatedBy: username,
	})
	if result.Error != nil {
		logCtx.WithField("reason", result.Error).Error("error create cart")
		c.AbortWithStatusJSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    cart,
		"message": "Sucess!",
	})
}

// EditQtyCart	goDocs
// @Summary      update qty of item in cart
// @Description  update qty of item in cart, delete it if qty 0, failed when qty item < qty, need credential
// @Tags         Cart
// @Param				 id path int true "edit by id"
// @Param				 Authorization	header		string	true	"Bearer {token}" default(Bearer {token})
// @Param        tags body PutEditCartRequest true "Body Request"
// @Produce      application/json
// @Router       /cart/{id} [put]
func (s *CartController) PutEditCart(c *gin.Context) {
	// bind data
	var req PutEditCartRequest
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
		"api": "PutEditCart",
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

	cartRepo := repository.NewCartRepository(s.db)
	cart, result := cartRepo.OneByIdAndAccountId(id, int(account.ID), "Item")
	if result.RowsAffected == 0 || result.Error != nil {
		err := errors.New("error find cart")
		if result.Error != nil {
			err = result.Error
		}
		logCtx.WithField("reason", err).Error("error find cart")
		c.AbortWithStatusJSON(http.StatusNotFound, nil)
		return
	}

	if *req.Qty > *cart.Item.Qty {
		logCtx.WithField("reason", "request qty > item qty").Error("error add qty")
		c.AbortWithStatusJSON(http.StatusPreconditionFailed, gin.H{"error": "qty > item qty"})
		return
	}

	if *req.Qty <= 0 {
		cartRepo.Delete(int(cart.ID), true)
		c.JSON(http.StatusOK, gin.H{"message": "Sucess!"})
		return
	}

	cart, result = cartRepo.Update(id, model.Cart{
		Qty:       *req.Qty,
		UpdatedBy: username,
	})
	if result.Error != nil {
		logCtx.WithField("reason", result.Error).Error("error update cart")
		c.AbortWithStatusJSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Sucess!",
		"data":    cart,
	})
}

// deleteCart	goDocs
// @Summary      delete item by own cart id
// @Description  delete item by own cart cart, need credential
// @Tags         Cart
// @Param				 id path int true "delete by id"
// @Param				 Authorization	header		string	true	"Bearer {token}" default(Bearer {token})
// @Produce      application/json
// @Router       /cart/{id} [delete]
func (s *CartController) DeleteCart(c *gin.Context) {
	// log
	logCtx := log.WithFields(log.Fields{
		"api": "PutEditCart",
	})

	idS := c.Param("id")
	id, err := strconv.Atoi(idS)
	if err != nil {
		logCtx.WithField("reason", err).Error("error parse id")
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid id"})
		return
	}

	username := c.GetString("username")

	if err := s.db.Transaction(func(tx *gorm.DB) error {
		cartRepo := repository.NewCartRepository(tx)
		_, result := cartRepo.Update(id, model.Cart{
			DeletedBy: &username,
		})
		if result.Error != nil {
			return result.Error
		}
		if result := cartRepo.Delete(id, true); result.Error != nil {
			return result.Error
		}
		return nil
	}); err != nil {
		logCtx.WithField("reason", err).Error("error delete cart")
		c.AbortWithStatusJSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Sucess!",
	})
}
