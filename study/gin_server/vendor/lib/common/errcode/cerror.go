//
// !!!! 此文件为生成,不要编辑
//
package errcode

const (
	OKCode int = iota // "OK"

	CommonBegin            = 1  // "通用开始"
	CommonErr              = 2  // "通用错误"
	ParamInvalid           = 3  // "请求参数错误"
	OperDbErr              = 4  // "操作数据库失败"
	ResourceRaceErr        = 5  // "资源操作中"
	NotAuthorized          = 6  // "越权操作"
	InterServerConnectFail = 7  // "内部服务请求错误"
	InvalidIPErr           = 8  // "非法的来源IP"
	TypeConvertErr         = 9  // "内部类型转换错误"
	ExternalRequestErr     = 10 // "请求外部资源错误"
	OperCacheErr           = 11 // "缓存失败"
	InnerConfigError       = 12 // "业务配置错误"
	TaskRunning            = 13 // "任务正在处理中"

	TokenInvalidErr        = 100  // "非法token"
	IssuerInvalidErr       = 101  // "非法的token签发者"
	PageSizeOverFlow       = 102  // "分页页大小超长"
	UserNotExist           = 103  // "用户名不存在"
	PwdNotRight            = 104  // "密码错误"
	GenTokenFaild          = 105  // "签发令牌失败"
	UserAlreadyExist       = 106  // "用户名已经存在"
	InnerServerFail        = 107  // "内部服务连接失败"
	InnerServerResErr      = 108  // "内部服务响应格式错误"
	TokenTimeOut           = 109  // "token授权超时"
	WechatGetTokenFail     = 110  // "获取微信access_token失败"
	WechatRefreshTokenFail = 111  // "刷新微信access_token失败"
	WechatCheckTokenFail   = 112  // "检查微信access_token失败"
	WechatGetUserInfoFail  = 113  // "获取微信用户信息失败"
	CommonEnd              = 1000 // "通用结束"

	GetUinfoFail     = 1001 // "获取用户信息失败"
	CreateUserFail   = 1002 // "创建用户失败"
	UnsupportChannel = 1003 // "未支持的用户渠道"

	QuestionNotEnough   = 1101 // 题库不足
	InvalidQuesCategory = 1102 // 非法题库类型

	InvalidRoomType    = 1201 // 非法玩法类型
	InvalidEntranceId  = 1202 // 无效入口
	RoomOperFail       = 1203 // 房间资产调度异常
	UserAlreadyIn      = 1204 // 用户已经在房间内
	CreateRoomFail     = 1205 // 创建房间失败
	RoomFull           = 1206 // 房间已满
	UserNotIn          = 1207 // 用户不在房间中
	RoomNotExist       = 1208 // 房间不存在
	RoomTokenFormatErr = 1209 // 房间令牌格式错误
)

func GetErrMsg(code int) string {
	val, ok := g_err_map[code]
	if ok {
		return val
	}
	return "内部错误"
}

var g_err_map = map[int]string{
	OKCode:                 "OK",
	TokenTimeOut:           "token授权超时",
	InnerConfigError:       "业务配置错误",
	TaskRunning:            "任务正在处理中",
	InnerServerResErr:      "内部服务响应格式错误",
	InterServerConnectFail: "内部服务请求错误",
	InnerServerFail:        "内部服务连接失败",
	TypeConvertErr:         "内部类型转换错误",
	PageSizeOverFlow:       "分页页大小超长",
	CreateUserFail:         "创建用户失败",
	WechatRefreshTokenFail: "刷新微信access_token失败",
	PwdNotRight:            "密码错误",
	OperDbErr:              "操作数据库失败",
	UnsupportChannel:       "未支持的用户渠道",
	WechatCheckTokenFail:   "检查微信access_token失败",
	UserNotExist:           "用户名不存在",
	UserAlreadyExist:       "用户名已经存在",
	GenTokenFaild:          "签发令牌失败",
	OperCacheErr:           "缓存失败",
	WechatGetTokenFail:     "获取微信access_token失败",
	WechatGetUserInfoFail:  "获取微信用户信息失败",
	GetUinfoFail:           "获取用户信息失败",
	ParamInvalid:           "请求参数错误",
	ExternalRequestErr:     "请求外部资源错误",
	ResourceRaceErr:        "资源操作中",
	NotAuthorized:          "越权操作",
	CommonBegin:            "通用开始",
	CommonEnd:              "通用结束",
	CommonErr:              "通用错误",
	TokenInvalidErr:        "非法token",
	IssuerInvalidErr:       "非法的token签发者",
	InvalidIPErr:           "非法的来源IP",
}
