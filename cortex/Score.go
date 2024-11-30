package cortex

import "errors"

var weights = []float64{
	-0.009542620, 0.051070686, 0.015278586, 0.001600832, 0.001284823,
	-0.008918919, 0.026382536, -0.004463617, 0.001020790, 0.001014553,
	0.050407484, 0.054981289, 0.040045738, 0.009324324, -0.001222453,
	-0.006530146, 0.070224532, 0.007835759, 0.010378378, 0.000694387,
	0.031748441, 0.163619543, 0.025359667, 0.005442827, -0.001164241,
}

const (
	ErrInvalidData = "invalid length of data"
)

func calculateScore(data []float64) (float64, error) {
	if len(data) != len(weights) {
		return 0, errors.New(ErrInvalidData)
	}

	var score float64
	for i := 0; i < len(data); i++ {
		score += data[i] * weights[i]
	}
	return score, nil
}
