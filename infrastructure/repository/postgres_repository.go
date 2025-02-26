package repository

import (
	"database/sql"

	"github.com/Magowtham/telephone_recharge_machine_http_server/domain/entity"
)

type PostgresRepository struct {
	database *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{
		database: db,
	}
}

func (repo *PostgresRepository) Init() error {
	query1 := `CREATE TABLE IF NOT EXISTS admins(
				admin_id VARCHAR(255) PRIMARY KEY,
				admin_name VARCHAR(255) NOT NULL UNIQUE,
				password VARCHAR(255) NOT NULL,
				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
			)`

	query2 := `CREATE TABLE IF NOT EXISTS machines(
				machine_id VARCHAR(255) PRIMARY KEY,
				admin_id VARCHAR(255) NOT NULL,
				label VARCHAR(255) NOT NULL,
				balance INTEGER NOT NULL,
				update_timestamp VARCHAR(255) NOT NULL,
				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				FOREIGN KEY (admin_id) REFERENCES admins(admin_id) ON DELETE CASCADE
			)`
	query3 := `CREATE TABLE IF NOT EXISTS users(
				user_id VARCHAR(255) PRIMARY KEY,
				admin_id VARCHAR(255) NOT NULL,
				machine_id VARCHAR(255) NOT NULL,
				email VARCHAR(255) NOT NULL UNIQUE, 
				user_name VARCHAR(255) NOT NULL UNIQUE,
				password VARCHAR(255) NOT NULL,
				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				FOREIGN KEY (machine_id) REFERENCES machines(machine_id) ON DELETE CASCADE,
				FOREIGN KEY (admin_id) REFERENCES admins(admin_id) ON DELETE CASCADE
			)`

	query4 := `CREATE TABLE IF NOT EXISTS recharge_history(
				machine_id VARCHAR(255) NOT NULL,
				amount INTEGER NOT NULL,
				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				FOREIGN KEY (machine_id) REFERENCES machines(machine_id) ON DELETE CASCADE
			)`

	query5 := `CREATE TABLE IF NOT EXISTS expense_history(
				machine_id VARCHAR(255) NOT NULL,
				amount INTEGER NOT NULL,
				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				FOREIGN KEY (machine_id) REFERENCES machines(machine_id) ON DELETE CASCADE
			)`

	tx, err := repo.database.Begin()

	if err != nil {
		return err
	}

	if _, err := tx.Exec(query1); err != nil {
		tx.Rollback()
		return err
	}

	if _, err := tx.Exec(query2); err != nil {
		tx.Rollback()
		return err
	}

	if _, err := tx.Exec(query3); err != nil {
		tx.Rollback()
		return err
	}

	if _, err := tx.Exec(query4); err != nil {
		tx.Rollback()
		return err
	}

	if _, err := tx.Exec(query5); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (repo *PostgresRepository) CheckAdminNameExists(adminName string) (bool, error) {
	var adminExists bool
	query := `SELECT EXISTS(SELECT 1 FROM admins WHERE admin_name = $1 )`
	err := repo.database.QueryRow(query, adminName).Scan(&adminExists)
	return adminExists, err
}

func (repo *PostgresRepository) CreateAdmin(admin *entity.Admin) error {
	query := `INSERT INTO admins (admin_id,admin_name,password) VALUES ($1,$2,$3)`
	_, err := repo.database.Exec(query, admin.AdminId, admin.AdminName, admin.Password)
	return err
}

func (repo *PostgresRepository) CheckAdminIdExists(adminId string) (bool, error) {
	var adminIdExists bool
	query := `SELECT EXISTS(SELECT 1 FROM admins WHERE admin_id = $1)`
	err := repo.database.QueryRow(query, adminId).Scan(&adminIdExists)
	return adminIdExists, err
}

func (repo *PostgresRepository) DeleteAdmin(adminId string) error {
	query := `DELETE FROM admins WHERE admin_id = $1`
	_, err := repo.database.Exec(query, adminId)
	return err
}

func (repo *PostgresRepository) GetAdminByAdminName(adminName string) (*entity.Admin, error) {
	var admin entity.Admin
	query := `SELECT admin_id,admin_name,password FROM admins WHERE admin_name = $1`
	err := repo.database.QueryRow(query, adminName).Scan(&admin.AdminId, &admin.AdminName, &admin.Password)
	return &admin, err
}

func (repo *PostgresRepository) CheckUserNameExists(userName string) (bool, error) {
	var userExists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE user_name = $1)`
	err := repo.database.QueryRow(query, userName).Scan(&userExists)
	return userExists, err
}

func (repo *PostgresRepository) CreateUser(user *entity.User) error {
	query := `INSERT INTO users (user_id,admin_id,machine_id,email,user_name,password) VALUES ($1,$2,$3,$4,$5,$6)`
	_, err := repo.database.Exec(query, user.UserId, user.AdminId, user.MachineId, user.Email, user.UserName, user.Password)
	return err
}

func (repo *PostgresRepository) CheckUserIdExists(userId string) (bool, error) {
	var userIdExists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE user_id=$1)`
	err := repo.database.QueryRow(query, userId).Scan(&userIdExists)
	return userIdExists, err
}

func (repo *PostgresRepository) DeleteUser(userId string) error {
	query := `DELETE FROM users WHERE user_id = $1`
	_, err := repo.database.Exec(query, userId)
	return err
}

