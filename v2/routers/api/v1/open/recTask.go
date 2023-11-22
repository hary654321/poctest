package open

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/projectdiscovery/nuclei/v2/cmd/nuclei"
	"github.com/projectdiscovery/nuclei/v2/core/slog"
	"github.com/projectdiscovery/nuclei/v2/lib/cache"
	"github.com/projectdiscovery/nuclei/v2/pkg/utils"
)

func RecTask(c *gin.Context) {

	target := c.PostForm("target")
	tempName := c.PostForm("tempName")
	tempContent := c.PostForm("tempContent")
	tempPath := c.PostForm("tempPath")

	taskId := c.PostForm("taskId")

	Validate := c.PostForm("validate")

	cache.Set(taskId, []byte("1"))

	dir, _ := os.Getwd()

	temPre := dir + "/vul"

	// slog.Println(slog.DEBUG, "temPre", temPre)
	tmp := ""
	if tempContent != "" {
		tmp = temPre + "/diy/" + tempName + ".yaml"
		utils.Write(tmp, tempContent)
	} else if tempPath != "" {
		tmp = temPre + tempPath

		slog.Println(slog.DEBUG, "===", tmp)
		if !utils.FileExists(tmp) {
			c.JSON(http.StatusOK, gin.H{
				"code": 400,
				"msg":  "模板不存在",
				"data": "",
			})
			return
		}
	}

	if target != "" {
		strArrayNew := strings.Split(target, ",")
		err := nuclei.Scan(strArrayNew, tmp, taskId, Validate == "1")

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 400,
				"msg":  err.Error(),
				"data": "",
			})
			return
		}
	}

	data := make(map[string]interface{})

	data["taskId"] = taskId

	data["startTime"] = utils.GetTime()

	data["res"] = utils.Read(utils.LogPath + taskId + ".json")

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "",
		"data": data,
	})
}
