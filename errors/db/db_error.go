package dbError

import "github.com/lib/pq"

func IsErrorCode(err error, errCode pq.ErrorCode) bool {
  if pgErr, ok := err.(*pq.Error); ok {
    return pgErr.Code == errCode
  }
  return false
}
