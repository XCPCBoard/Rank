package sort

import (
	"XCPCBoard/utils/keys"
	log "github.com/sirupsen/logrus"
	"rank/dao"
)

// getSiteId 获取用户网站ID
func getSiteId(site, userid string) string {
	res, err := dao.GetRedisData(keys.BuildKeyWithSiteID(site, userid))
	if err != nil {
		log.Errorf("get %s %s Error: %v", site, userid, err)
	}
	str, ok := res.(string)
	if !ok {
		log.Errorf("get %s %s  Error: type is wrong", site, userid)
	}
	return str
}

// getLastSiteKindIdNum 获取时效_网站_类型_用户id数据
func getLastSiteKindIdNum(last, site, kind, id string) int {
	res, err := dao.GetRedisData(keys.BuildKeyWithLastSiteKindID(last, site, kind, id))
	if err != nil {
		log.Errorf("get %s %s %s %s Error: %v", last, site, kind, id, err)
	}
	num, ok := res.(int)
	if !ok {
		log.Errorf("get %s %s %s %s Error: type is wrong", last, site, kind, id)
	}
	return num
}

// getLastSiteKindDifficultyIdNum 获取时效_网站_类型_难度_用户id数据
func getLastSiteKindDifficultyIdNu(last, site, kind, difficulty, id string) int {
	res, err := dao.GetRedisData(keys.BuildKeyWithLastSiteKindDifficultyID(last, site, kind, difficulty, id))
	if err != nil {
		log.Errorf("get %s %s %s %s %s Error: %v", last, site, kind, difficulty, id, err)
	}
	num, ok := res.(int)
	if !ok {
		log.Errorf("get %s %s %s %s %s Error: type is wrong", last, site, kind, difficulty, id)
	}
	return num
}

//getLastKindIDData 获取时效_类型_用户id数据
func getLastKindIDData(last, kind, id string) int {
	res, err := dao.GetRedisData(BuildKeyWithLastSiteID(last, kind, id))
	if err != nil {
		log.Errorf("get %s %s %s %s  Error: %v", last, kind, id, err)
	}
	num, ok := res.(int)
	if !ok {
		log.Errorf("get %s %s %s Error: type is wrong", last, kind, id)
	}
	return num
}

//getBlogData 读取db中Blog data 未知db_key
func getBlogData(Userid string) int {
	return 0
}
