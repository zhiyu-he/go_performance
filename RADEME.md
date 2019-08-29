### 每个目录的说明:

* ams: 汇编指令相关, 对基本的read、write等操作的汇编分析
* build_in: GO中内部基本方法性能测试
* lib: 对基础库的性能测试, 如json、http长链接
* simple_impl: 一些简单的组件实现, 如内存池、对象大小的度量


### 基本开发习惯

1. 对于如slice、map等容器类型数据结构, 指定**capacity**, 避免动态的容量调整
	```
		map = make(map[int64]int64, 124)
	```
2. 对于**只读的string、bytes**的, 可以采用原地的`String->Bytes, Bytes->String`

	```
		package util
		import "unsafe"
		
		func Str2Bytes(s string) []byte {
			x := (*[2]uintptr)(unsafe.Pointer(&s))
			h := [3]uintptr{x[0], x[1], x[1]}
			return *(*[]byte)(unsafe.Pointer(&h))
		}
		func Bytes2Str(b []byte) string {
			return *(*string)(unsafe.Pointer(&b))
		}
	```
3. 对于简单结构体的Marshal, 使用Struct, 不使用Map, 因为Map会被通过反射sort一次.
4. 对于fmt的使用, 要考虑具体的代码路径, 以及执行效率, 单纯的ItoA, 建议使用strconv.FormatInt()
5. 如果需要用map做去重(或者其他不需要value的场景),  建议使用struct{}空字符串作为value类型, 会比其他类型省内存, eg. map[string]struct{}
6. 对于需要随机的场景, 尽量使用time.Now().Nanosecond() 代替 rand.Int31n, 性能高出1倍+

