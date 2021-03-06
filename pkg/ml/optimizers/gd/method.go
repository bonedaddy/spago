// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gd

import (
	"github.com/nlpodyssey/spago/pkg/mat"
	"github.com/nlpodyssey/spago/pkg/ml/nn"
)

const (
	None int = iota
	SGD
	AdaGrad
	Adam
	RAdam
	RMSProp
)

// Empty interface implemented by the configuration structures of AdaGrad, Adam, RMSProp and SGD.
type MethodConfig interface{}

// Optimization Method
type Method interface {
	// Label can be None, SGD, AdaGrad, Adam, RMSProp
	Label() int
	// Delta returns the difference between the current params and where the method wants it to be.
	Delta(param *nn.Param) mat.Matrix
	// NewSupport returns a new support structure with the given dimensions
	NewSupport(r, c int) *nn.Payload
}

func GetOrSetPayload(param *nn.Param, m Method) *nn.Payload {
	payload := param.Payload()
	switch {
	case payload == nil:
		payload := m.NewSupport(param.Value().Dims())
		param.SetPayload(payload)
		return payload
	case payload.Label == None:
		payload := m.NewSupport(param.Value().Dims())
		param.SetPayload(payload)
		return payload
	case payload.Label == m.Label():
		return payload
	default:
		panic("gd: support structure non compatible with the optimization method")
	}
}
