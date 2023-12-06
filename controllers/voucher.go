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

type PostCreateVoucherRequest struct {
	Code        string  `json:"code" validate:"required"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Percentage  float64 `json:"percentage" validate:"required"`
	Max         float64 `json:"max"  validate:"required"`
}

type PutEditVoucherRequest struct {
	Code        string  `json:"code" validate:"required"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Percentage  float64 `json:"percentage" validate:"required"`
	Max         float64 `json:"max"  validate:"required"`
}

type VoucherController struct {
	db        *gorm.DB
	validator *util.Validator
}

func NewVoucherController(db *gorm.DB, validator *util.Validator) *VoucherController {
	return &VoucherController{
		db:        db,
		validator: validator,
	}
}

// GetAllVoucher	goDocs
// @Summary      get all voucher
// @Description  get all voucher need credentials
// @Tags         Voucher
// @Param				 Authorization	header		string	true	"Bearer {token}" default(Bearer {token})
// @Param				 code	query		string	false	"code voucher"
// @Produce      application/json
// @Router       /voucher/all [get]
func (s *VoucherController) GetVouchers(c *gin.Context) {
	// log
	logCtx := log.WithFields(log.Fields{
		"api":    "GetVouchers",
		"params": c.Request.URL.RawQuery,
	})

	voucherRepo := repository.NewVoucherRepository(s.db)
	voucher, result := voucherRepo.All(c.Request)
	if result.Error != nil {
		logCtx.WithField("reason", result.Error).Error("error find voucher")
		c.AbortWithStatusJSON(http.StatusNotFound, nil)
		return
	}

	meta := voucherRepo.MetaPaginate(c.Request)

	c.JSON(http.StatusOK, gin.H{
		"message": "Sucess!",
		"data":    voucher,
		"meta":    meta,
	})
}

// GetOneVoucherDetail	goDocs
// @Summary      get one voucher detail
// @Description  get one voucher detail, need credentials
// @Tags         Voucher
// @Param				 id path int true "get detail by id"
// @Param				 Authorization	header		string	true	"Bearer {token}" default(Bearer {token})
// @Produce      application/json
// @Router       /voucher/{id} [get]
func (s *VoucherController) GetVoucherDetail(c *gin.Context) {
	// log
	logCtx := log.WithFields(log.Fields{
		"api": "GetVoucher",
	})

	idS := c.Param("id")
	id, err := strconv.Atoi(idS)
	if err != nil {
		logCtx.WithField("reason", err).Error("error parse id")
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid id"})
		return
	}

	voucherRepo := repository.NewVoucherRepository(s.db)
	voucher, result := voucherRepo.OneById(id)
	if result.RowsAffected == 0 || result.Error != nil {
		err := errors.New("error find voucher")
		if result.Error != nil {
			err = result.Error
		}
		logCtx.WithField("reason", err).Error("error find voucher")
		c.AbortWithStatusJSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    voucher,
		"message": "Sucess!",
	})
}

// AddVoucher	goDocs
// @Summary      add voucher for admin user
// @Description  add voucher for admin user, need credential ADMIN user only
// @Tags         Voucher
// @Param				 Authorization	header		string	true	"Bearer {token}" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluQGV4YW1wbGUuY29tIiwiZW1haWwiOiJhZG1pbkBleGFtcGxlLmNvbSIsInR5cGUiOiJBRE1JTiIsImV4cCI6MTcyNTU5ODc3OX0.JamULnKlo3q38ZgIhfBOUI8U2WEv4nNfaLYvodtIx0c)
// @Param        tags body PostCreateVoucherRequest true "Body Request"
// @Produce      application/json
// @Router       /voucher [post]
func (s *VoucherController) PostCreateVoucher(c *gin.Context) {
	// bind data
	var req PostCreateVoucherRequest
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
		"api": "PostCreateVoucher",
	})

	username := c.GetString("username")

	voucherRepo := repository.NewVoucherRepository(s.db)
	voucher, result := voucherRepo.Create(model.Voucher{
		Code:        req.Code,
		Name:        req.Name,
		Description: req.Description,
		Percentage:  req.Percentage,
		Max:         req.Max,
		CreatedBy:   username,
	})
	if result.Error != nil {
		logCtx.WithField("reason", result.Error).Error("error create voucher")
		c.AbortWithStatusJSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    voucher,
		"message": "Sucess!",
	})
}

// EditVoucher	goDocs
// @Summary      edit voucher for admin user
// @Description  edit voucher for admin user, need credential ADMIN user only
// @Tags         Voucher
// @Param				 id path int true "edit by id"
// @Param				 Authorization	header		string	true	"Bearer {token}" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluQGV4YW1wbGUuY29tIiwiZW1haWwiOiJhZG1pbkBleGFtcGxlLmNvbSIsInR5cGUiOiJBRE1JTiIsImV4cCI6MTcyNTU5ODc3OX0.JamULnKlo3q38ZgIhfBOUI8U2WEv4nNfaLYvodtIx0c)
// @Param        tags body PostCreateVoucherRequest true "Body Request"
// @Produce      application/json
// @Router       /voucher/{id} [put]
func (s *VoucherController) PutEditVoucher(c *gin.Context) {
	// bind data
	var req PostCreateVoucherRequest
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
		"api": "PutEditVoucher",
	})

	idS := c.Param("id")
	id, err := strconv.Atoi(idS)
	if err != nil {
		logCtx.WithField("reason", err).Error("error parse id")
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid id"})
		return
	}

	username := c.GetString("username")

	voucherRepo := repository.NewVoucherRepository(s.db)
	voucher, result := voucherRepo.Update(id, model.Voucher{
		Code:        req.Code,
		Name:        req.Name,
		Description: req.Description,
		Percentage:  req.Percentage,
		Max:         req.Max,
		UpdatedBy:   username,
	})
	if result.Error != nil {
		logCtx.WithField("reason", result.Error).Error("error update voucher")
		c.AbortWithStatusJSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Sucess!",
		"data":    voucher,
	})
}

// DeleteVoucher	goDocs
// @Summary      delete voucher for admin user
// @Description  delete voucher for admin user, need credential ADMIN user only
// @Tags         Voucher
// @Param				 id path int true "delete by id"
// @Param				 Authorization	header		string	true	"Bearer {token}" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluQGV4YW1wbGUuY29tIiwiZW1haWwiOiJhZG1pbkBleGFtcGxlLmNvbSIsInR5cGUiOiJBRE1JTiIsImV4cCI6MTcyNTU5ODc3OX0.JamULnKlo3q38ZgIhfBOUI8U2WEv4nNfaLYvodtIx0c)
// @Produce      application/json
// @Router       /voucher/{id} [delete]
func (s *VoucherController) DeleteVoucher(c *gin.Context) {
	// log
	logCtx := log.WithFields(log.Fields{
		"api": "PutEditVoucher",
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
		voucherRepo := repository.NewVoucherRepository(tx)
		_, result := voucherRepo.Update(id, model.Voucher{
			DeletedBy: username,
		})
		if result.Error != nil {
			return result.Error
		}
		if result := voucherRepo.Delete(id, false); result.Error != nil {
			return result.Error
		}
		return nil
	}); err != nil {
		logCtx.WithField("reason", err).Error("error delete voucher")
		c.AbortWithStatusJSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Sucess!",
	})
}
