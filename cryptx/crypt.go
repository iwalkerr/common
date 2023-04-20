package cryptx

import (
	"common/gconv"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"sort"
	"time"

	"golang.org/x/crypto/scrypt"
)

const SignExpiredTime = 60

func PasswordEncrypt(salt, password string) string {
	dk, _ := scrypt.Key([]byte(password), []byte(salt), 32768, 8, 1, 32)
	return fmt.Sprintf("%x", string(dk))
}

// 后台接收到先验证
// http改为https
// 验证sign是否重复提交
// 0.查看此用户是否在黑名单中,如果在黑名单中，返回提示 ===> 此用户被拉黑,请联系客服
// 1.获取timestamp,比对时间是否过期(超过60s)，如果过期返回 ===> 接口超时
// 2.获取私钥
// 3.私钥+获取用户传的数据 ===> 加密成sign
// 4.比对sign是否一致

// 安全拦截
// 单用户单位时间内访问API次数限制
// 多用户限制API调用并发数
// 设置限制时间 ===> 设置黑名单

// 加密原则
// 1.请求来源的合法性、
// 2.请求参数被攥改
// 3.请求的唯一性(不可复制)(加时间戳60秒后过期)

// 注意中文需要url编码
func GetSign(jsonMap map[string]string, host string) string {
	var i int
	addressKey := make([]string, len(jsonMap))
	for k := range jsonMap {
		addressKey[i] = k
		i++
	}

	// 根据key的ASCII字符串排序
	sort.Strings(addressKey)

	// 字符串拼接
	var str string
	for _, k := range addressKey {
		str += k + fmt.Sprintf("%v", jsonMap[k])
	}
	str += host
	// md5加密
	return MD5([]byte(str))
}

func MD5(data []byte) string {
	_md5 := md5.New()
	_, _ = _md5.Write(data)
	return hex.EncodeToString(_md5.Sum([]byte("")))
}

// 参数验签中间件
func SignValid(mapdata map[string]string, host string) error {
	reqSign := mapdata["Sign"]

	// 1.验证接口是否超时
	sec := time.Now().UnixMilli() - gconv.Int64(mapdata["Ct"])
	if sec > SignExpiredTime*1000 {
		return errors.New("接口超时")
	}
	delete(mapdata, "Sign")
	mapdata["Salt"] = "4g89aojga#hhjj666777744@@ggjjj" // 此参数也不传

	// 2.验证参数合法性
	sign := GetSign(mapdata, host)
	if reqSign != sign {
		return errors.New("参数被串改")
	}
	return nil
}
