package servicenow

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

var httpRE = regexp.MustCompile("^https?://")

// Err represents a possible error message that came back from the server
type Err struct {
	Err    string `json:"error"`
	Reason string `json:"reason"`
}

func (e Err) Error() string {
	if e.Reason == "" {
		return e.Err
	}
	return fmt.Sprintf("%s: %s", e.Err, e.Reason)
}

// Client helps users interact with a ServiceNow instance.
type Client struct {
	Username, Password, Instance string
}

func sys(param string) string {
	return fmt.Sprintf("sysparm_%s", param)
}

// PerformFor creates and executes an authenticated HTTP request to an instance,
// for the given table, action and optional id, with the passed options, and
// unmarshal's the JSON into the passed output interface pointer, returning an
// error.
func (c *Client) PerformFor(table, action, id string, opts url.Values, body interface{}, out interface{}) error {
	inst := c.Instance

	if !httpRE.MatchString(inst) {
		inst = "https://" + inst
	}

	// This is a weird hack because the double encoding of url.Values onto the existing
	// url.Values causes issues with complex query values. For example:
	//  - sys_created_on>=javascript:gs.dateGenerate('2019-06-01','00:00:00')
	if opts != nil {
		values := make([]string, 0)
		for k, v := range opts {
			values = append(values, fmt.Sprintf("%s=%s", k, strings.Join(v, ",")))
			opts.Del(k)
		}
		opts.Add(sys("query"), strings.Join(values, "^"))
	} else {
		opts = url.Values{}
	}

	opts.Set("JSONv2", "")
	opts.Set("action", action)

	if id != "" {
		opts.Set(sys("sys_id"), id)
		opts.Set("displayvalue", "true")
	}

	meth := http.MethodGet
	buf := &bytes.Buffer{}

	if body != nil {
		meth = http.MethodPost
		if err := json.NewEncoder(buf).Encode(body); err != nil {
			return err
		}
	}

	req, err := http.NewRequest(meth, fmt.Sprintf("%s/%s.do?%s", inst, table, opts.Encode()), buf)
	if err != nil {
		return err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	req.SetBasicAuth(c.Username, c.Password)

	res, err := http.DefaultClient.Do(req)
	if res != nil {
		defer func() {
			if defErr := res.Body.Close(); defErr != nil {
				err = fmt.Errorf("%s::%s", err, defErr)
			}
		}()
	}
	if err != nil {
		return err
	}

	buf.Reset()

	// Store JSON so we can do a preliminary error check
	var echeck Err

	if err = json.NewDecoder(io.TeeReader(res.Body, buf)).Decode(&echeck); err == nil && echeck.Err != "" {
		return echeck
	}

	return json.NewDecoder(buf).Decode(out)
}

// GetFor performs a servicenow get to the specified table, with options, and
// unmarshal's JSON into the output parameter.
func (c Client) GetFor(table string, id string, opts url.Values, out interface{}) error {
	return c.PerformFor(table, "get", id, opts, nil, out)
}

// GetRecordsFor performs a servicenow getRecords to the specified table, with
// options, and unmarshal's JSON into the output parameter.
func (c Client) GetRecordsFor(table string, opts url.Values, out interface{}) error {
	return c.PerformFor(table, "getRecords", "", opts, nil, out)
}

// GetRecords performs a servicenow getRecords to the specified table, with
// options, and returns a map for each item
func (c Client) GetRecords(table string, opts url.Values) ([]map[string]interface{}, error) {
	var out struct {
		Records []map[string]interface{}
	}
	err := c.GetRecordsFor(table, opts, &out)
	return out.Records, err
}

// Insert creates a new record for the specified table, with the specified obj
// data, and takes a destination object out for the response data.
func (c Client) Insert(table string, obj, out interface{}) error {
	return c.PerformFor(table, "insert", "", nil, obj, out)
}

// Update creates a new record for the specified table, with the specified obj
// data, and takes a destination object out for the response data.
func (c Client) Update(table string, opts url.Values, obj, out interface{}) error {
	return c.PerformFor(table, "update", "", opts, obj, out)
}
