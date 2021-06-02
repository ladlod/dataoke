package dataoke

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/parnurzeal/gorequest"
)

type DaTaoKeApp struct {
	appKey    string
	appSecret string
	debug     bool

	clients *sync.Pool
}

// NewDaTaoKeApp
func NewDaTaoKeApp(appKey string, appSecret string, debug bool) *DaTaoKeApp {
	return &DaTaoKeApp{
		appKey:    appKey,
		appSecret: appSecret,
		debug:     debug,
		clients: &sync.Pool{
			New: func() interface{} {
				return gorequest.New()
			},
		},
	}
}

// postQuery
func (d *DaTaoKeApp) postQuery(params Params, uri string) (resp []byte, err error) {
	client := d.clients.Get().(*gorequest.SuperAgent)
	defer d.clients.Put(client)

	ts := time.Now().UnixNano()
	nonce := 100000 + rand.Int63n(900000)
	params.Set("appKey", d.appKey)
	params.Set("version", "v1.2.4")
	params.Set("timer", ts)
	params.Set("nonce", nonce)
	params.Set("signRan", d.createSign(nonce, ts))
	reqUrl := fmt.Sprintf("%v/%v?%v", ReqDomain, uri, params.GetQuery())

	_, respStr, errors := client.Get(reqUrl).End()
	if err = getErrorsError(errors); err != nil {
		return
	}
	resp = []byte(respStr)

	// log
	if d.debug {
		log.Printf("req:%v, resp:%v", reqUrl, respStr)
	}

	return
}

// createSign
func (d *DaTaoKeApp) createSign(nonce int64, ts int64) (sign string) {
	key := fmt.Sprintf("appKey=%v&timer=%v&nonce=%v&key=%v", d.appKey, ts, nonce, d.appSecret)
	h := md5.New()
	h.Write([]byte(signStr))
	cipherStr := h.Sum(nil)
	sign = strings.ToUpper(hex.EncodeToString(cipherStr))
	return
}
