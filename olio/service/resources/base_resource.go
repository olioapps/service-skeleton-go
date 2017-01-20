package resources

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rachoac/service-skeleton-go/olio/api"
	"github.com/rachoac/service-skeleton-go/olio/util"
)

type BaseResource struct {
	ginEngine *gin.Engine
}

func (self *BaseResource) returnJSON(c *gin.Context, status int, record interface{}) {
	w := c.Writer

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")

	if record != nil {
		c.IndentedJSON(status, record)
	}
}

func (self *BaseResource) returnXML(c *gin.Context, status int, record interface{}) {
	w := c.Writer

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "text/xml")

	if record != nil {
		output, err := xml.Marshal(record)
		if err == nil {
			fmt.Println(string(output))
			c.Writer.Write(output)
		}
	}
}

func (self *BaseResource) returnError(c *gin.Context, status int, message string) {
	var outboundStruct struct {
		Code  int    `json:"code"`
		Error string `json:"error"`
	}

	outboundStruct.Error = message
	outboundStruct.Code = status
	self.returnJSON(c, status, outboundStruct)
}

func (self *BaseResource) returnJSONException(c *gin.Context, exception *api.Exception) {
	self.returnError(c, exception.ErrorCode, exception.Err)
}

func (self *BaseResource) returnXMLError(c *gin.Context, status int, message string) {
	type Error struct {
		Code    int    `json:"code"`
		Message string `json:"error"`
	}

	outboundStruct := Error{}
	outboundStruct.Message = message
	outboundStruct.Code = status
	self.returnXML(c, status, outboundStruct)
}

func (self *BaseResource) parseString(c *gin.Context, paramName string) string {
	param := c.Query(paramName)
	if param == "" {
		param = c.Param(paramName)
	}

	return param
}

func (self *BaseResource) parseInt(c *gin.Context, paramName string) int64 {
	param := self.parseString(c, paramName)
	if param == "" {
		return 0
	}

	return util.StringToInt64(param)
}

func (self *BaseResource) parseDate(c *gin.Context, field string) *time.Time {
	if field == "" {
		return nil
	}
	param := c.Param(field)
	if param == "" {
		// try parameter
		param = c.Query(field)
	}

	if param == "" {
		return nil
	}

	if _, err := strconv.Atoi(param); err == nil {
		t, err := msToTime(param)
		if err != nil {
			return nil
		}
		return &t
	}

	startAt, e := time.Parse(time.RFC3339, param)
	startAt = startAt.UTC()
	if e != nil {
		return nil
	}
	return &startAt
}

func (self *BaseResource) parseArray(c *gin.Context, paramName string) []string {
	param := self.parseString(c, paramName)
	if param == "" {
		return nil
	}
	return strings.Split(param, ",")
}

func msToTime(ms string) (time.Time, error) {
	msInt, err := strconv.ParseInt(ms, 10, 64)
	if err != nil {
		return time.Time{}, err
	}

	return time.Unix(0, msInt*int64(time.Millisecond)), nil
}
