package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

// {"status":"success","data":{"resultType":"vector","result":[{"metric":{},"value":[1579065303.357,"17"]}]}}
type dataType struct {
	ResultType string    `json:"resultType"`
	Result     []resultT `json:"result"`
}

type resultT struct {
	Metric map[string]string `json:"metric"`
	Value  valueT            `json:"value"`
}

type valueT struct {
	TimeStamp float32
	Value     string
}

func (v *valueT) UnmarshalJSON(buf []byte) error {
	tmp := []interface{}{&v.TimeStamp, &v.Value}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}
	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields in Notification: %d != %d", g, e)
	}
	return nil
}

func prometheusFetcher(u *url.URL) (float64, error) {
	resp, err := http.Get(u.String())
	if err != nil {
		log.Fatalln(err)
		return 0, err
	}
	defer resp.Body.Close()

	var result struct {
		Status string   `json:"status"`
		Data   dataType `json:"data"`
	}

	json.NewDecoder(resp.Body).Decode(&result)
	val := result.Data.Result[0].Value.Value
	f, err := strconv.ParseFloat(val, 64)
	if err != nil {
		log.Fatalln(err)
		return 0, err
	}

	return f, nil
}
