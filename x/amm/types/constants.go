package types

import (
	"cosmossdk.io/math"
)

const (
	OneShareExponent = 18

	BalancerGasFeeForSwap = 10_000
)

var (
	// OneShare represents the amount of subshares in a single pool share.
	OneShare = math.NewIntWithDecimal(1, OneShareExponent)

	// InitPoolSharesSupply is the amount of new shares to initialize a pool with.
	InitPoolSharesSupply = OneShare.MulRaw(100)

	// GuaranteedWeightPrecision Scaling factor for every weight. The pool weight is:
	// weight_in_MsgCreateBalancerPool * GuaranteedWeightPrecision
	//
	// This is done so that smooth weight changes have enough precision to actually be smooth.
	GuaranteedWeightPrecision int64 = 1 << 30

	oneHalf           = math.LegacyMustNewDecFromStr("0.5")
	twoDec            = math.LegacyMustNewDecFromStr("2")
	ln2               = math.LegacyMustNewDecFromStr("0.693147180559945309")
	inverseLn2        = math.LegacyMustNewDecFromStr("1.442695040888963407")
	euler             = math.LegacyMustNewDecFromStr("2.718281828459045235")
	powIterationLimit = int64(150_000)

	// PowPrecision Don't EVER change after initializing
	// TODO: Analyze choice here.
	powPrecision = math.LegacyMustNewDecFromStr("0.00000001")
)
