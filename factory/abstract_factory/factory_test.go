package abstract_factory

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeInstance(t *testing.T) {
	tt := "un1"

	i := makeInstance("MySQLUser")
	u, ok := i.(IUser)
	assert.True(t, ok)
	u.SetUserName(tt)
	assert.Equal(t, u.GetUserName(), tt)

	i = makeInstance("NotExist")
	u, ok = i.(IUser)
	assert.True(t, !ok)
}

func TestDataAccess_CreateUser(t *testing.T) {

	tc := []struct {
		env string
		err error
		un  string
	}{
		{"MySQL", nil, "jira"},
		{"NotExist", ErrDBTypeNotSupport, ""},
	}

	for _, tt := range tc {
		os.Setenv(DBEnv, tt.env)
		da := NewDataAccess()

		u, err := da.CreateUser()

		if tt.err != nil {
			assert.ErrorIs(t, tt.err, err)
			continue
		}
		assert.NoError(t, err)
		u.SetUserName(tt.un)
		assert.Equal(t, u.GetUserName(), tt.un)
	}

}
