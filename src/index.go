package contoller

import (
	"net/http"
	"github.com/gin-gonic/gin"

)

func IndexGET(c *gin.Context){
  c,String(http.StatusOK, "Hello, world")
}

