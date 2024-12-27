package engine

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sensfo/server/domain"
	"github.com/sensfo/server/internal/logger"
)

// ─────────────────────────────────────────────────────────────────────────────

type ginContextWrapper struct {
	ginCtx *gin.Context
}

func (g *ginContextWrapper) JSON(statusCode int, obj interface{}) {
	g.ginCtx.JSON(statusCode, obj)
}

func (g *ginContextWrapper) String(statusCode int, msg string) {
	g.ginCtx.String(statusCode, msg)
}

func (g *ginContextWrapper) Param(key string) string {
	return g.ginCtx.Param(key)
}

func (g *ginContextWrapper) Query(key string) string {
	return g.ginCtx.Query(key)
}

func (g *ginContextWrapper) Bind(obj interface{}) error {
	return g.ginCtx.Bind(obj)
}

// ─────────────────────────────────────────────────────────────────────────────

type GinEngine struct {
	instance *gin.Engine
	server   *http.Server
	logger   logger.Logger
	routes   []domain.Route
}

func (g *GinEngine) ListenAndServe(port int) error {
	g.server = &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: g.instance,
	}
	go g.serve()

	return nil
}

func (g *GinEngine) serve() {
	g.bindRoutes()

	g.logger.Info("Server started", "url", "http://localhost"+g.server.Addr)

	if err := g.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	sig := <-quit

	g.logger.Info("Received signal", "signal", sig)

	g.Shutdown()
}

func (g *GinEngine) bindRoutes() {
	for _, route := range g.routes {
		route.Bind(g)
	}
}

func (g *GinEngine) Shutdown() error {
	if g.server == nil {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return g.server.Shutdown(ctx)
}

func (g *GinEngine) GET(pattern string, handlers ...domain.RouteHandlerFunc) {
	g.instance.GET(pattern, buildHandlerFuncs(handlers)...)
}

func (g *GinEngine) POST(pattern string, handlers ...domain.RouteHandlerFunc) {
	g.instance.POST(pattern, buildHandlerFuncs(handlers)...)
}

func (g *GinEngine) PUT(pattern string, handlers ...domain.RouteHandlerFunc) {
	g.instance.PUT(pattern, buildHandlerFuncs(handlers)...)
}
func (g *GinEngine) DELETE(pattern string, handlers ...domain.RouteHandlerFunc) {
	g.instance.DELETE(pattern, buildHandlerFuncs(handlers)...)
}

func (g *GinEngine) HEAD(pattern string, handlers ...domain.RouteHandlerFunc) {
	g.instance.HEAD(pattern, buildHandlerFuncs(handlers)...)
}

func (g *GinEngine) EnableReleaseMode() {
	gin.SetMode(gin.ReleaseMode)
}

func NewGinEngine(routes []domain.Route, logger logger.Logger) domain.Engine {
	instance := gin.New()
	instance.Use(gin.Logger(), gin.Recovery())

	return &GinEngine{
		instance: instance,
		logger:   logger,
		routes:   routes,
	}
}

func buildHandlerFuncs(handlers []domain.RouteHandlerFunc) []gin.HandlerFunc {
	handlerFuncs := make([]gin.HandlerFunc, len(handlers))

	for _, handler := range handlers {
		handlerFuncs = append(handlerFuncs, ginAdapter(handler))
	}

	return handlerFuncs
}

// ginAdapter adapts a generic RouteHandlerFunc to gin.HandlerFunc
func ginAdapter(handler domain.RouteHandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		wrappedContext := &ginContextWrapper{ginCtx: c}
		handler(wrappedContext)
	}
}
