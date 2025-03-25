// Adding a new property `search_path_database` to the Role struct

// Role represents a PostgreSQL role

type Role struct {
	...
	// The search path to be used within a specific database for this role
	SearchPathDatabase map[string]string `json:"search_path_database,omitempty"`
	...
}

// A helper function to handle the search path for a specific database in a role

func (r *Role) SetSearchPathForDatabase(db string, path string) error {
	if r.SearchPathDatabase == nil {
		r.SearchPathDatabase = make(map[string]string)
	}
	r.SearchPathDatabase[db] = path
	return nil
}