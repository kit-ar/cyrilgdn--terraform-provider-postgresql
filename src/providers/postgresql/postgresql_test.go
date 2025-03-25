// Adding tests for the new `search_path_database` feature

func TestSearchPathDatabase(t *testing.T) {
	// Test cases to cover:
	// - Roles with a large number of databases
	// - Roles with no databases
	// - Conflicts between globally and database-specific `search_path` settings
	...
}