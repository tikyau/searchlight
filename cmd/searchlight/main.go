//go:generate stringer -type=State ../../pkg/icinga/types.go
package main

import (
	"log"
	"os"

	logs "github.com/appscode/go/log/golog"
	_ "github.com/appscode/searchlight/client/scheme"
	"github.com/appscode/searchlight/pkg/cmds"
	_ "k8s.io/client-go/kubernetes/fake"
	_ "k8s.io/client-go/pkg/api/v1"
)

func main() {
	logs.InitLogs()
	defer logs.FlushLogs()

	if err := cmds.NewCmdSearchlight(Version).Execute(); err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}
