package constants

import . "project/internal/domain/types"

const (
	RepositoryUnknownError             RepositoryErrorReason = "unknown_error"
	RepositoryApplicationError         RepositoryErrorReason = "application_error"
	RepositoryNotFoundError            RepositoryErrorReason = "not_found_error"
	RepositoryForeignKeyViolationError RepositoryErrorReason = "foreign_key_violation_error"
	RepositoryUniqueConstraintError    RepositoryErrorReason = "unique_constraint_error"
	RepositoryIndexError               RepositoryErrorReason = "index_error"
	RepositoryQuerySyntaxError         RepositoryErrorReason = "query_syntax_error"
	RepositoryTransactionError         RepositoryErrorReason = "transaction_error"
	RepositoryTimeoutError             RepositoryErrorReason = "timeout_error"
)

const (
	DefaultRepositoryStackSkip   StackSkip   = 3
	DefaultRepositoryStackLength StackLength = 10
)
