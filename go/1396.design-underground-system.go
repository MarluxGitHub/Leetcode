package leetcode

/*
 * @lc app=leetcode id=1396 lang=golang
 *
 * [1396] Design Underground System
 */

// @lc code=start
type UndergroundSystem struct {
	TravelRecord map[int]Travel
}

type Travel struct {
	StartStation, EndStation string
	StartTime, EndTime int
	CustomerID int
}


func Constructor() UndergroundSystem {
	return UndergroundSystem{
		TravelRecord: make(map[int]Travel),
	}
}


func (this *UndergroundSystem) CheckIn(id int, stationName string, t int)  {
	this.TravelRecord[len(this.TravelRecord)] = Travel{
		StartStation: stationName,
		StartTime: t,
		CustomerID: id,
	}
}


func (this *UndergroundSystem) CheckOut(id int, stationName string, t int)  {
	for i := len(this.TravelRecord) - 1; i >= 0; i-- {
		if this.TravelRecord[i].CustomerID == id && this.TravelRecord[i].EndStation == "" {
			travel := this.TravelRecord[i]
			travel.EndStation = stationName
			travel.EndTime = t
			this.TravelRecord[i] = travel
			break
		}
	}
}


func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	var totalTime, count int
	for _, travel := range this.TravelRecord {
		if travel.StartStation == startStation && travel.EndStation == endStation {
			totalTime += travel.EndTime - travel.StartTime
			count++
		}
	}
	return float64(totalTime) / float64(count)
}


/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */
// @lc code=end

