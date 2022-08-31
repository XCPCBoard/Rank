package sort

import (
	"XCPCBoard/utils/keys"
	"fmt"
)

const (
	GxuRatingKey  = "gxu_rating"
	AttendanceKey = "Attendance"

	K = 32 //公式中的K因子
)

var (
	OJ              = [5]string{keys.NowcoderKey, keys.AtCoderKey, keys.LuoguKey, keys.VJudgeKey, keys.CodeforcesKey}
	Difficulty      = [5]string{keys.DifficultyEasy, keys.DifficultyBasic, keys.DifficultyAdvanced, keys.DifficultyHard, keys.DifficultyUnknown}
	DifficultyScore = [5]int{1, 2, 3, 4, 2} //难度所对应分值
	usersIDTable    = [...]string{}         //从mysql获取所有用户主键，未知key
)

// 新增keys工具类

// BuildKeyWithLastSiteID 时效_类型_用户id
func BuildKeyWithLastSiteID(last, site, id string) string {
	return fmt.Sprintf("%v %v_%v", last, site, id)
}
