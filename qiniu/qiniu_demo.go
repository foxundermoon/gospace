package main

import (
	"log"
	"os"

	. "github.com/qiniu/api/conf"
	"github.com/qiniu/api/io"
	"github.com/qiniu/api/rs"
)

func uptoken(bucketName string) string {
	putPolicy := rs.PutPolicy{
		Scope: bucketName,
		//CallbackUrl: callbackUrl,
		//CallbackBody:callbackBody,
		//ReturnUrl:   returnUrl,
		//ReturnBody:  returnBody,
		//AsyncOps:    asyncOps,
		//EndUser:     endUser,
		//Expires:     expires,
	}
	return putPolicy.Token(nil)
}

type PutExtra struct {
	Params map[string]string //可选，用户自定义参数，必须以 "x:" 开头
	//若不以x:开头，则忽略
	MimeType string //可选，当为 "" 时候，服务端自动判断
	Crc32    uint32
	CheckCrc uint32
	// CheckCrc == 0: 表示不进行 crc32 校验
	// CheckCrc == 1: 对于 Put 等同于 CheckCrc = 2；对于 PutFile 会自动计算 crc32 值
	// CheckCrc == 2: 表示进行 crc32 校验，且 crc32 值就是上面的 Crc32 变量
}

func uploadWithoutKey() error {

	return nil
}

var err error
var token string
var key string
var f *os.File
var extra = &io.PutExtra{
//Params:    params,
//MimeType:  mieType,
//Crc32:     crc32,
//CheckCrc:  CheckCrc,
}
var ret io.PutRet

func init() {

	ACCESS_KEY = "3EZOcwdTLi7H_GgckpOlibdMJIa36t1ZPSQe5zWh"
	SECRET_KEY = "pu3wChxncjD-oJvjqvSzkc5iCfgGqBCqrVLWEdAR"
	token = uptoken("vvfox-public")
	log.Println("upload token:", token)
	key = "upload/go/test/hello.txt"
	f, err = os.Open("hello.txt")
	if err != nil {
		log.Println("open file failed:", err)
		return
	}
}
func uploadByPut() {
	var err error

	// ret       变量用于存取返回的信息，详情见 io.PutRet
	// uptoken   为业务服务器端生成的上传口令
	// key       为文件存储的标识
	// r         为io.Reader类型，用于从其读取数据
	// extra     为上传文件的额外信息,可为空， 详情见 io.PutExtra, 可选

	err = io.Put(nil, &ret, token, key, f, extra)

	if err != nil {
		//上传产生错误
		log.Print("io.Put failed:", err)
		return
	}

	//上传成功，处理返回值
	log.Print("upload success", ret)
}
func main() {
	uploadByPut()
}
