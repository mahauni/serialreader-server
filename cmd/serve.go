package cmd

import (
	"os"
	"os/signal"
	"syscall"

	server "github.com/mahauni/serialreader-server/internal"
	"github.com/spf13/cobra"
)

var (
	port              int
	arduinoDevicePath string
)

func init() {
	serveCmd.Flags().StringVarP(&arduinoDevicePath, "arduino_path", "f", "/dev/cu.usbmodem14201", "The location of the connected arduino device on your computer.")
	serveCmd.MarkFlagRequired("arduino_path")

	serveCmd.Flags().IntVarP(&port, "port", "p", 50051, "The port to run this server on")

	rootCmd.AddCommand(serveCmd)
}

func doServe() {
	server := server.New(arduinoDevicePath, port)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		server.StopMainRuntimeLoop()
	}()
	server.RunMainRuntimeLoop()
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run the gRPC server",
	Long:  `Run the gRPC server to allow other services to access the serial reader`,
	Run: func(cmd *cobra.Command, args []string) {
		doServe()
	},
}
