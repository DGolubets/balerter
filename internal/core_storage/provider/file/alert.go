package file

import (
	"fmt"
	"github.com/balerter/balerter/internal/alert/alert"
	"go.etcd.io/bbolt"
	"go.uber.org/zap"
)

func (s *Storage) GetOrNew(name string) (*alert.Alert, error) {
	var v []byte

	err := s.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(bucketAlert)
		if b == nil {
			return errBucketNotFound
		}
		v = b.Get([]byte(name))

		return nil
	})

	if err != nil {
		s.logger.Error("bbolt: error get item", zap.ByteString("bucket", bucketAlert), zap.String("key", name), zap.Error(err))
		return nil, fmt.Errorf("error get item, %w", err)
	}

	a := alert.AcquireAlert()
	a.SetName(name)

	// if the buffer is empty, returns a new alert
	if len(v) == 0 {
		return a, nil
	}

	err = a.Unmarshal(v, a)
	if err != nil {
		return nil, fmt.Errorf("error unmarshal alert, %w", err)
	}

	return a, nil

}

func (s *Storage) All() []*alert.Alert {
	res := make([]*alert.Alert, 0)
	panic("not implemented")
	return res
}

func (s *Storage) Release(a *alert.Alert) {
	alert.ReleaseAlert(a)
}