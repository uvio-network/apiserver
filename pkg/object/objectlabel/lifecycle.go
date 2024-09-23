package objectlabel

type DesiredLifecycle string
type InterimLifecycle string

const (
	LifecycleOnchain DesiredLifecycle = "onchain"
	LifecyclePending InterimLifecycle = "pending"
)

const (
	LifecycleBalance DesiredLifecycle = "balance"
	LifecycleDispute DesiredLifecycle = "dispute"
	LifecyclePropose DesiredLifecycle = "propose"
	LifecycleResolve DesiredLifecycle = "resolve"
	LifecycleSettled DesiredLifecycle = "settled"
)
