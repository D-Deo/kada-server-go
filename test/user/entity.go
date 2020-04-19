package user

//User 用户数据
type User struct {
	ID           uint32 //用户ID
	UserName     string //用户名
	NickName     string //用户昵称
	Gold         int64  //携带金币
	GoldBank     int64  //银行金币
	Gem          int64  //钻石
	Ticket       int64  //点卷
	Face         string //头像
	Vip          uint32 //VIP等级
	Phone        string //手机号
	Password     string //登录密码
	PasswordBank string //银行密码
	Token        string //唯一特征码
	Device       uint32 //设备类型（0未知 1Android 2iOS 3Web 4Win 5Mac）
}
