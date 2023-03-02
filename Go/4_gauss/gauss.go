package main

import (
	. "fmt"
)

type Rational struct {
	numerator   int
	denominator int
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func simplify(r Rational) Rational {
	g := gcd(r.numerator, r.denominator)
	return Rational{r.numerator / g, r.denominator / g}
}

func (r Rational) String() string {
	return Sprintf("%d/%d", r.numerator, r.denominator)
}

func (r Rational) Add(s Rational) Rational {
	return simplify(Rational{
		r.numerator*s.denominator + s.numerator*r.denominator,
		r.denominator * s.denominator,
	})
}

func (r Rational) Sub(s Rational) Rational {
	return simplify(Rational{
		r.numerator*s.denominator - s.numerator*r.denominator,
		r.denominator * s.denominator,
	})
}

func (r Rational) Mul(s Rational) Rational {
	return simplify(Rational{
		r.numerator * s.numerator,
		r.denominator * s.denominator,
	})
}

func (r Rational) MulC(k int) Rational {
	return simplify(Rational{
		r.numerator * k,
		r.denominator,
	})
}


func (r Rational) Div(s Rational) Rational {
	return simplify(Rational{
		r.numerator * s.denominator,
		r.denominator * s.numerator,
	})
}

func ConvertIntToRational(x int) Rational {
	var r Rational
	r.numerator = x
	r.denominator = 1
	return r
}

// func ClearEmptyLine(N, Matrix [][]Rational) (int, Rational) {
// 	for i := 0; i < N; i++
// }

func main(){
	var N int
	Scanln(&N)
	var M = make([][]Rational, 0, N)
	for i := 0; i < N; i++ {
		line := make([]Rational, 0, N + 1)
		for j := 0; j < N + 1; j++ {
			x := 0
			Scan(&x)
			line = append(line, ConvertIntToRational(x))	
		} 
		M = append(M, line)
	}

	for i := 0; i < N - 1; i++ {
		for j := i+1; j < N; j++ {
			Print("Counting K: ")
			k := M[j][i].Div(M[i][i]) //Dopishi uslovie na nol
			Println("==> K = ",k)
			for q := i; q < N+1; q++ {
				M[j][q] = M[j][q].Sub( M[i][q].Mul(k) )
			}
		}  
	}
	for i := 0; i < N; i++ {
		Println(M[i])
	}

	Ans := make([]Rational, 0, N)
	warnExist := false 
	for i := N - 1; i > -1; i-- {
		Println("=> i = ", i, "ANS = ", Ans)
		rightVal := M[i][N]
		var leftVal Rational
		leftVal.numerator = 0
		leftVal.denominator = 1
		for j := i + 1; j < N; j++ {
			Println("===> i = ",i, "j = ", j)
			if len(Ans) + i == N - 1 {
				leftVal = Ans[N - 1 - j].Mul(M[i][j])
				rightVal = rightVal.Sub(leftVal)
			} else {
				warnExist = true
				Println("No solution")
			} 
			
		}
		rightVal = rightVal.Div(M[i][i]) 
		Ans = append(Ans, rightVal)
	}

	if !warnExist {
		for i := N - 1; i > -1; i-- {
			Println(Ans[i])
		}
	}
}