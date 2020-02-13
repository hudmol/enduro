// Code generated by goa v3.1.1, DO NOT EDIT.
//
// collection HTTP client CLI support package
//
// Command:
// $ goa gen github.com/artefactual-labs/enduro/internal/api/design -o
// internal/api

package client

import (
	"encoding/json"
	"fmt"
	"strconv"

	collection "github.com/artefactual-labs/enduro/internal/api/gen/collection"
	goa "goa.design/goa/v3/pkg"
)

// BuildListPayload builds the payload for the collection list endpoint from
// CLI flags.
func BuildListPayload(collectionListName string, collectionListOriginalID string, collectionListTransferID string, collectionListAipID string, collectionListPipelineID string, collectionListEarliestCreatedTime string, collectionListLatestCreatedTime string, collectionListStatus string, collectionListCursor string) (*collection.ListPayload, error) {
	var err error
	var name *string
	{
		if collectionListName != "" {
			name = &collectionListName
		}
	}
	var originalID *string
	{
		if collectionListOriginalID != "" {
			originalID = &collectionListOriginalID
		}
	}
	var transferID *string
	{
		if collectionListTransferID != "" {
			transferID = &collectionListTransferID
			if transferID != nil {
				err = goa.MergeErrors(err, goa.ValidateFormat("transferID", *transferID, goa.FormatUUID))
			}
			if err != nil {
				return nil, err
			}
		}
	}
	var aipID *string
	{
		if collectionListAipID != "" {
			aipID = &collectionListAipID
			if aipID != nil {
				err = goa.MergeErrors(err, goa.ValidateFormat("aipID", *aipID, goa.FormatUUID))
			}
			if err != nil {
				return nil, err
			}
		}
	}
	var pipelineID *string
	{
		if collectionListPipelineID != "" {
			pipelineID = &collectionListPipelineID
			if pipelineID != nil {
				err = goa.MergeErrors(err, goa.ValidateFormat("pipelineID", *pipelineID, goa.FormatUUID))
			}
			if err != nil {
				return nil, err
			}
		}
	}
	var earliestCreatedTime *string
	{
		if collectionListEarliestCreatedTime != "" {
			earliestCreatedTime = &collectionListEarliestCreatedTime
			if earliestCreatedTime != nil {
				err = goa.MergeErrors(err, goa.ValidateFormat("earliestCreatedTime", *earliestCreatedTime, goa.FormatDateTime))
			}
			if err != nil {
				return nil, err
			}
		}
	}
	var latestCreatedTime *string
	{
		if collectionListLatestCreatedTime != "" {
			latestCreatedTime = &collectionListLatestCreatedTime
			if latestCreatedTime != nil {
				err = goa.MergeErrors(err, goa.ValidateFormat("latestCreatedTime", *latestCreatedTime, goa.FormatDateTime))
			}
			if err != nil {
				return nil, err
			}
		}
	}
	var status *string
	{
		if collectionListStatus != "" {
			status = &collectionListStatus
			if status != nil {
				if !(*status == "new" || *status == "in progress" || *status == "done" || *status == "error" || *status == "unknown" || *status == "queued" || *status == "pending" || *status == "abandoned") {
					err = goa.MergeErrors(err, goa.InvalidEnumValueError("status", *status, []interface{}{"new", "in progress", "done", "error", "unknown", "queued", "pending", "abandoned"}))
				}
			}
			if err != nil {
				return nil, err
			}
		}
	}
	var cursor *string
	{
		if collectionListCursor != "" {
			cursor = &collectionListCursor
		}
	}
	v := &collection.ListPayload{}
	v.Name = name
	v.OriginalID = originalID
	v.TransferID = transferID
	v.AipID = aipID
	v.PipelineID = pipelineID
	v.EarliestCreatedTime = earliestCreatedTime
	v.LatestCreatedTime = latestCreatedTime
	v.Status = status
	v.Cursor = cursor

	return v, nil
}

// BuildShowPayload builds the payload for the collection show endpoint from
// CLI flags.
func BuildShowPayload(collectionShowID string) (*collection.ShowPayload, error) {
	var err error
	var id uint
	{
		var v uint64
		v, err = strconv.ParseUint(collectionShowID, 10, 64)
		id = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT")
		}
	}
	v := &collection.ShowPayload{}
	v.ID = id

	return v, nil
}

