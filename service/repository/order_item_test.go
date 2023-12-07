package repository

import (
	"net/http"
	"net/url"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/avarian/online-shopping-cart/model"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OrderItemNewMockDB() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Printf("An error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	if err != nil {
		log.Printf("An error '%s' was not expected when opening gorm database", err)
	}

	return gormDB, mock
}

func Test_OrderItemIndex(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}

	type args struct {
		r *http.Request
	}

	tests := []struct {
		name    string
		args    args
		wantErr error
		want    []model.OrderItem
		mockFn  func(a args) fields
	}{
		{
			name: "Success",
			args: args{
				&http.Request{
					URL: &url.URL{RawQuery: ""},
				},
			},
			want: []model.OrderItem{
				{
					ID:          1,
					OrderID:     1,
					ItemID:      1,
					Name:        "name",
					Description: "desc",
					Price:       1,
					Qty:         1,
					Total:       1,
				},
				{
					ID:          2,
					OrderID:     2,
					ItemID:      2,
					Name:        "name2",
					Description: "desc2",
					Price:       2,
					Qty:         2,
					Total:       2,
				},
			},
			mockFn: func(args) fields {
				db, mock := OrderItemNewMockDB()

				row := sqlmock.NewRows([]string{"id", "order_id", "item_id", "name", "description", "price", "qty", "total"}).
					AddRow(1, 1, 1, "name", "desc", 1, 1, 1).
					AddRow(2, 2, 2, "name2", "desc2", 2, 2, 2)
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `order_items` WHERE `order_items`.`deleted_at` IS NULL")).WillReturnRows(row)

				return fields{
					db: db,
				}
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			dep := tt.mockFn(tt.args)

			p := NewOrderItemRepository(dep.db)

			got, result := p.Index(tt.args.r)
			assert.Equal(t, tt.wantErr, result.Error)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_OrderItemAll(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}

	type args struct {
		r *http.Request
	}

	tests := []struct {
		name    string
		args    args
		wantErr error
		want    []model.OrderItem
		mockFn  func(a args) fields
	}{
		{
			name: "Success",
			args: args{
				&http.Request{
					URL: &url.URL{RawQuery: ""},
				},
			},
			want: []model.OrderItem{
				{
					ID:          1,
					OrderID:     1,
					ItemID:      1,
					Name:        "name",
					Description: "desc",
					Price:       1,
					Qty:         1,
					Total:       1,
				},
				{
					ID:          2,
					OrderID:     2,
					ItemID:      2,
					Name:        "name2",
					Description: "desc2",
					Price:       2,
					Qty:         2,
					Total:       2,
				},
			},
			mockFn: func(args) fields {
				db, mock := OrderItemNewMockDB()

				row := sqlmock.NewRows([]string{"id", "order_id", "item_id", "name", "description", "price", "qty", "total"}).
					AddRow(1, 1, 1, "name", "desc", 1, 1, 1).
					AddRow(2, 2, 2, "name2", "desc2", 2, 2, 2)
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `order_items` WHERE `order_items`.`deleted_at` IS NULL")).WillReturnRows(row)

				return fields{
					db: db,
				}
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			dep := tt.mockFn(tt.args)

			p := NewOrderItemRepository(dep.db)

			got, result := p.All(tt.args.r)
			assert.Equal(t, tt.wantErr, result.Error)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_OrderItemOne(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}

	type args struct {
		r *http.Request
	}

	tests := []struct {
		name    string
		args    args
		wantErr error
		want    model.OrderItem
		mockFn  func(a args) fields
	}{
		{
			name: "Success",
			args: args{
				&http.Request{
					URL: &url.URL{RawQuery: ""},
				},
			},
			want: model.OrderItem{
				ID:          1,
				OrderID:     1,
				ItemID:      1,
				Name:        "name",
				Description: "desc",
				Price:       1,
				Qty:         1,
				Total:       1,
			},
			mockFn: func(args) fields {
				db, mock := OrderItemNewMockDB()

				row := sqlmock.NewRows([]string{"id", "order_id", "item_id", "name", "description", "price", "qty", "total"}).
					AddRow(1, 1, 1, "name", "desc", 1, 1, 1).
					AddRow(2, 2, 2, "name2", "desc2", 2, 2, 2)
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `order_items` WHERE `order_items`.`deleted_at` IS NULL")).WillReturnRows(row)

				return fields{
					db: db,
				}
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			dep := tt.mockFn(tt.args)

			p := NewOrderItemRepository(dep.db)

			got, result := p.One(tt.args.r)
			assert.Equal(t, tt.wantErr, result.Error)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_OrderItemOneById(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}

	type args struct {
		id int
	}

	tests := []struct {
		name    string
		args    args
		wantErr error
		want    model.OrderItem
		mockFn  func(a args) fields
	}{
		{
			name: "Success",
			args: args{
				id: 1,
			},
			want: model.OrderItem{
				ID:          1,
				OrderID:     1,
				ItemID:      1,
				Name:        "name",
				Description: "desc",
				Price:       1,
				Qty:         1,
				Total:       1,
			},
			mockFn: func(args) fields {
				db, mock := OrderItemNewMockDB()

				row := sqlmock.NewRows([]string{"id", "order_id", "item_id", "name", "description", "price", "qty", "total"}).
					AddRow(1, 1, 1, "name", "desc", 1, 1, 1).
					AddRow(2, 2, 2, "name2", "desc2", 2, 2, 2)
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `order_items` WHERE id = ? AND `order_items`.`deleted_at` IS NULL")).WillReturnRows(row)

				return fields{
					db: db,
				}
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			dep := tt.mockFn(tt.args)

			p := NewOrderItemRepository(dep.db)

			got, result := p.OneById(tt.args.id)
			assert.Equal(t, tt.wantErr, result.Error)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_OrderItemCreate(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}

	type args struct {
		order_item model.OrderItem
	}

	tests := []struct {
		name    string
		args    args
		wantErr error
		want    model.OrderItem
		mockFn  func(a args) fields
	}{
		{
			name: "Success",
			args: args{
				order_item: model.OrderItem{
					ID:          1,
					OrderID:     1,
					ItemID:      1,
					Name:        "name",
					Description: "desc",
					Price:       1,
					Qty:         1,
					Total:       1,
				},
			},
			want: model.OrderItem{
				ID:          1,
				OrderID:     1,
				ItemID:      1,
				Name:        "name",
				Description: "desc",
				Price:       1,
				Qty:         1,
				Total:       1,
				CreatedBy:   "SYSTEM",
				UpdatedBy:   "SYSTEM",
			},
			mockFn: func(args) fields {
				db, mock := OrderItemNewMockDB()

				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `order_items` (`order_id`,`item_id`,`name`,`description`,`price`,`qty`,`total`,`created_by`,`updated_by`,`deleted_by`,`deleted_at`,`id`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)")).WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()

				return fields{
					db: db,
				}
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			dep := tt.mockFn(tt.args)

			p := NewOrderItemRepository(dep.db)

			got, result := p.Create(tt.args.order_item)
			assert.Equal(t, tt.wantErr, result.Error)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_OrderItemUpdate(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}

	type args struct {
		id         int
		order_item model.OrderItem
	}

	tests := []struct {
		name    string
		args    args
		wantErr error
		want    model.OrderItem
		mockFn  func(a args) fields
	}{
		{
			name: "Success",
			args: args{
				id: 1,
				order_item: model.OrderItem{
					ID:          1,
					OrderID:     1,
					ItemID:      1,
					Name:        "name",
					Description: "desc",
					Price:       1,
					Qty:         1,
					Total:       1,
				},
			},
			want: model.OrderItem{
				ID:          1,
				OrderID:     1,
				ItemID:      1,
				Name:        "name",
				Description: "desc",
				Price:       1,
				Qty:         1,
				Total:       1,
			},
			mockFn: func(args) fields {
				db, mock := OrderItemNewMockDB()

				row := sqlmock.NewRows([]string{"id", "order_id", "item_id", "name", "description", "price", "qty", "total"}).
					AddRow(1, 1, 1, "name", "desc", 1, 1, 1).
					AddRow(2, 2, 2, "name2", "desc2", 2, 2, 2)
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `order_items` WHERE id = ? AND `order_items`.`deleted_at` IS NULL")).WillReturnRows(row)

				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta("UPDATE `order_items` SET `order_id`=?,`item_id`=?,`name`=?,`description`=?,`price`=?,`qty`=?,`total`=?,`created_by`=?,`updated_by`=?,`deleted_by`=?,`created_at`=?,`updated_at`=?,`deleted_at`=? WHERE `order_items`.`deleted_at` IS NULL AND `id` = ?")).WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()

				return fields{
					db: db,
				}
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			dep := tt.mockFn(tt.args)

			p := NewOrderItemRepository(dep.db)

			got, result := p.Update(tt.args.id, tt.args.order_item)
			got.UpdatedAt = nil
			assert.Equal(t, tt.wantErr, result.Error)
			assert.Equal(t, tt.want, got)
		})
	}
}
