type pair struct {
    value string
    timestamp int
}
type TimeMap struct {
    m map[string][]pair
}


func Constructor() TimeMap {
    return TimeMap{
        map[string][]pair{},
    }
}


func (this *TimeMap) Set(key string, value string, timestamp int)  {
    this.m[key] = append(this.m[key], pair{value, timestamp})
}


func (this *TimeMap) Get(key string, timestamp int) string {
    list, exists := this.m[key]
    if !exists {
        return ""
    }

    if len(list) == 0 {
        return ""
    }

    p := 0
    q := len(list)-1

    for p <= q {
        mid := (p+q)/2
        if list[mid].timestamp > timestamp {
            q = mid - 1
        } else {
            p = mid + 1
        }
    }

    if q < 0 {
        return ""
    }
    return list[q].value
}
