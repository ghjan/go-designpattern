package main

import "fmt"

/* Facade
外观模式：为子系统中的一组接口提供一个一致的界面，此模式定义了一个高层接口，这个接口使得这一子系统更加容易使用。
https://github.com/jeanphorn/go-design-patterns/tree/master/structural_patterns
 */
type Facade struct {
	M Music
	V Video
	C Count
}

func (facade *Facade) GetRecommendVideos() error {
	facade.V.GetVideos()
	facade.C.GetCountByID(111)

	return nil
}

type Music struct {
}

func (music *Music) GetMusic() error {
	fmt.Println("get music material")
	// logic code here
	return nil
}

type Video struct {
	vid int64
}

func (video *Video) GetVideos() error {
	fmt.Println("get videos1")
	return nil
}

type Count struct {
	PraiseCnt  int64 //点赞数
	CommentCnt int64 //评论数
	CollectCnt int64 //收藏数
}

func (count *Count) GetCountByID(id int64) (*Count, error) {
	fmt.Println("get video counts")
	return count, nil
}

func main() {
	f := &Facade{}
	f.GetRecommendVideos()
}
