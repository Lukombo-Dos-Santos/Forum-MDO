package structure

type UserAccount struct {
	Id        int
	Name      string //Name of the user
	Image     string //Path src
	Email     string
	UUID      string
	Password  string
	Post      []Post
	Comment   []Comment
	Admin     bool //true: the user is the Admin
	Connected bool
}

type Comment struct {
	Name           string
	CommentID      string
	Message        string
	UserImage      string
	DateTime       string
	PostID         string
	CommentLike    int
	CommentDislike int
	Connected      bool
}

type HomeFeedPost struct {
	PostID           string
	Name             string
	Message          string
	UserImage        string
	DateTime         string
	Picture          string
	NumberOfComment  int
	NumberOfLikes    int
	NumberOfDislikes int
	Categories       string
	Categories2      string
}

type UserFeedPost struct {
	PostID           string
	Name             string
	Message          string
	UserImage        string
	DateTime         string
	Picture          string
	NumberOfComment  int
	NumberOfLikes    int
	NumberOfDislikes int
	Categories       string
	Categories2      string
}

type Categorie1FeedPost struct {
	PostID           string
	Name             string
	Message          string
	UserImage        string
	DateTime         string
	Picture          string
	NumberOfComment  int
	NumberOfLikes    int
	NumberOfDislikes int
	Categories       string
	Categories2      string
}

type Categorie2FeedPost struct {
	PostID           string
	Name             string
	Message          string
	UserImage        string
	DateTime         string
	Picture          string
	NumberOfComment  int
	NumberOfLikes    int
	NumberOfDislikes int
	Categories       string
	Categories2      string
}

type Categorie3FeedPost struct {
	PostID           string
	Name             string
	Message          string
	UserImage        string
	DateTime         string
	Picture          string
	NumberOfComment  int
	NumberOfLikes    int
	NumberOfDislikes int
	Categories       string
	Categories2      string
}

type Post struct {
	PostID      string
	Name        string
	Message     string
	UserImage   string
	DateTime    string
	Picture     string
	CountCom    int
	Count       int
	CountDis    int
	Comment     Comment
	Categories  string
	Categories2 string
	Connected   bool
}

type Likes struct {
	NumberOfLikes int
	UserNameLike  string
	PostId        string
	DateTime      string
	Connected     bool
}
type Dislikes struct {
	NumberOfDislikes int
	UserNameDislike  string
	PostID           string
	DateTime         string
	Connected        bool
}

type AuthGoogle struct {
	Access_Token  string `json:"access_token"`
	Expires_In    int    `json:"expires_in"`
	Refresh_Token string `json:"refresh_token"`
	Id_Token      string `json:"id_token"`
	Scope         string `json:"scope"`
	Token_Type    string `json:"token_type"`
}

type GoogleUser struct {
	Name           string `json:"name"`
	Picture        string `json:"picture"`
	Email          string `json:"email"`
	Email_Verified string `json:"email_verified"`
}

type AuthGitHub struct {
	Access_Token string `json:"access_token"`
	Scope        string `json:"scope"`
	Token_Type   string `json:"token_type"`
}

type GithubUser struct {
	Avatar_Url string `json:"avatar_url"`
	Name       string `json:"name"`
	Email      string `json:"email"`
}

type AuthDiscord struct {
	Access_Token  string `json:"access_token"`
	Token_Type    string `json:"token_type"`
	Expires_In    int    `json:"expires_in"`
	Refresh_Token string `json:"refresh_token"`
	Scope         string `json:"scope"`
}

type DiscordUser struct {
	Identify string `json:"identify"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
}
