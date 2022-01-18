package vtasample

import (
	"github.com/davecgh/go-spew/spew"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
	"golang.org/x/tools/go/callgraph/cha"
	"golang.org/x/tools/go/callgraph/vta"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

var Analyzer = &analysis.Analyzer{
	Name: "vtasample",
	Doc:  "vtasample",
	Run:  run,
	Requires: []*analysis.Analyzer{
		buildssa.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	s := pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA)
	allfns := ssautil.AllFunctions(s.Pkg.Prog)
	for k, v := range allfns {
		cg := vta.CallGraph(map[*ssa.Function]bool{
			k: v,
		}, cha.CallGraph(s.Pkg.Prog))
		pass.Reportf(k.Pos(), "%s", spew.Sdump(cg))
	}
	return nil, nil
}

