package api

import (
	"github.com/DdZ-Fred/go-twitter-clone/utils"
)

func Api(globals utils.Globals) {
	AuthRouter(globals)
}
