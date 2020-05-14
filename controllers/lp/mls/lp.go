/**
  create by yy on 2020/3/23
*/

package mls

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

type SubPage struct {
	beego.Controller
}

func (s *SubPage) Lp() {

	trackId := s.GetString("tid")

	if trackId == "" {
		// 自动获取trackId
		trackId = getTrackId()
		s.Redirect("http://deliciouskitchen.recipenice.com/mls/lp?tid="+trackId, 302)
		s.StopRun()
	}

	s.Data["TrackId"] = trackId

	s.TplName = "lp/mls/lp.html"
}

func (s *SubPage) Pin() {
	s.TplName = "lp/mls/pin.html"
}

func (s *SubPage) Thank() {
	s.TplName = "lp/mls/thank.html"
}

func getTrackId() string {
	var (
		err      error
		response *httplib.BeegoHTTPRequest
		result   []byte
	)

	postData := make(map[string]interface{})

	postData["service_id"] = "36881-RP"
	postData["service_name"] = "RP-36881-RP"
	postData["PromoterID"] = 3
	postData["offer_id"] = 192
	postData["camp_id"] = 29
	postData["aff_name"] = "AAA"
	postData["PostbackPrice"] = 0.8

	getTrackIdUrl := "http://ml.foxseekmedia.com/aff/track"

	request := httplib.Post(getTrackIdUrl)

	if response, err = request.JSONBody(postData); err != nil {
		fmt.Println(err)
	}

	if result, err = response.Bytes(); err != nil {
		fmt.Println(err)
	}

	return string(result)
}