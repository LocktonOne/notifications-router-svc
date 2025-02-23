/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package regources

type AtomicSwapQuoteAsset struct {
	Key
	Attributes AtomicSwapQuoteAssetAttributes `json:"attributes"`
}
type AtomicSwapQuoteAssetResponse struct {
	Data     AtomicSwapQuoteAsset `json:"data"`
	Included Included             `json:"included"`
}

type AtomicSwapQuoteAssetListResponse struct {
	Data     []AtomicSwapQuoteAsset `json:"data"`
	Included Included               `json:"included"`
	Links    *Links                 `json:"links"`
}

// MustAtomicSwapQuoteAsset - returns AtomicSwapQuoteAsset from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustAtomicSwapQuoteAsset(key Key) *AtomicSwapQuoteAsset {
	var atomicSwapQuoteAsset AtomicSwapQuoteAsset
	if c.tryFindEntry(key, &atomicSwapQuoteAsset) {
		return &atomicSwapQuoteAsset
	}
	return nil
}
