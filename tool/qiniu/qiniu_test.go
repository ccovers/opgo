package qiniu

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/qiniu/api.v7/storage"
)

func TestUpload(t *testing.T) {
	body, err := ioutil.ReadFile("img.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}
	url, err := Img().Upload(100, body, &storage.ZoneHuabei)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("url: ", url)

}

func TestHttpUrl(t *testing.T) {
	name := "/user/30/license.jpg"
	path := fmt.Sprintf("qiniu://test-ucket%s", name)
	url := Img().HttpUrl(path)
	fmt.Printf("path: %s\n", url)
}

func ExampleQiniuUrl() {
	u := "http://test.qiniu.xxx.com/user/195/416c2bf576e888c4d71a3c430906a882bca1c9e0.jpg?" +
		"e=1548839130&token=7-V8DyRk17v8_u71-a9fUh2R5HygnMXcgbzW07SQ:QixqCbVHr0VCahiNpOlGofz22jE="
	fmt.Println(Img().QiniuUrl(u))
	// Output:
	// qiniu://test-bucket/user/195/416c2bf576e888c4d71a3c430906a882bca1c9e0.jpg <nil>
}
