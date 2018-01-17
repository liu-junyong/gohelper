package gohelper

import (
	"github.com/liu-junyong/go-logger/logger"
	"strconv"
	"sync"
	"strings"
	"time"
)



func StringIpToInt(ipstring string) int {
	ipSegs := strings.Split(ipstring, ".")
	var ipInt int = 0
	var pos uint = 24
	for _, ipSeg := range ipSegs {
		tempInt, _ := strconv.Atoi(ipSeg)
		tempInt = tempInt << pos
		ipInt = ipInt | tempInt
		pos -= 8
	}
	return ipInt
}



type  LimitPolicy struct{
	AccessRlock sync.RWMutex
	AccessLimit []int
	AccessMap map[int]int
}

//var g_UserLimit	*LimitPolicy
var g_IPLimit	*LimitPolicy

func (this *LimitPolicy) Init(){
	this.AccessLimit = make([]int,0)
	this.AccessMap = make(map[int]int)
	go this.timer()
}



func (this *LimitPolicy)timer() {
	timer1 := time.NewTicker(6 * time.Second)
	for {
		select {
		case <-timer1.C:
			go this.UpdateAccess()
		}
	}
}



func (this *LimitPolicy)AddAccess(UserID string,weight int){
	uid,_ := strconv.Atoi( UserID )
	this.AccessRlock.Lock()
	defer this.AccessRlock.Unlock()

	for i := 0 ; i < weight ; i++  {
		this.AccessLimit = append( this.AccessLimit,uid )
	}
}



func (this *LimitPolicy)AddIP(address string,weight int){
	ip := StringIpToInt(address)
	this.AccessRlock.Lock()
	defer this.AccessRlock.Unlock()

	for i := 0 ; i < weight ; i++  {
		this.AccessLimit = append( this.AccessLimit,ip )
	}
}


func (this *LimitPolicy)UpdateAccess(){
	this.AccessRlock.Lock()
	defer this.AccessRlock.Unlock()
	lens := len(this.AccessLimit)
	target := 0
	if lens > 16000 {
		target = lens/5
	}
	if lens > 8000 {
		target = lens/10
	}
	if lens > 4000 {
		target = lens/20
	}
	if lens > 2000 {
		target = lens/30
	}
	if lens > 1000 {
		target = lens/40
	}
	if lens > 500 {
		target = lens/50
	}
	if lens > 100 {
		target = lens/100
	}


	if lens%10 == 0 {
		logger.Info("仓库长度",lens)
	}

	this.AccessLimit = this.AccessLimit[target:]

	maps := make(map[int]int)
	for _,v := range this.AccessLimit {
		x := maps[v]
		maps[v] = x+1
	}
	this.AccessMap = maps
}


func (this *LimitPolicy)CheckAccess(UserID string,limit int)bool{
	this.AccessRlock.Lock()
	defer this.AccessRlock.Unlock()

	uid,_ := strconv.Atoi(UserID)
	x := this.AccessMap[uid]
	logger.Info(x,uid)

	return x < limit
}


func (this *LimitPolicy)ChecIPkAccess(address string,limit int)bool{
	this.AccessRlock.Lock()
	defer this.AccessRlock.Unlock()

	ip  := StringIpToInt(address)
	x := this.AccessMap[ip]

	ret := x < limit
	if ret  {
		logger.Info(x,address)
	}else{
		logger.Error("访问次数：",x,"ip:",address,"长度：",len(this.AccessLimit))
	}
	return ret
}
