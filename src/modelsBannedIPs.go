package main

import (
	"database/sql"
)

type BannedIPs struct{}

func (*BannedIPs) Check(ip string) (bool, error) {
	var id int
	if err := db.QueryRow(`
		SELECT id
		FROM banned_ips
		WHERE ip = ?
	`, ip).Scan(&id); err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

func (*BannedIPs) Insert(ip string, userID int) error {
	var stmt *sql.Stmt
	if v, err := db.Prepare(`
		INSERT INTO banned_ips (ip, user_id)
		VALUES (?, ?)
	`); err != nil {
		return err
	} else {
		stmt = v
	}
	defer stmt.Close()

	_, err := stmt.Exec(ip, userID)
	return err
}
