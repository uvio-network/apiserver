package sample

import (
	"github.com/xh3b4sd/tracer"
)

var IndexNilError = &tracer.Error{
	Kind: "IndexNilError",
	Desc: "At least one of the provided indices was found to be nil, which should never happen. Calculating pagination cursors expects the zero value of any index to be an initialized *big.Int set to 0.",
	Docs: "https://github.com/uvio-network/contracts/blob/v0.5.0/contracts/Claims.sol#L1826-L1882",
}

var IndexOutOfRangeError = &tracer.Error{
	Kind: "IndexOutOfRangeError",
	Desc: "The provided cursor was found to be out of range, which should never happen. Calculating pagination cursors expects the provided paging pointer to be within either side of a market's index range.",
	Docs: "https://github.com/uvio-network/contracts/blob/v0.5.0/contracts/Claims.sol#L1826-L1882",
}

var IndexZeroError = &tracer.Error{
	Kind: "IndexZeroError",
	Desc: "Both of the provided total amounts at index 0 and 7 were found to be zero, which should never happen. Calculating pagination cursors expects at least one side of the market to contain at least one market participant to process.",
	Docs: "https://github.com/uvio-network/contracts/blob/v0.5.0/contracts/Claims.sol#L1826-L1882",
}
