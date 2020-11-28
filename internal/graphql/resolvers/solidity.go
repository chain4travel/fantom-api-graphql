// Package resolvers implements GraphQL resolvers to incoming API requests.
package resolvers

import "fantom-api-graphql/internal/solidity"

// SolidityVersions resolves a list of Solidity releases supported
// for smart contract validation.
func (rs *rootResolver) SolidityVersions() []string {
	return solidity.SolidityReleases[:]
}