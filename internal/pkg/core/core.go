package core

import (
	"fmt"

	"github.com/luongdinhkhanhvinh/go-gin/pkg/color"
)

const _UI = `
 ██████╗  ██████╗        █████╗ ██████╗ ██╗
██╔════╝ ██╔═══██╗      ██╔══██╗██╔══██╗██║
██║  ███╗██║   ██║█████╗███████║██████╔╝██║
██║   ██║██║   ██║╚════╝██╔══██║██╔═══╝ ██║
╚██████╔╝╚██████╔╝      ██║  ██║██║     ██║
 ╚═════╝  ╚═════╝       ╚═╝  ╚═╝╚═╝     ╚═╝
`

func main() {
	fmt.Println(color.Blue(_UI))
}
