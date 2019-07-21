package dzpk_read_file

import (
	"fmt"
	"net/http"
)

//读取大量的数据如果是十万条或者更多采用分割数据加载的方式
func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "<h1>hello word %v <h1>", request.FormValue("name"))
	})
	http.ListenAndServe(":9090", nil)
}
