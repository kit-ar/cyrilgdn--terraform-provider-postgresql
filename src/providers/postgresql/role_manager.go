// Updating the ManageRole function to handle the new `search_path_database` property

func (rm *RoleManager) ManageRole(r Role) error {
	...
	// Handling the search_path_database
	for db, path := range r.SearchPathDatabase {
		rm.db.Exec("SET search_path TO " + path + " FOR " + db)
	}
	...
}