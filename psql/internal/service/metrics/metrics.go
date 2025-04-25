package metrics

type IMetrics interface {
	IncCounter(string)
}
