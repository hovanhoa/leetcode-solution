func calPoints(operations []string) int {
    records := []int{}
    sum := 0
    for _, v := range operations {
        if v == "+" {
            twoLast := records[len(records)-1] + records[len(records)-2]
            records = append(records, twoLast)
            sum += twoLast
        } else if v == "D" {
            last := records[len(records)-1]
            records = append(records, last * 2)
            sum += last * 2            
        } else if v == "C" {
            sum -= records[len(records)-1]
            records = records[:len(records)-1]
        } else {
            i, _ := strconv.Atoi(v)
            records = append(records, i)
            sum += i
        }
    }

    return sum
}