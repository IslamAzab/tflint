package main

import (
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/wata727/tflint/rules"
	"github.com/wata727/tflint/tflint"
)

type CustomRule struct{}

func (r *CustomRule) Name() string {
	return "my_custom_rule"
}

func (r *CustomRule) Enabled() bool {
	return true
}

func (r *CustomRule) Severity() string {
	return tflint.NOTICE
}

func (r *CustomRule) Link() string {
	return ""
}

func (r *CustomRule) Check(runner *tflint.Runner) error {
	return runner.WalkResourceAttributes("aws_instance", "instance_type", func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			runner.EmitIssue(
				r,
				fmt.Sprintf("instance_type is %s", val),
				attribute.Expr.Range(),
			)
			return nil
		})
	})
}

func NewRules() []rules.Rule {
	return []rules.Rule{&CustomRule{}}
}
