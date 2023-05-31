package leetcode

type LRUCache struct {
	cache map[int]int
	hotKey []int
	cap int
}


func Constructor(capacity int) LRUCache {
	return LRUCache{
		cache: make(map[int]int, 0),
		cap: capacity,
		hotKey: make([]int, 0),
	}
}


func (this *LRUCache) Get(key int) int {
	_, ok := this.cache[key]
	if ok {
		this.hotKey = rebuildHotKey(this.hotKey, key)
		this.hotKey = append(this.hotKey, key)
		return this.cache[key]
	}

	return -1
}


func (this *LRUCache) Put(key int, value int)  {
	if this.cap == 0 {
		return
	}
	if len(this.cache) == this.cap {
		if len(this.hotKey) == 0 {
			return
		}
		coldKey := this.hotKey[0]
		if coldKey != key {
			delete(this.cache, coldKey)
		}
		this.cache[key] = value
		this.hotKey = rebuildHotKey(this.hotKey, key)
		this.hotKey = append(this.hotKey, key)
	} else {
		this.cache[key] = value
		this.hotKey = rebuildHotKey(this.hotKey, key)
		this.hotKey = append(this.hotKey, key)
	}
}

func rebuildHotKey(nums []int, key int) []int {
	for k, v := range nums {
		if v == key {
			nums1 := nums[:k]
			nums2 := nums[k+1:]
			nums1 = append(nums1, nums2...)
			return nums1
		}
	}
	return nums
}