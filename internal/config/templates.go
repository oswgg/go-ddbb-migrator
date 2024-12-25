package config

const MigratorRCFile = `config_folder_path=config/`

const ConfigFile = `development:
	username: root
	password: root
	database: database
	host: localhost
	port: 3306
	dialect: mysql
production:
	username: $USERNAME
	password: $PASSWORD
	database: $DATABASE
	host: $HOST
	port: $PORT
	dialect: $DIALECT
test:
	username: $TEST_USERNAME
	password: $TEST_PASSWORD
	database: $TEST_DATABASE
	host: $TEST_HOST
	port: $TEST_PORT
	dialect: $TEST_DIALECT
`