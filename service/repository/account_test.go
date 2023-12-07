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

func AccountNewMockDB() (*gorm.DB, sqlmock.Sqlmock) {
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

func Test_AccountIndex(t *testing.T) {
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
		want    []model.Account
		mockFn  func(a args) fields
	}{
		{
			name: "Success",
			args: args{
				&http.Request{
					URL: &url.URL{RawQuery: ""},
				},
			},
			want: []model.Account{
				{
					ID:          1,
					Name:        "Test",
					Email:       "test@mail.com",
					PhoneNumber: "088888888888",
					Password:    "$2a$12$9ikh2RxE5tRPwWSMHJvEmOtZ1ISgJnkdFVkmmkGikw2DzXyCc3g5q",
					Address:     "alamat",
					Type:        "admin",
				},
				{
					ID:          2,
					Name:        "Test2",
					Email:       "test2@mail.com",
					PhoneNumber: "088888888882",
					Password:    "$2a$12$9ikh2RxE5tRPwWSMHJvEmOtZ1ISgJnkdFVkmmkGikw2DzXyCc3g5q",
					Address:     "alamat2",
					Type:        "admin",
				},
			},
			mockFn: func(args) fields {
				db, mock := AccountNewMockDB()

				row := sqlmock.NewRows([]string{"id", "name", "email", "phone_number", "password", "address", "type"}).
					AddRow(1, "Test", "test@mail.com", "088888888888", "$2a$12$9ikh2RxE5tRPwWSMHJvEmOtZ1ISgJnkdFVkmmkGikw2DzXyCc3g5q", "alamat", "admin").
					AddRow(2, "Test2", "test2@mail.com", "088888888882", "$2a$12$9ikh2RxE5tRPwWSMHJvEmOtZ1ISgJnkdFVkmmkGikw2DzXyCc3g5q", "alamat2", "admin")
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `accounts` WHERE `accounts`.`deleted_at` IS NULL")).WillReturnRows(row)

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

			p := NewAccountRepository(dep.db)

			got, result := p.Index(tt.args.r)
			assert.Equal(t, tt.wantErr, result.Error)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_AccountAll(t *testing.T) {
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
		want    []model.Account
		mockFn  func(a args) fields
	}{
		{
			name: "Success",
			args: args{},
			want: []model.Account{
				{
					ID:          1,
					Name:        "Test",
					Email:       "test@mail.com",
					PhoneNumber: "088888888888",
					Password:    "$2a$12$9ikh2RxE5tRPwWSMHJvEmOtZ1ISgJnkdFVkmmkGikw2DzXyCc3g5q",
					Address:     "alamat",
					Type:        "admin",
				},
				{
					ID:          2,
					Name:        "Test2",
					Email:       "test2@mail.com",
					PhoneNumber: "088888888882",
					Password:    "$2a$12$9ikh2RxE5tRPwWSMHJvEmOtZ1ISgJnkdFVkmmkGikw2DzXyCc3g5q",
					Address:     "alamat2",
					Type:        "admin",
				},
			},
			mockFn: func(args) fields {
				db, mock := AccountNewMockDB()

				row := sqlmock.NewRows([]string{"id", "name", "email", "phone_number", "password", "address", "type"}).
					AddRow(1, "Test", "test@mail.com", "088888888888", "$2a$12$9ikh2RxE5tRPwWSMHJvEmOtZ1ISgJnkdFVkmmkGikw2DzXyCc3g5q", "alamat", "admin").
					AddRow(2, "Test2", "test2@mail.com", "088888888882", "$2a$12$9ikh2RxE5tRPwWSMHJvEmOtZ1ISgJnkdFVkmmkGikw2DzXyCc3g5q", "alamat2", "admin")
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `accounts` WHERE `accounts`.`deleted_at` IS NULL")).WillReturnRows(row)

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

			p := NewAccountRepository(dep.db)

			got, result := p.All(tt.args.r)
			assert.Equal(t, tt.wantErr, result.Error)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_AccountOne(t *testing.T) {
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
		want    model.Account
		mockFn  func(a args) fields
	}{
		{
			name: "Success",
			args: args{},
			want: model.Account{
				ID:          1,
				Name:        "Test",
				Email:       "test@mail.com",
				PhoneNumber: "088888888888",
				Password:    "$2a$12$9ikh2RxE5tRPwWSMHJvEmOtZ1ISgJnkdFVkmmkGikw2DzXyCc3g5q",
				Address:     "alamat",
				Type:        "admin",
			},
			mockFn: func(args) fields {
				db, mock := AccountNewMockDB()

				row := sqlmock.NewRows([]string{"id", "name", "email", "phone_number", "password", "address", "type"}).
					AddRow(1, "Test", "test@mail.com", "088888888888", "$2a$12$9ikh2RxE5tRPwWSMHJvEmOtZ1ISgJnkdFVkmmkGikw2DzXyCc3g5q", "alamat", "admin")
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `accounts` WHERE `accounts`.`deleted_at` IS NULL")).WillReturnRows(row)

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

			p := NewAccountRepository(dep.db)

			got, result := p.One(tt.args.r)
			assert.Equal(t, tt.wantErr, result.Error)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_AccountOneById(t *testing.T) {
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
		want    model.Account
		mockFn  func(a args) fields
	}{
		{
			name: "Success",
			args: args{
				id: 1,
			},
			want: model.Account{
				ID:          1,
				Name:        "Test",
				Email:       "test@mail.com",
				PhoneNumber: "088888888888",
				Password:    "$2a$12$9ikh2RxE5tRPwWSMHJvEmOtZ1ISgJnkdFVkmmkGikw2DzXyCc3g5q",
				Address:     "alamat",
				Type:        "admin",
			},
			mockFn: func(args) fields {
				db, mock := AccountNewMockDB()

				row := sqlmock.NewRows([]string{"id", "name", "email", "phone_number", "password", "address", "type"}).
					AddRow(1, "Test", "test@mail.com", "088888888888", "$2a$12$9ikh2RxE5tRPwWSMHJvEmOtZ1ISgJnkdFVkmmkGikw2DzXyCc3g5q", "alamat", "admin")
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `accounts` WHERE id = ? AND `accounts`.`deleted_at` IS NULL")).WillReturnRows(row)

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

			p := NewAccountRepository(dep.db)

			got, result := p.OneById(tt.args.id)
			assert.Equal(t, tt.wantErr, result.Error)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_AccountCreate(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}

	type args struct {
		account model.Account
	}

	tests := []struct {
		name    string
		args    args
		wantErr error
		want    model.Account
		mockFn  func(a args) fields
	}{
		{
			name: "Success",
			args: args{
				account: model.Account{
					ID:          1,
					Name:        "Test",
					Email:       "test@mail.com",
					PhoneNumber: "088888888888",
					Password:    "$2a$12$9ikh2RxE5tRPwWSMHJvEmOtZ1ISgJnkdFVkmmkGikw2DzXyCc3g5q",
					Address:     "alamat",
					Type:        "admin",
				},
			},
			want: model.Account{
				ID:          1,
				Name:        "Test",
				Email:       "test@mail.com",
				PhoneNumber: "088888888888",
				Password:    "$2a$12$9ikh2RxE5tRPwWSMHJvEmOtZ1ISgJnkdFVkmmkGikw2DzXyCc3g5q",
				Address:     "alamat",
				Type:        "admin",
				CreatedBy:   "SYSTEM",
				UpdatedBy:   "SYSTEM",
			},
			mockFn: func(args) fields {
				db, mock := AccountNewMockDB()

				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `accounts` (`name`,`email`,`phone_number`,`password`,`address`,`type`,`created_by`,`updated_by`,`deleted_by`,`deleted_at`,`id`) VALUES (?,?,?,?,?,?,?,?,?,?,?)")).WillReturnResult(sqlmock.NewResult(1, 1))
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

			p := NewAccountRepository(dep.db)

			got, result := p.Create(tt.args.account)
			assert.Equal(t, tt.wantErr, result.Error)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_AccountUpdate(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}

	type args struct {
		id      int
		account model.Account
	}

	tests := []struct {
		name    string
		args    args
		wantErr error
		want    model.Account
		mockFn  func(a args) fields
	}{
		{
			name: "Success",
			args: args{
				id: 1,
				account: model.Account{
					Name:        "Test",
					Email:       "test@mail.com",
					PhoneNumber: "088888888888",
					Password:    "$2a$12$9ikh2RxE5tRPwWSMHJvEmOtZ1ISgJnkdFVkmmkGikw2DzXyCc3g5q",
					Address:     "alamat",
					Type:        "admin",
				},
			},
			want: model.Account{
				ID:          1,
				Name:        "Test",
				Email:       "test@mail.com",
				PhoneNumber: "088888888888",
				Password:    "$2a$12$9ikh2RxE5tRPwWSMHJvEmOtZ1ISgJnkdFVkmmkGikw2DzXyCc3g5q",
				Address:     "alamat",
				Type:        "admin",
			},
			mockFn: func(args) fields {
				db, mock := AccountNewMockDB()

				row := sqlmock.NewRows([]string{"id", "name", "email", "phone_number", "password", "address", "type"}).
					AddRow(1, "Test3", "test3@mail.com", "088888888883", "$2a$12$9ikh2RxE5tRPwWSMHJvEmOtZ1ISgJnkdFVkmmkGikw2DzXyCc3g5q", "alamat3", "admin")
				mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `accounts` WHERE id = ? AND `accounts`.`deleted_at` IS NULL")).WillReturnRows(row)

				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta("UPDATE `accounts` SET `name`=?,`email`=?,`phone_number`=?,`password`=?,`address`=?,`type`=?,`created_by`=?,`updated_by`=?,`deleted_by`=?,`created_at`=?,`updated_at`=?,`deleted_at`=? WHERE `accounts`.`deleted_at` IS NULL AND `id` = ?")).WillReturnResult(sqlmock.NewResult(1, 1))
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

			p := NewAccountRepository(dep.db)

			got, result := p.Update(tt.args.id, tt.args.account)
			got.UpdatedAt = nil
			assert.Equal(t, tt.wantErr, result.Error)
			assert.Equal(t, tt.want, got)
		})
	}
}
