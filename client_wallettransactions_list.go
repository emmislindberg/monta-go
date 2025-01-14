package monta

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// ListWalletTransactionsRequest is the request input to the [Client.ListWalletTransactions] method.
type ListWalletTransactionsRequest struct {
	// Page number to request (starts with 1).
	Page int
	// Number of items PerPage (between 1 and 100, default 10).
	PerPage int
	// FromDate allows to filter to retrieve transactions where [WalletTransaction.CreatedAt] >= FromDate.
	FromDate time.Time
	// ToDate allows to filter to retrieve transactions where [WalletTransaction.CreatedAt] <= ToDate.
	ToDate time.Time
}

// ListWalletTransactionsResponse is the response output from the [Client.ListWalletTransactions] method.
type ListWalletTransactionsResponse struct {
	// WalletTransactions in the current page.
	WalletTransactions []*WalletTransaction `json:"data"`
	// PageMeta with metadata about the current page.
	PageMeta PageMeta `json:"meta"`
}

// ListWalletTransactions to retrieve your wallet transactions.
func (c *Client) ListWalletTransactions(
	ctx context.Context,
	request *ListWalletTransactionsRequest,
) (_ *ListWalletTransactionsResponse, err error) {
	const method, path = http.MethodGet, "/v1/wallet-transactions"
	defer func() {
		if err != nil {
			err = fmt.Errorf("%s %s: %w", method, path, err)
		}
	}()
	requestURL, err := url.Parse(apiHost + path)
	if err != nil {
		return nil, err
	}
	query := requestURL.Query()
	if request.Page > 0 {
		query.Set("page", strconv.Itoa(request.Page))
	}
	if request.PerPage > 0 {
		query.Set("perPage", strconv.Itoa(request.PerPage))
	}
	if !request.FromDate.IsZero() {
		query.Set("fromDate", request.FromDate.UTC().Format(time.RFC3339))
	}
	if !request.ToDate.IsZero() {
		query.Set("toDate", request.ToDate.UTC().Format(time.RFC3339))
	}
	requestURL.RawQuery = query.Encode()
	httpRequest, err := http.NewRequestWithContext(ctx, method, requestURL.String(), nil)
	if err != nil {
		return nil, err
	}
	if err := c.setAuthorization(ctx, httpRequest); err != nil {
		return nil, err
	}
	httpResponse, err := c.httpClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = httpResponse.Body.Close()
	}()
	if httpResponse.StatusCode != http.StatusOK {
		return nil, newStatusError(httpResponse)
	}
	var response ListWalletTransactionsResponse
	if err := json.NewDecoder(httpResponse.Body).Decode(&response); err != nil {
		return nil, err
	}
	return &response, nil
}
