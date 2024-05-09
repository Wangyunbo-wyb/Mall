package cache

import (
	"fmt"
	"strconv"
)

const (
	// RankKey 每日排名
	RankKey             = "rank"
	SkillProductKey     = "skill:product:%d"
	SkillProductListKey = "skill:product_list"
	SkillProductUserKey = "skill:user:%s"
)

// ProductViewKey 获取商品浏览数
func ProductViewKey(id uint) string {
	return fmt.Sprintf("view:product:%s", strconv.Itoa(int(id)))
}
