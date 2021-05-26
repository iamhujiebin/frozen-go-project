package main

import "encoding/json"

func SendRedRain_Bak(redId, goldId, num []int64, step float64, level int32, dur int64, token string, uid int64, liveid string, mark int64) []map[string]interface{} {
	ms := make([]map[string]interface{}, 0)
	ms = append(ms, map[string]interface{}{
		"tp": "red_rain",
		"data": map[string]interface{}{
			"red_id":     redId,
			"golden_id":  goldId,
			"red_num":    len(redId),
			"level":      level,
			"duration":   dur,
			"step":       step,
			"step_count": num,
			"token":      token,
			"mark_id":    mark,
		},
	})

	return []map[string]interface{}{
		{
			"b": map[string]interface{}{
				"ev": "s.m",
			},
			"dest":   2,
			"userid": uid,
			"liveid": liveid,
			"ms":     ms,
		},
	}
}

func main() {
	res := SendRedRain_Bak([]int64{1, 2, 3}, []int64{4, 5, 6}, []int64{1}, 1, 1, 60, "", 100246, "1916128840000010272", 1)
	j, _ := json.Marshal(res)
	println(string(j))
}
