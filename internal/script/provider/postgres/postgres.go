package postgres

import (
	"fmt"
	"github.com/balerter/balerter/internal/config/scripts/postgres"
	"github.com/balerter/balerter/internal/script/script"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	defaultTimeout = 3000
)

// Provider represents Postgres script provider
type Provider struct {
	name  string
	db    *sqlx.DB
	query string
}

// New creates new Postgres script provider
func New(cfg postgres.Postgres) (*Provider, error) {
	p := &Provider{
		name:  "postgres." + cfg.Name,
		query: cfg.Query,
	}

	if cfg.Timeout == 0 {
		cfg.Timeout = defaultTimeout
	}

	pgConnString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s&sslrootcert=%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
		cfg.SSLMode,
		cfg.SSLCertPath,
	)
	var err error

	p.db, err = sqlx.Open("postgres", pgConnString)
	if err != nil {
		return nil, err
	}

	if err := p.db.Ping(); err != nil {
		p.db.Close()
		return nil, err
	}

	return p, nil
}

// Get returns scripts from the provider
func (p *Provider) Get() ([]*script.Script, error) {
	rows, err := p.db.Query(p.query)
	if err != nil {
		return nil, fmt.Errorf("error db query, %w", err)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	defer rows.Close()

	var name string
	var body []byte

	ss := make([]*script.Script, 0)

	for rows.Next() {
		err = rows.Scan(&name, &body)
		if err != nil {
			return nil, err
		}

		s := script.New()
		s.Name = p.name + "." + name
		s.Body = body

		if err := s.ParseMeta(); err != nil {
			return nil, err
		}

		ss = append(ss, s)
	}

	return ss, nil
}

// Stop the provider
func (p *Provider) Stop() error {
	return p.db.Close()
}
