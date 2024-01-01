package http_server

import "log/slog"

func HTTPErr(err error) slog.Attr {
	er := slog.Attr{
		"err",
		slog.StringValue(err.Error()),
	}
	return er
}
