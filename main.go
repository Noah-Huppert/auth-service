package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/Noah-Huppert/auth-service/http/resp"
	"github.com/Noah-Huppert/auth-service/confloader"
)

conf := Conf{}

func routePublicConfig(c *gin.Context) {
	resp.Resp(gin.H { 
		"client_id": "111029956109-3fjs8mun37hds1p6h24cb82j01fdhavi.apps.googleusercontent.com",
		"scopes": "profile",
	}, 200, c)
}

type GoogleIdTokenVerifyResp struct {
	Email string `json:"email"`
	EmailVerfied string `json:"email_verified"`
	Name string `json:"name"`
	ProfilePicUri string `json:"picture"`
}

func routeVerify(c *gin.Context) {
	idToken := c.PostForm("id_token")

	url := fmt.Sprintf("https://www.googleapis.com/oauth2/v3/tokeninfo?id_token=%v", idToken)
	r, err := http.Get(url)
	defer r.Body.Close()

	if err != nil {
		resp.Resp(gin.H {
			"error": "error validating id token",
		}, 500, c)
		return
	}

	idTokenResp := GoogleIdTokenVerifyResp{}
	json.NewDecoder(r.Body).Decode(&idTokenResp)
	c.String(200, fmt.Sprintf("%#v", idTokenResp))
}

func main() {
	fmt.Println("Loading configuration");
	conf = 

	router := gin.Default()

	router.GET("/v1/auth/google/public_config", routePublicConfig)
	router.POST("/v1/auth/google/verify", routeVerify);

	router.Run(":80")
}
