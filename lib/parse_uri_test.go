package lib

import (
	"strings"
"net/url"
	"testing"
)

var uri_source = "ac=wifi&app_name=news_article&_rticket=1503149705764&ssmix=a&item_id=6455077430972383757&from_category=__all__&device_type=OPPO+R9+Plusm+A&aggr_type=1&os_api=22&uuid=861730035912252&openudid=353938b88c11ec0b&version_code=631&os_version=5.1.1&update_version_code=6310&latitude=30.838400681832415&channel=oppo-cpa&device_platform=android&ab_version=160531%2C164185%2C164424%2C161980%2C163248%2C163565%2C161068%2C157000%2C163761%2C159169%2C162397%2C134127%2C160021%2C162012%2C163294%2C162740%2C152027%2C162572%2C156262%2C163513%2C159223%2C157295%2C160967%2C161926%2C31210%2C164116%2C131207%2C145585%2C162593%2C161379%2C157524%2C162573%2C161720%2C156993%2C150352%2C164095&iid=13484115781&device_brand=OPPO&manifest_version_code=631&ab_client=a1%2Cc4%2Ce1%2Cf2%2Cg2%2Cf7&abflag=3&version_name=6.3.1&ab_feature=102749%2C94563&device_id=34639849336&language=zh&plugin=2431&longitude=119.91140224217966&article_page=1&flags=64&context=1&aid=13&group_id=6455077430972383757&resolution=1080%2A1920&dpi=480"



func parseUriByNormal(uri string) map[string]string {
	params := strings.Split(uri, "&")
	paramsMap := make(map[string]string)
	for _, param := range params {
		kv := strings.Split(param, "=")
		if len(kv) == 2 {
			paramsMap[kv[0]] = kv[1]
		}
	}
	return paramsMap
}


func parseUriByLib(uri string) map[string][]string {
	values, err := url.ParseQuery(uri)
	if err != nil {
		return nil
	}
	return map[string][]string(values)
}



func BenchmarkNormalParse(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		parseUriByLib(uri_source)
	}
}

func BenchmarkLibParse(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		parseUriByNormal(uri_source)
	}
}