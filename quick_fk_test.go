package main

import (
  //"fmt"
  "math"
)


type Polynomial struct {
  Coefficients []float64
}


func (p *Polynomial) Evaluate(x *Polynomial) float64 {
  result := 0.0

  for i := len(p.Coefficients) - 1; i >= 0; i-- {
    result = result*x.Coefficients[0] + p.Coefficients[i]
  }

  return result
}


func Residual(system []*Polynomial, degree int, polynomial *Polynomial) []float64 {
  residual := make([]float64, degree)


  for i, p := range system {
    residual[i] = p.Evaluate(polynomial)
  }

  return residual
}


func Correction(system []*Polynomial, degree int, polynomial *Polynomial, residual []float64) *Polynomial {
  correction := &Polynomial{make([]float64, degree)}
  for i := 0; i < degree; i++ {
    for j, p := range system {
      correction.Coefficients[i] += p.Coefficients[i] * residual[j]
    }
  }

  return correction
}


func Converged(correction *Polynomial, tolerance float64) bool {
  for _, c := range correction.Coefficients {
    if math.Abs(c) > tolerance {
      return false
    }
  }

  return true
}

func Solutions(polynomial *Polynomial) []float64 {
  return polynomial.Coefficients
}


func FK(system []*Polynomial, degree int, start *Polynomial, tolerance float64) []float64 {

  polynomial := start

  for {
    residual := Residual(system, degree, polynomial)
    correction := Correction(system, degree, polynomial, residual)
    if Converged(correction, tolerance) {
      return Solutions(polynomial)
    }
    polynomial = &Polynomial{make([]float64, degree)}
    for i := 0; i < degree; i++ {
      polynomial.Coefficients[i] = start.Coefficients[i] + correction.Coefficients[i]
    }
  }
}

func main() {
//system := []*Polynomial{
  //&Polynomial{[]float64{1, -3, 2}},
  //&Polynomial{[]float64{1, -1, -1}},
}
