/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2019-01-12
 * Time: 19:48
 */
package validate

import "github.com/gin-gonic/gin"

type V interface {
	MyValidate() gin.HandlerFunc
}

type SomeValidate struct {
	NewPostV        V
	NewCateV        V
	NewTagV         V
	NewSystemV      V
	NewLinkV        V
	NewAuthLoginV   V
	NewAuthRegister V
}

func NewValidate() *SomeValidate {
	return &SomeValidate{
		NewPostV:        &PostStoreV{},
		NewCateV:        &CateStoreV{},
		NewTagV:         &TagStoreV{},
		NewSystemV:      &SystemUpdateV{},
		NewLinkV:        &LinkStoreV{},
		NewAuthLoginV:   &AuthLoginV{},
		NewAuthRegister: &AuthRegisterV{},
	}
}
