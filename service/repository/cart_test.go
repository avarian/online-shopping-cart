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

func CartNewMockDB() (*gorm.DB, sqlmock.Sqlmock) {
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

func Test_CartIndex(t *testing.T) {
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
		want    []model.Cart
		mockFn  func(a args) fields
	}{
		{
			name: "Success",
			args: args{
				&http.Request{
					URL: &url.URL{RawQuery: ""},
				},
			},
			want: []model.Cart{
				{
					ID:        1,
					AccountID: 1,
					ItemID:    1,
					Qty:       1,
				},
				{
					ID:        2,
					AccountID: 2,
					ItemID:    2,
					Qty:       2,
				},
			},
			mockFn: func(args) fields {
				db, mock := CartNewMockDB()

				row := sqlmock.NewRows([]string{"id", "account_id", "item_id", "qty"}).
					AddRow(1, 1, 1, 1).
					AddRow(2, 2, 2, 2)
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `carts` WHERE `carts`.`deleted_at` IS NULL")).WillReturnRows(row)

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

			p := NewCartRepository(dep.db)

			got, result := p.Index(tt.args.r)
			assert.Equal(t, tt.wantErr, result.Error)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_CartAll(t *testing.T) {
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
		want    []model.Cart
		mockFn  func(a args) fields
	}{
		{
			name: "Success",
			args: args{},
			want: []model.Cart{
				{
					ID:        1,
					AccountID: 1,
					ItemID:    1,
					Qty:       1,
				},
				{
					ID:        2,
					AccountID: 2,
					ItemID:    2,
					Qty:       2,
				},
			},
			mockFn: func(args) fields {
				db, mock := CartNewMockDB()

				row := sqlmock.NewRows([]string{"id", "account_id", "item_id", "qty"}).
					AddRow(1, 1, 1, 1).
					AddRow(2, 2, 2, 2)
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `carts` WHERE `carts`.`deleted_at` IS NULL")).WillReturnRows(row)

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

			p := NewCartRepository(dep.db)

			got, result := p.All(tt.args.r)
			assert.Equal(t, tt.wantErr, result.Error)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_CartOne(t *testing.T) {
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
		want    model.Cart
		mockFn  func(a args) fields
	}{
		{
			name: "Success",
			args: args{},
			want: model.Cart{
				ID:        1,
				AccountID: 1,
				ItemID:    1,
				Qty:       1,
			},
			mockFn: func(args) fields {
				db, mock := CartNewMockDB()

				row := sqlmock.NewRows([]string{"id", "account_id", "item_id", "qty"}).
					AddRow(1, 1, 1, 1)
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `carts` WHERE `carts`.`deleted_at` IS NULL")).WillReturnRows(row)

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

			p := NewCartRepository(dep.db)

			got, result := p.One(tt.args.r)
			assert.Equal(t, tt.wantErr, result.Error)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_CartOneById(t *testing.T) {
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
		want    model.Cart
		mockFn  func(a args) fields
	}{
		{
			name: "Success",
			args: args{
				id: 1,
			},
			want: model.Cart{
				ID:        1,
				AccountID: 1,
				ItemID:    1,
				Qty:       1,
			},
			mockFn: func(args) fields {
				db, mock := CartNewMockDB()

				row := sqlmock.NewRows([]string{"id", "account_id", "item_id", "qty"}).
					AddRow(1, 1, 1, 1)
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `carts` WHERE id = ? AND `carts`.`deleted_at` IS NULL")).WillReturnRows(row)

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

			p := NewCartRepository(dep.db)

			got, result := p.OneById(tt.args.id)
			assert.Equal(t, tt.wantErr, result.Error)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_CartCreate(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}

	type args struct {
		cart model.Cart
	}

	tests := []struct {
		name    string
		args    args
		wantErr error
		want    model.Cart
		mockFn  func(a args) fields
	}{
		{
			name: "Success",
			args: args{
				cart: model.Cart{
					ID:        1,
					AccountID: 1,
					ItemID:    1,
					Qty:       1,
				},
			},
			want: model.Cart{
				ID:        1,
				AccountID: 1,
				ItemID:    1,
				Qty:       1,
				CreatedBy: "SYSTEM",
				UpdatedBy: "SYSTEM",
			},
			mockFn: func(args) fields {
				db, mock := CartNewMockDB()

				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `carts` (`account_id`,`item_id`,`qty`,`created_by`,`updated_by`,`deleted_by`,`deleted_at`,`id`) VALUES (?,?,?,?,?,?,?,?)")).WillReturnResult(sqlmock.NewResult(1, 1))
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

			p := NewCartRepository(dep.db)

			got, result := p.Create(tt.args.cart)
			assert.Equal(t, tt.wantErr, result.Error)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_CartUpdate(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}

	type args struct {
		id   int
		cart model.Cart
	}

	tests := []struct {
		name    string
		args    args
		wantErr error
		want    model.Cart
		mockFn  func(a args) fields
	}{
		{
			name: "Success",
			args: args{
				id: 1,
				cart: model.Cart{
					ID:        1,
					AccountID: 1,
					ItemID:    1,
					Qty:       1,
				},
			},
			want: model.Cart{
				ID:        1,
				AccountID: 1,
				ItemID:    1,
				Qty:       1,
			},
			mockFn: func(args) fields {
				db, mock := CartNewMockDB()

				row := sqlmock.NewRows([]string{"id", "account_id", "item_id", "qty"}).
					AddRow(1, 2, 2, 2)
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `carts` WHERE id = ? AND `carts`.`deleted_at` IS NULL")).WillReturnRows(row)

				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta("UPDATE `carts` SET `account_id`=?,`item_id`=?,`qty`=?,`created_by`=?,`updated_by`=?,`deleted_by`=?,`created_at`=?,`updated_at`=?,`deleted_at`=? WHERE `carts`.`deleted_at` IS NULL AND `id` = ?")).WillReturnResult(sqlmock.NewResult(1, 1))
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

			p := NewCartRepository(dep.db)

			got, result := p.Update(tt.args.id, tt.args.cart)
			got.UpdatedAt = nil
			assert.Equal(t, tt.wantErr, result.Error)
			assert.Equal(t, tt.want, got)
		})
	}
}
