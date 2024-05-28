package e

var MsgFlags = map[int]string{
	SUCCESS:                   "ok",
	ERROR:                     "fail",
	INVALID_PARAMS:            "Invalid parameters",
	ERROR_NOT_EXIST_ARTICLE:   "Article does not exist",
	ERROR_DELETE_ARTICLE_FAIL: "Failed to delete article",
	ERROR_EDIT_ARTICLE_FAIL:   "Failed to edit article",
	ERROR_GET_ARTICLE_FAIL:    "Failed to get article",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
