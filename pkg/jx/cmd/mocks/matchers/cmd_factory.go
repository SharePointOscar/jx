// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	cmd "github.com/jenkins-x/jx/pkg/jx/cmd"
	"github.com/petergtz/pegomock"
	"reflect"
)

func AnyCmdFactory() cmd.Factory {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(cmd.Factory))(nil)).Elem()))
	var nullValue cmd.Factory
	return nullValue
}

func EqCmdFactory(value cmd.Factory) cmd.Factory {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue cmd.Factory
	return nullValue
}
