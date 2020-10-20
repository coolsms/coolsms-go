package cash

import (
	"github.com/coolsms/coolsms-go/apirequest"
	"github.com/coolsms/coolsms-go/types"
)

// Cash struct
type Cash struct {
	Config map[string]string
}

// Balance get balance information
func (r *Cash) Balance() (types.Balance, error) {
	request := apirequest.NewAPIRequest()
	result := types.Balance{}
	setCustomConfigErr := request.SetCustomConfig(r.Config)
	if setCustomConfigErr != nil {
		return result, setCustomConfigErr
	}
	params := map[string]string{}
	err := request.GET("cash/v1/balance", params, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
