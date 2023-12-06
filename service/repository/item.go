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

type ItemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) *ItemRepository {
	return &ItemRepository{
		db: db,
	}
}

func (s *ItemRepository) FilterScope(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db
	}
}

func (s *ItemRepository) PaginateScope(r *http.Request) func(db *gorm.DB) *gorm.DB {
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

func (s *ItemRepository) MetaPaginate(r *http.Request) map[string]interface{} {
	q := r.URL.Query()
	var totalRows int64
	s.db.Model(model.Item{}).Scopes(s.FilterScope(r)).Count(&totalRows)

	pageSize, _ := strconv.Atoi(q.Get("page_size"))
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	totalPages := int(math.Ceil(float64(totalRows) / float64(pageSize)))
	meta := map[string]interface{}{
		"page_size":   pageSize,
		"total_rows":  totalRows,
		"total_pages": totalPages,
	}
	return meta
}

func (s *ItemRepository) Index(r *http.Request, preload ...string) ([]model.Item, *gorm.DB) {
	var table []model.Item
	tx := s.db.Scopes(s.FilterScope(r), s.PaginateScope(r))
	for _, v := range preload {
		tx.Preload(v)
	}
	query := tx.Find(&table)

	return table, query
}

func (s *ItemRepository) All(preload ...string) ([]model.Item, *gorm.DB) {
	var table []model.Item
	tx := s.db
	for _, v := range preload {
		tx.Preload(v)
	}
	query := tx.Find(&table)

	return table, query
}

func (s *ItemRepository) One(r *http.Request, preload ...string) (model.Item, *gorm.DB) {
	var table model.Item
	tx := s.db.Scopes(s.FilterScope(r))
	for _, v := range preload {
		tx.Preload(v)
	}
	query := tx.Find(&table)

	return table, query
}

func (s *ItemRepository) OneById(id int, preload ...string) (model.Item, *gorm.DB) {
	var table model.Item
	tx := s.db.Where("id = ?", id)
	for _, v := range preload {
		tx.Preload(v)
	}
	query := tx.Find(&table)

	return table, query
}

func (s *ItemRepository) Create(data model.Item) (model.Item, *gorm.DB) {
	var table model.Item
	s.AssignData(&table, data)
	query := s.db.Create(&table)
	return table, query
}

func (s *ItemRepository) Update(id int, data model.Item) (model.Item, *gorm.DB) {
	var table model.Item
	table, result := s.OneById(id)
	if result.RowsAffected == 0 {
		result.Error = errors.New(fmt.Sprintf("data not found with id = %d", id))
		return table, result
	}
	s.AssignData(&table, data)
	query := s.db.Save(&table)
	return table, query
}

func (s *ItemRepository) Delete(id int, isHard bool) *gorm.DB {
	tx := s.db
	if isHard {
		tx.Unscoped()
	}
	query := tx.Delete(&model.Item{}, id)
	return query
}

func (s *ItemRepository) AssignData(table *model.Item, data model.Item) {
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
