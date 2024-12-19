// This file was generated by github.com/nelsam/hel.  Do not
// edit this code by hand unless you *really* know what you're
// doing.  Expect any changes made manually to be overwritten
// the next time hel regenerates this file.

package metrics_test

import (
	"github.com/cloudfoundry/dropsonde/metricbatcher"
)

type mockMetricBatcher struct {
	BatchIncrementCounterCalled chan bool
	BatchIncrementCounterInput  struct {
		Name chan string
	}
	BatchAddCounterCalled chan bool
	BatchAddCounterInput  struct {
		Name  chan string
		Delta chan uint64
	}
	CloseCalled        chan bool
	BatchCounterCalled chan bool
	BatchCounterInput  struct {
		Name chan string
	}
	BatchCounterOutput struct {
		Ret0 chan metricbatcher.BatchCounterChainer
	}
}

func newMockMetricBatcher() *mockMetricBatcher {
	m := &mockMetricBatcher{}
	m.BatchIncrementCounterCalled = make(chan bool, 100)
	m.BatchIncrementCounterInput.Name = make(chan string, 100)
	m.BatchAddCounterCalled = make(chan bool, 100)
	m.BatchAddCounterInput.Name = make(chan string, 100)
	m.BatchAddCounterInput.Delta = make(chan uint64, 100)
	m.CloseCalled = make(chan bool, 100)
	m.BatchCounterCalled = make(chan bool, 100)
	m.BatchCounterInput.Name = make(chan string, 100)
	m.BatchCounterOutput.Ret0 = make(chan metricbatcher.BatchCounterChainer, 100)
	return m
}
func (m *mockMetricBatcher) BatchIncrementCounter(name string) {
	m.BatchIncrementCounterCalled <- true
	m.BatchIncrementCounterInput.Name <- name
}
func (m *mockMetricBatcher) BatchAddCounter(name string, delta uint64) {
	m.BatchAddCounterCalled <- true
	m.BatchAddCounterInput.Name <- name
	m.BatchAddCounterInput.Delta <- delta
}
func (m *mockMetricBatcher) Close() {
	m.CloseCalled <- true
}
func (m *mockMetricBatcher) BatchCounter(name string) metricbatcher.BatchCounterChainer {
	m.BatchCounterCalled <- true
	m.BatchCounterInput.Name <- name
	return <-m.BatchCounterOutput.Ret0
}

type mockBatchCounterChainer struct {
	SetTagCalled chan bool
	SetTagInput  struct {
		Key, Value chan string
	}
	SetTagOutput struct {
		Ret0 chan metricbatcher.BatchCounterChainer
	}
	IncrementCalled chan bool
	AddCalled       chan bool
	AddInput        struct {
		Value chan uint64
	}
}

func newMockBatchCounterChainer() *mockBatchCounterChainer {
	m := &mockBatchCounterChainer{}
	m.SetTagCalled = make(chan bool, 100)
	m.SetTagInput.Key = make(chan string, 100)
	m.SetTagInput.Value = make(chan string, 100)
	m.SetTagOutput.Ret0 = make(chan metricbatcher.BatchCounterChainer, 100)
	m.IncrementCalled = make(chan bool, 100)
	m.AddCalled = make(chan bool, 100)
	m.AddInput.Value = make(chan uint64, 100)
	return m
}
func (m *mockBatchCounterChainer) SetTag(key, value string) metricbatcher.BatchCounterChainer {
	m.SetTagCalled <- true
	m.SetTagInput.Key <- key
	m.SetTagInput.Value <- value
	return <-m.SetTagOutput.Ret0
}
func (m *mockBatchCounterChainer) Increment() {
	m.IncrementCalled <- true
}
func (m *mockBatchCounterChainer) Add(value uint64) {
	m.AddCalled <- true
	m.AddInput.Value <- value
}
