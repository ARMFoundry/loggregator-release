// This file was generated by github.com/nelsam/hel.  Do not
// edit this code by hand unless you *really* know what you're
// doing.  Expect any changes made manually to be overwritten
// the next time hel regenerates this file.

package grpcmanager_test

import (
	"doppler/grpcmanager"
	"plumbing"
	"time"

	"github.com/cloudfoundry/sonde-go/events"
	"golang.org/x/net/context"
)

type mockRegistrar struct {
	RegisterCalled chan bool
	RegisterInput  struct {
		ID         chan string
		IsFirehose chan bool
		Setter     chan grpcmanager.DataSetter
	}
	RegisterOutput struct {
		Ret0 chan func()
	}
}

func newMockRegistrar() *mockRegistrar {
	m := &mockRegistrar{}
	m.RegisterCalled = make(chan bool, 100)
	m.RegisterInput.ID = make(chan string, 100)
	m.RegisterInput.IsFirehose = make(chan bool, 100)
	m.RegisterInput.Setter = make(chan grpcmanager.DataSetter, 100)
	m.RegisterOutput.Ret0 = make(chan func(), 100)
	return m
}
func (m *mockRegistrar) Register(ID string, isFirehose bool, setter grpcmanager.DataSetter) func() {
	m.RegisterCalled <- true
	m.RegisterInput.ID <- ID
	m.RegisterInput.IsFirehose <- isFirehose
	m.RegisterInput.Setter <- setter
	return <-m.RegisterOutput.Ret0
}

type mockDataSetter struct {
	SetCalled chan bool
	SetInput  struct {
		Data chan []byte
	}
}

func newMockDataSetter() *mockDataSetter {
	m := &mockDataSetter{}
	m.SetCalled = make(chan bool, 100)
	m.SetInput.Data = make(chan []byte, 100)
	return m
}
func (m *mockDataSetter) Set(data []byte) {
	m.SetCalled <- true
	m.SetInput.Data <- data
}

type mockDataDumper struct {
	LatestContainerMetricsCalled chan bool
	LatestContainerMetricsInput  struct {
		AppID chan string
	}
	LatestContainerMetricsOutput struct {
		Ret0 chan []*events.Envelope
	}
	RecentLogsForCalled chan bool
	RecentLogsForInput  struct {
		AppID chan string
	}
	RecentLogsForOutput struct {
		Ret0 chan []*events.Envelope
	}
}

func newMockDataDumper() *mockDataDumper {
	m := &mockDataDumper{}
	m.LatestContainerMetricsCalled = make(chan bool, 100)
	m.LatestContainerMetricsInput.AppID = make(chan string, 100)
	m.LatestContainerMetricsOutput.Ret0 = make(chan []*events.Envelope, 100)
	m.RecentLogsForCalled = make(chan bool, 100)
	m.RecentLogsForInput.AppID = make(chan string, 100)
	m.RecentLogsForOutput.Ret0 = make(chan []*events.Envelope, 100)
	return m
}
func (m *mockDataDumper) LatestContainerMetrics(appID string) []*events.Envelope {
	m.LatestContainerMetricsCalled <- true
	m.LatestContainerMetricsInput.AppID <- appID
	return <-m.LatestContainerMetricsOutput.Ret0
}
func (m *mockDataDumper) RecentLogsFor(appID string) []*events.Envelope {
	m.RecentLogsForCalled <- true
	m.RecentLogsForInput.AppID <- appID
	return <-m.RecentLogsForOutput.Ret0
}

type mockSender struct {
	SendCalled chan bool
	SendInput  struct {
		Arg0 chan *plumbing.Response
	}
	SendOutput struct {
		Ret0 chan error
	}
	ContextCalled chan bool
	ContextOutput struct {
		Ret0 chan context.Context
	}
}

func newMockSender() *mockSender {
	m := &mockSender{}
	m.SendCalled = make(chan bool, 100)
	m.SendInput.Arg0 = make(chan *plumbing.Response, 100)
	m.SendOutput.Ret0 = make(chan error, 100)
	m.ContextCalled = make(chan bool, 100)
	m.ContextOutput.Ret0 = make(chan context.Context, 100)
	return m
}
func (m *mockSender) Send(arg0 *plumbing.Response) error {
	m.SendCalled <- true
	m.SendInput.Arg0 <- arg0
	return <-m.SendOutput.Ret0
}
func (m *mockSender) Context() context.Context {
	m.ContextCalled <- true
	return <-m.ContextOutput.Ret0
}

type mockContext struct {
	DeadlineCalled chan bool
	DeadlineOutput struct {
		Deadline chan time.Time
		Ok       chan bool
	}
	DoneCalled chan bool
	DoneOutput struct {
		Ret0 chan (<-chan struct{})
	}
	ErrCalled chan bool
	ErrOutput struct {
		Ret0 chan error
	}
	ValueCalled chan bool
	ValueInput  struct {
		Key chan interface{}
	}
	ValueOutput struct {
		Ret0 chan interface{}
	}
}

func newMockContext() *mockContext {
	m := &mockContext{}
	m.DeadlineCalled = make(chan bool, 100)
	m.DeadlineOutput.Deadline = make(chan time.Time, 100)
	m.DeadlineOutput.Ok = make(chan bool, 100)
	m.DoneCalled = make(chan bool, 100)
	m.DoneOutput.Ret0 = make(chan (<-chan struct{}), 100)
	m.ErrCalled = make(chan bool, 100)
	m.ErrOutput.Ret0 = make(chan error, 100)
	m.ValueCalled = make(chan bool, 100)
	m.ValueInput.Key = make(chan interface{}, 100)
	m.ValueOutput.Ret0 = make(chan interface{}, 100)
	return m
}
func (m *mockContext) Deadline() (deadline time.Time, ok bool) {
	m.DeadlineCalled <- true
	return <-m.DeadlineOutput.Deadline, <-m.DeadlineOutput.Ok
}
func (m *mockContext) Done() <-chan struct{} {
	m.DoneCalled <- true
	return <-m.DoneOutput.Ret0
}
func (m *mockContext) Err() error {
	m.ErrCalled <- true
	return <-m.ErrOutput.Ret0
}
func (m *mockContext) Value(key interface{}) interface{} {
	m.ValueCalled <- true
	m.ValueInput.Key <- key
	return <-m.ValueOutput.Ret0
}
