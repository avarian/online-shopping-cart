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

func OrderVoucherNewMockDB() (*gorm.DB, sqlmock.Sqlmock) {
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

func Test_OrderVoucherIndex(t *testing.T) {
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
		want    []model.OrderVoucher
		mockFn  func(a args) fields
	}{
		{
			name: "Success",
			args: args{
				&http.Request{
					URL: &url.URL{RawQuery: ""},
				},
			},
			want: []model.OrderVoucher{
				{
					ID:          1,
					OrderID:     1,
					VoucherID:   1,
					Code:        "code",
					Name:        "name",
					Description: "desc",
					Percentage:  1,
					Max:         1,
					Total:       1,
					Applied:     1,
				},
				{
					ID:          2,
					OrderID:     2,
					VoucherID:   2,
					Code:        "code2",
					Name:        "name2",
					Description: "desc2",
					Percentage:  2,
					Max:         2,
					Total:       2,
					Applied:     2,
				},
			},
			mockFn: func(args) fields {
				db, mock := OrderVoucherNewMockDB()

				row := sqlmock.NewRows([]string{"id", "order_id", "voucher_id", "code", "name", "description", "percentage", "max", "total", "applied"}).
					AddRow(1, 1, 1, "code", "name", "desc", 1, 1, 1, 1).
					AddRow(2, 2, 2, "code2", "name2", "desc2", 2, 2, 2, 2)
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `order_vouchers` WHERE `order_vouchers`.`deleted_at` IS NULL")).WillReturnRows(row)

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

			p := NewOrderVoucherRepository(dep.db)

			got, result := p.Index(tt.args.r)
			assert.Equal(t, tt.wantErr, result.Error)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_OrderVoucherAll(t *testing.T) {
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
		want    []model.OrderVoucher
		mockFn  func(a args) fields
	}{
		{
			name: "Success",
			args: args{
				&http.Request{
					URL: &url.URL{RawQuery: ""},
				},
			},
			want: []model.OrderVoucher{
				{
					ID:          1,
					OrderID:     1,
					VoucherID:   1,
					Code:        "code",
					Name:        "name",
					Description: "desc",
					Percentage:  1,
					Max:         1,
					Total:       1,
					Applied:     1,
				},
				{
					ID:          2,
					OrderID:     2,
					VoucherID:   2,
					Code:        "code2",
					Name:        "name2",
					Description: "desc2",
					Percentage:  2,
					Max:         2,
					Total:       2,
					Applied:     2,
				},
			},
			mockFn: func(args) fields {
				db, mock := OrderVoucherNewMockDB()

				row := sqlmock.NewRows([]string{"id", "order_id", "voucher_id", "code", "name", "description", "percentage", "max", "total", "applied"}).
					AddRow(1, 1, 1, "code", "name", "desc", 1, 1, 1, 1).
					AddRow(2, 2, 2, "code2", "name2", "desc2", 2, 2, 2, 2)
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `order_vouchers` WHERE `order_vouchers`.`deleted_at` IS NULL")).WillReturnRows(row)

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

			p := NewOrderVoucherRepository(dep.db)

			got, result := p.All(tt.args.r)
			assert.Equal(t, tt.wantErr, result.Error)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_OrderVoucherOne(t *testing.T) {
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
		want    model.OrderVoucher
		mockFn  func(a args) fields
	}{
		{
			name: "Success",
			args: args{
				&http.Request{
					URL: &url.URL{RawQuery: ""},
				},
			},
			want: model.OrderVoucher{
				ID:          1,
				OrderID:     1,
				VoucherID:   1,
				Code:        "code",
				Name:        "name",
				Description: "desc",
				Percentage:  1,
				Max:         1,
				Total:       1,
				Applied:     1,
			},
			mockFn: func(args) fields {
				db, mock := OrderVoucherNewMockDB()

				row := sqlmock.NewRows([]string{"id", "order_id", "voucher_id", "code", "name", "description", "percentage", "max", "total", "applied"}).
					AddRow(1, 1, 1, "code", "name", "desc", 1, 1, 1, 1).
					AddRow(2, 2, 2, "code2", "name2", "desc2", 2, 2, 2, 2)
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `order_vouchers` WHERE `order_vouchers`.`deleted_at` IS NULL")).WillReturnRows(row)

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

			p := NewOrderVoucherRepository(dep.db)

			got, result := p.One(tt.args.r)
			assert.Equal(t, tt.wantErr, result.Error)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_OrderVoucherOneById(t *testing.T) {
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
		want    model.OrderVoucher
		mockFn  func(a args) fields
	}{
		{
			name: "Success",
			args: args{
				id: 1,
			},
			want: model.OrderVoucher{
				ID:          1,
				OrderID:     1,
				VoucherID:   1,
				Code:        "code",
				Name:        "name",
				Description: "desc",
				Percentage:  1,
				Max:         1,
				Total:       1,
				Applied:     1,
			},
			mockFn: func(args) fields {
				db, mock := OrderVoucherNewMockDB()

				row := sqlmock.NewRows([]string{"id", "order_id", "voucher_id", "code", "name", "description", "percentage", "max", "total", "applied"}).
					AddRow(1, 1, 1, "code", "name", "desc", 1, 1, 1, 1).
					AddRow(2, 2, 2, "code2", "name2", "desc2", 2, 2, 2, 2)
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `order_vouchers` WHERE id = ? AND `order_vouchers`.`deleted_at` IS NULL")).WillReturnRows(row)

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

			p := NewOrderVoucherRepository(dep.db)

			got, result := p.OneById(tt.args.id)
			assert.Equal(t, tt.wantErr, result.Error)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_OrderVoucherCreate(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}

	type args struct {
		order_voucher model.OrderVoucher
	}

	tests := []struct {
		name    string
		args    args
		wantErr error
		want    model.OrderVoucher
		mockFn  func(a args) fields
	}{
		{
			name: "Success",
			args: args{
				order_voucher: model.OrderVoucher{
					ID:          1,
					OrderID:     1,
					VoucherID:   1,
					Code:        "code",
					Name:        "name",
					Description: "desc",
					Percentage:  1,
					Max:         1,
					Total:       1,
					Applied:     1,
				},
			},
			want: model.OrderVoucher{
				ID:          1,
				OrderID:     1,
				VoucherID:   1,
				Code:        "code",
				Name:        "name",
				Description: "desc",
				Percentage:  1,
				Max:         1,
				Total:       1,
				Applied:     1,
				CreatedBy:   "SYSTEM",
				UpdatedBy:   "SYSTEM",
			},
			mockFn: func(args) fields {
				db, mock := OrderVoucherNewMockDB()

				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `order_vouchers` (`order_id`,`voucher_id`,`code`,`name`,`description`,`percentage`,`max`,`total`,`applied`,`created_by`,`updated_by`,`deleted_by`,`deleted_at`,`id`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)")).WillReturnResult(sqlmock.NewResult(1, 1))
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

			p := NewOrderVoucherRepository(dep.db)

			got, result := p.Create(tt.args.order_voucher)
			assert.Equal(t, tt.wantErr, result.Error)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_OrderVoucherUpdate(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}

	type args struct {
		id            int
		order_voucher model.OrderVoucher
	}

	tests := []struct {
		name    string
		args    args
		wantErr error
		want    model.OrderVoucher
		mockFn  func(a args) fields
	}{
		{
			name: "Success",
			args: args{
				id: 1,
				order_voucher: model.OrderVoucher{
					ID:          1,
					OrderID:     1,
					VoucherID:   1,
					Code:        "code",
					Name:        "name",
					Description: "desc",
					Percentage:  1,
					Max:         1,
					Total:       1,
					Applied:     1,
				},
			},
			want: model.OrderVoucher{
				ID:          1,
				OrderID:     1,
				VoucherID:   1,
				Code:        "code",
				Name:        "name",
				Description: "desc",
				Percentage:  1,
				Max:         1,
				Total:       1,
				Applied:     1,
			},
			mockFn: func(args) fields {
				db, mock := OrderVoucherNewMockDB()

				row := sqlmock.NewRows([]string{"id", "order_id", "voucher_id", "code", "name", "description", "percentage", "max", "total", "applied"}).
					AddRow(1, 1, 1, "code", "name", "desc", 1, 1, 1, 1).
					AddRow(2, 2, 2, "code2", "name2", "desc2", 2, 2, 2, 2)
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `order_vouchers` WHERE id = ? AND `order_vouchers`.`deleted_at` IS NULL")).WillReturnRows(row)

				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta("UPDATE `order_vouchers` SET `order_id`=?,`voucher_id`=?,`code`=?,`name`=?,`description`=?,`percentage`=?,`max`=?,`total`=?,`applied`=?,`created_by`=?,`updated_by`=?,`deleted_by`=?,`created_at`=?,`updated_at`=?,`deleted_at`=? WHERE `order_vouchers`.`deleted_at` IS NULL AND `id` = ?")).WillReturnResult(sqlmock.NewResult(1, 1))
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

			p := NewOrderVoucherRepository(dep.db)

			got, result := p.Update(tt.args.id, tt.args.order_voucher)
			got.UpdatedAt = nil
			assert.Equal(t, tt.wantErr, result.Error)
			assert.Equal(t, tt.want, got)
		})
	}
}
