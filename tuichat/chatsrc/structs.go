package chatsrc

type post struct {
	username string
	message  string
	time     string
}

type room struct {
	roomid   string
	username string
}

var posts = []post{}

