package transform

import (
	"errors"
	"strings"

	"github.com/prometheus/prometheus/pkg/labels"
	pp "github.com/prometheus/prometheus/promql/parser"
)

func GetLabelMatchers(flags []string) (map[string]string, error) {
	inj := map[string]string{}
	for _, matcher := range flags {
		parts := strings.Split(matcher, "=")
		if len(parts) != 2 {
			return nil, errors.New("malformed label injector")
		}
		inj[parts[0]] = parts[1]
	}
	return inj, nil
}

func Transform(arg string, matchers *map[string]string) (string, error) {

	exp, err := pp.ParseExpr(arg)

	if err != nil {
		return "", err
	}

	if e, ok := exp.(*pp.VectorSelector); ok {
		InjectLabelMatcher(e, matchers)
	}

	TraverseNode(exp, matchers)
	return exp.String(), nil
}

func TraverseNode(exp pp.Node, matchers *map[string]string) {

	for _, c := range pp.Children(exp) {

		if e, ok := c.(*pp.VectorSelector); ok {
			InjectLabelMatcher(e, matchers)
		}
		TraverseNode(c, matchers)
	}
}

func InjectLabelMatcher(e *pp.VectorSelector, matchers *map[string]string) {
	for key, val := range *matchers {
		var found = false
		for _, existing := range e.LabelMatchers {
			if existing.Name == key {
				found = true
				break
			}
		}
		if found {
			continue
		}
		e.LabelMatchers = append(
			e.LabelMatchers,
			&labels.Matcher{
				Type:  labels.MatchEqual,
				Name:  key,
				Value: val,
			},
		)
	}
}
