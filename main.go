package main

import (
	. "ffAPI/controller"
	. "ffAPI/middleware"
	. "ffAPI/models"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	InitDB()
	defer CloseDB()

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "Cache-Control"}

	router.Use(static.Serve("/", static.LocalFile("./generated/build", true)))

	router.Use(cors.New(config))
	router.POST("/login", login)
	router.GET("/checkToken", AuthUser(), checkToken)

	router.GET("/pers", AuthUser(), pers)
	router.GET("/cities", AuthUser(), cities)

	router.GET("/file", file)

	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("ffwemding.dynv6.net"),
		Cache:      autocert.DirCache("/mytemp"),
	}

	autotls.RunWithManager(router, &m)
}

func login(c *gin.Context) {
	var login Login
	c.BindJSON(&login)
	retJWT := DoLogin(login, c)
	c.IndentedJSON(http.StatusOK, retJWT)
}

func checkToken(c *gin.Context) {
	tokenRes := CheckToken(c)
	c.IndentedJSON(http.StatusOK, tokenRes)
}

func pers(c *gin.Context) {
	persons := GetPersons()
	c.IndentedJSON(http.StatusOK, persons)
}

func cities(c *gin.Context) {
	cities := GetCities()
	c.IndentedJSON(http.StatusOK, cities)
}

func file(c *gin.Context) {
	numbers := [4]int{1, 2, 3, 4}
	pathZip, fileZip := CreateCityPDFs(numbers[:], 2022)
	c.Writer.Header().Set("Content-Disposition", "attachment; filename="+fileZip)
	c.Header("Content-Language", fileZip)
	c.File(pathZip + fileZip)
}
