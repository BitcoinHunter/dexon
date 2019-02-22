// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package params

import (
	"encoding/json"
	"math/big"

	"github.com/dexon-foundation/dexon/common"
	"github.com/dexon-foundation/dexon/common/math"
)

var _ = (*dexconConfigSpecMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (d DexconConfig) MarshalJSON() ([]byte, error) {
	type DexconConfig struct {
		GenesisCRSText    string                  `json:"genesisCRSText"`
		Owner             common.Address          `json:"owner"`
		MinStake          *math.HexOrDecimal256   `json:"minStake"`
		LockupPeriod      uint64                  `json:"lockupPeriod"`
		MiningVelocity    float32                 `json:"miningVelocity"`
		NextHalvingSupply *math.HexOrDecimal256   `json:"nextHalvingSupply"`
		LastHalvedAmount  *math.HexOrDecimal256   `json:"lastHalvedAmount"`
		BlockGasLimit     uint64                  `json:"blockGasLimit"`
		LambdaBA          uint64                  `json:"lambdaBA"`
		LambdaDKG         uint64                  `json:"lambdaDKG"`
		NotarySetSize     uint32                  `json:"notarySetSize"`
		DKGSetSize        uint32                  `json:"dkgSetSize"`
		RoundLength       uint64                  `json:"roundLength"`
		MinBlockInterval  uint64                  `json:"minBlockInterval"`
		FineValues        []*math.HexOrDecimal256 `json:"fineValues"`
	}
	var enc DexconConfig
	enc.GenesisCRSText = d.GenesisCRSText
	enc.Owner = d.Owner
	enc.MinStake = (*math.HexOrDecimal256)(d.MinStake)
	enc.LockupPeriod = d.LockupPeriod
	enc.MiningVelocity = d.MiningVelocity
	enc.NextHalvingSupply = (*math.HexOrDecimal256)(d.NextHalvingSupply)
	enc.LastHalvedAmount = (*math.HexOrDecimal256)(d.LastHalvedAmount)
	enc.BlockGasLimit = d.BlockGasLimit
	enc.LambdaBA = d.LambdaBA
	enc.LambdaDKG = d.LambdaDKG
	enc.NotarySetSize = d.NotarySetSize
	enc.DKGSetSize = d.DKGSetSize
	enc.RoundLength = d.RoundLength
	enc.MinBlockInterval = d.MinBlockInterval
	if d.FineValues != nil {
		enc.FineValues = make([]*math.HexOrDecimal256, len(d.FineValues))
		for k, v := range d.FineValues {
			enc.FineValues[k] = (*math.HexOrDecimal256)(v)
		}
	}
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (d *DexconConfig) UnmarshalJSON(input []byte) error {
	type DexconConfig struct {
		GenesisCRSText    *string                 `json:"genesisCRSText"`
		Owner             *common.Address         `json:"owner"`
		MinStake          *math.HexOrDecimal256   `json:"minStake"`
		LockupPeriod      *uint64                 `json:"lockupPeriod"`
		MiningVelocity    *float32                `json:"miningVelocity"`
		NextHalvingSupply *math.HexOrDecimal256   `json:"nextHalvingSupply"`
		LastHalvedAmount  *math.HexOrDecimal256   `json:"lastHalvedAmount"`
		BlockGasLimit     *uint64                 `json:"blockGasLimit"`
		LambdaBA          *uint64                 `json:"lambdaBA"`
		LambdaDKG         *uint64                 `json:"lambdaDKG"`
		NotarySetSize     *uint32                 `json:"notarySetSize"`
		DKGSetSize        *uint32                 `json:"dkgSetSize"`
		RoundLength       *uint64                 `json:"roundLength"`
		MinBlockInterval  *uint64                 `json:"minBlockInterval"`
		FineValues        []*math.HexOrDecimal256 `json:"fineValues"`
	}
	var dec DexconConfig
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.GenesisCRSText != nil {
		d.GenesisCRSText = *dec.GenesisCRSText
	}
	if dec.Owner != nil {
		d.Owner = *dec.Owner
	}
	if dec.MinStake != nil {
		d.MinStake = (*big.Int)(dec.MinStake)
	}
	if dec.LockupPeriod != nil {
		d.LockupPeriod = *dec.LockupPeriod
	}
	if dec.MiningVelocity != nil {
		d.MiningVelocity = *dec.MiningVelocity
	}
	if dec.NextHalvingSupply != nil {
		d.NextHalvingSupply = (*big.Int)(dec.NextHalvingSupply)
	}
	if dec.LastHalvedAmount != nil {
		d.LastHalvedAmount = (*big.Int)(dec.LastHalvedAmount)
	}
	if dec.BlockGasLimit != nil {
		d.BlockGasLimit = *dec.BlockGasLimit
	}
	if dec.LambdaBA != nil {
		d.LambdaBA = *dec.LambdaBA
	}
	if dec.LambdaDKG != nil {
		d.LambdaDKG = *dec.LambdaDKG
	}
	if dec.NotarySetSize != nil {
		d.NotarySetSize = *dec.NotarySetSize
	}
	if dec.DKGSetSize != nil {
		d.DKGSetSize = *dec.DKGSetSize
	}
	if dec.RoundLength != nil {
		d.RoundLength = *dec.RoundLength
	}
	if dec.MinBlockInterval != nil {
		d.MinBlockInterval = *dec.MinBlockInterval
	}
	if dec.FineValues != nil {
		d.FineValues = make([]*big.Int, len(dec.FineValues))
		for k, v := range dec.FineValues {
			d.FineValues[k] = (*big.Int)(v)
		}
	}
	return nil
}
