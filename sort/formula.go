package sort

import (
	"XCPCBoard/utils/keys"
	"math"
	"rank/dao"
	"sort"
)

//userid 用户主键
//siteId 网站id

//--------------------------------基础rating--------------------------------------------

//getBaseRating 获取基础rating
func getBaseRating(userid string) int {
	// problem Sum part
	problemScore := 0.0
	for _, site := range OJ {
		num := getLastSiteKindIdNum(keys.LastAll, site, keys.PassAmountKey, getSiteId(site, userid))
		problemScore = problemScore + float64(num)
	}
	// oj rating part
	ratingScore := float64(getLastSiteKindIdNum(keys.LastAll, keys.CodeforcesKey, keys.RatingKey, userid))*0.5 +
		float64(getLastSiteKindIdNum(keys.LastAll, keys.AtCoderKey, keys.RatingKey, userid))*0.5
	//blog num part 未知blog_key及其存储方式
	blogScore := 0.0
	//四舍五入后强转
	baseRating := int(math.Ceil(problemScore*0.5 + ratingScore*0.4 + blogScore*0.1 - 0.5))
	return baseRating
}

//--------------------------------周期模块分数----------------------------------------
// last 周期长度

//getProblemScore 根据权重计算过题数+难度所得score
func getProblemScore(last, userid string) float64 {
	ProblemScore := 0
	//oj_Difficulty 遍历所有网站所有难度题目
	for _, site := range OJ {
		for i := 0; i < len(Difficulty); i++ {
			ProblemScore = ProblemScore + DifficultyScore[i]*
				getLastSiteKindDifficultyIdNu(last, site, keys.PassAmountKey, Difficulty[i], getSiteId(site, userid))
		}
	}
	return float64(ProblemScore)
}

//getCodeforcesRatingScore cfRating 换算分值
func getCodeforcesRatingScore(last, siteId string) int {
	lastCfRating := getLastSiteKindIdNum(last, keys.CodeforcesKey, keys.RatingKey, siteId)
	//现在的rating值 key 不确定
	nowCfRating := getLastSiteKindIdNum(keys.LastAll, keys.CodeforcesKey, keys.RatingKey, siteId)
	//周期 rating 改变量
	d := nowCfRating - lastCfRating
	if d <= 0 {
		return 0
	}
	if lastCfRating <= 600 {
		return 1
	}
	cfRatingScore := int(math.Ceil(float64(lastCfRating)/400.0 + float64(lastCfRating*d)/20000.0 - 0.5))
	return cfRatingScore
}

//getAtcoderRatingScore atcRating 换算分值
func getAtcoderRatingScore(last, siteId string) int {
	lastAtcRating := getLastSiteKindIdNum(last, keys.AtCoderKey, keys.RatingKey, siteId)
	//现在的rating值 key 不确定
	nowAtcRating := getLastSiteKindIdNum(keys.LastAll, keys.AtCoderKey, keys.RatingKey, siteId)
	//周期 rating 改变量
	d := nowAtcRating - lastAtcRating
	if d <= 0 {
		return 0
	}
	if lastAtcRating <= 400 {
		return 1
	}
	p := 0.0
	if lastAtcRating >= 1000 {
		p = 20.0
	} else {
		p = 50.0
	}
	atcRatingScore := int(math.Ceil(float64(lastAtcRating)*(1.0+float64(d)/p) - 0.5))
	return atcRatingScore
}

//getAllRatingScore Rating Score part
func getAllRatingScore(last, userid string) float64 {
	return float64(getAtcoderRatingScore(last, getSiteId(keys.AtCoderKey, userid)) +
		getCodeforcesRatingScore(last, getSiteId(keys.CodeforcesKey, userid)))
}

//getBlogScore blog分值 未知key
func getBlogScore(last, userid string) float64 {
	return 0
}

//getAttendanceScore 考勤分值
func getAttendanceScore(last, userid string) float64 {
	return float64(getLastKindIDData(last, AttendanceKey, userid))
}

//--------------------------------周期表现rating计算----------------------------------------

