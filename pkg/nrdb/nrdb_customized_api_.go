// Code *NOT* generated by tutone
package nrdb

import "context"

// PerformNRQLQuery facilitates making an NRQL query using NerdGraph.
func (n *Nrdb) PerformNRQLQuery(accountID int, query NRQL) (*NRDBResultContainerMultiResultCustomized, error) {
	return n.PerformNRQLQueryWithContext(context.Background(), accountID, query)
}

// PerformNRQLQueryWithContext facilitates making a NRQL query.
func (n *Nrdb) PerformNRQLQueryWithContext(ctx context.Context, accountID int, query NRQL) (*NRDBResultContainerMultiResultCustomized, error) {
	respBody := gqlNRQLQueryResponseCustomized{}

	vars := map[string]interface{}{
		"accountId": accountID,
		"query":     query,
	}

	if err := n.client.NerdGraphQueryWithContext(ctx, gqlNrqlQuery, vars, &respBody); err != nil {
		return nil, err
	}

	return &respBody.Actor.Account.NRQL, nil
}

type gqlNRQLQueryResponseCustomized struct {
	Actor struct {
		Account struct {
			NRQL NRDBResultContainerMultiResultCustomized
		}
	}
}
