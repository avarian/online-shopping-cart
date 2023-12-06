package commands

import (
	"math/rand"
	"time"

	"github.com/avarian/online-shopping-cart/model"
	"github.com/spf13/cobra"
)

var (
	migrateCmd = &cobra.Command{
		Use:   "migrate",
		Short: "Migrate command",
		RunE: func(cmd *cobra.Command, args []string) error {
			return migrateCommand()
		},
		PreRun: func(cmd *cobra.Command, args []string) {
			rand.Seed(time.Now().UnixNano())
		},
	}
)

func migrateCommand() (err error) {
	// Postgres database
	db := newMysqlDB("mysql")
	db.AutoMigrate(
		&model.Account{},
		&model.Item{},
		&model.Cart{},
		&model.Order{},
		&model.Voucher{},
		&model.OrderItem{},
		&model.OrderVoucher{},
	)
	return nil
}
