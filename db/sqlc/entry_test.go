package db

import (
	"context"
	"database/sql"
	"github.com/pouyasadri/go-simple-bank/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func (q *Queries) CreateRandomEntry(t *testing.T) Entry {
	arg := CreateEntryParams{
		AccountID: 1,
		Amount:    util.RandomMoney(),
	}
	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)
	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreateAt)
	return entry

}
func TestQueries_CreateEntry(t *testing.T) {
	testQueries.CreateRandomEntry(t)
}

func TestQueries_GetEntry(t *testing.T) {
	entry1 := testQueries.CreateRandomEntry(t)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreateAt, entry2.CreateAt, time.Second)
}

func TestQueries_GetEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		testQueries.CreateRandomEntry(t)
	}
	arg := GetEntriesParams{
		AccountID: 1,
		Limit:     5,
		Offset:    5,
	}
	entries, err := testQueries.GetEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)
	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}

func TestQueries_UpdateEntry(t *testing.T) {
	entry1 := testQueries.CreateRandomEntry(t)
	arg := UpdateEntryParams{
		ID:     entry1.ID,
		Amount: util.RandomMoney(),
	}
	entry2, err := testQueries.UpdateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)
	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, arg.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreateAt, entry2.CreateAt, time.Second)
}

func TestQueries_DeleteEntry(t *testing.T) {
	entry1 := testQueries.CreateRandomEntry(t)
	err := testQueries.DeleteEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, entry2)
}