// BuildDeletePayload builds the payload for the collection delete endpoint
// from CLI flags.
func BuildDeletePayload(collectionDeleteID string) (*collection.DeletePayload, error) {
	var err error
	var id uint
	{
		var v uint64
		v, err = strconv.ParseUint(collectionDeleteID, 10, 64)
		id = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT")
		}
	}
	v := &collection.DeletePayload{}
	v.ID = id

	return v, nil
}

// BuildCancelPayload builds the payload for the collection cancel endpoint
// from CLI flags.
func BuildCancelPayload(collectionCancelID string) (*collection.CancelPayload, error) {
	var err error
	var id uint
	{
		var v uint64
		v, err = strconv.ParseUint(collectionCancelID, 10, 64)
		id = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT")
		}
	}
	v := &collection.CancelPayload{}
	v.ID = id

	return v, nil
}

// BuildRetryPayload builds the payload for the collection retry endpoint from
// CLI flags.
func BuildRetryPayload(collectionRetryID string) (*collection.RetryPayload, error) {
	var err error
	var id uint
	{
		var v uint64
		v, err = strconv.ParseUint(collectionRetryID, 10, 64)
		id = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT")
		}
	}
	v := &collection.RetryPayload{}
	v.ID = id

	return v, nil
}

// BuildWorkflowPayload builds the payload for the collection workflow endpoint
// from CLI flags.
func BuildWorkflowPayload(collectionWorkflowID string) (*collection.WorkflowPayload, error) {
	var err error
	var id uint
	{
		var v uint64
		v, err = strconv.ParseUint(collectionWorkflowID, 10, 64)
		id = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT")
		}
	}
	v := &collection.WorkflowPayload{}
	v.ID = id

	return v, nil
}

// BuildDownloadPayload builds the payload for the collection download endpoint
// from CLI flags.
func BuildDownloadPayload(collectionDownloadID string) (*collection.DownloadPayload, error) {
	var err error
	var id uint
	{
		var v uint64
		v, err = strconv.ParseUint(collectionDownloadID, 10, 64)
		id = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT")
		}
	}
	v := &collection.DownloadPayload{}
	v.ID = id

	return v, nil
}

// BuildDecidePayload builds the payload for the collection decide endpoint
// from CLI flags.
func BuildDecidePayload(collectionDecideBody string, collectionDecideID string) (*collection.DecidePayload, error) {
	var err error
	var body struct {
		// Decision option to proceed with
		Option *string
	}
	{
		err = json.Unmarshal([]byte(collectionDecideBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"option\": \"Recusandae laudantium quidem consequatur ducimus excepturi perferendis.\"\n   }'")
		}
	}
	var id uint
	{
		var v uint64
		v, err = strconv.ParseUint(collectionDecideID, 10, 64)
		id = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT")
		}
	}
	v := &collection.DecidePayload{}
	if body.Option != nil {
		v.Option = *body.Option
	}
	v.ID = id

	return v, nil
}

// BuildBulkPayload builds the payload for the collection bulk endpoint from
// CLI flags.
func BuildBulkPayload(collectionBulkBody string) (*collection.BulkPayload, error) {
	var err error
	var body BulkRequestBody
	{
		err = json.Unmarshal([]byte(collectionBulkBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"operation\": \"cancel\",\n      \"size\": 4963086957530573892,\n      \"status\": \"unknown\"\n   }'")
		}
		if !(body.Operation == "retry" || body.Operation == "cancel" || body.Operation == "abandon") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.operation", body.Operation, []interface{}{"retry", "cancel", "abandon"}))
		}
		if !(body.Status == "new" || body.Status == "in progress" || body.Status == "done" || body.Status == "error" || body.Status == "unknown" || body.Status == "queued" || body.Status == "pending" || body.Status == "abandoned") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.status", body.Status, []interface{}{"new", "in progress", "done", "error", "unknown", "queued", "pending", "abandoned"}))
		}
		if err != nil {
			return nil, err
		}
	}
	v := &collection.BulkPayload{
		Operation: body.Operation,
		Status:    body.Status,
		Size:      body.Size,
	}

	return v, nil
}
