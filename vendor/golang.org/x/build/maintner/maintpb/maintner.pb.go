// Code generated by protoc-gen-go4grpc; DO NOT EDIT
// source: maintner.proto

/*
Package maintpb is a generated protocol buffer package.

It is generated from these files:
	maintner.proto

It has these top-level messages:
	Mutation
	GithubMutation
	GithubIssueMutation
	BoolChange
	GithubLabel
	GithubMilestone
	GithubIssueEvent
	GithubDismissedReviewEvent
	GithubCommit
	GithubIssueSyncStatus
	GithubIssueCommentMutation
	GithubUser
	GitMutation
	GitRepo
	GitCommit
	GitDiffTree
	GitDiffTreeFile
	GerritMutation
	GitRef
*/
package maintpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Mutation struct {
	GithubIssue *GithubIssueMutation `protobuf:"bytes,1,opt,name=github_issue,json=githubIssue" json:"github_issue,omitempty"`
	Github      *GithubMutation      `protobuf:"bytes,3,opt,name=github" json:"github,omitempty"`
	Git         *GitMutation         `protobuf:"bytes,2,opt,name=git" json:"git,omitempty"`
	Gerrit      *GerritMutation      `protobuf:"bytes,4,opt,name=gerrit" json:"gerrit,omitempty"`
}

func (m *Mutation) Reset()                    { *m = Mutation{} }
func (m *Mutation) String() string            { return proto.CompactTextString(m) }
func (*Mutation) ProtoMessage()               {}
func (*Mutation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Mutation) GetGithubIssue() *GithubIssueMutation {
	if m != nil {
		return m.GithubIssue
	}
	return nil
}

func (m *Mutation) GetGithub() *GithubMutation {
	if m != nil {
		return m.Github
	}
	return nil
}

func (m *Mutation) GetGit() *GitMutation {
	if m != nil {
		return m.Git
	}
	return nil
}

func (m *Mutation) GetGerrit() *GerritMutation {
	if m != nil {
		return m.Gerrit
	}
	return nil
}

type GithubMutation struct {
	Owner string `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	Repo  string `protobuf:"bytes,2,opt,name=repo" json:"repo,omitempty"`
	// Updated labels. (All must have id set at least)
	Labels []*GithubLabel `protobuf:"bytes,3,rep,name=labels" json:"labels,omitempty"`
	// Updated milestones. (All must have id set at least)
	Milestones []*GithubMilestone `protobuf:"bytes,4,rep,name=milestones" json:"milestones,omitempty"`
}

func (m *GithubMutation) Reset()                    { *m = GithubMutation{} }
func (m *GithubMutation) String() string            { return proto.CompactTextString(m) }
func (*GithubMutation) ProtoMessage()               {}
func (*GithubMutation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GithubMutation) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *GithubMutation) GetRepo() string {
	if m != nil {
		return m.Repo
	}
	return ""
}

func (m *GithubMutation) GetLabels() []*GithubLabel {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *GithubMutation) GetMilestones() []*GithubMilestone {
	if m != nil {
		return m.Milestones
	}
	return nil
}

type GithubIssueMutation struct {
	Owner  string `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	Repo   string `protobuf:"bytes,2,opt,name=repo" json:"repo,omitempty"`
	Number int32  `protobuf:"varint,3,opt,name=number" json:"number,omitempty"`
	// not_exist is set true if the issue has been found to not exist.
	// If true, the owner/repo/number fields above must still be set.
	// If a future issue mutation for the same number arrives without
	// not_exist set, then the issue comes back to life.
	NotExist         bool                       `protobuf:"varint,13,opt,name=not_exist,json=notExist" json:"not_exist,omitempty"`
	Id               int64                      `protobuf:"varint,12,opt,name=id" json:"id,omitempty"`
	User             *GithubUser                `protobuf:"bytes,4,opt,name=user" json:"user,omitempty"`
	Assignees        []*GithubUser              `protobuf:"bytes,10,rep,name=assignees" json:"assignees,omitempty"`
	DeletedAssignees []int64                    `protobuf:"varint,11,rep,packed,name=deleted_assignees,json=deletedAssignees" json:"deleted_assignees,omitempty"`
	Created          *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=created" json:"created,omitempty"`
	Updated          *google_protobuf.Timestamp `protobuf:"bytes,6,opt,name=updated" json:"updated,omitempty"`
	Body             string                     `protobuf:"bytes,7,opt,name=body" json:"body,omitempty"`
	Title            string                     `protobuf:"bytes,9,opt,name=title" json:"title,omitempty"`
	NoMilestone      bool                       `protobuf:"varint,15,opt,name=no_milestone,json=noMilestone" json:"no_milestone,omitempty"`
	// When setting a milestone, only the milestone_id must be set.
	// TODO: allow num or title to be used if Github only returns those? So far unneeded.
	// The num and title, if non-zero, are treated as if they were a GithubMutation.Milestone update.
	MilestoneId    int64                         `protobuf:"varint,16,opt,name=milestone_id,json=milestoneId" json:"milestone_id,omitempty"`
	MilestoneNum   int64                         `protobuf:"varint,17,opt,name=milestone_num,json=milestoneNum" json:"milestone_num,omitempty"`
	MilestoneTitle string                        `protobuf:"bytes,18,opt,name=milestone_title,json=milestoneTitle" json:"milestone_title,omitempty"`
	Closed         *BoolChange                   `protobuf:"bytes,19,opt,name=closed" json:"closed,omitempty"`
	Locked         *BoolChange                   `protobuf:"bytes,25,opt,name=locked" json:"locked,omitempty"`
	PullRequest    bool                          `protobuf:"varint,28,opt,name=pull_request,json=pullRequest" json:"pull_request,omitempty"`
	ClosedAt       *google_protobuf.Timestamp    `protobuf:"bytes,21,opt,name=closed_at,json=closedAt" json:"closed_at,omitempty"`
	ClosedBy       *GithubUser                   `protobuf:"bytes,22,opt,name=closed_by,json=closedBy" json:"closed_by,omitempty"`
	RemoveLabel    []int64                       `protobuf:"varint,23,rep,packed,name=remove_label,json=removeLabel" json:"remove_label,omitempty"`
	AddLabel       []*GithubLabel                `protobuf:"bytes,24,rep,name=add_label,json=addLabel" json:"add_label,omitempty"`
	Comment        []*GithubIssueCommentMutation `protobuf:"bytes,8,rep,name=comment" json:"comment,omitempty"`
	CommentStatus  *GithubIssueSyncStatus        `protobuf:"bytes,14,opt,name=comment_status,json=commentStatus" json:"comment_status,omitempty"`
	Event          []*GithubIssueEvent           `protobuf:"bytes,26,rep,name=event" json:"event,omitempty"`
	EventStatus    *GithubIssueSyncStatus        `protobuf:"bytes,27,opt,name=event_status,json=eventStatus" json:"event_status,omitempty"`
}

