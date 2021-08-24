package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	var arg CreateAccountParams

	faker.FakeData(&arg)

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	randomAccount := createRandomAccount(t)

	account, err := testQueries.GetAccount(context.Background(), randomAccount.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, account.Owner, randomAccount.Owner)
	require.Equal(t, account.Balance, randomAccount.Balance)
	require.Equal(t, account.Currency, randomAccount.Currency)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}

func TestUpdateAccount(t *testing.T) {
	randomAccount := createRandomAccount(t)

	args := UpdateAccountParams{
		Balance: randomAccount.Balance + 1,
		ID:      randomAccount.ID,
	}

	account, err := testQueries.UpdateAccount(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, account.Owner, randomAccount.Owner)
	require.NotEqual(t, account.Balance, randomAccount.Balance)
	require.Equal(t, account.Balance, randomAccount.Balance+1)
	require.Equal(t, account.Currency, randomAccount.Currency)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}

func TestDeleteAccount(t *testing.T) {
	randomAccount := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), randomAccount.ID)

	require.NoError(t, err)

	deletedAccount, err := testQueries.GetAccount(context.Background(), randomAccount.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, deletedAccount)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	args := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), args)

	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
