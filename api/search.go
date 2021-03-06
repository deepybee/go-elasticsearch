// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package api

import (
	"net/http"
	"time"

	"github.com/elastic/go-elasticsearch/transport"
	"github.com/elastic/go-elasticsearch/util"
)

// SearchOption is a non-required Search option that gets applied to an HTTP request.
type SearchOption func(r *transport.Request)

// WithSearchIndex - a comma-separated list of index names to search; use "_all" or empty string to perform the operation on all indices.
func WithSearchIndex(index []string) SearchOption {
	return func(r *transport.Request) {
	}
}

// WithSearchType - a comma-separated list of document types to search; leave empty to perform the operation on all types.
func WithSearchType(documentType []string) SearchOption {
	return func(r *transport.Request) {
	}
}

// WithSearchSource - true or false to return the _source field or not, or a list of fields to return.
func WithSearchSource(source []string) SearchOption {
	return func(r *transport.Request) {
	}
}

// WithSearchSourceExclude - a list of fields to exclude from the returned _source field.
func WithSearchSourceExclude(sourceExclude []string) SearchOption {
	return func(r *transport.Request) {
	}
}

// WithSearchSourceInclude - a list of fields to extract and return from the _source field.
func WithSearchSourceInclude(sourceInclude []string) SearchOption {
	return func(r *transport.Request) {
	}
}

// WithSearchAllowNoIndices - whether to ignore if a wildcard indices expression resolves into no concrete indices. (This includes "_all" string or when no indices have been specified).
func WithSearchAllowNoIndices(allowNoIndices bool) SearchOption {
	return func(r *transport.Request) {
	}
}

// WithSearchAnalyzeWildcard - specify whether wildcard and prefix queries should be analyzed (default: false).
func WithSearchAnalyzeWildcard(analyzeWildcard bool) SearchOption {
	return func(r *transport.Request) {
	}
}

// WithSearchAnalyzer - the analyzer to use for the query string.
func WithSearchAnalyzer(analyzer string) SearchOption {
	return func(r *transport.Request) {
	}
}

// WithSearchBatchedReduceSize - the number of shard results that should be reduced at once on the coordinating node. This value should be used as a protection mechanism to reduce the memory overhead per search request if the potential number of shards in the request can be large.
func WithSearchBatchedReduceSize(batchedReduceSize int) SearchOption {
	return func(r *transport.Request) {
	}
}

// SearchDefaultOperator - the default operator for query string query (AND or OR).
type SearchDefaultOperator int

const (
	// SearchDefaultOperatorAND can be used to set SearchDefaultOperator to "AND"
	SearchDefaultOperatorAND = iota
	// SearchDefaultOperatorOR can be used to set SearchDefaultOperator to "OR"
	SearchDefaultOperatorOR = iota
)

// WithSearchDefaultOperator - the default operator for query string query (AND or OR).
func WithSearchDefaultOperator(defaultOperator SearchDefaultOperator) SearchOption {
	return func(r *transport.Request) {
	}
}

// WithSearchDf - the field to use as default where no field prefix is given in the query string.
func WithSearchDf(df string) SearchOption {
	return func(r *transport.Request) {
	}
}

// WithSearchDocvalueFields - a comma-separated list of fields to return as the docvalue representation of a field for each hit.
func WithSearchDocvalueFields(docvalueFields []string) SearchOption {
	return func(r *transport.Request) {
	}
}

// SearchExpandWildcards - whether to expand wildcard expression to concrete indices that are open, closed or both.
type SearchExpandWildcards int

const (
	// SearchExpandWildcardsOpen can be used to set SearchExpandWildcards to "open"
	SearchExpandWildcardsOpen = iota
	// SearchExpandWildcardsClosed can be used to set SearchExpandWildcards to "closed"
	SearchExpandWildcardsClosed = iota
	// SearchExpandWildcardsNone can be used to set SearchExpandWildcards to "none"
	SearchExpandWildcardsNone = iota
	// SearchExpandWildcardsAll can be used to set SearchExpandWildcards to "all"
	SearchExpandWildcardsAll = iota
)

// WithSearchExpandWildcards - whether to expand wildcard expression to concrete indices that are open, closed or both.
func WithSearchExpandWildcards(expandWildcards SearchExpandWildcards) SearchOption {
	return func(r *transport.Request) {
	}
}

// WithSearchExplain - specify whether to return detailed information about score computation as part of a hit.
func WithSearchExplain(explain bool) SearchOption {
	return func(r *transport.Request) {
	}
}

// WithSearchFielddataFields - a comma-separated list of fields to return as the docvalue representation of a field for each hit.
func WithSearchFielddataFields(fielddataFields []string) SearchOption {
	return func(r *transport.Request) {
	}
}

// WithSearchFrom - starting offset (default: 0).
func WithSearchFrom(from int) SearchOption {
	return func(r *transport.Request) {
	}
}

// WithSearchIgnoreUnavailable - whether specified concrete indices should be ignored when unavailable (missing or closed).
func WithSearchIgnoreUnavailable(ignoreUnavailable bool) SearchOption {
	return func(r *transport.Request) {
	}
}

// WithSearchLenient - specify whether format-based query failures (such as providing text to a numeric field) should be ignored.
func WithSearchLenient(lenient bool) SearchOption {
	return func(r *transport.Request) {
	}
}

// WithSearchPreference - specify the node or shard the operation should be performed on (default: random).
func WithSearchPreference(preference string) SearchOption {
	return func(r *transport.Request) {
	}
}

// WithSearchQ - query in the Lucene query string syntax.
func WithSearchQ(q string) SearchOption {
	return func(r *transport.Request) {
	}
}

// WithSearchRequestCache - specify if request cache should be used for this request or not, defaults to index level setting.
func WithSearchRequestCache(requestCache bool) SearchOption {
	return func(r *transport.Request) {
	}
}

