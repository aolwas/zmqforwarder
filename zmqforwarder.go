package main

import(
  "flag"
  "fmt"
  "os"
  tail "github.com/ActiveState/tail"
  zmq "github.com/pebbe/zmq4"
)

func checkErr(err error) {
  if err != nil {
    panic(err)
  }
}

func WatchFile(filepath string, shippingChan chan<- string) {
  tailed, err := tail.TailFile(filepath,
    tail.Config{
      Location: &tail.SeekInfo{Whence: os.SEEK_END},
      Follow: true,
      ReOpen: true})
  checkErr(err)

  for line := range(tailed.Lines) {
    if len(line.Text) > 0 {
      shippingChan <- line.Text
    }
  }
}

func LogSender(serverAddress string, shippingChan <-chan string) {
  server, err := zmq.NewSocket(zmq.PUB)
  defer server.Close()
  checkErr(err)
  server.Bind(serverAddress)

  for {
    line := <-shippingChan
    server.Send(line, 0)
  }
}

func main() {
  var serverAddress string
  flag.StringVar(&serverAddress, "server", "tcp://localhost:2120", "address of PULL socket")
  flag.Parse()
  files := flag.Args()

  if len(files) > 0 {
    shippingChan := make(chan string, 1000)

    for _, file := range(files) {
      go WatchFile(file, shippingChan)
    }

    LogSender(serverAddress, shippingChan)
  } else {
    fmt.Println("usage: zmqforwarder [-server tcp://server:port] logfile [logfile ...]")
  }
}
