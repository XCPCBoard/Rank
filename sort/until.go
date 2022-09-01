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

//----------------------------------基本数学公式及存储结构体-------------------------------------------------

type KV struct {
	uerId  string
	rating int
}

type userRating []KV

func (ur userRating) Len() int { return len(ur) }

func (ur userRating) Swap(i, j int) { ur[i], ur[j] = ur[j], ur[i] }

func (ur userRating) Less(i, j int) bool {
	return ur[i].rating > ur[j].rating
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
