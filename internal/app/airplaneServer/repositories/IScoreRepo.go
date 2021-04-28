package repositories

type IScoreRepo interface {
	Create(int64, int64) error
}
