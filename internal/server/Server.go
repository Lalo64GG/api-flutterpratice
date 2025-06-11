package server

import (
	"api/internal/config"
	"api/internal/user/infraestructure/http/routes"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine   *gin.Engine
	host     string
	port     string
	httpAddr string

	useHTTPS bool
	certFile string
	keyFile  string
}

func NewServer(host, port string, useHTTPS bool, certFile, keyFile string) Server {
	gin.SetMode(gin.ReleaseMode)

	srv := Server{
		engine:   gin.New(),
		host:     host,
		port:     port,
		httpAddr: host + ":" + port,

		useHTTPS: useHTTPS,
		certFile: certFile,
		keyFile:  keyFile,
	}

	srv.engine.Use(func(c *gin.Context) {
		// Seguridad básica
		c.Header("X-Frame-Options", "DENY")
		c.Header("Content-Security-Policy", "default-src 'self'; connect-src *; font-src *; script-src-elem * 'unsafe-inline'; img-src * data:; style-src * 'unsafe-inline';")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
		c.Header("Referrer-Policy", "strict-origin")
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("Permissions-Policy", "geolocation=(),midi=(),sync-xhr=(),microphone=(),camera=(),magnetometer=(),gyroscope=(),fullscreen=(self),payment=()")
		c.Next()
	})

	srv.engine.Use(gin.Recovery())
	srv.engine.Use(gin.Logger())
	srv.engine.Use(config.ConfigurationCors())

	srv.engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong!"})
	})

	userRoutes := srv.engine.Group("/v1/user")
	routes.Routes(userRoutes)

	return srv
}

func (s *Server) Run() error {
	log.Println("Starting server on " + s.httpAddr)
	if s.useHTTPS {
		log.Println("Serving with HTTPS")
		return s.engine.RunTLS(s.httpAddr, s.certFile, s.keyFile)
	} else {
		log.Println("Serving with HTTP")
		return s.engine.Run(s.httpAddr)
	}
}
