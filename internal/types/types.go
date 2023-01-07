package types

type UserID int64

type DaysActiveSet map[int]struct{}

type DaysActiveByUserID map[UserID]DaysActiveSet

type StreakResults [][]int
