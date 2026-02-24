package structure

var MIGRATIONS = Node{
	Type:     Folder,
	Name:     "migrations",
	Resource: createUsersTableMigrations,
}

var createUsersTableMigrations = []Node{
	{
		Type:     File,
		Name:     "000001_create_users_table.up.sql",
		Template: "templates/migrations/create_users_table_up.temp",
	},
	{
		Type:     File,
		Name:     "000001_create_users_table.down.sql",
		Template: "templates/migrations/create_users_table_down.temp",
	},
}
