package sort

import (
	"rank/dao"
)

// getRating 获取rating
func getRating(site, kind, id string) {
	return dao.GetRedisData(BuildKeyWithSiteKindID(site, kind, id))
}
