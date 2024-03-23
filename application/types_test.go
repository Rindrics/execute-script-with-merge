package application

import (
	"testing"

	"github.com/Rindrics/execute-script-with-merge/domain"
	mock "github.com/Rindrics/execute-script-with-merge/domain/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestAppHasRequiredLabel(t *testing.T) {
	app := New("required-label", "main", nil)

	t.Run("HasRequiredLabel", func(t *testing.T) {
		event := domain.ParsedEvent{
			Labels: domain.Labels{"required-label"},
		}

		assert.True(t, app.HasRequiredLabel(event))
	})
	t.Run("NotHasRequiredLabel", func(t *testing.T) {
		event := domain.ParsedEvent{
			Labels: domain.Labels{"other-label"},
		}
		assert.False(t, app.HasRequiredLabel(event))
	})
}

func TestAppIsDefaultBranch(t *testing.T) {
	app := New("required-label", "main", nil)

	t.Run("IsDefaultBranch", func(t *testing.T) {
		event := domain.ParsedEvent{
			Branch: "main",
		}
		assert.True(t, app.IsDefaultBranch(event))
	})
	t.Run("NotIsDefaultBranch", func(t *testing.T) {
		event := domain.ParsedEvent{
			Branch: "other-branch",
		}
		assert.False(t, app.IsDefaultBranch(event))
	})
}

func TestAppIsValid(t *testing.T) {
	app := New("required-label", "main", nil)

	t.Run("Valid", func(t *testing.T) {
		event := domain.ParsedEvent{
			Branch: "main",
			Labels: domain.Labels{"required-label"},
		}
		assert.True(t, app.IsValid(event))
	})
	t.Run("NoLabel", func(t *testing.T) {
		event := domain.ParsedEvent{
			Branch: "main",
			Labels: domain.Labels{},
		}
		assert.False(t, app.IsValid(event))
	})
	t.Run("NotDefaultBranch", func(t *testing.T) {
		event := domain.ParsedEvent{
			Branch: "other-branch",
			Labels: domain.Labels{"required-label"},
		}
		assert.False(t, app.IsValid(event))
	})
}

func TestAppLoadExecutionDirectiveList(t *testing.T) {
	expectedDirectives := []domain.ExecutionDirective{"foo.sh", "bar.sh"}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockParser := mock.NewMockEventParser(ctrl)
	mockParser.EXPECT().ParseExecutionDirectives().Return(expectedDirectives, nil).Times(1)

	app := New("required-label", "main", mockParser)

	t.Run("LoadExecutionDirectiveList", func(t *testing.T) {
		err := app.LoadExecutionDirectiveList()
		assert.Nil(t, err)
		assert.Equal(t, expectedDirectives, app.ExecutionDirectiveList.ExecutionDirectives,
			"The execution directives should match the expected values.")
	})
}
