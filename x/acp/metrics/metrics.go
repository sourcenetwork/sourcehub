package metrics

const (

	// counter
	// label for grpc method
	// not sure - host? weird because scraping multiple targets would skew the distribution since everyone executes this msg
	MsgTotal = "sourcehub_acp_msg_total"

	// counter
	// label for grpc method
	MsgErrors = "sourcehub_acp_msg_errors_total"

	// histogram
	// label for grpc method
	MsgSeconds = "sourcehub_acp_msg_seconds"

	// counter
	InvariantViolation = "sourcehub_acp_invariant_violation_total"

	// counter
	// label for grpc method
	QueryTotal = "sourcehub_acp_query_total"

	// counter
	// label for grpc method (?)
	QueryErrors = "sourcehub_acp_query_errors_total"

	// histogram
	// label for grpc method
	// label for error or not
	QuerySeconds = "sourcehub_acp_query_seconds"
)
