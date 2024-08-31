package main

// 假设每个数据项是一个包含帖子ID和评论ID的结构体
type DataItem struct {
	PostID    uint64
	CommentID uint64
}

func main() {
	// 假设这是你的数据集
	dataItems := []DataItem{
		{PostID: 245586098619678721, CommentID: 245618402763210753},
		{PostID: 245586098619678721, CommentID: 245618524398026753},
	}

	// 创建一个map，用于存储帖子ID和评论ID的映射
	postToIDsMap := make(map[uint64][]uint64)

	// 遍历数据集，填充map
	for _, item := range dataItems {
		// 检查map中是否已经有这个帖子ID的条目
		ids, exists := postToIDsMap[item.PostID]
		if !exists {
			// 如果没有，创建一个新的切片
			ids = []uint64{}
		}
		// 添加评论ID到对应的帖子ID下
		ids = append(ids, item.CommentID)
		// 更新map
		postToIDsMap[item.PostID] = ids
	}

}
