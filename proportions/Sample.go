package proportions

import (
  "math"

  "gonum.org/v1/gonum/stat/distuv"
)

/*
SampleProportion describes the values needed to calculate the statistics describing a difference in proportions
*/
type SampleProportion struct {
  Trials float64
  Successes float64
  Compare float64
  Mean float64
  Variance float64
  StandardDeviation float64
  BaseZScores []ZScore
  CompareZScore ZScore
}

// ZScore describes probability and mean at a given Z score value on a normal distribution
type ZScore struct {
  Probability float64
  Value float64
  Z float64
}

var defaultZScores = []float64{-3, -2, -1, 0, 1, 2, 3}

/*
NewSampleProportion initializes a SampleProportion
*/
func NewSampleProportion(
  trials float64,
  successes float64,
  compare float64,
) SampleProportion {

  sampleProportion := SampleProportion{
    Trials: trials,
    Successes: successes,
    Compare: compare,
  }

  sampleProportion.trial()

  return sampleProportion
}

// Perform calculates summary stats for a Binomial Experiment
func (p *SampleProportion) trial() {
  p.Mean = (p.Successes / p.Trials)
  p.Variance = (p.Mean * (1.0 - p.Mean)) / p.Trials
  p.StandardDeviation = math.Sqrt(p.Variance)
}

// Zscores calculates all Zscores for each default Zscore plus the comparison value if it's given
func (p *SampleProportion) Zscores() {
  normalDist := distuv.Normal{
    Mu: p.Mean,
    Sigma: p.StandardDeviation,
  }

  for _, score := range(defaultZScores) {
    value := p.Mean + (p.StandardDeviation * score)
    zScore := calculateZScore(p.StandardDeviation, value, score, normalDist)
    p.BaseZScores = append(p.BaseZScores, zScore)
  }

  if (p.Compare > 0) {
    compareValue := p.Compare / p.Trials
    compareZ := (compareValue - p.Mean) / p.StandardDeviation
    compareZScore := calculateZScore(p.StandardDeviation, compareValue, compareZ, normalDist)

    p.CompareZScore = compareZScore
  }
}

func calculateZScore(stDev float64,
  value float64,
  score float64,
  normalDist distuv.Normal,
) ZScore {

  zScore := ZScore{
    Value: value,
    Z: score,
    Probability: normalDist.CDF(value),
  }

  return zScore
}
