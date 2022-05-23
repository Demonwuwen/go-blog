package dao

import (
	"go-blog/models"
	"log"
)

func SavePost(post *models.Post) {
	ret, err := DB.Exec("insert into blog_post (title, content, markdown, category_id, user_id, view_count,"+
		"type, slug, create_at, update_at) values(?,?,?,?,?,?,?,?,?,?);", post.Title, post.Content, post.Markdown, post.CategoryId,
		post.UserId, post.ViewCount, post.Type, post.Slug, post.CreateAt, post.UpdateAt)
	if err != nil {
		log.Println(err)
	}
	pid, _ := ret.LastInsertId()
	post.Pid = int(pid)
	log.Println("SavePost done! pid = ",post.Pid)
}
func UpdatePost(post *models.Post) {
	_, err := DB.Exec("update  blog_post set  title=?, content=?, markdown=?, category_id=?,"+
		"type=?, slug=?, update_at=? where pid=?;", post.Title, post.Content, post.Markdown, post.CategoryId,
		 post.Type, post.Slug, post.UpdateAt, post.Pid)
	if err != nil {
		log.Println(err)
	}

	log.Println("SavePost done! pid = ",post.Pid)
}
func GetPostPageBySlug(slug string,page, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	rows, err := DB.Query("select * from blog_post where slug=? limit ?,?",slug, page, pageSize)
	if err != nil {

		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err = rows.Scan(&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func GetPostPage(page, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	rows, err := DB.Query("select * from blog_post limit ?,?", page, pageSize)
	if err != nil {

		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err = rows.Scan(&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func GetPostAll() ([]models.Post, error) {
	rows, err := DB.Query("select * from blog_post ")
	if err != nil {

		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err = rows.Scan(&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func GetPostById(pid int) (*models.Post, error) {
	p := &models.Post{}
	err := DB.QueryOne(p, "select * from blog_post where pid = ?", pid)
	return p,err
	//row := DB.QueryRow("select * from blog_post where pid = ?", pid)
	//var post models.Post
	//if row.Err() != nil {
	//	return post, row.Err()
	//}
	//err := row.Scan(
	//	&post.Pid,
	//	&post.Title,
	//	&post.Content,
	//	&post.Markdown,
	//	&post.CategoryId,
	//	&post.UserId,
	//	&post.ViewCount,
	//	&post.Type,
	//	&post.Slug,
	//	&post.CreateAt,
	//	&post.UpdateAt,
	//)
	//if err != nil {
	//	return post, err
	//}
	//return post, nil
}

func CountGetAllPostBySlug(cId int) (count int) {
	rows := DB.QueryRow("select count(1) from blog_post where slug = ?", cId)
	_ = rows.Scan(&count)
	return
}

func CountGetAllPostByCategoryId(cId int) (count int) {
	rows := DB.QueryRow("select count(1) from blog_post where category_id = ?", cId)
	rows.Scan(&count)
	return
}

func CountGetAllPost() (count int) {
	rows := DB.QueryRow("select count(1) from blog_post")
	rows.Scan(&count)
	return
}

func GetPostPageByCategoryId(cId, page, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	rows, err := DB.Query("select * from blog_post where category_id = ? limit ?,?", cId, page, pageSize)
	if err != nil {

		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err = rows.Scan(&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func GetSearchPost(val string) ([]models.SearchResp, error ){
	//page = (page - 1) * pageSize
	rows, err := DB.Query("select pid, title from blog_post where where title like", "%"+val+"%")
	if err != nil {
		return nil, err
	}
	var posts []models.SearchResp
	for rows.Next() {
		var post models.SearchResp
		err = rows.Scan(&post.Pid,
			&post.Pid,
			&post.Title,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}