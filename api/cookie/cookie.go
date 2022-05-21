package cookie

import (
	//"fmt"

	"reflect"

	"github.com/gin-contrib/sessions"

	"github.com/gin-gonic/gin"
)

type Ax47 interface {
}

func Cookieee_set(c *gin.Context, data interface{}) error {
	session := sessions.Default(c)
	v := reflect.ValueOf(data)
	typeOfS := v.Type()
	for i := 0; i < v.NumField(); i++ {

		session.Set(typeOfS.Field(i).Name, v.Field(i))
	}

	err := session.Save()

	return err

}

type res struct {
	Sso    interface{}
	status bool
}

func Cookieee_Get(c *gin.Context, data ...string) map[string]interface{} {

	result := make(map[string]interface{})
	session := sessions.Default(c)
	for _, v := range data {
		aa := session.Get(v)
		result[v] = aa
	}
	return result
}
