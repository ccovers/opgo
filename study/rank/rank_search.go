package main

import (
    "fmt"
)

type IFSearchSlice interface {
    Length() int
    Compare(index int, item interface{}) int // -1/0/1 小于/等于/大于  排行往后/当前排行/排行往前
    IsSameKey(index int, item interface{}) bool
}

// 二分查找
func Search(slice IFSearchSlice, left int, right int, item interface{}) int {
    return SearchWithInsert(slice, left, right, item, false)
}

// 二分查找
// isInsert
//    true 添加查找。有多个相同值插入末尾
//    false 有多个相同值，比较key是否相等
func SearchWithInsert(slice IFSearchSlice, left int, right int, item interface{}, isInsert bool) int {
    if left > right {
        if isInsert {
            if left < 0 {
                left = 0
            }
            return left
        }
        return -1
    }

    midIndex := (left + right) / 2
    ret := slice.Compare(midIndex, item)
    if ret == 1 {
        return SearchWithInsert(slice, left, midIndex-1, item, isInsert)
    } else if ret == -1 {
        return SearchWithInsert(slice, midIndex+1, right, item, isInsert)
    } else {
        if isInsert {
            // 需要插入时，分数相同时插入末尾
            for i := midIndex + 1; i < slice.Length(); i++ {
                if slice.Compare(i, item) != 0 {
                    return i
                }
            }
            return slice.Length()
        }
        // 查找到的值可能多个相同，还需要比较前后的key才能确定位置
        return searchNear(slice, item, midIndex)
    }
}

func searchNear(slice IFSearchSlice, item interface{}, index int) int {
    // 比较key
    if slice.IsSameKey(index, item) {
        return index
    }
    // 向前比较
    for i := index - 1; i >= 0; i-- {
        if slice.Compare(i, item) != 0 {
            break
        }
        if slice.IsSameKey(i, item) {
            return i
        }
    }
    // 向后比较
    for i := index + 1; i < slice.Length(); i++ {
        if slice.Compare(i, item) != 0 {
            break
        }
        if slice.IsSameKey(i, item) {
            return i
        }
    }
    return -1
}

// 排行榜的元素
type IFRankItem interface {
    Key() interface{}
    IsZero() bool
}

// 排行榜，有序包含多个排行榜元素
type IFRank interface {
    Length() int
    Compare(index int, item interface{}) int // -1/0/1 小于/等于/大于  排行往后/当前排行/排行往前
    IsSameKey(index int, item interface{}) bool

    GetItemByIndex(index int) IFRankItem
    GetItemByKey(key interface{}) (IFRankItem, bool)
    GetItems() []IFRankItem
    InsertToRank(index int, item IFRankItem, isOnRank bool)
    RemoveItem(index int, isOnRank bool) // 通过索引删除
    RemoveItemOutLimit(limitLen int)     // 删除超过长度的item
}

// 更新排行榜，返回排行榜更新后位置索引，如果为-1则表示未更新
func UpdateSearchRank(rank IFRank, item IFRankItem, MAX_RANK_NUM int) {
    _, err := updateRankWithInsertLeft(rank, item, MAX_RANK_NUM, 0)
    if err != nil {
        fmt.Printf("add item[%+v: %+v] err: %+v\n", item.Key(), item, err)
    }
}

// 更新排行榜，返回排行榜更新后位置索引，如果为-1则表示未更新
// left：
//    如果非有序插入，那么填0值
//    有序插入时，left表示插入时从第left个开始插入，left应该是上一次更新后位置索引
func updateRankWithInsertLeft(rank IFRank, item IFRankItem, MAX_RANK_NUM int, left int) (int, error) {
    if MAX_RANK_NUM <= 1 || rank == nil || item == nil || item.Key() == nil {
        return -1, fmt.Errorf("rank param invalid")
    }

    isOnRank := false
    rankItem, ok := rank.GetItemByKey(item.Key())
    if ok {
        isOnRank = true
        // 找到并删除原有排行的item（从头到尾找）
        delIndex := Search(rank, 0, rank.Length()-1, rankItem)
        if delIndex < 0 {
            return MAX_RANK_NUM + 1, fmt.Errorf("item not found: %d\n", rankItem.Key(), rankItem)
        }
        if item.IsZero() {
            isOnRank = false
        }
        rank.RemoveItem(delIndex, isOnRank)
    }
    // 判断是否上榜
    if item.IsZero() || (rank.Length() >= MAX_RANK_NUM &&
        rank.Compare(rank.Length()-1, item) < 1) {
        return MAX_RANK_NUM + 1, nil
    }
    // 找到适合的排行，插入item
    addIndex := SearchWithInsert(rank, left, rank.Length()-1, item, true)
    rank.InsertToRank(addIndex, item, isOnRank)
    // 长度超过限制时，比较超过的部分，比限制长度最后一个元素值更小的删除
    if rank.Length() > MAX_RANK_NUM {
        rank.RemoveItemOutLimit(MAX_RANK_NUM)
    }
    return addIndex, nil
}

func (self *TItem) Key() interface{} {
    return self.Id
}

func (self *TItem) IsZero() bool {
    return 0 == self.Score
}

func (self *TRank) Length() int {
    return len(self.RankItems)
}

