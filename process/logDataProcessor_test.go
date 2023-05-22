package process

import (
	"strings"
	"testing"

	"github.com/multiversx/mx-chain-core-go/core/check"
	"github.com/multiversx/mx-chain-core-go/data/outport"
	"github.com/multiversx/mx-chain-ws-connector-template-go/testscommon"
	"github.com/stretchr/testify/require"
)

func TestNewLogDataProcessor(t *testing.T) {
	t.Parallel()

	t.Run("nil marshaller, should return error", func(t *testing.T) {
		dp, err := NewLogDataProcessor(nil)
		require.True(t, check.IfNil(dp))
		require.Equal(t, errNilMarshaller, err)
	})

	t.Run("should work", func(t *testing.T) {
		dp, err := NewLogDataProcessor(&testscommon.MarshallerMock{})
		require.False(t, check.IfNil(dp))
		require.Nil(t, err)

		err = dp.Close()
		require.Nil(t, err)
	})
}

func TestLogDataProcessor_ProcessPayload(t *testing.T) {
	t.Parallel()

	marshaller := &testscommon.MarshallerMock{}
	dp, _ := NewLogDataProcessor(marshaller)

	t.Run("invalid topic, should return error", func(t *testing.T) {
		t.Parallel()

		err := dp.ProcessPayload([]byte("payload"), "invalid topic")
		require.True(t, strings.Contains(err.Error(), errInvalidOperationType.Error()))
		require.True(t, strings.Contains(err.Error(), "payload"))
		require.True(t, strings.Contains(err.Error(), "invalid topic"))
	})

	t.Run("save block", func(t *testing.T) {
		t.Parallel()

		err := dp.ProcessPayload([]byte("payload"), outport.TopicSaveBlock)
		require.NotNil(t, err)

		outportBlock := &outport.OutportBlock{}
		outportBlockBytes, err := marshaller.Marshal(outportBlock)
		require.Nil(t, err)

		err = dp.ProcessPayload(outportBlockBytes, outport.TopicSaveBlock)
		require.Nil(t, err)
	})

	t.Run("revert indexed block", func(t *testing.T) {
		t.Parallel()

		err := dp.ProcessPayload([]byte("payload"), outport.TopicRevertIndexedBlock)
		require.NotNil(t, err)

		blockData := &outport.BlockData{}
		blockDataBytes, err := marshaller.Marshal(blockData)
		require.Nil(t, err)

		err = dp.ProcessPayload(blockDataBytes, outport.TopicRevertIndexedBlock)
		require.Nil(t, err)
	})

	t.Run("save rounds", func(t *testing.T) {
		t.Parallel()

		err := dp.ProcessPayload([]byte("payload"), outport.TopicRevertIndexedBlock)
		require.NotNil(t, err)

		roundsInfo := &outport.RoundsInfo{}
		roundsInfoBytes, err := marshaller.Marshal(roundsInfo)
		require.Nil(t, err)

		err = dp.ProcessPayload(roundsInfoBytes, outport.TopicRevertIndexedBlock)
		require.Nil(t, err)
	})

	t.Run("save rounds", func(t *testing.T) {
		t.Parallel()

		err := dp.ProcessPayload([]byte("payload"), outport.TopicSaveRoundsInfo)
		require.NotNil(t, err)

		roundsInfo := &outport.RoundsInfo{}
		roundsInfoBytes, err := marshaller.Marshal(roundsInfo)
		require.Nil(t, err)

		err = dp.ProcessPayload(roundsInfoBytes, outport.TopicSaveRoundsInfo)
		require.Nil(t, err)
	})

	t.Run("save validators rating", func(t *testing.T) {
		t.Parallel()

		err := dp.ProcessPayload([]byte("payload"), outport.TopicSaveValidatorsRating)
		require.NotNil(t, err)

		ratingData := &outport.ValidatorsRating{}
		ratingDataBytes, err := marshaller.Marshal(ratingData)
		require.Nil(t, err)

		err = dp.ProcessPayload(ratingDataBytes, outport.TopicSaveValidatorsRating)
		require.Nil(t, err)
	})

	t.Run("save validators pubKeys", func(t *testing.T) {
		t.Parallel()

		err := dp.ProcessPayload([]byte("payload"), outport.TopicSaveValidatorsPubKeys)
		require.NotNil(t, err)

		validatorsPubKeys := &outport.ValidatorsPubKeys{}
		validatorsPubKeysBytes, err := marshaller.Marshal(validatorsPubKeys)
		require.Nil(t, err)

		err = dp.ProcessPayload(validatorsPubKeysBytes, outport.TopicSaveValidatorsPubKeys)
		require.Nil(t, err)
	})

	t.Run("save accounts", func(t *testing.T) {
		t.Parallel()

		err := dp.ProcessPayload([]byte("payload"), outport.TopicSaveAccounts)
		require.NotNil(t, err)

		accounts := &outport.Accounts{}
		accountsBytes, err := marshaller.Marshal(accounts)
		require.Nil(t, err)

		err = dp.ProcessPayload(accountsBytes, outport.TopicSaveAccounts)
		require.Nil(t, err)
	})

	t.Run("finalized block", func(t *testing.T) {
		t.Parallel()

		err := dp.ProcessPayload([]byte("payload"), outport.TopicFinalizedBlock)
		require.NotNil(t, err)

		finalizedBlock := &outport.FinalizedBlock{}
		finalizedBlockBytes, err := marshaller.Marshal(finalizedBlock)
		require.Nil(t, err)

		err = dp.ProcessPayload(finalizedBlockBytes, outport.TopicFinalizedBlock)
		require.Nil(t, err)
	})
}