func (m *GithubIssueMutation) Reset()                    { *m = GithubIssueMutation{} }
func (m *GithubIssueMutation) String() string            { return proto.CompactTextString(m) }
func (*GithubIssueMutation) ProtoMessage()               {}
func (*GithubIssueMutation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GithubIssueMutation) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *GithubIssueMutation) GetRepo() string {
	if m != nil {
		return m.Repo
	}
	return ""
}

func (m *GithubIssueMutation) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *GithubIssueMutation) GetNotExist() bool {
	if m != nil {
		return m.NotExist
	}
	return false
}

func (m *GithubIssueMutation) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GithubIssueMutation) GetUser() *GithubUser {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *GithubIssueMutation) GetAssignees() []*GithubUser {
	if m != nil {
		return m.Assignees
	}
	return nil
}

func (m *GithubIssueMutation) GetDeletedAssignees() []int64 {
	if m != nil {
		return m.DeletedAssignees
	}
	return nil
}

func (m *GithubIssueMutation) GetCreated() *google_protobuf.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *GithubIssueMutation) GetUpdated() *google_protobuf.Timestamp {
	if m != nil {
		return m.Updated
	}
	return nil
}

func (m *GithubIssueMutation) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *GithubIssueMutation) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *GithubIssueMutation) GetNoMilestone() bool {
	if m != nil {
		return m.NoMilestone
	}
	return false
}

func (m *GithubIssueMutation) GetMilestoneId() int64 {
	if m != nil {
		return m.MilestoneId
	}
	return 0
}

func (m *GithubIssueMutation) GetMilestoneNum() int64 {
	if m != nil {
		return m.MilestoneNum
	}
	return 0
}

func (m *GithubIssueMutation) GetMilestoneTitle() string {
	if m != nil {
		return m.MilestoneTitle
	}
	return ""
}

func (m *GithubIssueMutation) GetClosed() *BoolChange {
	if m != nil {
		return m.Closed
	}
	return nil
}

func (m *GithubIssueMutation) GetLocked() *BoolChange {
	if m != nil {
		return m.Locked
	}
	return nil
}

func (m *GithubIssueMutation) GetPullRequest() bool {
	if m != nil {
		return m.PullRequest
	}
	return false
}

func (m *GithubIssueMutation) GetClosedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.ClosedAt
	}
	return nil
}

func (m *GithubIssueMutation) GetClosedBy() *GithubUser {
	if m != nil {
		return m.ClosedBy
	}
	return nil
}

func (m *GithubIssueMutation) GetRemoveLabel() []int64 {
	if m != nil {
		return m.RemoveLabel
	}
	return nil
}

func (m *GithubIssueMutation) GetAddLabel() []*GithubLabel {
	if m != nil {
		return m.AddLabel
	}
	return nil
}

func (m *GithubIssueMutation) GetComment() []*GithubIssueCommentMutation {
	if m != nil {
		return m.Comment
	}
	return nil
}

func (m *GithubIssueMutation) GetCommentStatus() *GithubIssueSyncStatus {
	if m != nil {
		return m.CommentStatus
	}
	return nil
}

func (m *GithubIssueMutation) GetEvent() []*GithubIssueEvent {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *GithubIssueMutation) GetEventStatus() *GithubIssueSyncStatus {
	if m != nil {
		return m.EventStatus
	}
	return nil
}

// BoolChange represents a change to a boolean value.
type BoolChange struct {
	Val bool `protobuf:"varint,1,opt,name=val" json:"val,omitempty"`
}

func (m *BoolChange) Reset()                    { *m = BoolChange{} }
func (m *BoolChange) String() string            { return proto.CompactTextString(m) }
func (*BoolChange) ProtoMessage()               {}
func (*BoolChange) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *BoolChange) GetVal() bool {
	if m != nil {
		return m.Val
	}
	return false
}

type GithubLabel struct {
	Id   int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *GithubLabel) Reset()                    { *m = GithubLabel{} }
