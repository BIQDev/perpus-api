package modules

import (
	_ "github.com/BIQDev/perpus-api/internal/modules/booklist"
	"github.com/BIQDev/perpus-api/internal/router"
)

func init(){
	router.MiddlewareInit()
}
