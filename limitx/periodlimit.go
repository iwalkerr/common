package limitx

import (
	"github.com/zeromicro/go-zero/core/limit"
	"github.com/zeromicro/go-zero/core/logx"
)

func PeriodLimit(l *limit.PeriodLimit, key string) bool {
	code, err := l.Take(key)
	if err != nil {
		logx.Info(err)
		return true
	}

	switch code {
	case limit.OverQuota: // 拒绝
		logx.Infof("%s 请求拒绝了", key)
		return false
	case limit.Allowed: // 允许
		return true
	case limit.HitQuota: // 恰好达到了上限
		logx.Infof("%s 请求恰好达到上限了", key)
		return false
	default:
		logx.Infof("%s redis故障、过载", key)
		return true
	}
}