// getSaScore 根据各项权重计算周期内实际的获胜概率
func getSaScore(last, selfId, rivalId string) float64 {
	//获取 1V1 所需数据并计算 S_A
	blogSelf := getBlogScore(last, selfId)
	blogRival := getBlogScore(last, rivalId)
	problemSelf := getProblemScore(last, selfId)
	problemRival := getProblemScore(last, rivalId)
	ratingSelf := getAllRatingScore(last, selfId)
	ratingRival := getAllRatingScore(last, rivalId)
	attendanceSelf := getAttendanceScore(last, selfId)
	attendanceRival := getAttendanceScore(last, rivalId)
	//公式
	Sa := (problemSelf*0.4)/(problemSelf+problemRival) + (ratingSelf*0.3)/(ratingSelf+ratingRival) +
		(blogSelf*0.2)/(blogSelf+blogRival) + (attendanceSelf*0.1)/(attendanceRival+attendanceSelf)
	return Sa
}

// countOle OLE公式计算预测胜率
func countOle(last, user, rival string) float64 {
	//用户上一个周期 gxU_rating
	RSelf := getLastKindIDData(last, GxuRatingKey, user)
	//对手上一个周期 gxU_rating
	RRival := getLastKindIDData(last, GxuRatingKey, rival)
	res := 1.0 / (1.0 + math.Pow(10.0, (float64(RRival)-float64(RSelf))/400.0))
	return res
}

//countPa 计算几何平均数的差值
func countPa(upSum float64, downSum float64, upNum int, downNum int) float64 {
	//几何平均数：连乘后开n次方
	Pa := math.Pow(upSum, 1.0/float64(upNum)) - math.Pow(downSum, 1.0/float64(downNum))
	return Pa
}

func countAllAddRating(last, userid string) int {
	upSum := 1.0   //(S_Ai-E_Ai)大于0的乘积
	upNum := 0     //计数，用于求几何平均数
	downSum := 1.0 //(S_Ai-E_Ai)小于0的乘积
	downNum := 0   //计数，用于求几何平均数
	//用户上一个周期 gxU_rating
	//RSelf := getLastKindIDData(last, GxuRatingKey, userid)
	// 计算 1Vn 时的P_A，也是1V1的几何平均数
	for _, rival := range usersIDTable {
		if userid == rival {
			continue
		}
		ESelf := countOle(last, userid, rival)
		SSelf := getSaScore(last, userid, rival)
		if SSelf >= ESelf {
			upSum = upSum * (SSelf - ESelf)
			upNum = upNum + 1
		} else {
			downSum = downSum * (ESelf - SSelf)
			downNum = downNum + 1
		}
	}
	pa := countPa(upSum, downSum, upNum, downNum)
	gxuRatingNew := int(math.Ceil(float64(K)*pa) - 0.5)
	return gxuRatingNew
}

//--------------------------------rating数据修正----------------------------------------

//countAdjust 调整量公式
func countAdjust(RSum int) int {
	adjust := int((-1.0 - float64(K*RSum)) / float64(len(usersIDTable)))
	return adjust
}

// firstCorrectRating 第一次修正
func firstCorrectRating(last string) userRating {
	usersAddRating := make([]KV, 0)
	RSum := 0
	for _, user := range usersIDTable {
		addRating := countAllAddRating(last, user)
		RSum = RSum + addRating

		usersAddRating = append(usersAddRating, KV{user, addRating})
	}
	adjust := countAdjust(RSum)
	for i, _ := range usersAddRating {
		usersAddRating[i].rating = usersAddRating[i].rating + adjust
	}
	return usersAddRating
}

// secondCorrectRating 第二次修正
func secondCorrectRating(last string) userRating {
	usersAddRating := firstCorrectRating(last)
	sort.Sort(userRating(usersAddRating))
	L := len(usersAddRating)
	n := Min(L, 4*int(math.Sqrt(float64(L))))
	RSum := 0
	for i := 0; i < n; i++ {
		RSum = RSum + usersAddRating[i].rating
	}
	adjust := Min(Max(countAdjust(RSum), -10), 0)
	for i := 0; i < n; i++ {
		usersAddRating[i].rating = usersAddRating[i].rating + adjust
	}
	return usersAddRating
}

//-----------------------------------------------------------------------------

// Flush 周期更新Gxu_rating
func Flush(last string){
	usersAddRating:=secondCorrectRating(last)
	for _,user:=range usersAddRating{
		rating:=getLastKindIDData(last,GxuRatingKey,user.uerId)
		dao.UpdateRedis(BuildKeyWithLastSiteID(last,GxuRatingKey,user.uerId),rating+user.rating)
	}
}
