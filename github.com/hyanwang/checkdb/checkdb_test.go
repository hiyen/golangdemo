package checkdb

import (
	"fmt"
	"testing"
)

func TestRunPhp(t *testing.T) {
	res := RunPHP("http://10.100.55.111/dexter/2.php")
	fmt.Printf("body:%s", res)
	if string(res) == "1" {
		fmt.Println("right")
	} else {
		fmt.Println("wrong")
	}
	t.Log(res)
}
