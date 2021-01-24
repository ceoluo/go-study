package obj_pool

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

type ReusableObj struct {

}

type ObjPool struct {
	// 用于缓冲可重用对象
	buffChan chan *ReusableObj
}

func NewObjPool(numberOfObj int) *ObjPool {
	var objPool ObjPool
	objPool.buffChan = make(chan *ReusableObj, numberOfObj)
	for i:=0;i<numberOfObj;i++ {
		objPool.buffChan <- &ReusableObj{}
	}
	return &objPool
}

func (p *ObjPool) GetObj(timeout time.Duration) (*ReusableObj, error)  {
	select {
	case ret := <-p.buffChan:
		return ret, nil
	case <- time.After(timeout):
		return nil, errors.New("time out")
	}
}

func (p *ObjPool) ReleaseObj(obj *ReusableObj) error {
	select {
	case p.buffChan <- obj:
		return nil
	default:
		return errors.New("overflow")
	}
}

func TestObjPool(t *testing.T)  {
	// 创建对象池
	objPool := NewObjPool(10)
	//if err:=objPool.ReleaseObj(&ReusableObj{});err!=nil{
	//	t.Error(err)
	//}
	for i:=0;i<11;i++ {
		if v, err := objPool.GetObj(time.Second*1);err!=nil{
			t.Error(err)
		}else {
			fmt.Printf("%T\n", v)
			if err:=objPool.ReleaseObj(v); err!=nil{
				t.Error(err)
			}
		}
	}

	fmt.Println("Done")
}
