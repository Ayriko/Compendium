package db

import (
	"Compendium/globals"
)

func AddCampaignDB(name string, image string, description string) (bool, error) {
	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("INSERT INTO campaign (title, description, image) VALUES (?, ?, ?)")
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(name, description, image)
	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}

func UpCampaignDB(name string, image string, description string, id int) (bool, error) {
	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("update campaign set title=?, description=?, image=? where id=?")
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(name, description, image, id)
	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}

func GetCampaignDB() ([]globals.Campaign, error) {
	rows, err := DB.Query("SELECT * FROM campaign")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	campaigns := make([]globals.Campaign, 0) //ensemble de campagnes

	for rows.Next() {
		singleCampaign := globals.Campaign{} //on crée une instance d'une campagne
		err = rows.Scan(&singleCampaign.Id, &singleCampaign.Title, &singleCampaign.Description, &singleCampaign.Image)
		if err != nil {
			return nil, err
		}

		campaigns = append(campaigns, singleCampaign)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return campaigns, err
}

func GetCampaignIdDB(id int) ([]globals.Campaign, error) {
	rows, err := DB.Query("SELECT * FROM campaign WHERE id = ?", id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	campaigns := make([]globals.Campaign, 0) 

	for rows.Next() {
		singleCampaign := globals.Campaign{} 
		err = rows.Scan(&singleCampaign.Id, &singleCampaign.Title, &singleCampaign.Description, &singleCampaign.Image)
		if err != nil {
			return nil, err
		}

		campaigns = append(campaigns, singleCampaign)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return campaigns, err
}

func DeleteCampaignDB(id int) (bool, error) {
	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}

	stmt, err := DB.Prepare(`DELETE FROM campaign WHERE id = ?;`)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
} 

func AddWorldDB(name string, image string, description string, id int) (bool, error) {
	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("INSERT INTO world (title, content, image, compaign_id) VALUES (?, ?, ?, ?)")
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(name, description, image, id)
	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}

func GetWorldDB(campaignId int) ([]globals.World, error){
	rows, err := DB.Query("SELECT id, title, content, image FROM world WHERE compaign_id = ?", campaignId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	worlds := make([]globals.World, 0) 

	for rows.Next() {
		singleWorld := globals.World{} 
		err = rows.Scan(&singleWorld.Id, &singleWorld.Title, &singleWorld.Content, &singleWorld.Image)
		if err != nil {
			return nil, err
		}

		worlds = append(worlds, singleWorld)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return worlds, err
}

func UpWorldDB(name string, image string, description string, id int) (bool, error) {
	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("UPDATE world SET title = ? content = ? image = ? WHERE id = ?")
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(name, description, image, id)
	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}

func DeleteWorldDB(id int) (bool, error) {
	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}

	stmt, err := DB.Prepare(`DELETE FROM world WHERE id = ?;`)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}

func AddCharacterDB(name string, image string, description string, id int) (bool, error) {
	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("INSERT INTO characters (name, content, image, compaign_id) VALUES (?, ?, ?, ?)")
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(name, description, image, id)
	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}

func UpCharDB(name string, image string, description string, id int) (bool, error) {
	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("UPDATE characters SET name = ? content = ? image = ? WHERE id = ?")
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(name, description, image, id)
	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}

func GetCharactersDB(campaignId int) ([]globals.Characters, error){
	rows, err := DB.Query("SELECT id, name, content, image FROM characters WHERE compaign_id = ?", campaignId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	characters := make([]globals.Characters, 0) 

	for rows.Next() {
		singleCharacter := globals.Characters{} 
		err = rows.Scan(&singleCharacter.Id, &singleCharacter.Name, &singleCharacter.Content, &singleCharacter.Image)
		if err != nil {
			return nil, err
		}

		characters = append(characters, singleCharacter)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return characters, err
}

func DeleteCharacterDB(id int) (bool, error) {
	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}

	stmt, err := DB.Prepare(`DELETE FROM characters WHERE id = ?;`)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}

func AddLoreDB(name string, image string, description string, id int) (bool, error) {
	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("INSERT INTO lore (title, content, image, compaign_id) VALUES (?, ?, ?, ?)")
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(name, description, image, id)
	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}

func UpLoreDB(name string, image string, description string, id int) (bool, error) {
	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("UPDATE lore SET title = ? image = ? content = ? WHERE id = ?")
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(name, image, description, id)
	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}

func GetLoreDB(campaignId int) ([]globals.Lore, error){
	rows, err := DB.Query("SELECT id, title, image, content FROM lore WHERE compaign_id = ?", campaignId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	lores := make([]globals.Lore, 0) 

	for rows.Next() {
		singleLore := globals.Lore{} 
		err = rows.Scan(&singleLore.Id, &singleLore.Title, &singleLore.Image, &singleLore.Content)
		if err != nil {
			return nil, err
		}

		lores = append(lores, singleLore)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return lores, err
}

func DeleteLoreDB(id int) (bool, error) {
	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}

	stmt, err := DB.Prepare(`DELETE FROM lore WHERE id = ?;`)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}


