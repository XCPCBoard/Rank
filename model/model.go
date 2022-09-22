package model

const (
	userName = "LQY"
	userId   = "123"
)

type KV struct {
	UerId  string
	Rating float64
}

type userRating []KV

func (ur userRating) Len() int { return len(ur) }

func (ur userRating) Swap(i, j int) { ur[i], ur[j] = ur[j], ur[i] }

func (ur userRating) Less(i, j int) bool {
	return ur[i].Rating > ur[j].Rating
}

func Min(x, y float64) float64 {
	if x < y {
		return x
	}
	return y
}

func Max(x, y float64) float64 {
	if x > y {
		return x
	}
	return y
}
