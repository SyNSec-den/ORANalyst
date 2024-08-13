package exectest

//func Z3TestFunc() {
//	ctx := z3.NewContext(nil)
//
//	//x, y := ctx.BoolConst("x"), ctx.BoolConst("y")
//
//	p, _ := ctx.IntConst("p"), ctx.IntConst("q")
//
//	ival, _, _ := ctx.FromInt(5, ctx.IntSort()).(z3.Int).AsInt64()[0]
//
//	r := p.GE(ival)
//	s := p.LT(ctx.FromInt(5, ctx.IntSort()).(z3.Int))
//	t := r.And(s)
//
//	v := ctx.Simplify(t, z3.NewSimplifyConfig(ctx))
//	fmt.Printf("%s\n", v.String())
//
//	//s := z3.NewSolver(ctx)
//	//s.Assert(conjecture)
//	//sat, err := s.Check()
//	//if err != nil {
//	//	fmt.Printf("failed to compute satisfiability: %s", err)
//	//} else if sat {
//	//	fmt.Printf("disproved De Morgan's law\n")
//	//} else {
//	//	fmt.Printf("Good\n")
//	//}
//}
