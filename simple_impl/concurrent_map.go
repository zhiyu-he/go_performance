package simple_impl


type ConcurrentMap struct {
	dat map[interface{}]interface{}

	op struct {
		op int
		key interface{}
		dat interface{}
	}
	chan op
}


func (c ConcurrentMap) Get(key interface{}) (interface{}, bool) {

}

func (c ConcurrentMap) Del(key interface{}) {

}

func (c ConcurrentMap) Add(key, val interface{}) {

}