func (m *GithubLabel) String() string            { return proto.CompactTextString(m) }
func (*GithubLabel) ProtoMessage()               {}
func (*GithubLabel) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GithubLabel) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GithubLabel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GithubMilestone struct {
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// Following only need to be non-zero on changes:
	Title  string      `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Closed *BoolChange `protobuf:"bytes,3,opt,name=closed" json:"closed,omitempty"`
	Number int64       `protobuf:"varint,4,opt,name=number" json:"number,omitempty"`
}

func (m *GithubMilestone) Reset()                    { *m = GithubMilestone{} }
func (m *GithubMilestone) String() string            { return proto.CompactTextString(m) }
func (*GithubMilestone) ProtoMessage()               {}
func (*GithubMilestone) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GithubMilestone) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GithubMilestone) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *GithubMilestone) GetClosed() *BoolChange {
	if m != nil {
		return m.Closed
	}
	return nil
}

func (m *GithubMilestone) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

// See https://developer.github.com/v3/activity/events/types/#issuesevent
// for some info.
type GithubIssueEvent struct {
	// Required:
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// event_type can be one of "assigned", "unassigned", "labeled",
	// "unlabeled", "opened", "edited", "milestoned", "demilestoned",
	// "closed", "reopened", "referenced", "renamed" or anything else
	// that Github adds in the future.
	EventType string                     `protobuf:"bytes,2,opt,name=event_type,json=eventType" json:"event_type,omitempty"`
	ActorId   int64                      `protobuf:"varint,3,opt,name=actor_id,json=actorId" json:"actor_id,omitempty"`
	Created   *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=created" json:"created,omitempty"`
	// label is populated for "labeled" and "unlabeled" events.
	// The label will usually not have an ID, due to Github's API
	// not returning one.
	Label *GithubLabel `protobuf:"bytes,5,opt,name=label" json:"label,omitempty"`
	// milestone is populated for "milestoned" and "demilestoned" events.
	// The label will usually not have an ID, due to Github's API
	// not returning one.
	Milestone *GithubMilestone `protobuf:"bytes,6,opt,name=milestone" json:"milestone,omitempty"`
	// For "assigned", "unassigned":
	AssigneeId int64 `protobuf:"varint,7,opt,name=assignee_id,json=assigneeId" json:"assignee_id,omitempty"`
	AssignerId int64 `protobuf:"varint,8,opt,name=assigner_id,json=assignerId" json:"assigner_id,omitempty"`
	// For "referenced", "closed":
	Commit *GithubCommit `protobuf:"bytes,9,opt,name=commit" json:"commit,omitempty"`
	// For "renamed" events:
	RenameFrom        string `protobuf:"bytes,11,opt,name=rename_from,json=renameFrom" json:"rename_from,omitempty"`
	RenameTo          string `protobuf:"bytes,12,opt,name=rename_to,json=renameTo" json:"rename_to,omitempty"`
	ReviewerId        int64  `protobuf:"varint,13,opt,name=reviewer_id,json=reviewerId" json:"reviewer_id,omitempty"`
	ReviewRequesterId int64  `protobuf:"varint,14,opt,name=review_requester_id,json=reviewRequesterId" json:"review_requester_id,omitempty"`
	// Contents of a dismissed review event, see dismissed_review in
	// https://developer.github.com/v3/issues/events/ for more info
	DismissedReview *GithubDismissedReviewEvent `protobuf:"bytes,15,opt,name=dismissed_review,json=dismissedReview" json:"dismissed_review,omitempty"`
	// other_json is usually empty. If Github adds event types or fields
	// in the future, this captures those added fields. If non-empty it
	// will be a JSON object with the fields that weren't understood.
	OtherJson []byte `protobuf:"bytes,10,opt,name=other_json,json=otherJson,proto3" json:"other_json,omitempty"`
}

func (m *GithubIssueEvent) Reset()                    { *m = GithubIssueEvent{} }
func (m *GithubIssueEvent) String() string            { return proto.CompactTextString(m) }
func (*GithubIssueEvent) ProtoMessage()               {}
func (*GithubIssueEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *GithubIssueEvent) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GithubIssueEvent) GetEventType() string {
	if m != nil {
		return m.EventType
	}
	return ""
}

func (m *GithubIssueEvent) GetActorId() int64 {
	if m != nil {
		return m.ActorId
	}
	return 0
}

func (m *GithubIssueEvent) GetCreated() *google_protobuf.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *GithubIssueEvent) GetLabel() *GithubLabel {
	if m != nil {
		return m.Label
	}
	return nil
}

func (m *GithubIssueEvent) GetMilestone() *GithubMilestone {
	if m != nil {
		return m.Milestone
	}
	return nil
}

func (m *GithubIssueEvent) GetAssigneeId() int64 {
	if m != nil {
		return m.AssigneeId
	}
	return 0
}

func (m *GithubIssueEvent) GetAssignerId() int64 {
	if m != nil {
		return m.AssignerId
	}
	return 0
}

func (m *GithubIssueEvent) GetCommit() *GithubCommit {
	if m != nil {
		return m.Commit
	}
	return nil
}

func (m *GithubIssueEvent) GetRenameFrom() string {
	if m != nil {
		return m.RenameFrom
	}
	return ""
}

func (m *GithubIssueEvent) GetRenameTo() string {
	if m != nil {
		return m.RenameTo
	}
	return ""
}

func (m *GithubIssueEvent) GetReviewerId() int64 {
	if m != nil {
		return m.ReviewerId
	}
	return 0
}

func (m *GithubIssueEvent) GetReviewRequesterId() int64 {
	if m != nil {
		return m.ReviewRequesterId
	}
	return 0
}

func (m *GithubIssueEvent) GetDismissedReview() *GithubDismissedReviewEvent {
	if m != nil {
		return m.DismissedReview
	}
	return nil
}

func (m *GithubIssueEvent) GetOtherJson() []byte {
	if m != nil {
		return m.OtherJson
	}
	return nil
}

// Contents of a dismissed review event - when someone leaves a
// review requesting changes and someone else dismisses it. See
// https://developer.github.com/v3/issues/events for more information.
type GithubDismissedReviewEvent struct {
	ReviewId         int64  `protobuf:"varint,1,opt,name=review_id,json=reviewId" json:"review_id,omitempty"`
	DismissalMessage string `protobuf:"bytes,3,opt,name=dismissal_message,json=dismissalMessage" json:"dismissal_message,omitempty"`
	State            string `protobuf:"bytes,4,opt,name=state" json:"state,omitempty"`
}

func (m *GithubDismissedReviewEvent) Reset()                    { *m = GithubDismissedReviewEvent{} }
func (m *GithubDismissedReviewEvent) String() string            { return proto.CompactTextString(m) }
func (*GithubDismissedReviewEvent) ProtoMessage()               {}
func (*GithubDismissedReviewEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *GithubDismissedReviewEvent) GetReviewId() int64 {
	if m != nil {
		return m.ReviewId
	}
	return 0
}

func (m *GithubDismissedReviewEvent) GetDismissalMessage() string {
	if m != nil {
		return m.DismissalMessage
	}
	return ""
}

