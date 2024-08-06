package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/shirou/gopsutil/v3/process"
	"github.com/spf13/cobra"
)

var (
	cores int
	clock int
	ram   int
)

var runCmd = &cobra.Command{
	Use:   "run [flags] [command]",
	Short: "Run a command with limited resources defined in the given flags",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("No command provided")
			os.Exit(1)
		}

		command := args[0]
		argsForCommand := args[1:]

		fmt.Printf("Running command: %s with args: %v\n", command, argsForCommand)
		fmt.Printf("Limits - Cores: %d, Clock: %d MHz, RAM: %d MB\n", cores, clock, ram)

		runtime.GOMAXPROCS(cores)

		cmdToRun := exec.Command(command, argsForCommand...)
		cmdToRun.Stdin = os.Stdin
		cmdToRun.Stdout = os.Stdout
		cmdToRun.Stderr = os.Stderr
		cmdToRun.Env = os.Environ()

		sigChannel := make(chan os.Signal, 1)
		signal.Notify(sigChannel, syscall.SIGINT, syscall.SIGTERM)

		if err := cmdToRun.Start(); err != nil {
			fmt.Println("Error starting command:", err)
			return
		}

		go monitorUsage(cmdToRun.Process.Pid)

		go func() {
			for sig := range sigChannel {
				_ = cmdToRun.Process.Signal(sig)
			}
		}()

		if err := cmdToRun.Wait(); err != nil {
			fmt.Println("Command execution failed:", err)
		}

		waitStatus := cmdToRun.ProcessState.Sys().(syscall.WaitStatus)
		os.Exit(waitStatus.ExitStatus())
	},
}

func monitorUsage(pid int) {
	for {
		p, err := process.NewProcess(int32(pid))
		if err != nil {
			fmt.Printf("Error getting process: %v\n", err)
			return
		}

		cpuPercent, err := p.CPUPercent()
		if err != nil {
			fmt.Printf("Error getting CPU usage: %v\n", err)
			return
		}

		memInfo, err := p.MemoryInfo()
		if err != nil {
			fmt.Printf("Error getting memory usage: %v\n", err)
			return
		}
		memUsageMB := memInfo.RSS / 1024 / 1024

		fmt.Printf("CPU Usage: %.2f%%, RAM Usage: %dMB\n", cpuPercent, memUsageMB)

		if int(cpuPercent) > clock || int(memUsageMB) > ram {
			fmt.Println("Resource limits exceeded, terminating process...")
			if err := p.Kill(); err != nil {
				fmt.Println("Failed to kill process:", err)
			}
			return
		}

		time.Sleep(1 * time.Second)
	}
}

func init() {
	runCmd.Flags().IntVarP(&cores, "cores", "c", 1, "Number of CPU cores")
	runCmd.Flags().IntVarP(&clock, "clock", "", 1000, "CPU clock speed in MHz")
	runCmd.Flags().IntVarP(&ram, "ram", "r", 512, "RAM in MB")

	rootCmd.AddCommand(runCmd)
}
