package mysql

import (
	"gin-bluebell/models"
	"strings"

	"github.com/jmoiron/sqlx"
)

func CreatePost(p *models.Post) (err error) {
	sqlStr := `INSERT INTO  post(post_id, title, content, author_id, community_id) VALUES (?, ?, ?, ?, ?)`
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return
}

// GetPostDetailById
func GetPostDetailById(pid int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := `SELECT post_id, title, content, author_id, community_id, create_time from post WHERE post_id = ?`
	err = db.Get(post, sqlStr, pid)
	return
}

func GetPostList(page, pageSize int64) (posts []*models.Post, err error) {
	sqlStr := `SELECT post_id, title, content, author_id, community_id, create_time 
			   FROM post ORDER BY create_time DESC LIMIT ?, ?`
	posts = make([]*models.Post, 0, 2)
	err = db.Select(&posts, sqlStr, (page-1)*pageSize, pageSize)
	return
}

// GetPostListByIDs 根据给定的 id 列表查询帖子数据
func GetPostListByIDs(ids []string) (postList []*models.Post, err error) {
	sqlStr := `SELECT post_id, title, content, author_id, community_id, create_time
	FROM post WHERE post_id IN (?)
	ORDER BY FIND_IN_SET(post_id, ?)`

	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}

	query = db.Rebind(query)
	err = db.Select(&postList, query, args...)
	return
}