func (self *TRank) Compare(index int, item interface{}) int {
    if item.(*TItem).Score < self.RankItems[index].Score {
        return -1
    } else if item.(*TItem).Score > self.RankItems[index].Score {
        return 1
    }
    return 0
}

func (self *TRank) IsSameKey(index int, item interface{}) bool {
    return item.(*TItem).Id == self.RankItems[index].Id
}

func (self *TRank) GetItemByKey(key interface{}) (IFRankItem, bool) {
    item, ok := self.RankItemByKey[key.(int64)]
    return item, ok
}

func (self *TRank) GetItems() []IFRankItem {
    items := make([]IFRankItem, 0, self.Length())
    for _, item := range self.RankItems {
        items = append(items, item)
    }
    return items
}

func (self *TRank) GetItemByIndex(index int) IFRankItem {
    return self.RankItems[index]
}

func (self *TRank) InsertToRank(index int, item IFRankItem, isOnRank bool) {
    rankItems := append([]*TItem{item.(*TItem)}, self.RankItems[index:]...)
    self.RankItems = append(self.RankItems[:index], rankItems...)
    self.RankItemByKey[item.Key().(int64)] = item.(*TItem)
}

func (self *TRank) RemoveItem(index int, isOnRank bool) {
    delete(self.RankItemByKey, self.RankItems[index].Id)
    self.RankItems = append(self.RankItems[:index], self.RankItems[index+1:]...)
}

func (self *TRank) RemoveItemOutLimit(limitLen int) {
    for i := limitLen; i < self.Length(); i++ {
        delete(self.RankItemByKey, self.RankItems[i].Key().(int64))
    }
    self.RankItems = self.RankItems[:limitLen]
}

/*

// 更新排行榜，返回排行榜更新后位置索引，如果为-1则表示未更新
func UpdateSearchRankEx(rank *TRank, item *TItem, MAX_RANK_NUM int) {
    _, err := updateRankWithInsertLeftEx(rank, item, MAX_RANK_NUM)
    if err != nil {
        fmt.Printf("add item[%+v: %+v] err: %+v\n", item.Id, *item, err)
    }
}

func updateRankWithInsertLeftEx(rank *TRank, item *TItem, MAX_RANK_NUM int) (int, error) {
    rankItem, _ := rank.RankItemByKey[item.Id]
    if rankItem != nil {
        delIndex := SearchWithInsertEx(rank, 0, rank.Length()-1, rankItem, false)
        if delIndex < 0 {
            return MAX_RANK_NUM + 1, fmt.Errorf("item not found: %d\n", rankItem.Id)
        }
        rank.RankItems = append(rank.RankItems[:delIndex], rank.RankItems[delIndex+1:]...)
        delete(rank.RankItemByKey, rankItem.Id)
    }
    // 判断是否上榜
    if len(rank.RankItems) >= MAX_RANK_NUM && item.Score <= rank.RankItems[rank.Length()-1].Score {
        return MAX_RANK_NUM + 1, nil
    }
    // 找到适合的排行，插入item
    addIndex := SearchWithInsertEx(rank, 0, rank.Length()-1, item, true)
    rankItems := append([]*TItem{item}, rank.RankItems[addIndex:]...)
    rank.RankItems = append(rank.RankItems[:addIndex], rankItems...)
    rank.RankItemByKey[item.Id] = item

    if len(rank.RankItems) > MAX_RANK_NUM {
        for i := MAX_RANK_NUM; i < len(rank.RankItems); i++ {
            delete(rank.RankItemByKey, rank.RankItems[i].Id)
        }
        rank.RankItems = rank.RankItems[:MAX_RANK_NUM]
    }
    return addIndex, nil
}

// 二分查找
// isInsert
//    true 添加查找。有多个相同值插入末尾
//    false 有多个相同值，比较key是否相等
func SearchWithInsertEx(slice *TRank, left int, right int, item *TItem, isInsert bool) int {
    if left > right {
        if isInsert {
            if left < 0 {
                left = 0
            }
            return left
        }
        return -1
    }

    midIndex := (left + right) / 2
    if item.Score > slice.RankItems[midIndex].Score {
        return SearchWithInsertEx(slice, left, midIndex-1, item, isInsert)
    } else if item.Score < slice.RankItems[midIndex].Score {
        return SearchWithInsertEx(slice, midIndex+1, right, item, isInsert)
    } else {
        if isInsert {
            // 需要插入时，分数相同时插入末尾
            for i := midIndex + 1; i < len(slice.RankItems); i++ {
                if item.Score != slice.RankItems[i].Score {
                    return i
                }
            }
            return len(slice.RankItems)
        }
        // 查找到的值可能多个相同，还需要比较前后的key才能确定位置
        return searchNearEx(slice, item, midIndex)
    }
}

func searchNearEx(slice *TRank, item *TItem, index int) int {
    // 比较key
    if slice.RankItems[index].Id == item.Id {
        return index
    }
    // 向前比较
    for i := index - 1; i >= 0; i-- {
        if slice.RankItems[i].Score != item.Score {
            break
        }
        if slice.RankItems[i].Id == item.Id {
            return i
        }
    }
    // 向后比较
    for i := index + 1; i < len(slice.RankItems); i++ {
        if slice.RankItems[i].Score != item.Score {
            break
        }
        if slice.RankItems[i].Id == item.Id {
            return i
        }
    }
    return -1
}
*/
