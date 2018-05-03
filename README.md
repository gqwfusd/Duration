# Duration
本小程序的目的是为了对工作中使用的比较常用的接口请求耗时的算法的一个小结，这里面遇到的问题，得出的时间查elpsedTime是个复数
一开始使用的是
1.elapsedS := int32(endTime.Sub(beginTime).Nanoseconds()) / 1000000
发现得出的结果总是有负数，经过修改后，变为另一种更常用的方式
2.elapseS  := int32(time.Since(beginTime).Nanoseconds()) / 1000000
结果还是不行

这个时候，没有更好的办法只能去查看time包下的time.go的源码查看Duration返回值到底是如何定义的
结果恍然大悟，通过查看源码能够得知，Duration是个有符号的int64，如果转为int32后，最大数值2的32次方，最大是2s，超过2s就会溢出
// A Duration represents the elapsed time between two instants
// as an int64 nanosecond count. The representation limits the
// largest representable duration to approximately 290 years.
type Duration int64

const (
	minDuration Duration = -1 << 63
	maxDuration Duration = 1<<63 - 1
)

修改方法：将int32修改为int64搞定，或者换种写法都是可以的.
int(time.Since(beginTime).Seconds() * 1000)
