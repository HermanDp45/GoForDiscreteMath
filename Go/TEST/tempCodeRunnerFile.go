
func (r Rational) String() string {
	return fmt.Sprintf("%d/%d", r.numerator, r.denominator)
}