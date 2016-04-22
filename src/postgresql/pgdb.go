package postgresql

import (
	"database/sql"
	"fmt"
	"util"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

func init() {
	host := "localhost"
	port := 5432
	dbID := "tantan"
	user := "postgres"
	pwd := "1111"
	sourceName := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable",
		user, pwd, host, port, dbID)
	db, err = sql.Open("postgres", sourceName)
	util.CheckErr(err)
	fmt.Println("connect ok")
}

func GetAllUsers() (u interface{}) {
	var users []util.UserType
	sql := "select id, name from tbl_user"

	// query
	rows, err := db.Query(sql)
	util.CheckErr(err)
	for rows.Next() {
		var id, name string
		rows.Scan(&id, &name)
		user := util.UserType{util.User{id, name}, "user"}
		users = append(users, user)
	}
	return users
}

func InsertUser(u util.User) string {
	var userid string
	err := db.QueryRow(
		fmt.Sprintf(`INSERT INTO tbl_user(name)VALUES('%s') RETURNING id`, u.Name),
	).Scan(&userid)
	util.CheckErr(err)

	return userid
}

func GetRelationships(id string) (u interface{}) {
	//	var rel []util.UserType
	//sql := fmt."select id, name from tbl_user"
	var relationShips []util.RelationshipType

	//userb=id
	rows, err := db.Query(fmt.Sprintf("select usera, state from tbl_relation where userb=%s", id))
	util.CheckErr(err)
	for rows.Next() {
		var id, state string
		rows.Scan(&id, &state)
		//		relation := util.Relationship{id, state}
		//		(&relation).ConvertState()
		relationShip := util.RelationshipType{util.Relationship{id, state}, "relationship"}
		relationShips = append(relationShips, relationShip)
	}

	// usera=id
	rows, err = db.Query(fmt.Sprintf("select userb, state from tbl_relation where usera=%s", id))
	util.CheckErr(err)
	for rows.Next() {
		var id, state string
		rows.Scan(&id, &state)
		//		relation := util.Relationship{id, state}
		//		(&relation).ConvertState()
		relationShip := util.RelationshipType{util.Relationship{id, state}, "relationship"}
		relationShips = append(relationShips, relationShip)
	}
	return relationShips
}

func PutRelationships(usera, userb string, state string) string {
	var state_new string
	//query relationship between usera with userb
	rows, err := db.Query(
		fmt.Sprintf(
			"select state from tbl_relation where (usera=%s and userb=%s) or (userb=%s and usera=%s)",
			usera, userb, usera, userb))
	util.CheckErr(err)

	// is alread have relationship
	if rows.Next() {
		var state_now string
		var sql string
		rows.Scan(&state_now)
		fmt.Println("state_now=", state_now)
		switch state_now {
		case "disliked":
			sql = fmt.Sprintf(
				"update tbl_relation set state='%s' where (usera='%s' and userb='%s') or (userb='%s' and usera='%s')",
				state, usera, userb, usera, userb)
			state_new = state
		case "liked", "matched":
			if state == "disliked" {
				sql = fmt.Sprintf(
					"update tbl_relation set state='%s' where (usera='%s' and userb='%s') or (userb='%s' and usera='%s')",
					"disliked", usera, userb, usera, userb)
				state_new = "disliked"
			} else {
				sql = fmt.Sprintf(
					"update tbl_relation set state='%s' where (usera='%s' and userb='%s') or (userb='%s' and usera='%s')",
					"matched", usera, userb, usera, userb)
				state_new = "matched"
			}
		default:
			fmt.Println("not ming zhong ")
		}
		fmt.Println(sql)
		db.QueryRow(sql)

	} else {
		db.QueryRow(
			fmt.Sprintf(`INSERT INTO tbl_relation(usera,userb,state)VALUES('%s','%s','%s')`,
				usera, userb, state))

		state_new = state
	}
	return state_new

}
