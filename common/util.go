package common

import (
	"cat/common/consts"
	"encoding/json"
	"fmt"
	"github.com/dylanpeng/golib/coder"
	"github.com/gin-gonic/gin"
	"runtime/debug"
	"strings"
)

func GetKey(prefix string, items ...interface{}) string {
	format := prefix + strings.Repeat(":%v", len(items))
	return fmt.Sprintf(format, items...)
}

func ConvertStruct(a interface{}, b interface{}) (err error) {
	defer func() {
		if err != nil {
			Logger.Debugf("convert data failed | data: %s | error: %s", a, err)
		}
	}()

	data, err := json.Marshal(a)

	if err != nil {
		return
	}

	return json.Unmarshal(data, b)
}

func ConvertStructs(items ...fmt.Stringer) (err error) {
	for i := 0; i < len(items)-1; i += 2 {
		if err = ConvertStruct(items[i], items[i+1]); err != nil {
			return
		}
	}

	return
}

func CatchPanic() {
	if err := recover(); err != nil {
		Logger.Fatalf("catch panic | %s\n%s", err, debug.Stack())
	}
}

func GetCtxCoder(ctx *gin.Context) coder.ICoder {
	name := ctx.GetString(consts.CtxCoderKey)

	if name == coder.EncodingProtobuf {
		return coder.ProtoCoder
	} else if name == coder.EncodingJson {
		return coder.JsonCoder
	} else {
		return coder.JsonCoder
	}
}

func SetCtxCoder(ctx *gin.Context, encoding string) {
	if encoding == coder.EncodingProtobuf || encoding == coder.EncodingJson {
		ctx.Set(consts.CtxCoderKey, encoding)
	}
}
