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

func VoucherNewMockDB() (*gorm.DB, sqlmock.Sqlmock) {
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

func Test_VoucherIndex(t *testing.T) {
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
		want    []model.Voucher
		mockFn  func(a args) fields
	}{
		{
			name: "Success",
			args: args{
				&http.Request{
					URL: &url.URL{RawQuery: ""},
				},
			},
			want: []model.Voucher{
				{
					ID:          1,
					Code:        "1",
					Name:        "name",
					Description: "desc",
					Percentage:  1,
					Max:         1,
				},
				{
					ID:          2,
					Code:        "2",
					Name:        "name2",
					Description: "desc2",
					Percentage:  2,
					Max:         2,
				},
			},
			mockFn: func(args) fields {
				db, mock := VoucherNewMockDB()

				row := sqlmock.NewRows([]string{"id", "code", "name", "description", "percentage", "max"}).
					AddRow(1, "1", "name", "desc", 1, 1).
					AddRow(2, "2", "name2", "desc2", 2, 2)
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `vouchers` WHERE `vouchers`.`deleted_at` IS NULL")).WillReturnRows(row)

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

			p := NewVoucherRepository(dep.db)

			got, result := p.Index(tt.args.r)
			assert.Equal(t, tt.wantErr, result.Error)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_VoucherAll(t *testing.T) {
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
		want    []model.Voucher
		mockFn  func(a args) fields
	}{
		{
			name: "Success",
			args: args{
				&http.Request{
					URL: &url.URL{RawQuery: ""},
				},
			},
			want: []model.Voucher{
				{
					ID:          1,
					Code:        "1",
					Name:        "name",
					Description: "desc",
					Percentage:  1,
					Max:         1,
				},
				{
					ID:          2,
					Code:        "2",
					Name:        "name2",
					Description: "desc2",
					Percentage:  2,
					Max:         2,
				},
			},
			mockFn: func(args) fields {
				db, mock := VoucherNewMockDB()

				row := sqlmock.NewRows([]string{"id", "code", "name", "description", "percentage", "max"}).
					AddRow(1, "1", "name", "desc", 1, 1).
					AddRow(2, "2", "name2", "desc2", 2, 2)
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `vouchers` WHERE `vouchers`.`deleted_at` IS NULL")).WillReturnRows(row)

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

			p := NewVoucherRepository(dep.db)

			got, result := p.All(tt.args.r)
			assert.Equal(t, tt.wantErr, result.Error)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_VoucherOne(t *testing.T) {
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
		want    model.Voucher
		mockFn  func(a args) fields
	}{
		{
			name: "Success",
			args: args{
				&http.Request{
					URL: &url.URL{RawQuery: ""},
				},
			},
			want: model.Voucher{
				ID:          1,
				Code:        "1",
				Name:        "name",
				Description: "desc",
				Percentage:  1,
				Max:         1,
			},
			mockFn: func(args) fields {
				db, mock := VoucherNewMockDB()

				row := sqlmock.NewRows([]string{"id", "code", "name", "description", "percentage", "max"}).
					AddRow(1, "1", "name", "desc", 1, 1).
					AddRow(2, "2", "name2", "desc2", 2, 2)
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `vouchers` WHERE `vouchers`.`deleted_at` IS NULL")).WillReturnRows(row)

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

			p := NewVoucherRepository(dep.db)

			got, result := p.One(tt.args.r)
			assert.Equal(t, tt.wantErr, result.Error)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_VoucherOneById(t *testing.T) {
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
		want    model.Voucher
		mockFn  func(a args) fields
	}{
		{
			name: "Success",
			args: args{
				id: 1,
			},
			want: model.Voucher{
				ID:          1,
				Code:        "1",
				Name:        "name",
				Description: "desc",
				Percentage:  1,
				Max:         1,
			},
			mockFn: func(args) fields {
				db, mock := VoucherNewMockDB()

				row := sqlmock.NewRows([]string{"id", "code", "name", "description", "percentage", "max"}).
					AddRow(1, "1", "name", "desc", 1, 1).
					AddRow(2, "2", "name2", "desc2", 2, 2)
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `vouchers` WHERE id = ? AND `vouchers`.`deleted_at` IS NULL")).WillReturnRows(row)

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

			p := NewVoucherRepository(dep.db)

			got, result := p.OneById(tt.args.id)
			assert.Equal(t, tt.wantErr, result.Error)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_VoucherCreate(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}

	type args struct {
		voucher model.Voucher
	}

	tests := []struct {
		name    string
		args    args
		wantErr error
		want    model.Voucher
		mockFn  func(a args) fields
	}{
		{
			name: "Success",
			args: args{
				voucher: model.Voucher{
					ID:          1,
					Code:        "1",
					Name:        "name",
					Description: "desc",
					Percentage:  1,
					Max:         1,
				},
			},
			want: model.Voucher{
				ID:          1,
				Code:        "1",
				Name:        "name",
				Description: "desc",
				Percentage:  1,
				Max:         1,
				CreatedBy:   "SYSTEM",
				UpdatedBy:   "SYSTEM",
			},
			mockFn: func(args) fields {
				db, mock := VoucherNewMockDB()

				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `vouchers` (`code`,`name`,`description`,`percentage`,`max`,`created_by`,`updated_by`,`deleted_by`,`deleted_at`,`id`) VALUES (?,?,?,?,?,?,?,?,?,?)")).WillReturnResult(sqlmock.NewResult(1, 1))
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

			p := NewVoucherRepository(dep.db)

			got, result := p.Create(tt.args.voucher)
			assert.Equal(t, tt.wantErr, result.Error)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_VoucherUpdate(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}

	type args struct {
		id      int
		voucher model.Voucher
	}

	tests := []struct {
		name    string
		args    args
		wantErr error
		want    model.Voucher
		mockFn  func(a args) fields
	}{
		{
			name: "Success",
			args: args{
				id: 1,
				voucher: model.Voucher{
					ID:          1,
					Code:        "1",
					Name:        "name",
					Description: "desc",
					Percentage:  1,
					Max:         1,
				},
			},
			want: model.Voucher{
				ID:          1,
				Code:        "1",
				Name:        "name",
				Description: "desc",
				Percentage:  1,
				Max:         1,
			},
			mockFn: func(args) fields {
				db, mock := VoucherNewMockDB()

				row := sqlmock.NewRows([]string{"id", "code", "name", "description", "percentage", "max"}).
					AddRow(1, "1", "name", "desc", 1, 1).
					AddRow(2, "2", "name2", "desc2", 2, 2)
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `vouchers` WHERE id = ? AND `vouchers`.`deleted_at` IS NULL")).WillReturnRows(row)

				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta("UPDATE `vouchers` SET `code`=?,`name`=?,`description`=?,`percentage`=?,`max`=?,`created_by`=?,`updated_by`=?,`deleted_by`=?,`created_at`=?,`updated_at`=?,`deleted_at`=? WHERE `vouchers`.`deleted_at` IS NULL AND `id` = ?")).WillReturnResult(sqlmock.NewResult(1, 1))
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

			p := NewVoucherRepository(dep.db)

			got, result := p.Update(tt.args.id, tt.args.voucher)
			got.UpdatedAt = nil
			assert.Equal(t, tt.wantErr, result.Error)
			assert.Equal(t, tt.want, got)
		})
	}
}
