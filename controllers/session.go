package controllers

import(
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	
)

func GetSession(c *gin.Context) bool {
	
    session := sessions.Default(c)
    msg := session.Get("msg")
    
    if msg != nil {
        return true
    } else {
        return false
    }
}