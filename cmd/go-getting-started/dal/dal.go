package dal

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"gopkg.in/gorp.v1"
	"os"
	"time"
)

var (
	//db    *sql.DB     = nil
	dbmap *gorp.DbMap = nil
)

type User struct {
	Id      uint64
	Email   string `binding:"required"`
	Updated time.Time
	Created time.Time
}

type Entry struct {
	Id      uint64
	UserId  uint64
	Name    string    `binding:"required"`
	Price   float64   `binding:"required"`
	Tags    []string  `binding:"required"`
	When    time.Time `binding:"required"`
	Updated time.Time
	Created time.Time
}

type Tag struct {
	Id   uint32
	Name string `binding:"required"`
}

type EntryTag struct {
	EntryId uint64
	TagId   uint32
}

func Connect() (err error) {
	var db *sql.DB
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return
	}
	dbmap = &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	dbmap.AddTable(User{}).
		SetKeys(true, "Id").
		ColMap("Email").MaxSize = 100

	tm := dbmap.AddTable(Entry{})
	tm.SetKeys(true, "Id")
	tm.ColMap("Tags").Transient = true
	tm.ColMap("Name").MaxSize = 200

	dbmap.AddTable(Tag{}).
		SetKeys(true, "Id").
		ColMap("Name").MaxSize = 50

	dbmap.AddTable(EntryTag{}).
		SetUniqueTogether("EntryId", "TagId")

	err = dbmap.CreateTablesIfNotExists()
	if err != nil {
		return
	}
	return
}

func FindUser(email string) (user User, err error) {
	dbmap.SelectOne(&user, "SELECT * FROM User where email=?", email)
	return
}
func AddUser(user_in User) (user_out User, err error) {
	// TODO: check email
	user_out = user_in
	err = dbmap.Insert(&user_out)
	/*if err == nil {
	    user_out, err = FindUser(user_in.Email)
	}*/
	return
}

func GetOrAddUser(email string) (user User, err error) {
	user, err = FindUser(email)

	if err == sql.ErrNoRows {
		user.Email = email
		user, err = AddUser(user)
	}
	return
}

func AddEntry(entry_in Entry) (entry_out Entry, err error) {
	// TODO: check entry
	fmt.Printf("AddEntry: %+v\n", entry_in)
	entry_out = entry_in
	err = dbmap.Insert(&entry_out)
	if err == nil {
		for _, tag_name := range entry_out.Tags {
			var tag Tag
			tag, err = GetOrAddTag(tag_name)
			entry_tag := EntryTag{EntryId: entry_out.Id, TagId: tag.Id}
			dbmap.Insert(&entry_tag)
		}
	}
	//if err == nil {
	//  user_out, err = FindUser(c, user_in.Email)
	//}
	return
}

func FindTag(name string) (tag Tag, err error) {
	dbmap.SelectOne(&tag, "SELECT * FROM Tag where name=?", name)
	return
}
func AddTag(tag_in Tag) (tag_out Tag, err error) {
	// TODO: check name
	tag_out = tag_in
	err = dbmap.Insert(&tag_out)
	/*if err == nil {
	    user_out, err = FindUser(user_in.Email)
	}*/
	return
}

func GetOrAddTag(name string) (tag Tag, err error) {
	tag, err = FindTag(name)

	if err == sql.ErrNoRows {
		tag.Name = name
		tag, err = AddTag(tag)
	}
	return
}

/*
func dbFunc(c *gin.Context) {
    if _, err := db.Exec("CREATE TABLE IF NOT EXISTS ticks (tick timestamp)"); err != nil {
        c.String(http.StatusInternalServerError,
            fmt.Sprintf("Error creating database table: %q", err))
        return
    }

    if _, err := db.Exec("INSERT INTO ticks VALUES (now())"); err != nil {
        c.String(http.StatusInternalServerError,
            fmt.Sprintf("Error incrementing tick: %q", err))
        return
    }

    rows, err := db.Query("SELECT tick FROM ticks")
    if err != nil {
        c.String(http.StatusInternalServerError,
            fmt.Sprintf("Error reading ticks: %q", err))
        return
    }

    defer rows.Close()
    for rows.Next() {
        var tick time.Time
        if err := rows.Scan(&tick); err != nil {
            c.String(http.StatusInternalServerError,
                fmt.Sprintf("Error scanning ticks: %q", err))
            return
        }
        c.String(http.StatusOK, fmt.Sprintf("Read from DB: %s\n", tick.String()))
    }
}
*/