func (repo *PostgresRepository) GetAllUsers(adminId string) ([]*entity.User, error) {
	query := `SELECT user_id,machine_id,email,user_name FROM users WHERE admin_id=$1`
	rows, err := repo.database.Query(query, adminId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []*entity.User

	for rows.Next() {
		var user entity.User
		if err := rows.Scan(&user.UserId, &user.MachineId, &user.Email, &user.UserName); err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}

func (repo *PostgresRepository) GetUserByUserName(userName string) (*entity.User, error) {
	var user entity.User
	query := `SELECT user_id,machine_id,user_name,password FROM users WHERE user_name = $1`
	err := repo.database.QueryRow(query, userName).Scan(&user.UserId, &user.MachineId, &user.UserName, &user.Password)
	return &user, err
}

func (repo *PostgresRepository) CheckMachineIdExists(machineId string) (bool, error) {
	var machineIdExists bool

	query := `SELECT EXISTS(SELECT 1 FROM machines WHERE machine_id = $1)`

	err := repo.database.QueryRow(query, machineId).Scan(&machineIdExists)

	return machineIdExists, err
}

func (repo *PostgresRepository) CreateMachine(machine *entity.Machine) error {
	query := `INSERT INTO machines (machine_id,admin_id,label,balance,update_timestamp) VALUES ($1,$2,$3,$4,$5)`
	_, err := repo.database.Exec(query, machine.MachineId, machine.AdminId, machine.Label, machine.Balance, machine.UpdateTimestamp)
	return err
}

func (repo *PostgresRepository) DeleteMachine(machineId string) error {
	query := `DELETE FROM machines WHERE machine_id = $1`
	_, err := repo.database.Exec(query, machineId)
	return err
}

func (repo *PostgresRepository) GetMachinesByAdminId(adminId string) ([]*entity.Machine, error) {
	var machines []*entity.Machine

	query := `SELECT machine_id,label,balance,update_timestamp FROM machines WHERE admin_id = $1`
	rows, err := repo.database.Query(query, adminId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var machine entity.Machine

		if err := rows.Scan(&machine.MachineId, &machine.Label, &machine.Balance, &machine.UpdateTimestamp); err != nil {
			return nil, err
		}

		machines = append(machines, &machine)
	}

	return machines, nil
}

func (repo *PostgresRepository) GetMachineIdsByAdminId(adminId string) ([]string, error) {
	query := `SELECT machine_id FROM machines WHERE admin_id=$1`

	rows, err := repo.database.Query(query, adminId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var machineIds []string

	for rows.Next() {
		var machineId string
		if err := rows.Scan(&machineId); err != nil {
			return machineIds, nil
		}

		machineIds = append(machineIds, machineId)
	}

	return machineIds, nil
}

func (repo *PostgresRepository) RechargeMachine(machineId string, amount int32) error {
	query1 := `UPDATE machines SET balance = balance + $2 WHERE machine_id = $1`
	query2 := `INSERT INTO recharge_history (machine_id,amounrt) VALUES ($1,$2)`

	tx, err := repo.database.Begin()

	if err != nil {
		return err
	}

	if _, err := tx.Exec(query1, machineId, amount); err != nil {
		tx.Rollback()
		return err
	}

	if _, err := tx.Exec(query2, machineId, amount); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (repo *PostgresRepository) GetMachineBalance(machineId string) (int32, error) {
	var balance int32
	query := `SELECT balance FROM machines WHERE machine_id = $1`
	err := repo.database.QueryRow(query, machineId).Scan(&balance)
	return balance, err
}

func (repo *PostgresRepository) GetRechargeHistoryByMachineId(machineId string) ([]*entity.RechargeHistory, error) {
	var rechargeHistories []*entity.RechargeHistory

	query := `SELECT amount,created_at FROM recharge_history WHERE machine_id = $1`
	rows, err := repo.database.Query(query, machineId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var rechargeHistory entity.RechargeHistory

		if err := rows.Scan(&rechargeHistory.Amount, &rechargeHistory.Timestamp); err != nil {
			return nil, err
		}

		rechargeHistories = append(rechargeHistories, &rechargeHistory)
	}

	return rechargeHistories, nil
}

func (repo *PostgresRepository) DeductMachineBalance(machineId string, amount int32) error {
	query1 := `UPDATE machines SET balance = balance - $2 WHERE machine_id=$1 AND balance >= $2`
	query2 := `INSERT INTO expense_history (machine_id,amount) VALUES ($1,$2)`

	tx, err := repo.database.Begin()

	if err != nil {
		return err
	}

	if _, err := tx.Exec(query1, machineId, amount); err != nil {
		tx.Rollback()
		return err
	}

	if _, err := tx.Exec(query2, machineId, amount); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (repo *PostgresRepository) GetExpenseHistoryByMachineId(machineId string) ([]*entity.ExpenseHistory, error) {
	var expenseHistories []*entity.ExpenseHistory

	query := `SELECT amount,created_at FROM expense_history WHERE machine_id = $1`

	rows, err := repo.database.Query(query, machineId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var expenseHistory entity.ExpenseHistory

		if err := rows.Scan(&expenseHistory.Amount, &expenseHistory.Timestamp); err != nil {
			return nil, err
		}

		expenseHistories = append(expenseHistories, &expenseHistory)
	}

	return expenseHistories, nil
}
