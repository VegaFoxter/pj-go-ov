package database

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/upper/db/v4"
)

const UsersTableName = "users"

type user struct {
	Id          uint64      `db:"id,omitempty"`
	FirstName   string      `db:"first_name"`
	SecondName  string      `db:"second_name"`
	Password    string      `db:"password"`
	Email       string      `db:"email"`
	Role        domain.Role `db:"role"`
	CreatedDate time.Time   `db:"created_date,omitempty"`
	UpdatedDate time.Time   `db:"updated_date,omitempty"`
	DeletedDate *time.Time  `db:"deleted_date,omitempty"`
}

type UserRepository interface {
	FindByEmail(phone string) (domain.User, error)
	FindById(id uint64) (domain.User, error)
	Find(id uint64) (interface{}, error)
	Save(user domain.User) (domain.User, error)
	Update(user domain.User) (domain.User, error)
	Delete(id uint64) error
}

type userRepository struct {
	coll db.Collection
	sess db.Session
}

func NewUserRepository(dbSession db.Session) UserRepository {
	return userRepository{
		coll: dbSession.Collection(UsersTableName),
		sess: dbSession,
	}
}

func (r userRepository) FindByEmail(email string) (domain.User, error) {
	var u user
	err := r.coll.Find(db.Cond{"email": email, "deleted_date": nil}).One(&u)
	if err != nil {
		return domain.User{}, err
	}

	return r.mapModelToDomain(u), nil
}

func (r userRepository) FindById(id uint64) (domain.User, error) {
	var usr user
	err := r.coll.Find(db.Cond{"id": id}).One(&usr)
	if err != nil {
		return domain.User{}, err
	}

	return r.mapModelToDomain(usr), nil
}

func (r userRepository) Find(id uint64) (interface{}, error) {
	var usr user
	err := r.coll.Find(db.Cond{"id": id}).One(&usr)
	if err != nil {
		return domain.User{}, err
	}

	return r.mapModelToDomain(usr), nil
}

func (r userRepository) Save(user domain.User) (domain.User, error) {
	u := r.mapDomainToModel(user)
	u.CreatedDate, u.UpdatedDate = time.Now(), time.Now()
	err := r.coll.InsertReturning(&u)
	if err != nil {
		return domain.User{}, err
	}
	return r.mapModelToDomain(u), nil
}

func (r userRepository) Update(user domain.User) (domain.User, error) {
	u := r.mapDomainToModel(user)
	u.UpdatedDate = time.Now()
	err := r.coll.Find(db.Cond{"id": u.Id, "deleted_date": nil}).Update(&u)
	if err != nil {
		return domain.User{}, err
	}
	return r.mapModelToDomain(u), nil
}

func (r userRepository) Delete(id uint64) error {
	return r.coll.Find(db.Cond{"id": id, "deleted_date": nil}).Update(map[string]interface{}{"deleted_date": time.Now()})
}

func (r userRepository) mapDomainToModel(d domain.User) user {
	return user{
		Id:          d.Id,
		Email:       d.Email,
		Password:    d.Password,
		FirstName:   d.FirstName,
		SecondName:  d.SecondName,
		Role:        d.Role,
		CreatedDate: d.CreatedDate,
		UpdatedDate: d.UpdatedDate,
		DeletedDate: d.DeletedDate,
	}
}

func (r userRepository) mapModelToDomain(m user) domain.User {
	return domain.User{
		Id:          m.Id,
		Email:       m.Email,
		Password:    m.Password,
		FirstName:   m.FirstName,
		SecondName:  m.SecondName,
		Role:        m.Role,
		CreatedDate: m.CreatedDate,
		UpdatedDate: m.UpdatedDate,
		DeletedDate: m.DeletedDate,
	}
}
