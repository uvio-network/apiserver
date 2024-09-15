package objectlabel

type DesiredLifecycle string
type InterimLifecycle string

const (
	LifecyclePending InterimLifecycle = "pending"
)

const (
	LifecycleAdjourn DesiredLifecycle = "adjourn"
	LifecycleBalance DesiredLifecycle = "balance"
	LifecycleDispute DesiredLifecycle = "dispute"
	LifecycleNullify DesiredLifecycle = "nullify"
	LifecycleOnchain DesiredLifecycle = "onchain"
	LifecyclePropose DesiredLifecycle = "propose"
	LifecycleResolve DesiredLifecycle = "resolve"
)
