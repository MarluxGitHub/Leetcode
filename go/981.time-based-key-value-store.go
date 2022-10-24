package leetcode

/*
 * @lc app=leetcode id=981 lang=golang
 *
 * [981] Time Based Key-Value Store
 */

// @lc code=start
type TimeMap struct {
	Data map[string][]TimeMapEntry
}

type TimeMapEntry struct {
	Value string
	Timestamp int
}

func Constructor() TimeMap {
	return TimeMap{Data: make(map[string][]TimeMapEntry)}
}


func (this *TimeMap) Set(key string, value string, timestamp int)  {
	this.Data[key] = append(this.Data[key], TimeMapEntry{Value: value, Timestamp: timestamp})
}


func (this *TimeMap) Get(key string, timestamp int) string {
	entries, ok := this.Data[key]
	if !ok {
		return ""
	}
	return binarySearch(entries, timestamp)
}

func binarySearch(entries []TimeMapEntry, timestamp int) string {
	low, high := 0, len(entries)
	for low < high {
		mid := low + (high - low) / 2
		if entries[mid].Timestamp == timestamp {
			return entries[mid].Value
		}
		if entries[mid].Timestamp < timestamp {
			low = mid + 1
		} else {
			high = mid
		}
	}
	if low > 0 {
		return entries[low-1].Value
	}
	return ""
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
// @lc code=end

