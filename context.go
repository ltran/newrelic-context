package nrcontext

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/ltran/newrelic-context/nrgorm"
	"github.com/newrelic/go-agent"
)

type contextKey int

const txnKey contextKey = 0

// Set NewRelic transaction to context
func ContextWithTxn(c context.Context, txn newrelic.Transaction) context.Context {
	return context.WithValue(c, txnKey, txn)
}

// Get NewRelic transaction from context anywhere
func GetTnxFromContext(c context.Context) newrelic.Transaction {
	if tnx := c.Value(txnKey); tnx != nil {
		return tnx.(newrelic.Transaction)
	}
	return nil
}

// Sets transaction from Context to gorm settings, returns cloned DB
func SetTxnToGorm(ctx context.Context, db *gorm.DB) *gorm.DB {
	txn := GetTnxFromContext(ctx)
	return nrgorm.SetTxnToGorm(txn, db)
}
