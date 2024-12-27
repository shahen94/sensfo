package domain

type Context interface {
	JSON(statusCode int, obj interface{})
	String(statusCode int, msg string)
	Param(key string) string
	Query(key string) string
	Bind(obj interface{}) error
}

type Engine interface {
	Router
	ListenAndServe(int) error
	EnableReleaseMode()
	Shutdown() error
}

type Router interface {
	GET(string, ...RouteHandlerFunc)
	POST(string, ...RouteHandlerFunc)
	PUT(string, ...RouteHandlerFunc)
	DELETE(string, ...RouteHandlerFunc)
	HEAD(string, ...RouteHandlerFunc)
}

type Route interface {
	Bind(Router)
}

type RouteHandlerFunc = func(Context)
