// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	versioned1 "k8s.io/metrics/pkg/client/clientset/versioned"
)

func AnyPtrToVersioned1Clientset() *versioned1.Clientset {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*versioned1.Clientset))(nil)).Elem()))
	var nullValue *versioned1.Clientset
	return nullValue
}

func EqPtrToVersioned1Clientset(value *versioned1.Clientset) *versioned1.Clientset {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *versioned1.Clientset
	return nullValue
}
