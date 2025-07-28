package extension

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/umefy/go-web-app-template/internal/infrastructure/database"
	"github.com/umefy/go-web-app-template/internal/infrastructure/logger"
	"github.com/vektah/gqlparser/v2/ast"
)

type TransactionExtension struct {
	DbQuery *database.Query
	Logger  logger.Logger
}

func (e *TransactionExtension) ExtensionName() string {
	return "TransactionExtension"
}

func (e *TransactionExtension) Validate(schema graphql.ExecutableSchema) error {
	return nil
}

func (e *TransactionExtension) InterceptOperation(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
	oc := graphql.GetOperationContext(ctx)

	if oc.Operation != nil && oc.Operation.Operation == ast.Mutation {
		var resp *graphql.Response
		return func(_ context.Context) *graphql.Response {
			_, err := database.WithTx(ctx, e.DbQuery, e.Logger, func(ctx context.Context, tx *database.QueryTx) (any, error) {
				ctx = context.WithValue(ctx, database.TransactionCtxKey, tx)
				resp = next(ctx)(ctx)
				return nil, nil
			})
			if err != nil {
				return graphql.ErrorResponse(ctx, "transaction error: %v", err)
			}
			return resp
		}
	}

	return next(ctx)
}
