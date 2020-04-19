package user

import (
	"kada/server/core"
	"kada/server/service/logger"
)

//Handler 用户控制器
type Handler int

//Create 创建用户
func (o *Handler) Create(args *CreateArgs, back *CreateBack) error {
	logger.Info("[user] create args: %v", args)
	
	user := new(User)
	user.ID = core.CreateUID()
	user.Token = args.Token
	user.UserName = args.UserName
	user.Phone = args.Phone
	user.Password = args.Password
	user.Device = args.Device
	user.Gold = 10000
	user.GoldBank = 10
	user.Gem = 1000
	user.Ticket = 100
	
	back.User = user
	logger.Info("[user] create back: %v", back)
	return nil
}
