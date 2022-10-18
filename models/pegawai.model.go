package models

import (
	"myapp/db"
	"net/http"
)

type Pegawai struct {
	Id      int    `json:"id"`
	Nama    string `json:"nama"`
	Alamat  string `json:"alamat"`
	Telepon string `json:"telepon"`
}

func FetchAllPegawai() (Response, error) {
	var obj Pegawai
	var arrObjs []Pegawai
	var res Response

	con := db.CreateCon()

	sql := "SELECT * FROM pegawai"
	rows, err := con.Query(sql)

	defer rows.Close()
	if err != nil {
		return res, err
	}

	for rows.Next() {
		err := rows.Scan(&obj.Id, &obj.Nama, &obj.Alamat, &obj.Telepon)
		if err != nil {
			return res, err
		}

		arrObjs = append(arrObjs, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success!"
	res.Data = arrObjs

	return res, nil
}

func StorePegawai(nama string, alamat string, telepon string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sql := "INSERT INTO pegawai (nama, alamat, telepon) VALUES (?, ?, ?)"

	stmt, err := con.Prepare(sql)

	if err != nil {
		return res, err
	}
	result, err := stmt.Exec(nama, alamat, telepon)

	if err != nil {
		return res, err
	}

	lastInsertId, err := result.LastInsertId()

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success create data"
	res.Data = map[string]int64{
		"last_inserted_id": lastInsertId,
	}

	return res, nil
}

func UpdatePegawai(id int, nama string, alamat string, telepon string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sql := "UPDATE pegawai SET nama = ?, alamat = ?, telepon = ? WHERE id = ?"

	stmt, err := con.Prepare(sql)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nama, alamat, telepon, id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success update data"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}

func DeletePegawai(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sql := "DELETE FROM pegawai WHERE id = ?"

	stmt, err := con.Prepare(sql)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success update data"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}
