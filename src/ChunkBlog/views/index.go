package views

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
	"github.com/Cgo"
)

func Index (h *Cgo.RouterHandler) {
	url := "https://ren.cbim.org.cn/api/config/v1/light/model/state?id=33457"
	url = "https://ren.cbim.org.cn/api/schedule/v1/prepare/33457"
	url = "https://ren.cbim.org.cn/api/schedule/v1/instance/33457/model-33457-gaciftfshbjndorosscadfrgzfyebighaiedddhkjuifdbnmavijfrdbgdgajhhd"
	url = "https://ren.cbim.org.cn/api/schedule/v1/deploy/model-33457-gaciftfshbjndorosscadfrgzfyebighaiedddhkjuifdbnmavijfrdbgdgajhhd/dea230b1-99c3-4a79-bfd1-6827d3ecad39"
	url = "https://ren.cbim.org.cn/api/schedule/v1/instance"
	url = "https://ren.cbim.org.cn/api/schedule/v1/queryPort/model-33457-gaciftfshbjndorosscadfrgzfyebighaiedddhkjuifdbnmavijfrdbgdgajhhd"
	url = "https://ren.cbim.org.cn/api/config/v1/light/model/tree?id=33457"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Print(err)
	}
	reqRes, err := client.Do(req)
	if err != nil {
		log.Print(err)
	}

	defer reqRes.Body.Close()
	body, err := ioutil.ReadAll(reqRes.Body)

	log.Println(string(body))
	//ret := make(map[string]interface{})
	//json.Unmarshal(body, &ret)

	fmt.Fprintf(h.W, string(body))

}
