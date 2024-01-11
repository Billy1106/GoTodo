package repository

import (
	"memo-go/domain"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMapRepository(t *testing.T) {
	repo := NewTodoRepository()

	todo := domain.Todo{
		Title:     "todo1",
		Completed: false,
	}
	todo2 := domain.Todo{
		Title:     "todo2",
		Completed: true,
	}
	t.Run("AllGet Todo Test", func(t *testing.T) {
		r1, _ := repo.AllGet()
		require.Empty(t, r1)
	})

	t.Run("Store Todo Test", func(t *testing.T) {
		repo.Store(todo)
		r2, _ := repo.AllGet()
		require.Equal(t, r2[0], todo)
	})

	t.Run("AllGet Todo Test", func(t *testing.T) {
		repo.Store(todo2)
		todos := [2]domain.Todo{todo, todo2}
		r2, _ := repo.AllGet()
		for i := 0; i < len(r2); i++ {
			require.Equal(t, r2[i].Title, todos[i].Title)
			require.Equal(t, r2[i].Completed, todos[i].Completed)
			require.Equal(t, r2[i].ID, i)
		}
	})

	// t.Run("StatusUpdate Todo Test", func(t *testing.T) {
	// 	repo.StatusUpdate(0, true)
	// 	r2, _ := repo.AllGet()
	// 	require.Equal(t, r2[0].Completed, true)
	// 	require.Equal(t, r2[1].Completed, true)
	// })

	t.Run("Delete Todo Test", func(t *testing.T) {
		repo.Delete(1)
		r2, _ := repo.AllGet()
		require.Equal(t, len(r2), 1)
		require.Equal(t, r2[0], todo)

	})
}
