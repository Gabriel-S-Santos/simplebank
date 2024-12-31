package db

import (
	"context"
	"testing"

	"github.com/Gabriel-S-Santos/simplebank/tree/master/util"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	util.SeedRandom() //garante que os dados serão gerados diferentes do teste anterior
	
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomAmmount(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQuerries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}
