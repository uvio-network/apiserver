package objectlabel

type DesiredLifecycle string
type InterimLifecycle string

const (
	LifecycleExpired InterimLifecycle = "expired"
	LifecycleOnchain DesiredLifecycle = "onchain"
	LifecyclePending InterimLifecycle = "pending"
)

const (
	LifecycleBalance DesiredLifecycle = "balance"
	LifecycleDispute DesiredLifecycle = "dispute"
	LifecyclePropose DesiredLifecycle = "propose"
	LifecycleResolve DesiredLifecycle = "resolve"
)
