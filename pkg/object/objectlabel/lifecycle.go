package objectlabel

type DesiredLifecycle string
type InterimLifecycle string

const (
	LifecycleOnchain DesiredLifecycle = "onchain"
	LifecyclePending InterimLifecycle = "pending"
)

const (
	LifecycleAdjourn DesiredLifecycle = "adjourn"
	LifecycleBalance DesiredLifecycle = "balance"
	LifecycleDispute DesiredLifecycle = "dispute"
	LifecycleNullify DesiredLifecycle = "nullify"
	LifecyclePropose DesiredLifecycle = "propose"
	LifecycleResolve DesiredLifecycle = "resolve"
	LifecycleSettled DesiredLifecycle = "settled"
)