func (m *GithubDismissedReviewEvent) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

type GithubCommit struct {
	Owner    string `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	Repo     string `protobuf:"bytes,2,opt,name=repo" json:"repo,omitempty"`
	CommitId string `protobuf:"bytes,3,opt,name=commit_id,json=commitId" json:"commit_id,omitempty"`
}

func (m *GithubCommit) Reset()                    { *m = GithubCommit{} }
func (m *GithubCommit) String() string            { return proto.CompactTextString(m) }
func (*GithubCommit) ProtoMessage()               {}
func (*GithubCommit) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *GithubCommit) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *GithubCommit) GetRepo() string {
	if m != nil {
		return m.Repo
	}
	return ""
}

func (m *GithubCommit) GetCommitId() string {
	if m != nil {
		return m.CommitId
	}
	return ""
}

// GithubIssueSyncStatus notes where syncing is at for comments
// on an issue,
// This mutation type is only made at/after the same top-level mutation
// which created the corresponding comments.
type GithubIssueSyncStatus struct {
	// server_date is the "Date" response header from Github for the
	// final HTTP response.
	ServerDate *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=server_date,json=serverDate" json:"server_date,omitempty"`
}

func (m *GithubIssueSyncStatus) Reset()                    { *m = GithubIssueSyncStatus{} }
func (m *GithubIssueSyncStatus) String() string            { return proto.CompactTextString(m) }
func (*GithubIssueSyncStatus) ProtoMessage()               {}
func (*GithubIssueSyncStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *GithubIssueSyncStatus) GetServerDate() *google_protobuf.Timestamp {
	if m != nil {
		return m.ServerDate
	}
	return nil
}

type GithubIssueCommentMutation struct {
	Id      int64                      `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	User    *GithubUser                `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
	Body    string                     `protobuf:"bytes,3,opt,name=body" json:"body,omitempty"`
	Created *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=created" json:"created,omitempty"`
	Updated *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=updated" json:"updated,omitempty"`
}

func (m *GithubIssueCommentMutation) Reset()                    { *m = GithubIssueCommentMutation{} }
func (m *GithubIssueCommentMutation) String() string            { return proto.CompactTextString(m) }
func (*GithubIssueCommentMutation) ProtoMessage()               {}
func (*GithubIssueCommentMutation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *GithubIssueCommentMutation) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GithubIssueCommentMutation) GetUser() *GithubUser {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *GithubIssueCommentMutation) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *GithubIssueCommentMutation) GetCreated() *google_protobuf.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *GithubIssueCommentMutation) GetUpdated() *google_protobuf.Timestamp {
	if m != nil {
		return m.Updated
	}
	return nil
}

type GithubUser struct {
	Id    int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Login string `protobuf:"bytes,2,opt,name=login" json:"login,omitempty"`
}

func (m *GithubUser) Reset()                    { *m = GithubUser{} }
func (m *GithubUser) String() string            { return proto.CompactTextString(m) }
func (*GithubUser) ProtoMessage()               {}
func (*GithubUser) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *GithubUser) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GithubUser) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

type GitMutation struct {
	Repo *GitRepo `protobuf:"bytes,1,opt,name=repo" json:"repo,omitempty"`
	// commit adds a commit, or adds new information to a commit if fields
	// are added in the future.
	Commit *GitCommit `protobuf:"bytes,2,opt,name=commit" json:"commit,omitempty"`
}

func (m *GitMutation) Reset()                    { *m = GitMutation{} }
func (m *GitMutation) String() string            { return proto.CompactTextString(m) }
func (*GitMutation) ProtoMessage()               {}
func (*GitMutation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *GitMutation) GetRepo() *GitRepo {
	if m != nil {
		return m.Repo
	}
	return nil
}

func (m *GitMutation) GetCommit() *GitCommit {
	if m != nil {
		return m.Commit
	}
	return nil
}

// GitRepo identifies a git repo being mutated.
type GitRepo struct {
	// If go_repo is set, it identifies a go.googlesource.com/<go_repo> repo.
	GoRepo string `protobuf:"bytes,1,opt,name=go_repo,json=goRepo" json:"go_repo,omitempty"`
}

func (m *GitRepo) Reset()                    { *m = GitRepo{} }
func (m *GitRepo) String() string            { return proto.CompactTextString(m) }
func (*GitRepo) ProtoMessage()               {}
func (*GitRepo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *GitRepo) GetGoRepo() string {
	if m != nil {
		return m.GoRepo
	}
	return ""
}

type GitCommit struct {
	Sha1 string `protobuf:"bytes,1,opt,name=sha1" json:"sha1,omitempty"`
	// raw is the "git cat-file commit $sha1" output.
	Raw      []byte       `protobuf:"bytes,2,opt,name=raw,proto3" json:"raw,omitempty"`
	DiffTree *GitDiffTree `protobuf:"bytes,3,opt,name=diff_tree,json=diffTree" json:"diff_tree,omitempty"`
}

func (m *GitCommit) Reset()                    { *m = GitCommit{} }
func (m *GitCommit) String() string            { return proto.CompactTextString(m) }
func (*GitCommit) ProtoMessage()               {}
func (*GitCommit) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *GitCommit) GetSha1() string {
	if m != nil {
		return m.Sha1
	}
	return ""
}

func (m *GitCommit) GetRaw() []byte {
	if m != nil {
		return m.Raw
	}
	return nil
}

func (m *GitCommit) GetDiffTree() *GitDiffTree {
	if m != nil {
		return m.DiffTree
	}
	return nil
}

// git diff-tree --numstat oldtree newtree
type GitDiffTree struct {
	File []*GitDiffTreeFile `protobuf:"bytes,1,rep,name=file" json:"file,omitempty"`
}

func (m *GitDiffTree) Reset()                    { *m = GitDiffTree{} }
func (m *GitDiffTree) String() string            { return proto.CompactTextString(m) }
func (*GitDiffTree) ProtoMessage()               {}
func (*GitDiffTree) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *GitDiffTree) GetFile() []*GitDiffTreeFile {
	if m != nil {
		return m.File
	}
	return nil
}

