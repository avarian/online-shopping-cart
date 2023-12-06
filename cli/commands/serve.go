package commands

import (
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/avarian/online-shopping-cart/controllers"
	"github.com/avarian/online-shopping-cart/delivery/http"
	"github.com/avarian/online-shopping-cart/util"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Start server",
		RunE: func(cmd *cobra.Command, args []string) error {
			return serveCommand()
		},
		PreRun: func(cmd *cobra.Command, args []string) {
			rand.Seed(time.Now().UnixNano())
		},
	}
)

func serveCommand() (err error) {
	//
	// Initialize connections

	// Mysql database
	db := newMysqlDB("mysql")

	// Redis client
	// redis := newRedisClient(viper.GetString("redis.url"))
	// defer redis.Close()
	// jobs.SetRedisQueue(work.NewRedisQueue(redis))

	// validatorTranslate
	validator := util.ValidatorTranslate()

	//
	// Initialize Controllers
	//
	home := controllers.NewHomeController()
	account := controllers.NewAccountController(db, validator, viper.GetString("jwt_secret"))
	item := controllers.NewItemController(db, validator)
	cart := controllers.NewCartController(db, validator)
	voucher := controllers.NewVoucherController(db, validator)

	server := http.NewServer(viper.GetString("listen_address"),
		home,
		account,
		item,
		cart,
		voucher,
	)

	//
	// Start the process
	//
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		server.Start()
	}()

	done := make(chan os.Signal, 10)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	log.WithFields(log.Fields{
		"listen": viper.GetString("listen_address"),
		"now":    time.Now().Format("2006-01-02 15:04:05"),
	}).Info("server ready")
	<-done

	err = server.Stop()

	log.Info("waiting for remaining process to exit...")
	wg.Wait()
	log.Info("done.")

	return err
}
