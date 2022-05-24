// Copyright 2019 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package webhook

import (
	"testing"

	webhook_model "code.gitea.io/gitea/models/webhook"
	api "code.gitea.io/gitea/modules/structs"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSlackPayload(t *testing.T) {
	t.Run("Create", func(t *testing.T) {
		p := createTestPayload()

		d := new(SlackPayload)
		pl, err := d.Create(p)
		require.NoError(t, err)
		require.NotNil(t, pl)
		require.IsType(t, &SlackPayload{}, pl)

		assert.Equal(t, "[<http://localhost:3000/test/repo|test/repo>:<http://localhost:3000/test/repo/src/branch/test|test>] branch created by user1", pl.(*SlackPayload).Text)
	})

	t.Run("Delete", func(t *testing.T) {
		p := deleteTestPayload()

		d := new(SlackPayload)
		pl, err := d.Delete(p)
		require.NoError(t, err)
		require.NotNil(t, pl)
		require.IsType(t, &SlackPayload{}, pl)

		assert.Equal(t, "[<http://localhost:3000/test/repo|test/repo>:test] branch deleted by user1", pl.(*SlackPayload).Text)
	})

	t.Run("Fork", func(t *testing.T) {
		p := forkTestPayload()

		d := new(SlackPayload)
		pl, err := d.Fork(p)
		require.NoError(t, err)
		require.NotNil(t, pl)
		require.IsType(t, &SlackPayload{}, pl)

		assert.Equal(t, "<http://localhost:3000/test/repo2|test/repo2> is forked to <http://localhost:3000/test/repo|test/repo>", pl.(*SlackPayload).Text)
	})

	t.Run("Push", func(t *testing.T) {
		p := pushTestPayload()

		d := new(SlackPayload)
		pl, err := d.Push(p)
		require.NoError(t, err)
		require.NotNil(t, pl)
		require.IsType(t, &SlackPayload{}, pl)

		assert.Equal(t, "[<http://localhost:3000/test/repo|test/repo>:<http://localhost:3000/test/repo/src/branch/test|test>] 2 new commits pushed by user1", pl.(*SlackPayload).Text)
	})

	t.Run("Issue", func(t *testing.T) {
		p := issueTestPayload()

		d := new(SlackPayload)
		p.Action = api.HookIssueOpened
		pl, err := d.Issue(p)
		require.NoError(t, err)
		require.NotNil(t, pl)
		require.IsType(t, &SlackPayload{}, pl)

		assert.Equal(t, "[<http://localhost:3000/test/repo|test/repo>] Issue opened: <http://localhost:3000/test/repo/issues/2|#2 crash> by <https://try.gitea.io/user1|user1>", pl.(*SlackPayload).Text)

		p.Action = api.HookIssueClosed
		pl, err = d.Issue(p)
		require.NoError(t, err)
		require.NotNil(t, pl)
		require.IsType(t, &SlackPayload{}, pl)

		assert.Equal(t, "[<http://localhost:3000/test/repo|test/repo>] Issue closed: <http://localhost:3000/test/repo/issues/2|#2 crash> by <https://try.gitea.io/user1|user1>", pl.(*SlackPayload).Text)
	})

	t.Run("IssueComment", func(t *testing.T) {
		p := issueCommentTestPayload()

		d := new(SlackPayload)
		pl, err := d.IssueComment(p)
		require.NoError(t, err)
		require.NotNil(t, pl)
		require.IsType(t, &SlackPayload{}, pl)

		assert.Equal(t, "[<http://localhost:3000/test/repo|test/repo>] New comment on issue <http://localhost:3000/test/repo/issues/2|#2 crash> by <https://try.gitea.io/user1|user1>", pl.(*SlackPayload).Text)
	})

	t.Run("PullRequest", func(t *testing.T) {
		p := pullRequestTestPayload()

		d := new(SlackPayload)
		pl, err := d.PullRequest(p)
		require.NoError(t, err)
		require.NotNil(t, pl)
		require.IsType(t, &SlackPayload{}, pl)

		assert.Equal(t, "[<http://localhost:3000/test/repo|test/repo>] Pull request opened: <http://localhost:3000/test/repo/pulls/12|#12 Fix bug> by <https://try.gitea.io/user1|user1>", pl.(*SlackPayload).Text)
	})

	t.Run("PullRequestComment", func(t *testing.T) {
		p := pullRequestCommentTestPayload()

		d := new(SlackPayload)
		pl, err := d.IssueComment(p)
		require.NoError(t, err)
		require.NotNil(t, pl)
		require.IsType(t, &SlackPayload{}, pl)

		assert.Equal(t, "[<http://localhost:3000/test/repo|test/repo>] New comment on pull request <http://localhost:3000/test/repo/pulls/12|#12 Fix bug> by <https://try.gitea.io/user1|user1>", pl.(*SlackPayload).Text)
	})

	t.Run("Review", func(t *testing.T) {
		p := pullRequestTestPayload()
		p.Action = api.HookIssueReviewed

		d := new(SlackPayload)
		pl, err := d.Review(p, webhook_model.HookEventPullRequestReviewApproved)
		require.NoError(t, err)
		require.NotNil(t, pl)
		require.IsType(t, &SlackPayload{}, pl)

		assert.Equal(t, "[<http://localhost:3000/test/repo|test/repo>] Pull request review approved: <http://localhost:3000/test/repo/pulls/12|#12 Fix bug> by <https://try.gitea.io/user1|user1>", pl.(*SlackPayload).Text)
	})

	t.Run("ReviewComment", func(t *testing.T) {
		p := pullRequestTestPayload()
		p.Action = api.HookIssueReviewed

		d := new(SlackPayload)
		pl, err := d.Review(p, webhook_model.HookEventPullRequestReviewComment)
		require.NoError(t, err)
		require.NotNil(t, pl)
		require.IsType(t, &SlackPayload{}, pl)

		slackPl := pl.(*SlackPayload)
		assert.Equal(t, "[<http://localhost:3000/test/repo|test/repo>] Pull request review comment: <http://localhost:3000/test/repo/pulls/12|#12 Fix bug> by <https://try.gitea.io/user1|user1>", slackPl.Text)
		assert.Len(t, slackPl.Attachments, 2)
		assert.Equal(t, p.Review.Content, slackPl.Attachments[0].Text)
		assert.Equal(t, p.Review.Comments[0].Body, slackPl.Attachments[1].Text)
	})

	t.Run("Repository", func(t *testing.T) {
		p := repositoryTestPayload()

		d := new(SlackPayload)
		pl, err := d.Repository(p)
		require.NoError(t, err)
		require.NotNil(t, pl)
		require.IsType(t, &SlackPayload{}, pl)

		assert.Equal(t, "[<http://localhost:3000/test/repo|test/repo>] Repository created by <https://try.gitea.io/user1|user1>", pl.(*SlackPayload).Text)
	})

	t.Run("Release", func(t *testing.T) {
		p := pullReleaseTestPayload()

		d := new(SlackPayload)
		pl, err := d.Release(p)
		require.NoError(t, err)
		require.NotNil(t, pl)
		require.IsType(t, &SlackPayload{}, pl)

		assert.Equal(t, "[<http://localhost:3000/test/repo|test/repo>] Release created: <http://localhost:3000/test/repo/src/v1.0|v1.0> by <https://try.gitea.io/user1|user1>", pl.(*SlackPayload).Text)
	})
}

func TestSlackJSONPayload(t *testing.T) {
	p := pushTestPayload()

	pl, err := new(SlackPayload).Push(p)
	require.NoError(t, err)
	require.NotNil(t, pl)
	require.IsType(t, &SlackPayload{}, pl)

	json, err := pl.JSONPayload()
	require.NoError(t, err)
	assert.NotEmpty(t, json)
}