// GitDiffTreeFile represents one line of `git diff-tree --numstat` output.
type GitDiffTreeFile struct {
	File    string `protobuf:"bytes,1,opt,name=file" json:"file,omitempty"`
	Added   int64  `protobuf:"varint,2,opt,name=added" json:"added,omitempty"`
	Deleted int64  `protobuf:"varint,3,opt,name=deleted" json:"deleted,omitempty"`
	Binary  bool   `protobuf:"varint,4,opt,name=binary" json:"binary,omitempty"`
}

func (m *GitDiffTreeFile) Reset()                    { *m = GitDiffTreeFile{} }
func (m *GitDiffTreeFile) String() string            { return proto.CompactTextString(m) }
func (*GitDiffTreeFile) ProtoMessage()               {}
func (*GitDiffTreeFile) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *GitDiffTreeFile) GetFile() string {
	if m != nil {
		return m.File
	}
	return ""
}

func (m *GitDiffTreeFile) GetAdded() int64 {
	if m != nil {
		return m.Added
	}
	return 0
}

func (m *GitDiffTreeFile) GetDeleted() int64 {
	if m != nil {
		return m.Deleted
	}
	return 0
}

func (m *GitDiffTreeFile) GetBinary() bool {
	if m != nil {
		return m.Binary
	}
	return false
}

type GerritMutation struct {
	// Project is the Gerrit server and project, without scheme (https implied) or
	// trailing slash.
	Project string `protobuf:"bytes,1,opt,name=project" json:"project,omitempty"`
	// Commits to add.
	Commits []*GitCommit `protobuf:"bytes,2,rep,name=commits" json:"commits,omitempty"`
	// git refs to update.
	Refs []*GitRef `protobuf:"bytes,3,rep,name=refs" json:"refs,omitempty"`
}

func (m *GerritMutation) Reset()                    { *m = GerritMutation{} }
func (m *GerritMutation) String() string            { return proto.CompactTextString(m) }
func (*GerritMutation) ProtoMessage()               {}
func (*GerritMutation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *GerritMutation) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *GerritMutation) GetCommits() []*GitCommit {
	if m != nil {
		return m.Commits
	}
	return nil
}

func (m *GerritMutation) GetRefs() []*GitRef {
	if m != nil {
		return m.Refs
	}
	return nil
}

type GitRef struct {
	// ref is the git ref name, such as:
	//    HEAD
	//    refs/heads/master
	//    refs/changes/00/14700/1
	//    refs/changes/00/14700/meta
	//    refs/meta/config
	Ref string `protobuf:"bytes,1,opt,name=ref" json:"ref,omitempty"`
	// sha1 is the lowercase hex sha1
	Sha1 string `protobuf:"bytes,2,opt,name=sha1" json:"sha1,omitempty"`
}

func (m *GitRef) Reset()                    { *m = GitRef{} }
func (m *GitRef) String() string            { return proto.CompactTextString(m) }
func (*GitRef) ProtoMessage()               {}
func (*GitRef) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

func (m *GitRef) GetRef() string {
	if m != nil {
		return m.Ref
	}
	return ""
}

func (m *GitRef) GetSha1() string {
	if m != nil {
		return m.Sha1
	}
	return ""
}

func init() {
	proto.RegisterType((*Mutation)(nil), "maintpb.Mutation")
	proto.RegisterType((*GithubMutation)(nil), "maintpb.GithubMutation")
	proto.RegisterType((*GithubIssueMutation)(nil), "maintpb.GithubIssueMutation")
	proto.RegisterType((*BoolChange)(nil), "maintpb.BoolChange")
	proto.RegisterType((*GithubLabel)(nil), "maintpb.GithubLabel")
	proto.RegisterType((*GithubMilestone)(nil), "maintpb.GithubMilestone")
	proto.RegisterType((*GithubIssueEvent)(nil), "maintpb.GithubIssueEvent")
	proto.RegisterType((*GithubDismissedReviewEvent)(nil), "maintpb.GithubDismissedReviewEvent")
	proto.RegisterType((*GithubCommit)(nil), "maintpb.GithubCommit")
	proto.RegisterType((*GithubIssueSyncStatus)(nil), "maintpb.GithubIssueSyncStatus")
	proto.RegisterType((*GithubIssueCommentMutation)(nil), "maintpb.GithubIssueCommentMutation")
	proto.RegisterType((*GithubUser)(nil), "maintpb.GithubUser")
	proto.RegisterType((*GitMutation)(nil), "maintpb.GitMutation")
	proto.RegisterType((*GitRepo)(nil), "maintpb.GitRepo")
	proto.RegisterType((*GitCommit)(nil), "maintpb.GitCommit")
	proto.RegisterType((*GitDiffTree)(nil), "maintpb.GitDiffTree")
	proto.RegisterType((*GitDiffTreeFile)(nil), "maintpb.GitDiffTreeFile")
	proto.RegisterType((*GerritMutation)(nil), "maintpb.GerritMutation")
	proto.RegisterType((*GitRef)(nil), "maintpb.GitRef")
}

