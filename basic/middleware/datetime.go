package middleware

import (
	"math"
	"time"
)

/*	@package Timelag 计算当前时间和传入时间百分比
*	@param   options 数据类型map
*	@param	 NowTime 数据单位秒
*	@param	 TTL 	 数据单位秒
 */
func TimeLag(options configMap) int64 {
	timestamp := options["NowTime"]
	slp := options["TTl"]
	stamp := timestamp.(int64)
	slpn := slp.(int64)
	now := time.Now()
	lag := now.Unix() - stamp
	residue := lag * 100 / slpn
	num := math.Abs(float64(residue))
	return int64(num)
}
