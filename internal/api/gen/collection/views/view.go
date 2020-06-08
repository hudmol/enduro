// Code generated by goa v3.1.3, DO NOT EDIT.
//
// collection views
//
// Command:
// $ goa gen github.com/artefactual-labs/enduro/internal/api/design -o
// internal/api

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// EnduroStoredCollection is the viewed result type that is projected based on
// a view.
type EnduroStoredCollection struct {
	// Type to project
	Projected *EnduroStoredCollectionView
	// View to render
	View string
}

// EnduroCollectionWorkflowStatus is the viewed result type that is projected
// based on a view.
type EnduroCollectionWorkflowStatus struct {
	// Type to project
	Projected *EnduroCollectionWorkflowStatusView
	// View to render
	View string
}

// EnduroStoredCollectionView is a type that runs validations on a projected
// type.
type EnduroStoredCollectionView struct {
	// Identifier of collection
	ID *uint
	// Name of the collection
	Name *string
	// Status of the collection
	Status *string
	// Identifier of processing workflow
	WorkflowID *string
	// Identifier of latest processing workflow run
	RunID *string
	// Identifier of Archivematica transfer
	TransferID *string
	// Identifier of Archivematica AIP
	AipID *string
	// Identifier provided by the client
	OriginalID *string
	// Identifier of Archivematica pipeline
	PipelineID *string
	// Creation datetime
	CreatedAt *string
	// Start datetime
	StartedAt *string
	// Completion datetime
	CompletedAt *string
}

// EnduroCollectionWorkflowStatusView is a type that runs validations on a
// projected type.
type EnduroCollectionWorkflowStatusView struct {
	Status  *string
	History EnduroCollectionWorkflowHistoryCollectionView
}

// EnduroCollectionWorkflowHistoryCollectionView is a type that runs
// validations on a projected type.
type EnduroCollectionWorkflowHistoryCollectionView []*EnduroCollectionWorkflowHistoryView

// EnduroCollectionWorkflowHistoryView is a type that runs validations on a
// projected type.
type EnduroCollectionWorkflowHistoryView struct {
	// Identifier of collection
	ID *uint
	// Type of the event
	Type *string
	// Contents of the event
	Details interface{}
}

var (
	// EnduroStoredCollectionMap is a map of attribute names in result type
	// EnduroStoredCollection indexed by view name.
	EnduroStoredCollectionMap = map[string][]string{
		"default": []string{
			"id",
			"name",
			"status",
			"workflow_id",
			"run_id",
			"transfer_id",
			"aip_id",
			"original_id",
			"pipeline_id",
			"created_at",
			"started_at",
			"completed_at",
		},
	}
	// EnduroCollectionWorkflowStatusMap is a map of attribute names in result type
	// EnduroCollectionWorkflowStatus indexed by view name.
	EnduroCollectionWorkflowStatusMap = map[string][]string{
		"default": []string{
			"status",
			"history",
		},
	}
	// EnduroCollectionWorkflowHistoryCollectionMap is a map of attribute names in
	// result type EnduroCollectionWorkflowHistoryCollection indexed by view name.
	EnduroCollectionWorkflowHistoryCollectionMap = map[string][]string{
		"default": []string{
			"id",
			"type",
			"details",
		},
	}
	// EnduroCollectionWorkflowHistoryMap is a map of attribute names in result
	// type EnduroCollectionWorkflowHistory indexed by view name.
	EnduroCollectionWorkflowHistoryMap = map[string][]string{
		"default": []string{
			"id",
			"type",
			"details",
		},
	}
)

// ValidateEnduroStoredCollection runs the validations defined on the viewed
// result type EnduroStoredCollection.
func ValidateEnduroStoredCollection(result *EnduroStoredCollection) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateEnduroStoredCollectionView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateEnduroCollectionWorkflowStatus runs the validations defined on the
// viewed result type EnduroCollectionWorkflowStatus.
func ValidateEnduroCollectionWorkflowStatus(result *EnduroCollectionWorkflowStatus) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateEnduroCollectionWorkflowStatusView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateEnduroStoredCollectionView runs the validations defined on
// EnduroStoredCollectionView using the "default" view.
func ValidateEnduroStoredCollectionView(result *EnduroStoredCollectionView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "result"))
	}
	if result.CreatedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("created_at", "result"))
	}
	if result.Status != nil {
		if !(*result.Status == "new" || *result.Status == "in progress" || *result.Status == "done" || *result.Status == "error" || *result.Status == "unknown" || *result.Status == "queued" || *result.Status == "pending" || *result.Status == "abandoned") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.status", *result.Status, []interface{}{"new", "in progress", "done", "error", "unknown", "queued", "pending", "abandoned"}))
		}
	}
	if result.WorkflowID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.workflow_id", *result.WorkflowID, goa.FormatUUID))
	}
	if result.RunID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.run_id", *result.RunID, goa.FormatUUID))
	}
	if result.TransferID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.transfer_id", *result.TransferID, goa.FormatUUID))
	}
	if result.AipID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.aip_id", *result.AipID, goa.FormatUUID))
	}
	if result.PipelineID != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.pipeline_id", *result.PipelineID, goa.FormatUUID))
	}
	if result.CreatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.created_at", *result.CreatedAt, goa.FormatDateTime))
	}
	if result.StartedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.started_at", *result.StartedAt, goa.FormatDateTime))
	}
	if result.CompletedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.completed_at", *result.CompletedAt, goa.FormatDateTime))
	}
	return
}

// ValidateEnduroCollectionWorkflowStatusView runs the validations defined on
// EnduroCollectionWorkflowStatusView using the "default" view.
func ValidateEnduroCollectionWorkflowStatusView(result *EnduroCollectionWorkflowStatusView) (err error) {

	if result.History != nil {
		if err2 := ValidateEnduroCollectionWorkflowHistoryCollectionView(result.History); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateEnduroCollectionWorkflowHistoryCollectionView runs the validations
// defined on EnduroCollectionWorkflowHistoryCollectionView using the "default"
// view.
func ValidateEnduroCollectionWorkflowHistoryCollectionView(result EnduroCollectionWorkflowHistoryCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateEnduroCollectionWorkflowHistoryView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateEnduroCollectionWorkflowHistoryView runs the validations defined on
// EnduroCollectionWorkflowHistoryView using the "default" view.
func ValidateEnduroCollectionWorkflowHistoryView(result *EnduroCollectionWorkflowHistoryView) (err error) {

	return
}