func init() { proto.RegisterFile("maintner.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x57, 0xdb, 0x8e, 0xdb, 0x36,
	0x13, 0x86, 0xcf, 0xd2, 0xd8, 0xeb, 0x75, 0xb8, 0x39, 0x30, 0x9b, 0xfc, 0xf9, 0xf7, 0x57, 0x7e,
	0x34, 0x8b, 0x24, 0xf5, 0x36, 0x69, 0xd1, 0x16, 0x08, 0x8a, 0x22, 0xc7, 0xc2, 0x41, 0x93, 0x0b,
	0x66, 0x73, 0x2d, 0xc8, 0x26, 0xed, 0x55, 0x2a, 0x89, 0xae, 0x48, 0x6f, 0x6a, 0xa0, 0xbd, 0xe9,
	0xa3, 0xf4, 0x9d, 0xfa, 0x08, 0xbd, 0xee, 0x2b, 0x14, 0x1c, 0x52, 0x92, 0xed, 0xf5, 0x66, 0x93,
	0xde, 0x71, 0x66, 0xbe, 0xa1, 0xe6, 0x3c, 0x14, 0xf4, 0xd3, 0x28, 0xce, 0x74, 0x26, 0xf2, 0xe1,
	0x3c, 0x97, 0x5a, 0x92, 0x0e, 0xd2, 0xf3, 0xf1, 0xfe, 0xa3, 0x59, 0xac, 0x4f, 0x16, 0xe3, 0xe1,
	0x44, 0xa6, 0x47, 0x33, 0x99, 0x44, 0xd9, 0xec, 0x08, 0x11, 0xe3, 0xc5, 0xf4, 0x68, 0xae, 0x97,
	0x73, 0xa1, 0x8e, 0x74, 0x9c, 0x0a, 0xa5, 0xa3, 0x74, 0x5e, 0x9d, 0xec, 0x2d, 0xc1, 0x9f, 0x35,
	0xf0, 0x5e, 0x2d, 0x74, 0xa4, 0x63, 0x99, 0x91, 0xef, 0xa1, 0x67, 0xef, 0x0a, 0x63, 0xa5, 0x16,
	0x82, 0xd6, 0x0e, 0x6a, 0x87, 0xdd, 0x87, 0x37, 0x87, 0xee, 0x4b, 0xc3, 0x1f, 0x50, 0x38, 0x32,
	0xb2, 0x42, 0x87, 0x75, 0x67, 0x15, 0x93, 0x1c, 0x41, 0xdb, 0x92, 0xb4, 0x81, 0xaa, 0xd7, 0x36,
	0x54, 0x4b, 0x2d, 0x07, 0x23, 0x9f, 0x41, 0x63, 0x16, 0x6b, 0x5a, 0x47, 0xf4, 0xe5, 0x55, 0x74,
	0x09, 0x35, 0x00, 0xbc, 0x58, 0xe4, 0x79, 0xac, 0x69, 0x73, 0xf3, 0x62, 0x64, 0xaf, 0x5c, 0x8c,
	0x74, 0xf0, 0x47, 0x0d, 0xfa, 0xeb, 0xdf, 0x24, 0x97, 0xa1, 0x25, 0xdf, 0x67, 0x22, 0x47, 0xb7,
	0x7c, 0x66, 0x09, 0x42, 0xa0, 0x99, 0x8b, 0xb9, 0x44, 0x13, 0x7c, 0x86, 0x67, 0x72, 0x1f, 0xda,
	0x49, 0x34, 0x16, 0x89, 0xa2, 0x8d, 0x83, 0xc6, 0xa6, 0x61, 0x27, 0x8b, 0xf1, 0x8f, 0x46, 0xc8,
	0x1c, 0x86, 0x7c, 0x0b, 0x90, 0xc6, 0x89, 0x50, 0x5a, 0x66, 0x42, 0xd1, 0x26, 0x6a, 0xd0, 0x4d,
	0xc7, 0x0b, 0x00, 0x5b, 0xc1, 0x06, 0x7f, 0x7b, 0xb0, 0xb7, 0x25, 0xa6, 0x9f, 0x60, 0xe9, 0x55,
	0x68, 0x67, 0x8b, 0x74, 0x2c, 0x72, 0x0c, 0x78, 0x8b, 0x39, 0x8a, 0xdc, 0x00, 0x3f, 0x93, 0x3a,
	0x14, 0xbf, 0xc4, 0x4a, 0xd3, 0x9d, 0x83, 0xda, 0xa1, 0xc7, 0xbc, 0x4c, 0xea, 0xe7, 0x86, 0x26,
	0x7d, 0xa8, 0xc7, 0x9c, 0xf6, 0x0e, 0x6a, 0x87, 0x0d, 0x56, 0x8f, 0x39, 0xb9, 0x03, 0xcd, 0x85,
	0x12, 0xb9, 0x0b, 0xed, 0xde, 0x86, 0xe9, 0x6f, 0x95, 0xc8, 0x19, 0x02, 0xc8, 0x03, 0xf0, 0x23,
	0xa5, 0xe2, 0x59, 0x26, 0x84, 0xa2, 0x80, 0x8e, 0x6e, 0x45, 0x57, 0x28, 0x72, 0x0f, 0x2e, 0x71,
	0x91, 0x08, 0x2d, 0x78, 0x58, 0xa9, 0x76, 0x0f, 0x1a, 0x87, 0x0d, 0x36, 0x70, 0x82, 0xc7, 0x25,
	0xf8, 0x2b, 0xe8, 0x4c, 0x72, 0x11, 0x69, 0xc1, 0x69, 0x0b, 0x6d, 0xd9, 0x1f, 0xce, 0xa4, 0x9c,
	0x25, 0x62, 0x58, 0x14, 0xf4, 0xf0, 0xb8, 0xa8, 0x5f, 0x56, 0x40, 0x8d, 0xd6, 0x62, 0xce, 0x51,
	0xab, 0x7d, 0xb1, 0x96, 0x83, 0x9a, 0x68, 0x8e, 0x25, 0x5f, 0xd2, 0x8e, 0x8d, 0xa6, 0x39, 0x9b,
	0xb8, 0xeb, 0x58, 0x27, 0x82, 0xfa, 0x36, 0xee, 0x48, 0x90, 0xff, 0x41, 0x2f, 0x93, 0x61, 0x99,
	0x36, 0xba, 0x8b, 0xe1, 0xec, 0x66, 0xb2, 0x4c, 0xaa, 0x81, 0x94, 0xf2, 0x30, 0xe6, 0x74, 0x80,
	0xb1, 0xed, 0x96, 0xbc, 0x11, 0x27, 0xb7, 0x61, 0xa7, 0x82, 0x64, 0x8b, 0x94, 0x5e, 0x42, 0x4c,
	0xa5, 0xf7, 0x7a, 0x91, 0x92, 0x3b, 0xb0, 0x5b, 0x81, 0xac, 0x29, 0x04, 0x4d, 0xe9, 0x97, 0xec,
	0x63, 0xb4, 0xe9, 0x1e, 0xb4, 0x27, 0x89, 0x54, 0x82, 0xd3, 0xbd, 0x8d, 0xa4, 0x3d, 0x91, 0x32,
	0x79, 0x7a, 0x12, 0x65, 0x33, 0xc1, 0x1c, 0xc4, 0x80, 0x13, 0x39, 0xf9, 0x49, 0x70, 0x7a, 0xfd,
	0x03, 0x60, 0x0b, 0x31, 0xae, 0xcc, 0x17, 0x49, 0x12, 0xe6, 0xe2, 0xe7, 0x85, 0x50, 0x9a, 0xde,
	0xb4, 0xde, 0x1a, 0x1e, 0xb3, 0x2c, 0xf2, 0x0d, 0xf8, 0xf6, 0xe6, 0x30, 0xd2, 0xf4, 0xca, 0x85,
	0x21, 0xf7, 0x2c, 0xf8, 0xb1, 0x26, 0x5f, 0x94, 0x8a, 0xe3, 0x25, 0xbd, 0x7a, 0x7e, 0xb5, 0x39,
	0x8d, 0x27, 0x4b, 0x63, 0x4d, 0x2e, 0x52, 0x79, 0x2a, 0x42, 0x6c, 0x36, 0x7a, 0x0d, 0x2b, 0xa7,
	0x6b, 0x79, 0xd8, 0x86, 0x58, 0x94, 0x9c, 0x3b, 0x39, 0xfd, 0x40, 0xbf, 0x7a, 0x11, 0xe7, 0x56,
	0xe5, 0x3b, 0xe8, 0x4c, 0x64, 0x9a, 0x8a, 0x4c, 0x53, 0x0f, 0x15, 0x6e, 0x6f, 0x1b, 0x71, 0x4f,
	0x2d, 0xa4, 0x1c, 0x2d, 0x85, 0x0e, 0x79, 0x0e, 0x7d, 0x77, 0x0c, 0x95, 0x8e, 0xf4, 0x42, 0xd1,
	0x3e, 0xfa, 0x72, 0x6b, 0xdb, 0x2d, 0x6f, 0x96, 0xd9, 0xe4, 0x0d, 0xa2, 0xd8, 0x8e, 0xd3, 0xb2,
	0x24, 0x39, 0x82, 0x96, 0x38, 0x35, 0x36, 0xec, 0xa3, 0x0d, 0xd7, 0xb7, 0x69, 0x3f, 0x37, 0x00,
	0x66, 0x71, 0xe4, 0x31, 0xf4, 0xf0, 0x50, 0x7c, 0xf5, 0xc6, 0x47, 0x7d, 0xb5, 0x8b, 0x3a, 0x96,
	0x08, 0x6e, 0x01, 0x54, 0x39, 0x27, 0x03, 0x68, 0x9c, 0x46, 0x09, 0x4e, 0x19, 0x8f, 0x99, 0x63,
	0xf0, 0x00, 0xba, 0x2b, 0x21, 0x73, 0x93, 0xa2, 0x56, 0x4e, 0x0a, 0x02, 0xcd, 0x2c, 0x4a, 0x45,
	0x31, 0x82, 0xcc, 0x39, 0xf8, 0x15, 0x76, 0x37, 0x66, 0xdc, 0x19, 0xb5, 0xb2, 0xaf, 0xea, 0xab,
	0x7d, 0x55, 0xd5, 0x70, 0xe3, 0xe2, 0x1a, 0xae, 0x06, 0x5d, 0x13, 0xaf, 0x75, 0x54, 0xf0, 0x57,
	0x13, 0x06, 0x9b, 0xf1, 0x3a, 0xf3, 0xfd, 0xff, 0x00, 0xd8, 0xc0, 0x99, 0x6d, 0xe8, 0x8c, 0xf0,
	0x91, 0x73, 0xbc, 0x9c, 0x0b, 0x72, 0x1d, 0xbc, 0x68, 0xa2, 0x65, 0x6e, 0x3a, 0xb7, 0x81, 0x4a,
	0x1d, 0xa4, 0x47, 0x7c, 0x75, 0x22, 0x35, 0x3f, 0x7e, 0x22, 0xdd, 0x85, 0x96, 0x2d, 0xc7, 0xd6,
	0xd9, 0xbd, 0x56, 0x96, 0xa3, 0x85, 0x90, 0xaf, 0xc1, 0xaf, 0x46, 0x8b, 0x9d, 0x5f, 0xe7, 0x2f,
	0x8f, 0x0a, 0x4a, 0xfe, 0x0b, 0xdd, 0x62, 0xa0, 0x1a, 0xbb, 0x3b, 0x68, 0x37, 0x14, 0xac, 0x11,
	0x5f, 0x01, 0xa0, 0x63, 0xde, 0x1a, 0xc0, 0xf8, 0xf6, 0x39, 0xb4, 0x4d, 0x41, 0xc6, 0x1a, 0xc7,
	0x5d, 0xf7, 0xe1, 0x95, 0x8d, 0xcf, 0x3e, 0x45, 0x21, 0x73, 0x20, 0x73, 0x5f, 0x2e, 0x4c, 0xc6,
	0xc3, 0x69, 0x2e, 0x53, 0xda, 0xc5, 0x28, 0x82, 0x65, 0xbd, 0xc8, 0x65, 0x6a, 0x76, 0x8e, 0x03,
	0x68, 0x89, 0xdb, 0xc5, 0x67, 0x9e, 0x65, 0x1c, 0x4b, 0xab, 0x7d, 0x1a, 0x8b, 0xf7, 0xd6, 0x9a,
	0x1d, 0x6b, 0x4d, 0xc1, 0x1a, 0x71, 0x32, 0x84, 0x3d, 0x4b, 0x15, 0x93, 0xc7, 0x02, 0xfb, 0x08,
	0xbc, 0x64, 0x45, 0xac, 0x90, 0x8c, 0x38, 0x79, 0x0d, 0x03, 0x1e, 0xab, 0x34, 0x56, 0x66, 0x9c,
	0x58, 0x31, 0x4e, 0xe6, 0xb3, 0xcd, 0xfc, 0xac, 0x80, 0x31, 0x44, 0xd9, 0x96, 0xda, 0xe5, 0xeb,
	0x5c, 0x53, 0x23, 0x52, 0x9f, 0x88, 0x3c, 0x7c, 0xa7, 0x64, 0x46, 0xe1, 0xa0, 0x76, 0xd8, 0x63,
	0x3e, 0x72, 0x5e, 0x2a, 0x99, 0x05, 0xbf, 0xd7, 0x60, 0xff, 0xfc, 0xeb, 0xac, 0xef, 0x68, 0x7d,
	0x59, 0x78, 0x9e, 0x65, 0x8c, 0x38, 0xee, 0x40, 0xab, 0x14, 0x25, 0x61, 0x2a, 0x94, 0x8a, 0x66,
	0x02, 0x0b, 0xcd, 0x67, 0x83, 0x52, 0xf0, 0xca, 0xf2, 0x4d, 0xaf, 0x98, 0xf6, 0x16, 0x58, 0x6f,
	0x3e, 0xb3, 0xc4, 0xcb, 0xa6, 0x57, 0x1f, 0x34, 0x82, 0xb7, 0xd0, 0x5b, 0x4d, 0xcd, 0x27, 0xbc,
	0x13, 0x6e, 0x80, 0x6f, 0xd3, 0x58, 0xd4, 0xb8, 0xcf, 0x3c, 0xcb, 0x18, 0xf1, 0xe0, 0x18, 0xae,
	0x6c, 0x1d, 0x1d, 0xe4, 0x11, 0x74, 0x95, 0xc8, 0x4f, 0x45, 0x1e, 0x9a, 0x9d, 0xe9, 0x9e, 0x83,
	0x1f, 0xea, 0x00, 0xb0, 0xf0, 0x67, 0x91, 0x16, 0xe6, 0x65, 0xb9, 0x7f, 0xfe, 0x34, 0x3d, 0xd3,
	0xa3, 0xc5, 0x23, 0xa4, 0x7e, 0xd1, 0x23, 0xa4, 0x58, 0xdc, 0x8d, 0x95, 0xc5, 0xfd, 0xef, 0xda,
	0x74, 0xe5, 0xe1, 0xd0, 0xfa, 0xe8, 0x87, 0x43, 0xf0, 0x10, 0xa0, 0xb2, 0x69, 0xdb, 0xa8, 0x4b,
	0xe4, 0x2c, 0xce, 0x8a, 0x51, 0x87, 0x44, 0x10, 0xe2, 0x58, 0x2d, 0x7d, 0xff, 0xbf, 0xcb, 0x90,
	0x0d, 0xe8, 0x60, 0xd5, 0x57, 0x26, 0xe6, 0xd2, 0xe5, 0xec, 0x6e, 0xd9, 0x9f, 0x36, 0x26, 0x64,
	0x15, 0xb7, 0xde, 0x9c, 0x41, 0x00, 0x1d, 0xa7, 0x4c, 0xae, 0x41, 0x67, 0x26, 0xc3, 0xf2, 0x7e,
	0x9f, 0xb5, 0x67, 0xd2, 0x08, 0x02, 0x0e, 0x7e, 0xa9, 0x68, 0xa2, 0xa8, 0x4e, 0xa2, 0x07, 0x0e,
	0x82, 0x67, 0xb3, 0x0e, 0xf2, 0xe8, 0x3d, 0x7e, 0xad, 0xc7, 0xcc, 0xd1, 0xec, 0x56, 0x1e, 0x4f,
	0xa7, 0xa1, 0xce, 0x85, 0x70, 0x53, 0x7a, 0x6d, 0x98, 0x3d, 0x8b, 0xa7, 0xd3, 0xe3, 0x5c, 0x08,
	0xe6, 0x71, 0x77, 0x0a, 0x1e, 0xa1, 0xab, 0x85, 0x80, 0xdc, 0x87, 0xe6, 0x34, 0x4e, 0x4c, 0xed,
	0x9c, 0x79, 0x16, 0x17, 0x98, 0x17, 0x71, 0x22, 0x18, 0xa2, 0x82, 0x14, 0x77, 0xc9, 0xaa, 0xc0,
	0x18, 0xea, 0x2e, 0x40, 0x43, 0xcd, 0xd9, 0x04, 0x39, 0xe2, 0x5c, 0x70, 0x34, 0xb5, 0xc1, 0x2c,
	0x41, 0x28, 0x74, 0xdc, 0x8b, 0xb2, 0x98, 0xe2, 0x8e, 0x34, 0xcb, 0x63, 0x1c, 0x67, 0x51, 0xbe,
	0xc4, 0xea, 0xf0, 0x98, 0xa3, 0x82, 0xdf, 0xa0, 0xbf, 0xfe, 0xfb, 0x60, 0xee, 0x98, 0xe7, 0xf2,
	0x9d, 0x98, 0x68, 0xf7, 0xc1, 0x82, 0x24, 0xf7, 0xed, 0x9b, 0x21, 0xd6, 0x8a, 0xd6, 0xd1, 0x97,
	0x6d, 0xe9, 0x28, 0x20, 0xe4, 0xb6, 0xc9, 0xf0, 0xb4, 0xf8, 0x7f, 0xd8, 0x5d, 0xcf, 0xf0, 0x94,
	0xa1, 0x30, 0x18, 0x42, 0xdb, 0xd2, 0x18, 0x79, 0x31, 0x75, 0x9f, 0x34, 0xc7, 0x32, 0x3f, 0xf5,
	0x2a, 0x3f, 0xe3, 0x36, 0x96, 0xe5, 0x97, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xbb, 0x24, 0x15,
	0xd7, 0x0a, 0x0e, 0x00, 0x00,
}