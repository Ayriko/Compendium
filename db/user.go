package db

import (
	"Compendium/globals"
	"database/sql"
)

func GetUserById(id string) (globals.User, error) {

	stmt, err := DB.Prepare("SELECT id, username, password FROM user WHERE id = ?")
	//prepare et non direct query pour sécurité face aux injections, stmt = statement

	if err != nil {
		return globals.User{}, err //on renvoie une instance vide si l'id ne mène à rien
	}

	user := globals.User{} //instance de user

	//si on a pu créer le stmt alors on peut faire la query avec cet id
	sqlErr := stmt.QueryRow(id).Scan(&user.Id, &user.Username, &user.Password)

	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return globals.User{}, nil
		}
		return globals.User{}, sqlErr
	}
	return user, nil
}

func GetUserByUsername(name string) (globals.User, error) {

	stmt, err := DB.Prepare("SELECT id, username, password FROM user WHERE username = ?")
	//prepare et non direct query pour sécurité face aux injections, stmt = statement

	if err != nil {
		return globals.User{}, err //on renvoie une instance vide si l'id ne mène à rien
	}

	user := globals.User{} //instance de user

	//si on a pu créer le stmt alors on peut faire la query avec cet id
	sqlErr := stmt.QueryRow(name).Scan(&user.Id, &user.Username, &user.Password)

	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return globals.User{}, nil
		}
		return globals.User{}, sqlErr
	}
	return user, nil
}

func AddUser(newUser globals.User) (bool, error) {
	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("INSERT INTO user (id, username, password) VALUES (?, ?, ?)")
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(newUser.Id, newUser.Username, newUser.Password)
	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}

func GetCampaignTitleImage(id int) (string,string, error) {

	stmt, err := DB.Prepare("SELECT title, image FROM campaign WHERE id = ?")

	if err != nil {
		return "", "" ,err 
	}

	title := "" 
	image := ""

	sqlErr := stmt.QueryRow(id).Scan(&title, &image)

	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return "","", nil
		}
		return "", "", sqlErr
	}
	return title, image, nil
}