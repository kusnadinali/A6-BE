package datastruct

type HelloWorldRequest struct {
	NAME string `json:"name"`
}

type HelloWorldResponse struct {
	MESSAGE string `json:"message"`
}

type guidelinesActive struct {
	GuidelinesName string
	GuidelinesDesc string
	GuidelinesType string
	GuidelinesLink string
}

// type Postingan struct {
// 	postingan_id      string    //`json:"post_id"`
// 	username          string    //`json:"username"`
// 	profile_image     string    //`json:"profile_image"`
// 	postingan_caption string    //`json:"postingan_caption"`
// 	postingan_image   string    //`json:"postingan_image"`
// 	date_post         time.Time //`json:"date_post"`
// 	date_time         time.Time //`json:"date_time"`
// 	total_like        int
// 	total_comment     int
// }
