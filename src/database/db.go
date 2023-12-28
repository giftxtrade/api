// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package database

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createCategoryStmt, err = db.PrepareContext(ctx, createCategory); err != nil {
		return nil, fmt.Errorf("error preparing query CreateCategory: %w", err)
	}
	if q.createEventStmt, err = db.PrepareContext(ctx, createEvent); err != nil {
		return nil, fmt.Errorf("error preparing query CreateEvent: %w", err)
	}
	if q.createParticipantStmt, err = db.PrepareContext(ctx, createParticipant); err != nil {
		return nil, fmt.Errorf("error preparing query CreateParticipant: %w", err)
	}
	if q.createProductStmt, err = db.PrepareContext(ctx, createProduct); err != nil {
		return nil, fmt.Errorf("error preparing query CreateProduct: %w", err)
	}
	if q.createUserStmt, err = db.PrepareContext(ctx, createUser); err != nil {
		return nil, fmt.Errorf("error preparing query CreateUser: %w", err)
	}
	if q.filterProductsStmt, err = db.PrepareContext(ctx, filterProducts); err != nil {
		return nil, fmt.Errorf("error preparing query FilterProducts: %w", err)
	}
	if q.findAllEventsWithUserStmt, err = db.PrepareContext(ctx, findAllEventsWithUser); err != nil {
		return nil, fmt.Errorf("error preparing query FindAllEventsWithUser: %w", err)
	}
	if q.findCategoryByNameStmt, err = db.PrepareContext(ctx, findCategoryByName); err != nil {
		return nil, fmt.Errorf("error preparing query FindCategoryByName: %w", err)
	}
	if q.findEventByIdStmt, err = db.PrepareContext(ctx, findEventById); err != nil {
		return nil, fmt.Errorf("error preparing query FindEventById: %w", err)
	}
	if q.findEventForUserStmt, err = db.PrepareContext(ctx, findEventForUser); err != nil {
		return nil, fmt.Errorf("error preparing query FindEventForUser: %w", err)
	}
	if q.findEventForUserAsOrganizerStmt, err = db.PrepareContext(ctx, findEventForUserAsOrganizer); err != nil {
		return nil, fmt.Errorf("error preparing query FindEventForUserAsOrganizer: %w", err)
	}
	if q.findEventForUserAsParticipantStmt, err = db.PrepareContext(ctx, findEventForUserAsParticipant); err != nil {
		return nil, fmt.Errorf("error preparing query FindEventForUserAsParticipant: %w", err)
	}
	if q.findEventInvitesStmt, err = db.PrepareContext(ctx, findEventInvites); err != nil {
		return nil, fmt.Errorf("error preparing query FindEventInvites: %w", err)
	}
	if q.findProductByIdStmt, err = db.PrepareContext(ctx, findProductById); err != nil {
		return nil, fmt.Errorf("error preparing query FindProductById: %w", err)
	}
	if q.findProductByProductKeyStmt, err = db.PrepareContext(ctx, findProductByProductKey); err != nil {
		return nil, fmt.Errorf("error preparing query FindProductByProductKey: %w", err)
	}
	if q.findUserByEmailStmt, err = db.PrepareContext(ctx, findUserByEmail); err != nil {
		return nil, fmt.Errorf("error preparing query FindUserByEmail: %w", err)
	}
	if q.findUserByIdStmt, err = db.PrepareContext(ctx, findUserById); err != nil {
		return nil, fmt.Errorf("error preparing query FindUserById: %w", err)
	}
	if q.findUserByIdAndEmailStmt, err = db.PrepareContext(ctx, findUserByIdAndEmail); err != nil {
		return nil, fmt.Errorf("error preparing query FindUserByIdAndEmail: %w", err)
	}
	if q.findUserByIdOrEmailStmt, err = db.PrepareContext(ctx, findUserByIdOrEmail); err != nil {
		return nil, fmt.Errorf("error preparing query FindUserByIdOrEmail: %w", err)
	}
	if q.setUserAsAdminStmt, err = db.PrepareContext(ctx, setUserAsAdmin); err != nil {
		return nil, fmt.Errorf("error preparing query SetUserAsAdmin: %w", err)
	}
	if q.updateProductStmt, err = db.PrepareContext(ctx, updateProduct); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateProduct: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createCategoryStmt != nil {
		if cerr := q.createCategoryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createCategoryStmt: %w", cerr)
		}
	}
	if q.createEventStmt != nil {
		if cerr := q.createEventStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createEventStmt: %w", cerr)
		}
	}
	if q.createParticipantStmt != nil {
		if cerr := q.createParticipantStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createParticipantStmt: %w", cerr)
		}
	}
	if q.createProductStmt != nil {
		if cerr := q.createProductStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createProductStmt: %w", cerr)
		}
	}
	if q.createUserStmt != nil {
		if cerr := q.createUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createUserStmt: %w", cerr)
		}
	}
	if q.filterProductsStmt != nil {
		if cerr := q.filterProductsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing filterProductsStmt: %w", cerr)
		}
	}
	if q.findAllEventsWithUserStmt != nil {
		if cerr := q.findAllEventsWithUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findAllEventsWithUserStmt: %w", cerr)
		}
	}
	if q.findCategoryByNameStmt != nil {
		if cerr := q.findCategoryByNameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findCategoryByNameStmt: %w", cerr)
		}
	}
	if q.findEventByIdStmt != nil {
		if cerr := q.findEventByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findEventByIdStmt: %w", cerr)
		}
	}
	if q.findEventForUserStmt != nil {
		if cerr := q.findEventForUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findEventForUserStmt: %w", cerr)
		}
	}
	if q.findEventForUserAsOrganizerStmt != nil {
		if cerr := q.findEventForUserAsOrganizerStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findEventForUserAsOrganizerStmt: %w", cerr)
		}
	}
	if q.findEventForUserAsParticipantStmt != nil {
		if cerr := q.findEventForUserAsParticipantStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findEventForUserAsParticipantStmt: %w", cerr)
		}
	}
	if q.findEventInvitesStmt != nil {
		if cerr := q.findEventInvitesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findEventInvitesStmt: %w", cerr)
		}
	}
	if q.findProductByIdStmt != nil {
		if cerr := q.findProductByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findProductByIdStmt: %w", cerr)
		}
	}
	if q.findProductByProductKeyStmt != nil {
		if cerr := q.findProductByProductKeyStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findProductByProductKeyStmt: %w", cerr)
		}
	}
	if q.findUserByEmailStmt != nil {
		if cerr := q.findUserByEmailStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findUserByEmailStmt: %w", cerr)
		}
	}
	if q.findUserByIdStmt != nil {
		if cerr := q.findUserByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findUserByIdStmt: %w", cerr)
		}
	}
	if q.findUserByIdAndEmailStmt != nil {
		if cerr := q.findUserByIdAndEmailStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findUserByIdAndEmailStmt: %w", cerr)
		}
	}
	if q.findUserByIdOrEmailStmt != nil {
		if cerr := q.findUserByIdOrEmailStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findUserByIdOrEmailStmt: %w", cerr)
		}
	}
	if q.setUserAsAdminStmt != nil {
		if cerr := q.setUserAsAdminStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing setUserAsAdminStmt: %w", cerr)
		}
	}
	if q.updateProductStmt != nil {
		if cerr := q.updateProductStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateProductStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                                DBTX
	tx                                *sql.Tx
	createCategoryStmt                *sql.Stmt
	createEventStmt                   *sql.Stmt
	createParticipantStmt             *sql.Stmt
	createProductStmt                 *sql.Stmt
	createUserStmt                    *sql.Stmt
	filterProductsStmt                *sql.Stmt
	findAllEventsWithUserStmt         *sql.Stmt
	findCategoryByNameStmt            *sql.Stmt
	findEventByIdStmt                 *sql.Stmt
	findEventForUserStmt              *sql.Stmt
	findEventForUserAsOrganizerStmt   *sql.Stmt
	findEventForUserAsParticipantStmt *sql.Stmt
	findEventInvitesStmt              *sql.Stmt
	findProductByIdStmt               *sql.Stmt
	findProductByProductKeyStmt       *sql.Stmt
	findUserByEmailStmt               *sql.Stmt
	findUserByIdStmt                  *sql.Stmt
	findUserByIdAndEmailStmt          *sql.Stmt
	findUserByIdOrEmailStmt           *sql.Stmt
	setUserAsAdminStmt                *sql.Stmt
	updateProductStmt                 *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                                tx,
		tx:                                tx,
		createCategoryStmt:                q.createCategoryStmt,
		createEventStmt:                   q.createEventStmt,
		createParticipantStmt:             q.createParticipantStmt,
		createProductStmt:                 q.createProductStmt,
		createUserStmt:                    q.createUserStmt,
		filterProductsStmt:                q.filterProductsStmt,
		findAllEventsWithUserStmt:         q.findAllEventsWithUserStmt,
		findCategoryByNameStmt:            q.findCategoryByNameStmt,
		findEventByIdStmt:                 q.findEventByIdStmt,
		findEventForUserStmt:              q.findEventForUserStmt,
		findEventForUserAsOrganizerStmt:   q.findEventForUserAsOrganizerStmt,
		findEventForUserAsParticipantStmt: q.findEventForUserAsParticipantStmt,
		findEventInvitesStmt:              q.findEventInvitesStmt,
		findProductByIdStmt:               q.findProductByIdStmt,
		findProductByProductKeyStmt:       q.findProductByProductKeyStmt,
		findUserByEmailStmt:               q.findUserByEmailStmt,
		findUserByIdStmt:                  q.findUserByIdStmt,
		findUserByIdAndEmailStmt:          q.findUserByIdAndEmailStmt,
		findUserByIdOrEmailStmt:           q.findUserByIdOrEmailStmt,
		setUserAsAdminStmt:                q.setUserAsAdminStmt,
		updateProductStmt:                 q.updateProductStmt,
	}
}
