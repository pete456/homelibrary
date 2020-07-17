package daemon

import (
	"log"
	"os"
	"syscall"
)

func StartDaemon(files []uintptr) (int,error) {
	if _,isDaemon := os.LookupEnv("DAEMON"); !isDaemon {
		daemonENV := []string{"DAEMON=true"}
		_,err := syscall.ForkExec(os.Args[0], os.Args, &syscall.ProcAttr {
			Env: append(os.Environ(),daemonENV...),
			Sys: &syscall.SysProcAttr {
				Setsid: true,
			},
			Files: files,
		},
	)
	if err != nil {
		log.Panic("Error starting daemon: %s",err)
			return os.Getpid(),err
		}

	}
	return os.Getpid(),nil
}
