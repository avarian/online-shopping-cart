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

func OrderNewMockDB() (*gorm.DB, sqlmock.Sqlmock) {
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

func Test_OrderIndex(t *testing.T) {
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
		want    []model.Order
		mockFn  func(a args) fields
	}{
		{
			name: "Success",
			args: args{
				&http.Request{
					URL: &url.URL{RawQuery: ""},
				},
			},
			want: []model.Order{
				{
					ID:          1,
					AccountID:   1,
					Address:     "address",
					PhoneNumber: "0888888",
					Total:       1,
					Status:      "ORDERED",
				},
				{
					ID:          2,
					AccountID:   2,
					Address:     "address2",
					PhoneNumber: "08888882",
					Total:       2,
					Status:      "ORDERED2",
				},
			},
			mockFn: func(args) fields {
				db, mock := OrderNewMockDB()

				row := sqlmock.NewRows([]string{"id", "account_id", "address", "phone_number", "total", "status"}).
					AddRow(1, 1, "address", "0888888", 1, "ORDERED").
					AddRow(2, 2, "address2", "08888882", 2, "ORDERED2")
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `orders` WHERE `orders`.`deleted_at` IS NULL")).WillReturnRows(row)

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

			p := NewOrderRepository(dep.db)

			got, result := p.Index(tt.args.r)
			assert.Equal(t, tt.wantErr, result.Error)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_OrderAll(t *testing.T) {
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
		want    []model.Order
		mockFn  func(a args) fields
	}{
		{
			name: "Success",
			args: args{
				&http.Request{
					URL: &url.URL{RawQuery: ""},
				},
			},
			want: []model.Order{
				{
					ID:          1,
					AccountID:   1,
					Address:     "address",
					PhoneNumber: "0888888",
					Total:       1,
					Status:      "ORDERED",
				},
				{
					ID:          2,
					AccountID:   2,
					Address:     "address2",
					PhoneNumber: "08888882",
					Total:       2,
					Status:      "ORDERED2",
				},
			},
			mockFn: func(args) fields {
				db, mock := OrderNewMockDB()

				row := sqlmock.NewRows([]string{"id", "account_id", "address", "phone_number", "total", "status"}).
					AddRow(1, 1, "address", "0888888", 1, "ORDERED").
					AddRow(2, 2, "address2", "08888882", 2, "ORDERED2")
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `orders` WHERE `orders`.`deleted_at` IS NULL")).WillReturnRows(row)

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

			p := NewOrderRepository(dep.db)

			got, result := p.All(tt.args.r)
			assert.Equal(t, tt.wantErr, result.Error)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_OrderOne(t *testing.T) {
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
		want    model.Order
		mockFn  func(a args) fields
	}{
		{
			name: "Success",
			args: args{
				&http.Request{
					URL: &url.URL{RawQuery: ""},
				},
			},
			want: model.Order{
				ID:          1,
				AccountID:   1,
				Address:     "address",
				PhoneNumber: "0888888",
				Total:       1,
				Status:      "ORDERED",
			},
			mockFn: func(args) fields {
				db, mock := OrderNewMockDB()

				row := sqlmock.NewRows([]string{"id", "account_id", "address", "phone_number", "total", "status"}).
					AddRow(1, 1, "address", "0888888", 1, "ORDERED").
					AddRow(2, 2, "address2", "08888882", 2, "ORDERED2")
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `orders` WHERE `orders`.`deleted_at` IS NULL")).WillReturnRows(row)

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

			p := NewOrderRepository(dep.db)

			got, result := p.One(tt.args.r)
			assert.Equal(t, tt.wantErr, result.Error)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_OrderOneById(t *testing.T) {
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
		want    model.Order
		mockFn  func(a args) fields
	}{
		{
			name: "Success",
			args: args{
				id: 1,
			},
			want: model.Order{
				ID:          1,
				AccountID:   1,
				Address:     "address",
				PhoneNumber: "0888888",
				Total:       1,
				Status:      "ORDERED",
			},
			mockFn: func(args) fields {
				db, mock := OrderNewMockDB()

				row := sqlmock.NewRows([]string{"id", "account_id", "address", "phone_number", "total", "status"}).
					AddRow(1, 1, "address", "0888888", 1, "ORDERED").
					AddRow(2, 2, "address2", "08888882", 2, "ORDERED2")
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `orders` WHERE id = ? AND `orders`.`deleted_at` IS NULL")).WillReturnRows(row)

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

			p := NewOrderRepository(dep.db)

			got, result := p.OneById(tt.args.id)
			assert.Equal(t, tt.wantErr, result.Error)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_OrderCreate(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}

	type args struct {
		order model.Order
	}

	tests := []struct {
		name    string
		args    args
		wantErr error
		want    model.Order
		mockFn  func(a args) fields
	}{
		{
			name: "Success",
			args: args{
				order: model.Order{
					ID:          1,
					AccountID:   1,
					Address:     "address",
					PhoneNumber: "0888888",
					Total:       1,
					Status:      "ORDERED",
				},
			},
			want: model.Order{
				ID:          1,
				AccountID:   1,
				Address:     "address",
				PhoneNumber: "0888888",
				Total:       1,
				Status:      "ORDERED",
				CreatedBy:   "SYSTEM",
				UpdatedBy:   "SYSTEM",
			},
			mockFn: func(args) fields {
				db, mock := OrderNewMockDB()

				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `orders` (`account_id`,`address`,`phone_number`,`total`,`status`,`created_by`,`updated_by`,`deleted_by`,`deleted_at`,`id`) VALUES (?,?,?,?,?,?,?,?,?,?)")).WillReturnResult(sqlmock.NewResult(1, 1))
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

			p := NewOrderRepository(dep.db)

			got, result := p.Create(tt.args.order)
			assert.Equal(t, tt.wantErr, result.Error)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_OrderUpdate(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}

	type args struct {
		id    int
		order model.Order
	}

	tests := []struct {
		name    string
		args    args
		wantErr error
		want    model.Order
		mockFn  func(a args) fields
	}{
		{
			name: "Success",
			args: args{
				id: 1,
				order: model.Order{
					ID:          1,
					AccountID:   1,
					Address:     "address",
					PhoneNumber: "0888888",
					Total:       1,
					Status:      "ORDERED",
				},
			},
			want: model.Order{
				ID:          1,
				AccountID:   1,
				Address:     "address",
				PhoneNumber: "0888888",
				Total:       1,
				Status:      "ORDERED",
			},
			mockFn: func(args) fields {
				db, mock := OrderNewMockDB()

				row := sqlmock.NewRows([]string{"id", "account_id", "address", "phone_number", "total", "status"}).
					AddRow(1, 1, "address", "0888888", 1, "ORDERED").
					AddRow(2, 2, "address2", "08888882", 2, "ORDERED2")
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `orders` WHERE id = ? AND `orders`.`deleted_at` IS NULL")).WillReturnRows(row)

				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta("UPDATE `orders` SET `account_id`=?,`address`=?,`phone_number`=?,`total`=?,`status`=?,`created_by`=?,`updated_by`=?,`deleted_by`=?,`created_at`=?,`updated_at`=?,`deleted_at`=? WHERE `orders`.`deleted_at` IS NULL AND `id` = ?")).WillReturnResult(sqlmock.NewResult(1, 1))
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

			p := NewOrderRepository(dep.db)

			got, result := p.Update(tt.args.id, tt.args.order)
			got.UpdatedAt = nil
			assert.Equal(t, tt.wantErr, result.Error)
			assert.Equal(t, tt.want, got)
		})
	}
}
