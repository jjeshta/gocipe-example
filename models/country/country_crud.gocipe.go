// generated by gocipe 2684eeb0295ee5f2f4177b7c406ae5c8553e74d3b44b524a3eeb29e37d1363f0; DO NOT EDIT

package country

import (
	"context"
	"database/sql"
	"strconv"
	"strings"

	"github.com/fluxynet/gocipe-example/models"
	"github.com/gobuffalo/uuid"
)

var db *sql.DB

// Inject allows injection of services into the package
func Inject(database *sql.DB) {
	db = database
}

// Get returns a single Country from database by primary key
func Get(ctx context.Context, id string) (*Country, error) {
	var (
		rows   *sql.Rows
		err    error
		entity = New()
	)

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		break
	}

	rows, err = db.Query("SELECT id, name, continent FROM  WHERE id = $1 ORDER BY .id ASC", id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	if rows.Next() {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			err = rows.Scan(entity.ID, entity.Name, entity.Continent)
			if err != nil {
				return nil, err
			}
		}
	}

	return entity, nil
}

// List returns a slice containing Country records
func List(ctx context.Context, filters []models.ListFilter) ([]*Country, error) {
	var (
		list     []*Country
		segments []string
		values   []interface{}
		err      error
		rows     *sql.Rows
	)

	query := "SELECT id, name, continent FROM "

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		break
	}

	for i, filter := range filters {
		segments = append(segments, filter.Field+" "+filter.Operation+" $"+strconv.Itoa(i+1))
		values = append(values, filter.Value)
	}

	if len(segments) != 0 {
		query += " WHERE " + strings.Join(segments, " AND ")
	}

	rows, err = db.Query(query+" ORDER BY id ASC", values...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			break
		}

		entity := New()
		err = rows.Scan(entity.ID, entity.Name, entity.Continent)
		if err != nil {
			return nil, err
		}

		list = append(list, entity)
	}

	return list, nil
}

// Delete deletes a Country record from database and sets id to nil
func (entity *Country) Delete(ctx context.Context, tx *sql.Tx, autocommit bool) error {
	var (
		err  error
		stmt *sql.Stmt
	)
	id := *entity.ID

	if tx == nil {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			tx, err = db.Begin()
			if err != nil {
				return err
			}
		}
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		stmt, err = tx.Prepare("DELETE FROM  WHERE id = $1")
		if err != nil {
			return err
		}
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		_, err = stmt.Exec(id)
		if err == nil {
			entity.ID = nil
		} else {
			tx.Rollback()
			return err
		}
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		if autocommit {
			err = tx.Commit()
		}
	}

	return nil
}

// Delete deletes many Country records from database using filter
func Delete(ctx context.Context, filters []models.ListFilter, tx *sql.Tx, autocommit bool) error {
	var (
		err      error
		stmt     *sql.Stmt
		segments []string
		values   []interface{}
	)

	if tx == nil {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			tx, err = db.Begin()
			if err != nil {
				return err
			}
		}
	}

	query := "DELETE FROM "

	select {
	case <-ctx.Done():
		tx.Rollback()
		return ctx.Err()
	default:
		break
	}

	for i, filter := range filters {
		segments = append(segments, filter.Field+" "+filter.Operation+" $"+strconv.Itoa(i+1))
		values = append(values, filter.Value)
	}

	if len(segments) != 0 {
		query += " WHERE " + strings.Join(segments, " AND ")
	}

	stmt, err = db.Prepare(query)
	if err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		_, err = stmt.Exec(values...)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	select {
	case <-ctx.Done():
		tx.Rollback()
		return ctx.Err()
	default:
		if autocommit {
			err = tx.Commit()
		}
	}

	return err
}

// Save either inserts or updates a Country record based on whether or not id is nil
func (entity *Country) Save(ctx context.Context, tx *sql.Tx, autocommit bool) error {
	if entity.ID == nil {
		return entity.Insert(ctx, tx, autocommit)
	}
	return entity.Update(ctx, tx, autocommit)
}

// Insert performs an SQL insert for Country record and update instance with inserted id.
func (entity *Country) Insert(ctx context.Context, tx *sql.Tx, autocommit bool) error {
	var (
		id   string
		err  error
		stmt *sql.Stmt
	)

	if tx == nil {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			tx, err = db.Begin()
			if err != nil {
				return err
			}
		}
	}

	select {
	case <-ctx.Done():
		tx.Rollback()
		return ctx.Err()
	default:
		break
	}

	stmt, err = tx.Prepare("INSERT INTO  (id, name, continent) VALUES ($0, $1, $1)")
	if err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		break
	}

	idUUID, err := uuid.NewV4()

	if err == nil {
		id = idUUID.String()
	} else {
		tx.Rollback()
		return err
	}
	*entity.ID = id

	select {
	case <-ctx.Done():
		tx.Rollback()
		return ctx.Err()
	default:
		break
	}

	_, err = stmt.Exec(*entity.ID, *entity.Name, *entity.Continent)
	if err != nil {
		tx.Rollback()
		return err
	}

	select {
	case <-ctx.Done():
		tx.Rollback()
		return ctx.Err()
	default:
		if autocommit {
			err = tx.Commit()
		}
	}

	return nil
}

// Update Will execute an SQLUpdate Statement for Country in the database. Prefer using Save instead of Update directly.
func (entity *Country) Update(ctx context.Context, tx *sql.Tx, autocommit bool) error {
	var (
		err  error
		stmt *sql.Stmt
	)

	if tx == nil {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			tx, err = db.Begin()
			if err != nil {
				return err
			}
		}
	}

	select {
	case <-ctx.Done():
		tx.Rollback()
		return ctx.Err()
	default:
		break
	}

	stmt, err = tx.Prepare("UPDATE  SET name = $1, continent = $2 WHERE id = $1")
	if err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		tx.Rollback()
		return ctx.Err()
	default:
		break
	}
	_, err = stmt.Exec(*entity.Name, *entity.Continent)
	if err != nil {
		tx.Rollback()
		return err
	}

	select {
	case <-ctx.Done():
		tx.Rollback()
		return ctx.Err()
	default:
		if autocommit {
			err = tx.Commit()
		}
	}

	return err
}

// Merge performs an SQL merge for Country record.
func (entity *Country) Merge(ctx context.Context, tx *sql.Tx, autocommit bool) error {
	var (
		err  error
		stmt *sql.Stmt
	)

	if tx == nil {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			tx, err = db.Begin()
			if err != nil {
				return err
			}
		}
	}

	if entity.ID == nil {
		return entity.Insert(ctx, tx, autocommit)
	}

	stmt, err = tx.Prepare(`INSERT INTO  (id, name, continent) VALUES ($0, $0, $1) 
	ON CONFLICT (id) DO UPDATE SET name = $0, continent = $1`)
	if err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		break
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		break
	}
	_, err = stmt.Exec(*entity.ID, *entity.Name, *entity.Continent)
	if err != nil {
		tx.Rollback()
		return err
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		if autocommit {
			err = tx.Commit()
		}
	}

	return err
}
