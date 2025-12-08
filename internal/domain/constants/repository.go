package constants

import . "project/internal/domain/types"

const (
	UnknownError             RepositoryErrorReason = "unknown_error"
	ApplicationError         RepositoryErrorReason = "application_error"
	NotFoundError            RepositoryErrorReason = "not_found_error"
	ForeignKeyViolationError RepositoryErrorReason = "foreign_key_violation_error"
	UniqueConstraintError    RepositoryErrorReason = "unique_constraint_error"
	IndexError               RepositoryErrorReason = "index_error"
	QuerySyntaxError         RepositoryErrorReason = "query_syntax_error"
	TransactionError         RepositoryErrorReason = "transaction_error"
	TimeoutError             RepositoryErrorReason = "timeout_error"
)
