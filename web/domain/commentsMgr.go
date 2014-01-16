package domain

import (
	"github.com/m4tty/palaver/data/comments"
	"github.com/m4tty/palaver/web/resources"
	"time"
)

type CommentsMgr struct {
	CommentsDataManager *data.CommentDataManager
}

func (dm CommentsMgr) GetCommentById(id string) (comment *resources.CommentResource, err error) {
	dComment, err := dm.GetCommentById(id)
	if err != nil {
		return nil, err
	}
	var commentResource *resources.CommentResource = new(resources.CommentResource)
	mapDataToResource(*dComment, commentResource)
	return commentResource, nil
}

func (dm CommentsMgr) GetComments() (bundles []*resources.CommentResource, err error) {
	dComments, err := dm.GetComments()
	if err != nil {
		return nil, err
	}

	var comments []*resources.CommentResource
	comments = make([]*resources.CommentResource, len(dComments))
	for j, comment := range dComments {
		var commentResource *resources.CommentResource = new(resources.CommentResource)
		comments[j] = mapDataToResource(*dComment, commentResource)
	}

	return comments, nil
}

func (dm CommentsMgr) SaveComment(bundle *resources.CommentResource) (key string, err error) {
	return
}

func (dm CommentsMgr) DeleteComment(id string) (err error) {
	return
}

// mapper...
func mapResourceToData(commentResource *resources.CommentResource, commentData *data.Comment) {
	commentData.Id = commentResource.Id
	commentData.Text = commentResource.Text
	commentData.CreatedDate = commentResource.CreatedDate
	commentData.LastModified = time.Now().UTC()
	commentData.TargetURN = commentResource.TargetURN
	commentData.ParentURN = commentResource.ParentURN
	commentData.Likes = commentResource.Likes
	commentData.Dislikes = commentResource.Dislikes
	commentData.LikedBy = commentResource.LikedBy
	commentData.DislikedBy = commentResource.DislikedBy

	var a *data.Author = new(data.Author)
	commentData.Author = *a
	commentData.Author.Id = commentResource.Author.Id
	commentData.Author.DisplayName = commentResource.Author.DisplayName
	commentData.Author.Email = commentResource.Author.Email
	commentData.Author.ProfileUrl = commentResource.Author.ProfileUrl

	var av *data.Avatar = new(data.Avatar)
	commentData.Author.Avatar = *av
	commentData.Author.Avatar.Url = commentResource.Author.Avatar.Url
}

func mapDataToResource(commentData *data.Comment, commentResource *resources.CommentResource) {
	commentResource.Id = commentData.Id

	commentResource.Text = commentData.Text
	commentResource.CreatedDate = commentData.CreatedDate
	commentResource.LastModified = commentData.LastModified
	commentResource.TargetURN = commentData.TargetURN
	commentResource.ParentURN = commentData.ParentURN
	commentResource.Likes = commentData.Likes
	commentResource.Dislikes = commentData.Dislikes
	commentResource.LikedBy = commentData.LikedBy
	commentResource.DislikedBy = commentData.DislikedBy

	var a *resource.Author = new(resource.Author)
	commentResource.Author = *a
	commentResource.Author.Id = commentData.Author.Id
	commentResource.Author.DisplayName = commentData.Author.DisplayName
	commentResource.Author.Email = commentData.Author.Email
	commentResource.Author.ProfileUrl = commentData.Author.ProfileUrl

	var av *resource.Avatar = new(resource.Avatar)
	commentResource.Author.Avatar = *av
	commentResource.Author.Avatar.Url = commentData.Author.Avatar.Url
}
