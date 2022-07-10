package archivedraft_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/thomasduchatelle/dphoto/domain/catalog"
	"github.com/thomasduchatelle/dphoto/mocks"
	"testing"
)

func TestRelocateMovedMedias_full(t *testing.T) {
	a := assert.New(t)

	// given
	operator := new(mocks.MoveMediaOperator)
	repository := new(mocks.RepositoryPort)
	catalog.db = repository

	const transactionId = "move-transaction-1"
	const owner = "stark"

	movedMedias := []*catalog.MovedMedia{
		{Signature: catalog.MediaSignature{}, SourceFolderName: "A", SourceFilename: "001_before", TargetFolderName: "B", TargetFilename: "001"},
		{Signature: catalog.MediaSignature{}, SourceFolderName: "C", SourceFilename: "002", TargetFolderName: "B", TargetFilename: "002"},
		{Signature: catalog.MediaSignature{}, SourceFolderName: "A", SourceFilename: "003", TargetFolderName: "D", TargetFilename: "003"},
	}

	repository.On("FindReadyMoveTransactions", owner).Return([]*catalog.MoveTransaction{
		{TransactionId: transactionId, Count: 42},
	}, nil)
	repository.On("FindFilesToMove", transactionId, "").Return(movedMedias[:2], "next-page-1", nil)
	repository.On("FindFilesToMove", transactionId, "next-page-1").Return(movedMedias[2:], "", nil)

	repository.On("UpdateMediasLocation", owner, transactionId, movedMedias[:2]).Return(nil)
	repository.On("UpdateMediasLocation", owner, transactionId, movedMedias[2:]).Return(nil)

	repository.On("DeleteEmptyMoveTransaction", transactionId).Once().Return(nil)

	operator.On("Continue").Return(true)
	operator.On("UpdateStatus", 0, 42).Return(nil)
	operator.On("UpdateStatus", 2, 42).Return(nil)
	operator.On("UpdateStatus", 3, 42).Return(nil)

	operator.On("Move", catalog.MediaLocation{FolderName: "A", Filename: "001_before"}, catalog.MediaLocation{FolderName: "B", Filename: "001"}).Return("001", nil)
	operator.On("Move", catalog.MediaLocation{FolderName: "C", Filename: "002"}, catalog.MediaLocation{FolderName: "B", Filename: "002"}).Return("002", nil)
	operator.On("Move", catalog.MediaLocation{FolderName: "A", Filename: "003"}, catalog.MediaLocation{FolderName: "D", Filename: "003"}).Return("003", nil)

	// when
	got, err := catalog.RelocateMovedMedias(owner, operator, transactionId)

	if a.NoError(err) {
		a.Equal(3, got)
		repository.AssertExpectations(t)
		operator.AssertExpectations(t)
	}
}

func TestRelocateMovedMedias_interrupt(t *testing.T) {
	a := assert.New(t)

	// given
	const owner = "stark"
	operator := new(mocks.MoveMediaOperator)
	repository := new(mocks.RepositoryPort)
	catalog.db = repository

	const transactionId = "move-transaction-1"

	movedMedias := []*catalog.MovedMedia{
		{Signature: catalog.MediaSignature{}, SourceFolderName: "A", SourceFilename: "001", TargetFolderName: "B", TargetFilename: "001"},
		{Signature: catalog.MediaSignature{}, SourceFolderName: "C", SourceFilename: "002", TargetFolderName: "B", TargetFilename: "002"},
		{Signature: catalog.MediaSignature{}, SourceFolderName: "A", SourceFilename: "003", TargetFolderName: "D", TargetFilename: "003"},
	}

	repository.On("FindReadyMoveTransactions", owner).Return([]*catalog.MoveTransaction{
		{TransactionId: transactionId, Count: 42},
	}, nil)
	repository.On("FindFilesToMove", transactionId, "").Return(movedMedias[:2], "next-page-1", nil)

	repository.On("UpdateMediasLocation", owner, transactionId, movedMedias[:2]).Return(nil)

	repository.On("DeleteEmptyMoveTransaction", transactionId).Once().Return(nil)

	operator.On("Continue").Return(true).Once()
	operator.On("Continue").Return(false)
	operator.On("UpdateStatus", 0, 42).Return(nil)
	operator.On("UpdateStatus", 2, 42).Return(nil)

	operator.On("Move", catalog.MediaLocation{FolderName: "A", Filename: "001"}, catalog.MediaLocation{FolderName: "B", Filename: "001"}).Return("001", nil)
	operator.On("Move", catalog.MediaLocation{FolderName: "C", Filename: "002"}, catalog.MediaLocation{FolderName: "B", Filename: "002"}).Return("002", nil)

	// when
	got, err := catalog.RelocateMovedMedias(owner, operator, transactionId)

	if a.NoError(err) {
		a.Equal(2, got)
		operator.AssertExpectations(t)
	}
}

func TestRelocateMovedMedias_empty(t *testing.T) {
	a := assert.New(t)

	// given
	const owner = "ironman"
	operator := new(mocks.MoveMediaOperator)
	repository := new(mocks.RepositoryPort)
	catalog.db = repository

	repository.On("FindReadyMoveTransactions", owner).Return(nil, nil)

	// when
	got, err := catalog.RelocateMovedMedias(owner, operator, "no-transaction-exist")

	if a.NoError(err) {
		a.Equal(0, got)
		operator.AssertExpectations(t)
	}
}