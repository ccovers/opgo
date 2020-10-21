package main

// 插入最后位置或更新原有元素，向前向后一次比较大小并排序
func UpdateLastRank(rank *TRank, item *TItem, MAX_RANK_NUM int) {
    rankItem, _ := rank.RankItemByKey[item.Id]
    if nil == rankItem {
        if len(rank.RankItems) >= MAX_RANK_NUM && rank.RankItems[len(rank.RankItems)-1].Score < item.Score {
            lastItem := rank.RankItems[len(rank.RankItems)-1]
            delete(rank.RankItemByKey, lastItem.Id)
            item.Index = lastItem.Index
            rank.RankItems = append(rank.RankItems[:len(rank.RankItems)-1], item)
            rank.RankItemByKey[item.Id] = item
            rankItem = item
        } else if len(rank.RankItems) < MAX_RANK_NUM {
            // 排行榜未满，直接入榜
            item.Index = len(rank.RankItems)
            rank.RankItems = append(rank.RankItems, item)
            rank.RankItemByKey[item.Id] = item
            rankItem = item
        } else {
            return
        }
    } else {
        if item.Score == rankItem.Score {
            return
        }
        // 更新周活跃，并将排名置为当前索引
        rankItem.Score = item.Score
    }

    // 从公会当前排名开始，提升排名到公会对应位置
    for i := rankItem.Index; i >= 1; i-- {
        frontItem := rank.RankItems[i-1]
        nowItem := rank.RankItems[i]
        if nowItem.Score <= frontItem.Score {
            break
        }
        // 调整附近公会的排名和位置
        nowItem.Index, frontItem.Index = frontItem.Index, nowItem.Index
        rank.RankItems[nowItem.Index] = nowItem
        rank.RankItems[frontItem.Index] = frontItem
    }

    // 从公会当前排名开始，降低排名到公会对应位置
    for i := rankItem.Index; i < len(rank.RankItems)-1; i++ {
        nowItem := rank.RankItems[i]
        nextItem := rank.RankItems[i+1]
        if nowItem.Score >= nextItem.Score {
            break
        }
        // 调整附近公会的排名和位置
        nowItem.Index, nextItem.Index = nextItem.Index, nowItem.Index
        rank.RankItems[nowItem.Index] = nowItem
        rank.RankItems[nextItem.Index] = nextItem
    }
    return
}
