package main

/*"repository":{
  "name": "Diaspora",
  "url": "git@example.com:mike/diaspora.git",
  "description": "",
  "homepage": "http://example.com/mike/diaspora",
  "git_http_url":"http://example.com/mike/diaspora.git",
  "git_ssh_url":"git@example.com:mike/diaspora.git",
  "visibility_level":0
},
*/

type RepositoryType struct {
	Name        string
	Url         string
	Description string
	HomePage    string
}

type CommitAuthor struct {
	Name  string
	Email string
}

type CommitType struct {
	Id        string
	Message   string
	Timestamp string
	Url       string
	Author    CommitAuthor
	Added     []string
	Modified  []string
	Removed   []string
}

type HeaderType struct {
	ObjectKind string `json:"object_kind"`
	Repository RepositoryType
}
