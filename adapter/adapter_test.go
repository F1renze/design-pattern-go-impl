package adapter

import "testing"

func TestAdapter(t *testing.T) {
	var adapter IPlatformAdapter

	adapter = NewAWSClientAdapter()

	adapter.CreateServer(1, 2)

	adapter = NewAliYunClientAdapter()

	adapter.CreateServer(1, 2)
}
