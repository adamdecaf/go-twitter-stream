package main
import "io"
import "fmt"
import "log"
import "bytes"
import "net/http"

func maybePrintResponse(resp http.Response) {
	if resp != nil {
		unsafeReadAndPrint("resp", resp.Body)
	}
}

func unsafeReadAndPrint(prefix string, reader io.ReadCloser) {
	buffer := new(bytes.Buffer)
	buffer.ReadFrom(reader)
	fmt.Sprintf("%s - %s", prefix, buffer.String())
}

func printError(err error) {
	if err != nil {
		log.Printf(err.Error())
	}
}

func main() {
	client := &http.Client{}
	// todo: how to make this monadic?
	resp1,err1 := client.Get("http://banno.com")
	resp2,err2 := client.Get("http://thisdonedoesnotexisstallprobably.com")

	printError(err1)
	printError(err2)

	fmt.Println("start")
	maybePrintResponse(resp1)
	maybePrintResponse(resp2)
	fmt.Println("end")
}
