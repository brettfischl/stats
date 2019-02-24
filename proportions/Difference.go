package proportions

import (
  "math"

  "gonum.org/v1/gonum/stat/distuv"
)

/*
DifferenceOfProportions describes the values statistical likelihood that two SampleProportions are the same
*/
type DifferenceOfProportions struct {
  S1 SampleProportion
  S2 SampleProportion
  Difference float64
  Variance float64
  StandardDeviation float64
  Probability float64
}

// Test calculates the zscore of the difference between two SampleProportions
func (dp DifferenceOfProportions) Test() DifferenceOfProportions {
  dp.Difference = dp.S1.Mean - dp.S2.Mean
  dp.Variance = dp.S1.Variance + dp.S2.Variance
  dp.StandardDeviation = math.Sqrt(dp.Variance)


  normalDist := distuv.Normal{
    Mu: 0,
    Sigma: dp.StandardDeviation,
  }

  dp.Probability = normalDist.CDF(dp.Difference)

  return dp
}
