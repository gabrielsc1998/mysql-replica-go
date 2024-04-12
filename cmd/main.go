package main

import (
	"fmt"
	"time"

	"github.com/gabrielsc1998/mysql-replica-go/configs"
	"github.com/gabrielsc1998/mysql-replica-go/infra/mysql"
	"gorm.io/gorm"
)

func main() {
	config, err := configs.LoadConfig(".env")
	if err != nil {
		panic(err)
	}

	// ----- Write - Master ----- \\

	mysqlMaster := connectToMysqlMaster(config)
	defer mysqlMaster.Close()
	createUserInMaster(mysqlMaster)

	// ----- Read - Slave ----- \\

	time.Sleep(100 * time.Millisecond)

	mysqlSlave := connectToMysqlSlave(config)
	defer mysqlSlave.Close()
	users := getAllUsersFromSlave(mysqlSlave)

	if len(users) != 3 {
		fmt.Println("No users found in slave")
	} else {
		fmt.Println("\n----- Users found in slave -----")
		for i, user := range users {
			fmt.Printf(
				"User %d -> id: %d, name: %s, email: %s \n",
				i,
				user.ID,
				user.Name,
				user.Email,
			)
		}
	}
}

type Users struct {
	gorm.Model
	Name  string
	Email string
}

func connectToMysqlMaster(config *configs.Conf) *mysql.MySQLDB {
	mysqlMaster := mysql.NewMySQLDBConnection()
	err := mysqlMaster.Connect(mysql.MySQLConnectionOptions{
		Host:     config.DatabaseMasterHost,
		Port:     config.DatabaseMasterPort,
		User:     config.DatabaseMasterUser,
		Password: config.DatabaseMasterPass,
		Database: config.DatabaseMasterName,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to MySQL Master")

	// ----- Drop ----- \\

	err = mysqlMaster.DB.Migrator().DropTable(&Users{})
	if err != nil {
		panic(err)
	}

	// ----- Migrate ----- \\
	err = mysqlMaster.DB.AutoMigrate(&Users{})
	if err != nil {
		panic(err)
	}

	return mysqlMaster
}

func createUserInMaster(mysqlMaster *mysql.MySQLDB) {
	mysqlMaster.DB.Create(&Users{Name: "Gabriel", Email: "gabriel@gmail"})
	mysqlMaster.DB.Create(&Users{Name: "Lucas", Email: "lucas@gmail"})
	mysqlMaster.DB.Create(&Users{Name: "João", Email: "joao@gmail"})
	fmt.Println(" * Users created in Master *")
}

func connectToMysqlSlave(config *configs.Conf) *mysql.MySQLDB {
	mysqlSlave := mysql.NewMySQLDBConnection()
	err := mysqlSlave.Connect(mysql.MySQLConnectionOptions{
		Host:     config.DatabaseSlaveHost,
		Port:     config.DatabaseSlavePort,
		User:     config.DatabaseSlaveUser,
		Password: config.DatabaseSlavePass,
		Database: config.DatabaseSlaveName,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to MySQL Slave")
	return mysqlSlave
}

func getAllUsersFromSlave(mysqlSlave *mysql.MySQLDB) []Users {
	fmt.Println(" * Getting users from Slave... *")
	var users []Users
	mysqlSlave.DB.Find(&users)
	return users
}
