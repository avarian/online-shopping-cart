package repository

import (
	"errors"
	"fmt"
	"math"
	"net/http"
	"reflect"
	"strconv"

	"github.com/avarian/online-shopping-cart/model"
	"gorm.io/gorm"
)

type VoucherRepository struct {
	db *gorm.DB
}

func NewVoucherRepository(db *gorm.DB) *VoucherRepository {
	return &VoucherRepository{
		db: db,
	}
}

func (s *VoucherRepository) FilterScope(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		q := r.URL.Query()
		code := q.Get("code")
		if code != "" {
			if code != "" {
				db = db.Where("LOWER(code) = LOWER(?)", code)
			}
		}
		return db
	}
}

func (s *VoucherRepository) PaginateScope(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		q := r.URL.Query()
		page, _ := strconv.Atoi(q.Get("page"))
		if page == 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(q.Get("page_size"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		sortBy := q.Get("sort_by")
		if sortBy == "" {
			sortBy = "id"
		}

		direction := q.Get("direction")
		if direction == "" {
			direction = "desc"
		}

		sort := sortBy + " " + direction

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize).Order(sort)
	}
}

func (s *VoucherRepository) MetaPaginate(r *http.Request) map[string]interface{} {
	q := r.URL.Query()
	var totalRows int64
	s.db.Model(model.Voucher{}).Scopes(s.FilterScope(r)).Count(&totalRows)

	pageSize, _ := strconv.Atoi(q.Get("page_size"))
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	totalPages := int(math.Ceil(float64(totalRows) / float64(pageSize)))
	page, _ := strconv.Atoi(q.Get("page"))
	if page == 0 {
		page = 1
	}
	meta := map[string]interface{}{
		"page":        page,
		"page_size":   pageSize,
		"total_rows":  totalRows,
		"total_pages": totalPages,
	}
	return meta
}

func (s *VoucherRepository) Index(r *http.Request, preload ...string) ([]model.Voucher, *gorm.DB) {
	var table []model.Voucher
	tx := s.db.Scopes(s.FilterScope(r), s.PaginateScope(r))
	for _, v := range preload {
		tx = tx.Preload(v)
	}
	query := tx.Find(&table)

	return table, query
}

func (s *VoucherRepository) All(r *http.Request, preload ...string) ([]model.Voucher, *gorm.DB) {
	var table []model.Voucher
	tx := s.db.Scopes(s.FilterScope(r))
	for _, v := range preload {
		tx = tx.Preload(v)
	}
	query := tx.Find(&table)

	return table, query
}

func (s *VoucherRepository) One(r *http.Request, preload ...string) (model.Voucher, *gorm.DB) {
	var table model.Voucher
	tx := s.db.Scopes(s.FilterScope(r))
	for _, v := range preload {
		tx = tx.Preload(v)
	}
	query := tx.Find(&table)

	return table, query
}

func (s *VoucherRepository) OneById(id int, preload ...string) (model.Voucher, *gorm.DB) {
	var table model.Voucher
	tx := s.db.Where("id = ?", id)
	for _, v := range preload {
		tx = tx.Preload(v)
	}
	query := tx.Find(&table)

	return table, query
}

func (s *VoucherRepository) OneByCode(code string, preload ...string) (model.Voucher, *gorm.DB) {
	var table model.Voucher
	tx := s.db.Where("code = ?", code)
	for _, v := range preload {
		tx = tx.Preload(v)
	}
	query := tx.Find(&table)

	return table, query
}

func (s *VoucherRepository) Create(data model.Voucher) (model.Voucher, *gorm.DB) {
	var table model.Voucher
	s.AssignData(&table, data)
	query := s.db.Create(&table)
	return table, query
}

func (s *VoucherRepository) Update(id int, data model.Voucher) (model.Voucher, *gorm.DB) {
	var table model.Voucher
	table, result := s.OneById(id)
	if result.RowsAffected == 0 {
		result.Error = errors.New(fmt.Sprintf("data not found with id = %d", id))
		return table, result
	}
	s.AssignData(&table, data)
	query := s.db.Save(&table)
	return table, query
}

func (s *VoucherRepository) Delete(id int, isHard bool) *gorm.DB {
	tx := s.db
	if isHard {
		tx = tx.Unscoped()
	}
	query := tx.Delete(&model.Voucher{}, id)
	return query
}

func (s *VoucherRepository) AssignData(table *model.Voucher, data model.Voucher) {
	dataRV := reflect.ValueOf(data)
	tableRV := reflect.ValueOf(table)
	tableRVE := tableRV.Elem()

	for i := 0; i < dataRV.NumField(); i++ {
		if !dataRV.Field(i).IsZero() && (tableRVE.Field(i) != dataRV.Field(i)) {
			fv := tableRVE.FieldByName(dataRV.Type().Field(i).Name)
			fv.Set(dataRV.Field(i))
		}
	}
}
