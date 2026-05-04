package domain

import "time"

var (
	UninitializedString = ""
	UninitializedTime, _  = time.Parse("2006-01-02 15:04:05", "2002-11-21 11:20:00")
	UnitializedRating int16 = 1000
	UninitializedRole = "user"
)
