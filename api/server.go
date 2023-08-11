package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	codes "github.com/romsnl/httpepe.back/pkg"
)

type Server struct {
	router *gin.Engine
}

func (s *Server) Serve(addr string) error {
	return s.router.Run(addr)
}

func NewServer() *Server {
	router := gin.Default()
	s := &Server{
		router,
	}

	router.GET("/api/codes", s.getCodes)

	return s
}

func (s *Server) getCodes(c *gin.Context) {
	codes, err := codes.GetAllCodes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while fetching httpcodes"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": codes,
	})
}
