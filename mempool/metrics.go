package mempool

import (
	"sync"

	"github.com/cometbft/cometbft/types"
	"github.com/go-kit/kit/metrics"
)

const (
	// MetricsSubsystem is a subsystem shared by all metrics exposed by this
	// package.
	MetricsSubsystem = "mempool"
)

//go:generate go run ../scripts/metricsgen -struct=Metrics

// Metrics contains metrics exposed by this package.
// see MetricsProvider for descriptions.
type Metrics struct {
	// Number of uncommitted transactions in the mempool.
	Size metrics.Gauge

	// Histogram of transaction sizes in bytes.
	TxSizeBytes metrics.Histogram `metrics_buckettype:"exp" metrics_bucketsizes:"1,3,7"`

	// Number of failed transactions.
	FailedTxs metrics.Counter

	// RejectedTxs defines the number of rejected transactions. These are
	// transactions that passed CheckTx but failed to make it into the mempool
	// due to resource limits, e.g. mempool is full and no lower priority
	// transactions exist in the mempool.
	//metrics:Number of rejected transactions.
	RejectedTxs metrics.Counter

	// EvictedTxs defines the number of evicted transactions. These are valid
	// transactions that passed CheckTx and existed in the mempool but were later
	// evicted to make room for higher priority valid transactions that passed
	// CheckTx.
	//metrics:Number of evicted transactions.
	EvictedTxs metrics.Counter

	// Number of times transactions are rechecked in the mempool.
	RecheckTimes metrics.Counter

	// Histogram of times a transaction was received.
	TimesTxsWereReceived metrics.Histogram `metrics_buckettype:"exp" metrics_bucketsizes:"1,2,5"`

	// For keeping track of the number of times each transaction in the mempool was received.
	timesTxWasReceived sync.Map

	// Number of times transactions were received more than once.
	TxsReceivedMoreThanOnce metrics.Counter
}

func (m *Metrics) countOneTimeTxWasReceived(tx types.TxKey) {
	value, _ := m.timesTxWasReceived.LoadOrStore(tx, uint64(0))
	m.timesTxWasReceived.Store(tx, value.(uint64)+1)
}

func (m *Metrics) resetTimesTxWasReceived(tx types.TxKey) {
	m.timesTxWasReceived.Delete(tx)
}

func (m *Metrics) observeTimesTxsWereReceived() {
	m.timesTxWasReceived.Range(func(_, value interface{}) bool {
		m.TimesTxsWereReceived.Observe(float64(value.(uint64)))
		return true
	})
}
