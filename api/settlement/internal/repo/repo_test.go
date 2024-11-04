package repo_test

import (
	"context"
	"github.com/failuretoload/datamonster/settlement/domain"
	"github.com/failuretoload/datamonster/settlement/internal/repo"
	"github.com/failuretoload/datamonster/settlement/internal/sql"
	"github.com/failuretoload/datamonster/store/sqlite"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSettlementRepo_Integration(t *testing.T) {
	// Setup test database
	dao, err := sqlite.NewDAO()
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, dao.Close())
	})
	db := repo.New(dao, sql.SQLite())
	ctx := context.Background()

	t.Run("insert and get settlement", func(t *testing.T) {
		settlement := domain.Settlement{
			Owner:               "test-user",
			Name:                "Lantern Hoard",
			SurvivalLimit:       3,
			DepartingSurvival:   2,
			CollectiveCognition: 4,
			CurrentYear:         1,
		}

		// Test Insert
		id, err := db.Insert(ctx, settlement)
		require.NoError(t, err)
		assert.Greater(t, id, int32(0))

		// Test Get
		retrieved, err := db.Get(ctx, id, settlement.Owner)
		require.NoError(t, err)

		// Verify all fields
		assert.Equal(t, id, retrieved.ID)
		assert.Equal(t, settlement.Owner, retrieved.Owner)
		assert.Equal(t, settlement.Name, retrieved.Name)
		assert.Equal(t, settlement.SurvivalLimit, retrieved.SurvivalLimit)
		assert.Equal(t, settlement.DepartingSurvival, retrieved.DepartingSurvival)
		assert.Equal(t, settlement.CollectiveCognition, retrieved.CollectiveCognition)
		assert.Equal(t, settlement.CurrentYear, retrieved.CurrentYear)
	})

	t.Run("get non-existent settlement", func(t *testing.T) {
		_, err := db.Get(ctx, 999, "test-user")
		assert.Error(t, err)
	})

	t.Run("unauthorized access", func(t *testing.T) {
		// Create a settlement
		settlement := domain.Settlement{
			Owner:               "user-1",
			Name:                "White Lion",
			SurvivalLimit:       1,
			DepartingSurvival:   1,
			CollectiveCognition: 1,
			CurrentYear:         1,
		}

		id, err := db.Insert(ctx, settlement)
		require.NoError(t, err)

		// Try to access with wrong owner
		_, err = db.Get(ctx, id, "user-2")
		assert.Error(t, err)
	})

	t.Run("select settlements", func(t *testing.T) {
		owner := "test-user-select"
		settlements := []domain.Settlement{
			{
				Owner:               owner,
				Name:                "First Settlement",
				SurvivalLimit:       1,
				DepartingSurvival:   1,
				CollectiveCognition: 1,
				CurrentYear:         1,
			},
			{
				Owner:               owner,
				Name:                "Second Settlement",
				SurvivalLimit:       2,
				DepartingSurvival:   2,
				CollectiveCognition: 2,
				CurrentYear:         2,
			},
			{
				Owner:               "different-owner",
				Name:                "Other Settlement",
				SurvivalLimit:       3,
				DepartingSurvival:   3,
				CollectiveCognition: 3,
				CurrentYear:         3,
			},
		}

		// Insert all settlements
		for _, s := range settlements {
			_, err := db.Insert(ctx, s)
			require.NoError(t, err)
		}

		// Test Select
		retrieved, err := db.Select(ctx, owner)
		require.NoError(t, err)

		// Verify we got only the owner's settlements
		assert.Len(t, retrieved, 2)
		for _, s := range retrieved {
			assert.Equal(t, owner, s.Owner)
		}

		// Verify we got the correct names
		names := make(map[string]bool)
		for _, s := range retrieved {
			names[s.Name] = true
		}
		assert.True(t, names["First Settlement"])
		assert.True(t, names["Second Settlement"])
		assert.False(t, names["Other Settlement"])
	})

	t.Run("select returns empty slice when no settlements", func(t *testing.T) {
		retrieved, err := db.Select(ctx, "non-existent-user")
		require.NoError(t, err)
		assert.Empty(t, retrieved)
	})
}
