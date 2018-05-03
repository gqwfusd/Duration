func trace(t *C.int) func() {
   beginTime := time.Now()
   return func() {
        //endTime := time.Now()
        //elapsedS := int32(endTime.Sub(beginTime).Nanoseconds()) / 1000000
        elapsedS := int(time.Since(beginTime).Seconds() * 1000)
        if *t == 0 {
            util.RollLog().Print("|"+"success"+"|", elapsedS, "|login result:0")
        } else {
            util.RollLog().Print("|"+"failed"+"|", elapsedS, "|", *t)
        }
    }
}

func request(wg *sync.WaitGroup){
	defer wg.Done()
	host := "182.150.63.123"
	jparams := fmt.Sprintf("{\"username\":\"xxx\", \"password\":\"123456\"}")
	resp := C.int(0)
	defer trace(&resp)()
	resp = C.SPLogin(C.CString(host), C.uint16_t(10443), C.CString(jparams))

}