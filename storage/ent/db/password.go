// Code generated by entc, DO NOT EDIT.

package db

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/dexidp/dex/v2/storage/ent/db/password"
)

// Password is the model entity for the Password schema.
type Password struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Hash holds the value of the "hash" field.
	Hash []byte `json:"hash,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID string `json:"user_id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Password) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case password.FieldHash:
			values[i] = new([]byte)
		case password.FieldID:
			values[i] = new(sql.NullInt64)
		case password.FieldEmail, password.FieldUsername, password.FieldUserID:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Password", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Password fields.
func (pa *Password) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case password.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pa.ID = int(value.Int64)
		case password.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				pa.Email = value.String
			}
		case password.FieldHash:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field hash", values[i])
			} else if value != nil {
				pa.Hash = *value
			}
		case password.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				pa.Username = value.String
			}
		case password.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				pa.UserID = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Password.
// Note that you need to call Password.Unwrap() before calling this method if this Password
// was returned from a transaction, and the transaction was committed or rolled back.
func (pa *Password) Update() *PasswordUpdateOne {
	return (&PasswordClient{config: pa.config}).UpdateOne(pa)
}

// Unwrap unwraps the Password entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pa *Password) Unwrap() *Password {
	tx, ok := pa.config.driver.(*txDriver)
	if !ok {
		panic("db: Password is not a transactional entity")
	}
	pa.config.driver = tx.drv
	return pa
}

// String implements the fmt.Stringer.
func (pa *Password) String() string {
	var builder strings.Builder
	builder.WriteString("Password(")
	builder.WriteString(fmt.Sprintf("id=%v", pa.ID))
	builder.WriteString(", email=")
	builder.WriteString(pa.Email)
	builder.WriteString(", hash=")
	builder.WriteString(fmt.Sprintf("%v", pa.Hash))
	builder.WriteString(", username=")
	builder.WriteString(pa.Username)
	builder.WriteString(", user_id=")
	builder.WriteString(pa.UserID)
	builder.WriteByte(')')
	return builder.String()
}

// Passwords is a parsable slice of Password.
type Passwords []*Password

func (pa Passwords) config(cfg config) {
	for _i := range pa {
		pa[_i].config = cfg
	}
}
