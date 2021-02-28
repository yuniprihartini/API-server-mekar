package utils 

const (
	CREATE_USER = "INSERT INTO user(id_user,nik,username,tgl_lahir,id_pekerjaan,id_pendidikan) values (?,?,?,?,?,?)"
	SELECT_USERS=`SELECT 
	u.id_user,
	u.nik,
    u.username,
    u.tgl_lahir,
    p.id_pendidikan,
    p.jenjang_pendidikan,
    u.id_pekerjaan,
    pk.nama_pekerjaan,
    u.user_status,
    u.created_date,
    u.update_date
    FROM user as u INNER JOIN pendidikan as p
    ON u.id_pendidikan = p.id_pendidikan inner join
    pekerjaan as pk on u.id_pekerjaan = pk.id_pekerjaan where u.user_status = 1 LIMIT ?,?`
    SELECT_USER_BY_ID = `SELECT 
	u.id_user,
    u.nik,
    u.username,
    u.tgl_lahir,
    p.id_pendidikan,
    p.jenjang_pendidikan,
    u.id_pekerjaan,
    pk.nama_pekerjaan,
    u.user_status,
    u.created_date,
    u.update_date
    FROM user as u INNER JOIN pendidikan as p
    ON u.id_pendidikan = p.id_pendidikan inner join
    pekerjaan as pk on u.id_pekerjaan = pk.id_pekerjaan
	WHERE u.id_user = ?`
	UPDATE_USER = `UPDATE  user set nik=?,username=?,tgl_lahir=?,id_pekerjaan=?,id_pendidikan=? WHERE id_user = ?`
	DELETE_USER = `delete from user where id_user=?`
	SELECT_COUNT_DATA_USER = `SELECT count(*) as total_data FROM db_mekar.user`
	SELECT_PEKERJAAN = `SELECT * FROM pekerjaan`
	SELECT_PENDIDIKAN = `SELECT * FROM pendidikan`

	CREATE_ACCOUNT = `INSERT INTO akun(id_akun,email,password) values (?,?,?)`
	SELECT_ACCOUNT_BY_EMAIL = `SELECT * FROM akun where email=?`

)