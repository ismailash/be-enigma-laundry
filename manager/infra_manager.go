package manager

import (
	"database/sql"
	"fmt"

	"github.com/ismailash/be-enigma-laundry/config"
	_ "github.com/lib/pq"
)

type InfraManager interface {
	Conn() *sql.DB
}

type infraManager struct {
	db  *sql.DB
	cfg *config.Config
}

func NewInfraManager(cfg *config.Config) (InfraManager, error) {
	conn := &infraManager{cfg: cfg}
	if err := conn.openConn(); err != nil {
		return nil, err
	}

	return conn, nil
}

func (m *infraManager) openConn() error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", m.cfg.Host, m.cfg.Port, m.cfg.User, m.cfg.Password, m.cfg.Name)

	db, err := sql.Open(m.cfg.Driver, dsn)
	if err != nil {
		return fmt.Errorf("failed to open connection %v", err.Error())
	}

	m.db = db
	return nil
}

func (m *infraManager) Conn() *sql.DB {
	return m.db
}
