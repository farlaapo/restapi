package db

import (
	"Somali-Newsletter-App/pkg/config"
	"database/sql"
	"fmt"
	"log"
)


func ConnectDB(cfg *config.DBConfig)  (*sql.DB, error ) {
	constr := cfg.ConnectionString()

	db, err := sql.Open("postgres", constr)
	if err != nil {
		return nil, fmt.Errorf("failed to connnect the database: %v", err)

	}

	// ping 
	if err := db.Ping(); err != nil {

		return nil, fmt.Errorf("failed to ping the database: %v", err)
	}

	log.Println("succesfully connected to the database")
	return db, nil
}

func CreateTables (db *sql.DB) error {



	// Create token table 
	tokenTable := `CREATE TABLE IF NOT EXISTS tokens (
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		user_id UUID REFERENCES users(id) ON DELETE CASCADE,
		token VARCHAR(255) NOT NULL,
		expires_at TIMESTAMP NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		deleted_at TIMESTAMP
	

	
	)`

	//  Create permission table
	permissionTable := `CREATE TABLE IF NOT EXIST permissions (
		id UUID PRIMARY KEY,
		name VARCHAR(255) UINQUE NOT NULL
	);`

	// create userPermission table
	userPermissionTable := `CREATE TABLE IF NOT EXISTS user_permissions (
	 user_id UUID REFERENCES user(id) ON DELETE CASCADE,
	 permission_id INT REFERENCES permissions(id) ON DELETE CASCADE,
	 PRIMARY KEY (user_id, permission_id)
	);`

	roleTable := `CREATE TABLE IF NOT EXISTS roles (
		id UUID PRIMARY KEY,
		name VARCHAR(255) UNIQUE NOT NULL

	);`

	userRoleTable := `CREATE TABLE F NOT EXISTS user_roles (
	 user_id UUID REFERENCES user(id) ON DELETE CASCADE,
	 role_id UUID REFERENCES roles(id) ON DELETE CASCADE,
	 PRIMARY KEY (user_id, role_id)
	
	)`














	queries := []string{ permissionTable, userPermissionTable, roleTable, userRoleTable, tokenTable}
	for _, query := range queries {
		if _, err := db.Exec(query); err != nil {
			return  fmt.Errorf("failed to create al tables") 
		}
	}


	log.Println("successfully created all tables ")
	return nil

}