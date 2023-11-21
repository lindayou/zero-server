package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"zero-server/server/internal/types"
	"zero-server/server/model/admin_operation"
)

type OperationRecordMiddleware struct {
}

func NewOperationRecordMiddleware() *OperationRecordMiddleware {
	return &OperationRecordMiddleware{}
}

func (m *OperationRecordMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body []byte
		//var userId int
		if r.Method != http.MethodGet {
			var err error
			body, err = io.ReadAll(r.Body)
			if err != nil {
				log.Println("error occur", err)
			} else {
				r.Body = io.NopCloser(bytes.NewBuffer(body))
			}
		} else {
			query := r.URL.RawQuery
			query, _ = url.QueryUnescape(query)
			split := strings.Split(query, "&")
			m := make(map[string]string)
			for _, v := range split {
				kv := strings.Split(v, "=")
				if len(kv) == 2 {
					m[kv[0]] = kv[1]
				}
			}
			body, _ = json.Marshal(&m)
		}

		record := &types.CreateOperationReq{
			Ip:     r.RemoteAddr,
			Method: r.Method,
			Path:   r.URL.Path,
			Agent:  r.UserAgent(),
			Body:   string(body),
		}

		//上传文件  中间件日志对上传进行裁剪避免记录太大

		// 插入数据库接口数据
		operate := &admin_operation.SysOperationRecords{
			Ip:           record.Ip,
			Method:       record.Method,
			Path:         record.Path,
			Status:       0,
			Latency:      0,
			Agent:        record.Agent,
			ErrorMessage: "",
			Body:         record.Body,
			Resp:         "",
			UserId:       0,
		}
		conn := sqlx.NewMysql("root:qwe123-=@tcp(127.0.0.1:3306)/go-zero?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai")
		operation := admin_operation.NewSysOperationRecordsModel(conn)
		_, err := operation.Insert(context.Background(), operate)
		if err != nil {
			log.Println(err)
		}

		next(w, r)
		header := w.Header()

		fmt.Println(header)
	}
}
