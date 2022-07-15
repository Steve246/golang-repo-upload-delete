package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type FileResponse struct {
	c              *gin.Context
	httpStatusCode int
	fileName       string
}

func (f *FileResponse) Send() {
	f.c.File(f.fileName)
}

func NewSuccessFileResponse(c *gin.Context, fileName string) AppHttpResponse {
	return &FileResponse{
		c,
		http.StatusOK,
		fileName,
	}
}
