package steamid

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var accountID uint32
var steamID uint64
var steam2ID string
var steam3ID string

func init() {
	accountID = 103803916
	steamID = 76561198064069644
	steam2ID = "STEAM_1:1:51901958"
	steam3ID = "[U:103803916]"
}

func TestNewId(t *testing.T) {
	fmt.Printf("%v %v %v %v", accountID, steamID, steam2ID, steam3ID)
	steamid := NewIdAdv(accountID, 1, 1, 1)

	assert.Equal(t, accountID, steamid.GetAccountId())
	assert.Equal(t, steamID, steamid.ToUint64())

	steamid2 := NewAccountID(accountID)

	assert.Equal(t, accountID, steamid2.GetAccountId())
	assert.Equal(t, steamID, steamid2.ToUint64())
}
