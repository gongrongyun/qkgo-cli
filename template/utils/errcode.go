package utils

type ErrCode int
type ErrContent struct {
	HttpCode       int
	ErrMsg         string
	CauseStableErr bool
}
const (
	NoLogin           ErrCode = 100000

	NoPermGetPprof    ErrCode = 100001

)

var ErrMsgInfo = map[ErrCode]ErrContent{
	NoLogin: {
		HttpCode: 401,
		ErrMsg:   "暂未登陆，请先登陆",
	},
	NoPermGetPprof: {
		HttpCode: 403,
		ErrMsg:   "token error, no perm view pprof",
	},

}