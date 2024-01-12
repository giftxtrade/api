// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

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
	if q.acceptEventInviteStmt, err = db.PrepareContext(ctx, acceptEventInvite); err != nil {
		return nil, fmt.Errorf("error preparing query AcceptEventInvite: %w", err)
	}
	if q.createCategoryStmt, err = db.PrepareContext(ctx, createCategory); err != nil {
		return nil, fmt.Errorf("error preparing query CreateCategory: %w", err)
	}
	if q.createEventStmt, err = db.PrepareContext(ctx, createEvent); err != nil {
		return nil, fmt.Errorf("error preparing query CreateEvent: %w", err)
	}
	if q.createLinkStmt, err = db.PrepareContext(ctx, createLink); err != nil {
		return nil, fmt.Errorf("error preparing query CreateLink: %w", err)
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
	if q.declineEventInviteStmt, err = db.PrepareContext(ctx, declineEventInvite); err != nil {
		return nil, fmt.Errorf("error preparing query DeclineEventInvite: %w", err)
	}
	if q.deleteEventStmt, err = db.PrepareContext(ctx, deleteEvent); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteEvent: %w", err)
	}
	if q.deleteParticipantByIdAndEventIdStmt, err = db.PrepareContext(ctx, deleteParticipantByIdAndEventId); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteParticipantByIdAndEventId: %w", err)
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
	if q.findEventInvitesStmt, err = db.PrepareContext(ctx, findEventInvites); err != nil {
		return nil, fmt.Errorf("error preparing query FindEventInvites: %w", err)
	}
	if q.findEventSimpleStmt, err = db.PrepareContext(ctx, findEventSimple); err != nil {
		return nil, fmt.Errorf("error preparing query FindEventSimple: %w", err)
	}
	if q.findLinkByCodeStmt, err = db.PrepareContext(ctx, findLinkByCode); err != nil {
		return nil, fmt.Errorf("error preparing query FindLinkByCode: %w", err)
	}
	if q.findLinkWithEventByCodeStmt, err = db.PrepareContext(ctx, findLinkWithEventByCode); err != nil {
		return nil, fmt.Errorf("error preparing query FindLinkWithEventByCode: %w", err)
	}
	if q.findParticipantFromEventIdAndUserStmt, err = db.PrepareContext(ctx, findParticipantFromEventIdAndUser); err != nil {
		return nil, fmt.Errorf("error preparing query FindParticipantFromEventIdAndUser: %w", err)
	}
	if q.findParticipantUserWithFullEventByIdStmt, err = db.PrepareContext(ctx, findParticipantUserWithFullEventById); err != nil {
		return nil, fmt.Errorf("error preparing query FindParticipantUserWithFullEventById: %w", err)
	}
	if q.findParticipantWithIdAndEventIdStmt, err = db.PrepareContext(ctx, findParticipantWithIdAndEventId); err != nil {
		return nil, fmt.Errorf("error preparing query FindParticipantWithIdAndEventId: %w", err)
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
	if q.patchParticipantStmt, err = db.PrepareContext(ctx, patchParticipant); err != nil {
		return nil, fmt.Errorf("error preparing query PatchParticipant: %w", err)
	}
	if q.setUserAsAdminStmt, err = db.PrepareContext(ctx, setUserAsAdmin); err != nil {
		return nil, fmt.Errorf("error preparing query SetUserAsAdmin: %w", err)
	}
	if q.updateEventStmt, err = db.PrepareContext(ctx, updateEvent); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateEvent: %w", err)
	}
	if q.updateParticipantStatusStmt, err = db.PrepareContext(ctx, updateParticipantStatus); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateParticipantStatus: %w", err)
	}
	if q.updateProductStmt, err = db.PrepareContext(ctx, updateProduct); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateProduct: %w", err)
	}
	if q.verifyEventForUserAsOrganizerStmt, err = db.PrepareContext(ctx, verifyEventForUserAsOrganizer); err != nil {
		return nil, fmt.Errorf("error preparing query VerifyEventForUserAsOrganizer: %w", err)
	}
	if q.verifyEventForUserAsParticipantStmt, err = db.PrepareContext(ctx, verifyEventForUserAsParticipant); err != nil {
		return nil, fmt.Errorf("error preparing query VerifyEventForUserAsParticipant: %w", err)
	}
	if q.verifyEventWithEmailOrUserStmt, err = db.PrepareContext(ctx, verifyEventWithEmailOrUser); err != nil {
		return nil, fmt.Errorf("error preparing query VerifyEventWithEmailOrUser: %w", err)
	}
	if q.verifyEventWithParticipantIdStmt, err = db.PrepareContext(ctx, verifyEventWithParticipantId); err != nil {
		return nil, fmt.Errorf("error preparing query VerifyEventWithParticipantId: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.acceptEventInviteStmt != nil {
		if cerr := q.acceptEventInviteStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing acceptEventInviteStmt: %w", cerr)
		}
	}
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
	if q.createLinkStmt != nil {
		if cerr := q.createLinkStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createLinkStmt: %w", cerr)
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
	if q.declineEventInviteStmt != nil {
		if cerr := q.declineEventInviteStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing declineEventInviteStmt: %w", cerr)
		}
	}
	if q.deleteEventStmt != nil {
		if cerr := q.deleteEventStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteEventStmt: %w", cerr)
		}
	}
	if q.deleteParticipantByIdAndEventIdStmt != nil {
		if cerr := q.deleteParticipantByIdAndEventIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteParticipantByIdAndEventIdStmt: %w", cerr)
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
	if q.findEventInvitesStmt != nil {
		if cerr := q.findEventInvitesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findEventInvitesStmt: %w", cerr)
		}
	}
	if q.findEventSimpleStmt != nil {
		if cerr := q.findEventSimpleStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findEventSimpleStmt: %w", cerr)
		}
	}
	if q.findLinkByCodeStmt != nil {
		if cerr := q.findLinkByCodeStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findLinkByCodeStmt: %w", cerr)
		}
	}
	if q.findLinkWithEventByCodeStmt != nil {
		if cerr := q.findLinkWithEventByCodeStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findLinkWithEventByCodeStmt: %w", cerr)
		}
	}
	if q.findParticipantFromEventIdAndUserStmt != nil {
		if cerr := q.findParticipantFromEventIdAndUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findParticipantFromEventIdAndUserStmt: %w", cerr)
		}
	}
	if q.findParticipantUserWithFullEventByIdStmt != nil {
		if cerr := q.findParticipantUserWithFullEventByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findParticipantUserWithFullEventByIdStmt: %w", cerr)
		}
	}
	if q.findParticipantWithIdAndEventIdStmt != nil {
		if cerr := q.findParticipantWithIdAndEventIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findParticipantWithIdAndEventIdStmt: %w", cerr)
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
	if q.patchParticipantStmt != nil {
		if cerr := q.patchParticipantStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing patchParticipantStmt: %w", cerr)
		}
	}
	if q.setUserAsAdminStmt != nil {
		if cerr := q.setUserAsAdminStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing setUserAsAdminStmt: %w", cerr)
		}
	}
	if q.updateEventStmt != nil {
		if cerr := q.updateEventStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateEventStmt: %w", cerr)
		}
	}
	if q.updateParticipantStatusStmt != nil {
		if cerr := q.updateParticipantStatusStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateParticipantStatusStmt: %w", cerr)
		}
	}
	if q.updateProductStmt != nil {
		if cerr := q.updateProductStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateProductStmt: %w", cerr)
		}
	}
	if q.verifyEventForUserAsOrganizerStmt != nil {
		if cerr := q.verifyEventForUserAsOrganizerStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing verifyEventForUserAsOrganizerStmt: %w", cerr)
		}
	}
	if q.verifyEventForUserAsParticipantStmt != nil {
		if cerr := q.verifyEventForUserAsParticipantStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing verifyEventForUserAsParticipantStmt: %w", cerr)
		}
	}
	if q.verifyEventWithEmailOrUserStmt != nil {
		if cerr := q.verifyEventWithEmailOrUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing verifyEventWithEmailOrUserStmt: %w", cerr)
		}
	}
	if q.verifyEventWithParticipantIdStmt != nil {
		if cerr := q.verifyEventWithParticipantIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing verifyEventWithParticipantIdStmt: %w", cerr)
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
	db                                       DBTX
	tx                                       *sql.Tx
	acceptEventInviteStmt                    *sql.Stmt
	createCategoryStmt                       *sql.Stmt
	createEventStmt                          *sql.Stmt
	createLinkStmt                           *sql.Stmt
	createParticipantStmt                    *sql.Stmt
	createProductStmt                        *sql.Stmt
	createUserStmt                           *sql.Stmt
	declineEventInviteStmt                   *sql.Stmt
	deleteEventStmt                          *sql.Stmt
	deleteParticipantByIdAndEventIdStmt      *sql.Stmt
	filterProductsStmt                       *sql.Stmt
	findAllEventsWithUserStmt                *sql.Stmt
	findCategoryByNameStmt                   *sql.Stmt
	findEventByIdStmt                        *sql.Stmt
	findEventInvitesStmt                     *sql.Stmt
	findEventSimpleStmt                      *sql.Stmt
	findLinkByCodeStmt                       *sql.Stmt
	findLinkWithEventByCodeStmt              *sql.Stmt
	findParticipantFromEventIdAndUserStmt    *sql.Stmt
	findParticipantUserWithFullEventByIdStmt *sql.Stmt
	findParticipantWithIdAndEventIdStmt      *sql.Stmt
	findProductByIdStmt                      *sql.Stmt
	findProductByProductKeyStmt              *sql.Stmt
	findUserByEmailStmt                      *sql.Stmt
	findUserByIdStmt                         *sql.Stmt
	findUserByIdAndEmailStmt                 *sql.Stmt
	findUserByIdOrEmailStmt                  *sql.Stmt
	patchParticipantStmt                     *sql.Stmt
	setUserAsAdminStmt                       *sql.Stmt
	updateEventStmt                          *sql.Stmt
	updateParticipantStatusStmt              *sql.Stmt
	updateProductStmt                        *sql.Stmt
	verifyEventForUserAsOrganizerStmt        *sql.Stmt
	verifyEventForUserAsParticipantStmt      *sql.Stmt
	verifyEventWithEmailOrUserStmt           *sql.Stmt
	verifyEventWithParticipantIdStmt         *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                                       tx,
		tx:                                       tx,
		acceptEventInviteStmt:                    q.acceptEventInviteStmt,
		createCategoryStmt:                       q.createCategoryStmt,
		createEventStmt:                          q.createEventStmt,
		createLinkStmt:                           q.createLinkStmt,
		createParticipantStmt:                    q.createParticipantStmt,
		createProductStmt:                        q.createProductStmt,
		createUserStmt:                           q.createUserStmt,
		declineEventInviteStmt:                   q.declineEventInviteStmt,
		deleteEventStmt:                          q.deleteEventStmt,
		deleteParticipantByIdAndEventIdStmt:      q.deleteParticipantByIdAndEventIdStmt,
		filterProductsStmt:                       q.filterProductsStmt,
		findAllEventsWithUserStmt:                q.findAllEventsWithUserStmt,
		findCategoryByNameStmt:                   q.findCategoryByNameStmt,
		findEventByIdStmt:                        q.findEventByIdStmt,
		findEventInvitesStmt:                     q.findEventInvitesStmt,
		findEventSimpleStmt:                      q.findEventSimpleStmt,
		findLinkByCodeStmt:                       q.findLinkByCodeStmt,
		findLinkWithEventByCodeStmt:              q.findLinkWithEventByCodeStmt,
		findParticipantFromEventIdAndUserStmt:    q.findParticipantFromEventIdAndUserStmt,
		findParticipantUserWithFullEventByIdStmt: q.findParticipantUserWithFullEventByIdStmt,
		findParticipantWithIdAndEventIdStmt:      q.findParticipantWithIdAndEventIdStmt,
		findProductByIdStmt:                      q.findProductByIdStmt,
		findProductByProductKeyStmt:              q.findProductByProductKeyStmt,
		findUserByEmailStmt:                      q.findUserByEmailStmt,
		findUserByIdStmt:                         q.findUserByIdStmt,
		findUserByIdAndEmailStmt:                 q.findUserByIdAndEmailStmt,
		findUserByIdOrEmailStmt:                  q.findUserByIdOrEmailStmt,
		patchParticipantStmt:                     q.patchParticipantStmt,
		setUserAsAdminStmt:                       q.setUserAsAdminStmt,
		updateEventStmt:                          q.updateEventStmt,
		updateParticipantStatusStmt:              q.updateParticipantStatusStmt,
		updateProductStmt:                        q.updateProductStmt,
		verifyEventForUserAsOrganizerStmt:        q.verifyEventForUserAsOrganizerStmt,
		verifyEventForUserAsParticipantStmt:      q.verifyEventForUserAsParticipantStmt,
		verifyEventWithEmailOrUserStmt:           q.verifyEventWithEmailOrUserStmt,
		verifyEventWithParticipantIdStmt:         q.verifyEventWithParticipantIdStmt,
	}
}
