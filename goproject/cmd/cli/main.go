package main
import (
	"fmt"
	"github.com/aliciarojovera/beginners-guide-to-go/goproject/internal/network"

)

func main() {
	fmt.Println("our app entrypoint")
	network.Ping("127.0.0.1")
}