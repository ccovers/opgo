package main

import (
    "sort"
)

// 通过sort排序
type TItemSlice []*TItem

func (p TItemSlice) Len() int           { return len(p) }
func (p TItemSlice) Less(i, j int) bool { return p[i].Score > p[j].Score }
func (p TItemSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func UpdateSortRank(rank *TRank, item *TItem, MAX_RANK_NUM int) {
    rankItem, _ := rank.RankItemByKey[item.Id]
    if nil == rankItem {
        if len(rank.RankItems) >= MAX_RANK_NUM && rank.RankItems[len(rank.RankItems)-1].Score < item.Score {
            delete(rank.RankItemByKey, rank.RankItems[len(rank.RankItems)-1].Id)
            rank.RankItems = append(rank.RankItems[:len(rank.RankItems)-1], item)
            rank.RankItemByKey[item.Id] = item
        } else if len(rank.RankItems) < MAX_RANK_NUM {
            rank.RankItems = append(rank.RankItems, item)
            rank.RankItemByKey[item.Id] = item
        } else {
            return
        }
    } else {
        if item.Score == rankItem.Score {
            return
        }
        rankItem.Score = item.Score
    }
    sort.Sort(TItemSlice(rank.RankItems))
}
