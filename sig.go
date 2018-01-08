package gohelper  
   
import (  
    "fmt"  
    "os"
    "syscall"
    "os/signal"
    "time"
    "log"
)
   
type SignalHandler func(s os.Signal, arg interface{})  
   
type SignalSet struct {  
    m map[os.Signal]SignalHandler  
}  
   
func SignalSetNew() *SignalSet {  
    ss := new(SignalSet)  
    ss.m = make(map[os.Signal]SignalHandler)  
    return ss  
}  
   
func (set *SignalSet) Register(s os.Signal, handler SignalHandler) {  
    if _, found := set.m[s]; !found {  
        set.m[s] = handler  
    }  
}  
   
func (set *SignalSet) Handle(sig os.Signal, arg interface{}) (err error) {  
    if _, found := set.m[sig]; found {  
        set.m[sig](sig, arg)  
        return nil  
    } else {  
        return fmt.Errorf("No handler available for signal %v", sig)  
    }  
    panic("won't reach here")  
}



var stopFlag_ bool

func DoSig() {

    sigHandler := SignalSetNew()
    sigHandler.Register(syscall.SIGQUIT, sigHandlerFunc)
    sigHandler.Register(syscall.SIGUSR1, sigHandlerFunc)
    sigHandler.Register(syscall.SIGINT, sigHandlerFunc)
    sigChan := make(chan os.Signal, 10)
    signal.Notify(sigChan)

    for true {
        select {
        case sig := <-sigChan:
            err := sigHandler.Handle(sig, nil)
            if err != nil {
                log.Println("[ERROR] unknown signal received: ", sig)
            }
            if stopFlag_ {
                os.Exit(0)
            }
        default:
            time.Sleep(time.Duration(1) * time.Second)
        }
    }
}

func sigHandlerFunc(s os.Signal, arg interface{}) {
    log.Println("signal recv:",s)
    switch s {
    case syscall.SIGUSR1: // check
        stopFlag_ = true
    case syscall.SIGQUIT: // stop
        stopFlag_ = true
    case syscall.SIGINT: // stop
        stopFlag_ = true
    }
}