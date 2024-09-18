package logger

import "log"

type ApplicationLogger interface{
  Log(string) 
}

type SysoutLogger struct{
  Logger *log.Logger
}

func (syslog *SysoutLogger) Log(s string){
  syslog.Logger.Printf(s)
}
