package main

import (
	"fmt"
	"go/constant"
	"os"

	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

func main() {
	// Load the package.
	pkgs, err := packages.Load(&packages.Config{
		Mode: packages.LoadAllSyntax,
	}, os.Args[1])
	if err != nil {
		panic(err)
	}
	pkg := pkgs[0]
	if len(pkg.Errors) != 0 {
		panic(pkg.Errors[0])
	}
	//ast.Print(pkg.Fset, pkg.Syntax[0])

	// Load the SSA.
	program, _ := ssautil.AllPackages(pkgs, 0)
	program.Build()
	for _, pkg := range program.AllPackages() {
		main := pkg.Members["main"].(*ssa.Function)
		//main.WriteTo(os.Stdout)
		instr := main.Blocks[0].Instrs[0].(*ssa.BinOp)
		fmt.Println("instruction:", instr)
		value := instr.Y.(*ssa.Const)
		fmt.Println("value:", value)
		n, exact := constant.Uint64Val(value.Value)
		fmt.Println("result:", n, exact)
	}
}
