package services

import (
	"context"
	"net/http"

	"github.com/asim/go-micro/v3/errors"
	"github.com/asim/go-micro/v3/metadata"
	"github.com/google/uuid"
	"github.com/nuveo/log"
	"github.com/slory7/angulargo/src/infrastructure/business/contracts"
	"golang.org/x/net/trace"
)

const (
	ServiceNameTrending string = "srv.trending"
	ServiceNameGather   string = "srv.gather"
	ServiceNameApi      string = "api"
	ServiceNameApiWeb   string = "api.web"
)

func GetTrace(ctx context.Context) (traceID, fromName string) {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		return
	}
	traceID = md["Traceid"]
	fromName = md["Fromname"]
	return
}

func PrintTrace(ctx context.Context, localMethodName string) {
	traceID, fromName := GetTrace(ctx)
	if tr, ok := trace.FromContext(ctx); ok {
		tr.LazyPrintf("fromName: %s", fromName)
		tr.LazyPrintf("traceID: %s", traceID)
		tr.LazyPrintf("localMethodName: %s", localMethodName)
	}
	log.Printf("fromName: %s\n", fromName)
	log.Printf("traceID: %s\n", traceID)
	log.Printf("localMethodName: %s\n", localMethodName)
}

func GetContextWithTrace(ctx context.Context, fromNameIfEmpty string) context.Context {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		md = metadata.Metadata{}
	}
	traceID := md["X-Request-ID"]
	if len(traceID) == 0 {
		traceID = md["Traceid"]
	}
	if len(traceID) == 0 {
		traceID = uuid.New().String()
	}
	if len(md["Fromname"]) == 0 {
		md["Fromname"] = fromNameIfEmpty
	}
	md["Traceid"] = traceID

	ctx = metadata.NewContext(ctx, md)

	return ctx
}

func ToMicroError(err *contracts.BizError) error {
	switch err.Status {
	case contracts.NotFound:
		return errors.NotFound("BizError", "%s", err.Message)
	case contracts.Forbidden:
		return errors.Forbidden("BizError", "%s", err.Message)
	case contracts.Unauthorized:
		return errors.Unauthorized("BizError", "%s", err.Message)
	case contracts.Conflict:
		return errors.Conflict("BizError", "%s", err.Message)
	case contracts.BadData:
		return errors.BadRequest("BizError", "%s", err.Message)
	case contracts.BadLogic:
		return &errors.Error{Id: "BizError", Code: http.StatusUnprocessableEntity, Detail: err.Message, Status: http.StatusText(http.StatusUnprocessableEntity)}
	case contracts.Error:
		return errors.InternalServerError("BizError", "%s", err.Message)
	case contracts.Timeout:
		return errors.Timeout("BizError", "%s", err.Message)
	}
	return nil
}
