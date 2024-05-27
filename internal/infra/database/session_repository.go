package database

import (
	"fmt"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/google/uuid"
	"github.com/upper/db/v4"
)

const SessionsTableName = "sessions"

type sessions struct {
	UserId uint64    `db:"user_id"`
	UUID   uuid.UUID `db:"uuid"`
}

type SessionRepository interface {
	Save(sess domain.Session) error
	Exists(sess domain.Session) error
	Delete(sess domain.Session) error
}

type sessionRepository struct {
	coll db.Collection
}

func NewSessRepository(dbSession db.Session) SessionRepository {
	return sessionRepository{
		coll: dbSession.Collection(SessionsTableName),
	}
}

func (r sessionRepository) Save(sess domain.Session) error {
	a := r.mapDomainToModel(sess)
	err := r.coll.InsertReturning(&a)
	if err != nil {
		return err
	}

	return nil
}

func (r sessionRepository) Exists(sess domain.Session) error {
	exists, err := r.coll.Find(db.Cond{"user_id": sess.UserId, "uuid": sess.UUID}).Exists()
	if !exists {
		err = fmt.Errorf("sess not found")
	}
	return err
}

func (r sessionRepository) Delete(sess domain.Session) error {
	return r.coll.Find(db.Cond{"user_id": sess.UserId, "uuid": sess.UUID}).Delete()
}

func (r sessionRepository) mapDomainToModel(d domain.Session) sessions {
	return sessions{
		UserId: d.UserId,
		UUID:   d.UUID,
	}
}
