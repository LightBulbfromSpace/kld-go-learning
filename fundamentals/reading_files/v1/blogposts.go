package blogposts

import (
	"io"
	"io/fs"
)

type Post struct {
	Title string
}

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")

	if err != nil {
		return nil, err
	}

	var posts []Post
	for range dir {
		posts = append(posts, Post{})
	}

	return posts, nil
}

func getPost(fileSystem fs.FS, f fs.DirEntry) (Post, error) {
	postFile, err := fileSystem.Open(f.Name())

	if err != nil {
		return Post{}, err
	}

	defer postFile.Close()

	postData, err := io.ReadAll(postFile)

	if err != nil {
		return Post{}, err
	}
	post := Post{Title: string(postData)[7:]}
	return post, nil
}
