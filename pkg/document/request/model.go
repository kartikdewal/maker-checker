package request

import (
	"encoding/json"
	"github.com/jmoiron/sqlx/types"
	"time"
)

type ApprovalStatus string

const (
	Pending  ApprovalStatus = "Pending"
	Approved ApprovalStatus = "Approved"
	Rejected ApprovalStatus = "Rejected"
)

type Row struct {
	ID             string         `db:"id"`
	DocumentID     string         `db:"document_id"`
	CreatorID      string         `db:"creator_id"`
	ApproversJSON  types.JSONText `db:"approvers"`
	Approvers      []*Approver    `db:"-"`
	ApproverCount  int            `db:"approver_count"`
	Status         ApprovalStatus `db:"status"`
	RecipientEmail string         `db:"recipient_email"`
	CreatedAt      time.Time      `db:"created_at"`
	UpdatedAt      time.Time      `db:"updated_at"`
}

type Approver struct {
	ID         string         `json:"id"`
	Status     ApprovalStatus `json:"status"`
	ApprovedAt string         `json:"approved_at"`
}

func (r *Row) Unmarshal() error {
	var approvers []*Approver
	if err := r.ApproversJSON.Unmarshal(&approvers); err != nil {
		return err
	}
	r.Approvers = approvers
	return nil
}

func (r *Row) Marshal() (map[string]interface{}, error) {
	approversJSON, err := json.Marshal(r.Approvers)
	if err != nil {

		return nil, err
	}
	return map[string]interface{}{
		"id":              r.ID,
		"document_id":     r.DocumentID,
		"creator_id":      r.CreatorID,
		"approvers":       approversJSON,
		"approver_count":  r.ApproverCount,
		"status":          r.Status,
		"recipient_email": r.RecipientEmail,
		"created_at":      r.CreatedAt,
		"updated_at":      r.UpdatedAt,
	}, nil
}
