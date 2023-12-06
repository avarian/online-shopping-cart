package http

import (
	"context"
	"net/http"
	"time"

	"github.com/avarian/online-shopping-cart/controllers"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	listenAddress string
	http          *http.Server
}

func NewServer(listenAddress string,
	home *controllers.HomeController,
	account *controllers.AccountController,
	item *controllers.ItemController,
	cart *controllers.CartController,
	voucher *controllers.VoucherController,
) *Server {

	router := gin.Default()
	//
	// Http Routings
	//
	router.GET("/", home.GetHome)
	router.POST("/register", account.PostRegister)
	router.POST("/login", account.PostLogin)

	itemRoute := router.Group("/item").Use(Auth())
	{
		itemRoute.GET("/all", item.GetItems)
		itemRoute.GET("/:id", item.GetItemDetail)
		itemRoute.Use(Admin()).POST("/", item.PostCreateItem)
		itemRoute.Use(Admin()).PUT("/:id", item.PutEditItem)
		itemRoute.Use(Admin()).DELETE("/:id", item.DeleteItem)
	}

	cartRoute := router.Group("/cart").Use(Auth())
	{
		cartRoute.GET("/all", cart.GetCarts)
		cartRoute.GET("/:id", cart.GetCartDetail)
		cartRoute.POST("/", cart.PostCreateCartFromItem)
		cartRoute.PUT("/:id", cart.PutEditCart)
		cartRoute.DELETE("/:id", cart.DeleteCart)
	}

	voucherRoute := router.Group("/voucher").Use(Auth())
	{
		voucherRoute.GET("/all", voucher.GetVouchers)
		voucherRoute.GET("/:id", voucher.GetVoucherDetail)
		voucherRoute.Use(Admin()).POST("/", voucher.PostCreateVoucher)
		voucherRoute.Use(Admin()).PUT("/:id", voucher.PutEditVoucher)
		voucherRoute.Use(Admin()).DELETE("/:id", voucher.DeleteVoucher)
	}

	httpServer := &http.Server{
		Addr:              listenAddress,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		Handler:           router,
	}
	httpServer.SetKeepAlivesEnabled(true)

	srv := &Server{
		listenAddress: listenAddress,
		http:          httpServer,
	}
	return srv
}

func (s *Server) StartStopByContext(ctx context.Context) error {
	logCtx := log.WithField("listen", s.listenAddress)

	go func() {
		if err := s.http.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logCtx.Fatal(err)
		}
	}()

	logCtx.Info("server ready")
	<-ctx.Done()

	ctxDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := s.http.Shutdown(ctxDown); err != nil {
		logCtx.WithError(err).Error("http server shutdown failed")
		return err
	}

	logCtx.Info("http server shutdown gracefully")
	return nil

}

func (s *Server) Start() error {
	return s.http.ListenAndServe()
}

func (s *Server) Stop() error {
	logCtx := log.WithField("listen", s.listenAddress)

	ctxDown, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	if err := s.http.Shutdown(ctxDown); err != nil {
		logCtx.WithError(err).Error("http server shutdown failed")
		return err
	}

	logCtx.Info("http server shutdown gracefully")
	return nil
}
