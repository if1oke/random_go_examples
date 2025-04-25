package metrics

import "fmt"

type CLMetrics struct {
	transferSuccessCnt int
	transferFailedCnt  int
	reserveSuccessCnt  int
	reserveFailedCnt   int
}

func NewCLMetrics() *CLMetrics {
	return &CLMetrics{
		transferSuccessCnt: 0,
		transferFailedCnt:  0,
		reserveSuccessCnt:  0,
		reserveFailedCnt:   0,
	}
}

func (m *CLMetrics) IncCounter(label string) {
	switch label {
	case "transferSuccess":
		m.transferSuccessCnt++
		fmt.Printf("metrics: %s + 1, total %d\n", label, m.transferSuccessCnt)
	case "transferFailed":
		m.transferFailedCnt++
		fmt.Printf("metrics: %s + 1, total %d\n ", label, m.transferFailedCnt)
	case "reserveSuccess":
		m.reserveSuccessCnt++
		fmt.Printf("metrics: %s + 1, total %d\n ", label, m.reserveSuccessCnt)
	case "reserveFailed":
		m.reserveFailedCnt++
		fmt.Printf("metrics: %s + 1, total %d\n ", label, m.reserveFailedCnt)
	}
}
