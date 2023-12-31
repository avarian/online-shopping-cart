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

type PostCreateItemRequest struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float64 `json:"price"  validate:"required"`
	Qty         *int    `json:"qty"  validate:"required"`
}

type PutEditItemRequest struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float64 `json:"price"  validate:"required"`
	Qty         *int    `json:"qty"  validate:"required"`
}

type ItemController struct {
	db        *gorm.DB
	validator *util.Validator
}

func NewItemController(db *gorm.DB, validator *util.Validator) *ItemController {
	return &ItemController{
		db:        db,
		validator: validator,
	}
}

// GetAllItem	goDocs
// @Summary      get all item
// @Description  get all item need credentials
// @Tags         Item
// @Param				 Authorization	header		string	true	"Bearer {token}" default(Bearer {token})
// @Param				 name	query		string	false	"name item"
// @Param				 min_price	query		string	false	"min price item"
// @Param				 max_price	query		string	false	"max price item"
// @Produce      application/json
// @Router       /item/all [get]
func (s *ItemController) GetItems(c *gin.Context) {
	// log
	logCtx := log.WithFields(log.Fields{
		"api":    "GetItems",
		"params": c.Request.URL.RawQuery,
	})

	itemRepo := repository.NewItemRepository(s.db)
	item, result := itemRepo.Index(c.Request)
	if result.Error != nil {
		logCtx.WithField("reason", result.Error).Error("error find item")
		c.AbortWithStatusJSON(http.StatusNotFound, nil)
		return
	}

	meta := itemRepo.MetaPaginate(c.Request)

	c.JSON(http.StatusOK, gin.H{
		"message": "Sucess!",
		"data":    item,
		"meta":    meta,
	})
}

// GetOneItemDetail	goDocs
// @Summary      get one item detail
// @Description  get one item detail, need credentials
// @Tags         Item
// @Param				 id path int true "get detail by id"
// @Param				 Authorization	header		string	true	"Bearer {token}" default(Bearer {token})
// @Produce      application/json
// @Router       /item/{id} [get]
func (s *ItemController) GetItemDetail(c *gin.Context) {
	// log
	logCtx := log.WithFields(log.Fields{
		"api": "GetItem",
	})

	idS := c.Param("id")
	id, err := strconv.Atoi(idS)
	if err != nil {
		logCtx.WithField("reason", err).Error("error parse id")
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid id"})
		return
	}

	itemRepo := repository.NewItemRepository(s.db)
	item, result := itemRepo.OneById(id)
	if result.RowsAffected == 0 || result.Error != nil {
		err := errors.New("error find item")
		if result.Error != nil {
			err = result.Error
		}
		logCtx.WithField("reason", err).Error("error find item")
		c.AbortWithStatusJSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    item,
		"message": "Sucess!",
	})
}

// AddItem	goDocs
// @Summary      add item for admin user
// @Description  add item for admin user, need credential ADMIN user only
// @Tags         Item
// @Param				 Authorization	header		string	true	"Bearer {token}" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluQGV4YW1wbGUuY29tIiwiZW1haWwiOiJhZG1pbkBleGFtcGxlLmNvbSIsInR5cGUiOiJBRE1JTiIsImV4cCI6MTcyNTU5ODc3OX0.JamULnKlo3q38ZgIhfBOUI8U2WEv4nNfaLYvodtIx0c)
// @Param        tags body PostCreateItemRequest true "Body Request"
// @Produce      application/json
// @Router       /item [post]
func (s *ItemController) PostCreateItem(c *gin.Context) {
	// bind data
	var req PostCreateItemRequest
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
		"api": "PostCreateItem",
	})

	username := c.GetString("username")

	itemRepo := repository.NewItemRepository(s.db)
	item, result := itemRepo.Create(model.Item{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Qty:         req.Qty,
		CreatedBy:   username,
	})
	if result.Error != nil {
		logCtx.WithField("reason", result.Error).Error("error create item")
		c.AbortWithStatusJSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    item,
		"message": "Sucess!",
	})
}

// EditItem	goDocs
// @Summary      edit item for admin user
// @Description  edit item for admin user, need credential ADMIN user only
// @Tags         Item
// @Param				 id path int true "edit by id"
// @Param				 Authorization	header		string	true	"Bearer {token}" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluQGV4YW1wbGUuY29tIiwiZW1haWwiOiJhZG1pbkBleGFtcGxlLmNvbSIsInR5cGUiOiJBRE1JTiIsImV4cCI6MTcyNTU5ODc3OX0.JamULnKlo3q38ZgIhfBOUI8U2WEv4nNfaLYvodtIx0c)
// @Param        tags body PutEditItemRequest true "Body Request"
// @Produce      application/json
// @Router       /item/{id} [put]
func (s *ItemController) PutEditItem(c *gin.Context) {
	// bind data
	var req PutEditItemRequest
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
		"api": "PutEditItem",
	})

	idS := c.Param("id")
	id, err := strconv.Atoi(idS)
	if err != nil {
		logCtx.WithField("reason", err).Error("error parse id")
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid id"})
		return
	}

	username := c.GetString("username")

	itemRepo := repository.NewItemRepository(s.db)
	item, result := itemRepo.Update(id, model.Item{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Qty:         req.Qty,
		UpdatedBy:   username,
	})
	if result.Error != nil {
		logCtx.WithField("reason", result.Error).Error("error update item")
		c.AbortWithStatusJSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Sucess!",
		"data":    item,
	})
}

// DeleteItem	goDocs
// @Summary      delete item for admin user
// @Description  delete item for admin user, need credential ADMIN user only
// @Tags         Item
// @Param				 id path int true "delete by id"
// @Param				 Authorization	header		string	true	"Bearer {token}" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluQGV4YW1wbGUuY29tIiwiZW1haWwiOiJhZG1pbkBleGFtcGxlLmNvbSIsInR5cGUiOiJBRE1JTiIsImV4cCI6MTcyNTU5ODc3OX0.JamULnKlo3q38ZgIhfBOUI8U2WEv4nNfaLYvodtIx0c)
// @Produce      application/json
// @Router       /item/{id} [delete]
func (s *ItemController) DeleteItem(c *gin.Context) {
	// log
	logCtx := log.WithFields(log.Fields{
		"api": "PutEditItem",
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
		itemRepo := repository.NewItemRepository(tx)
		_, result := itemRepo.Update(id, model.Item{
			DeletedBy: &username,
		})
		if result.Error != nil {
			return result.Error
		}
		if result := itemRepo.Delete(id, false); result.Error != nil {
			return result.Error
		}
		return nil
	}); err != nil {
		logCtx.WithField("reason", err).Error("error delete item")
		c.AbortWithStatusJSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Sucess!",
	})
}
