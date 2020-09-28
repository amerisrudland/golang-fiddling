package iteration

const repeatBase = 5

func Repeat(c string, times int) string {
	var repeated string
	var repeatCount int
	if times > 0 {
		repeatCount = times
	} else {
		repeatCount = repeatBase
	}
	for i := 0; i < repeatCount; i++ {
		repeated += c
	}
	return repeated
}