// WithSearchRouting - a comma-separated list of specific routing values.
func WithSearchRouting(routing []string) SearchOption {
	return func(r *transport.Request) {
	}
}

// WithSearchScroll - specify how long a consistent view of the index should be maintained for scrolled search.
func WithSearchScroll(scroll time.Duration) SearchOption {
	return func(r *transport.Request) {
	}
}

// SearchSearchType - search operation type.
type SearchSearchType int

const (
	// SearchSearchTypeQueryThenFetch can be used to set SearchSearchType to "query_then_fetch"
	SearchSearchTypeQueryThenFetch = iota
	// SearchSearchTypeDfsQueryThenFetch can be used to set SearchSearchType to "dfs_query_then_fetch"
	SearchSearchTypeDfsQueryThenFetch = iota
)

// WithSearchSearchType - search operation type.
func WithSearchSearchType(searchType SearchSearchType) SearchOption {
	return func(r *transport.Request) {
	}
}

// WithSearchSize - number of hits to return (default: 10).
func WithSearchSize(size int) SearchOption {
	return func(r *transport.Request) {
	}
}

// WithSearchSort - a comma-separated list of <field>:<direction> pairs.
func WithSearchSort(sort []string) SearchOption {
	return func(r *transport.Request) {
	}
}

// WithSearchStats - specific 'tag' of the request for logging and statistical purposes.
func WithSearchStats(stats []string) SearchOption {
	return func(r *transport.Request) {
	}
}

// WithSearchStoredFields - a comma-separated list of stored fields to return as part of a hit.
func WithSearchStoredFields(storedFields []string) SearchOption {
	return func(r *transport.Request) {
	}
}

// WithSearchSuggestField - specify which field to use for suggestions.
func WithSearchSuggestField(suggestField string) SearchOption {
	return func(r *transport.Request) {
	}
}

// SearchSuggestMode - specify suggest mode.
type SearchSuggestMode int

const (
	// SearchSuggestModeMissing can be used to set SearchSuggestMode to "missing"
	SearchSuggestModeMissing = iota
	// SearchSuggestModePopular can be used to set SearchSuggestMode to "popular"
	SearchSuggestModePopular = iota
	// SearchSuggestModeAlways can be used to set SearchSuggestMode to "always"
	SearchSuggestModeAlways = iota
)

// WithSearchSuggestMode - specify suggest mode.
func WithSearchSuggestMode(suggestMode SearchSuggestMode) SearchOption {
	return func(r *transport.Request) {
	}
}

// WithSearchSuggestSize - how many suggestions to return in response.
func WithSearchSuggestSize(suggestSize int) SearchOption {
	return func(r *transport.Request) {
	}
}

// WithSearchSuggestText - the source text for which the suggestions should be returned.
func WithSearchSuggestText(suggestText string) SearchOption {
	return func(r *transport.Request) {
	}
}

// WithSearchTerminateAfter - the maximum number of documents to collect for each shard, upon reaching which the query execution will terminate early.
func WithSearchTerminateAfter(terminateAfter int) SearchOption {
	return func(r *transport.Request) {
	}
}

// WithSearchTimeout - explicit operation timeout.
func WithSearchTimeout(timeout time.Duration) SearchOption {
	return func(r *transport.Request) {
	}
}

// WithSearchTrackScores - whether to calculate and return scores even if they are not used for sorting.
func WithSearchTrackScores(trackScores bool) SearchOption {
	return func(r *transport.Request) {
	}
}

// WithSearchTypedKeys - specify whether aggregation and suggester names should be prefixed by their respective types in the response.
func WithSearchTypedKeys(typedKeys bool) SearchOption {
	return func(r *transport.Request) {
	}
}

// WithSearchVersion - specify whether to return document version as part of a hit.
func WithSearchVersion(version bool) SearchOption {
	return func(r *transport.Request) {
	}
}

// WithSearchBody - the search definition using the Query DSL.
func WithSearchBody(body map[string]interface{}) SearchOption {
	return func(r *transport.Request) {
	}
}

// WithSearchErrorTrace - include the stack trace of returned errors.
func WithSearchErrorTrace(errorTrace bool) SearchOption {
	return func(r *transport.Request) {
	}
}

// WithSearchFilterPath - a comma-separated list of filters used to reduce the respone.
func WithSearchFilterPath(filterPath []string) SearchOption {
	return func(r *transport.Request) {
	}
}

// WithSearchHuman - return human readable values for statistics.
func WithSearchHuman(human bool) SearchOption {
	return func(r *transport.Request) {
	}
}

// WithSearchIgnore - ignores the specified HTTP status codes.
func WithSearchIgnore(ignore []int) SearchOption {
	return func(r *transport.Request) {
	}
}

// WithSearchPretty - pretty format the returned JSON response.
func WithSearchPretty(pretty bool) SearchOption {
	return func(r *transport.Request) {
	}
}

// WithSearchSourceParam - the URL-encoded request definition. Useful for libraries that do not accept a request body for non-POST requests.
func WithSearchSourceParam(sourceParam string) SearchOption {
	return func(r *transport.Request) {
	}
}

// Search allows you to execute a search query and get back search hits that match the query. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/search-search.html for more info.
//
// options: optional parameters.
func (a *API) Search(options ...SearchOption) (*SearchResponse, error) {
	req := a.transport.NewRequest("GET")
	for _, option := range options {
		option(req)
	}
	resp, err := a.transport.Do(req)
	return &SearchResponse{resp}, err
}

// SearchResponse is the response for Search.
type SearchResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *SearchResponse) DecodeBody() (util.MapStr, error) {
	return transport.DecodeResponseBody(r.Response)
}
