# bugsnag-stacktracer

`bugsnag-stacktracer` helps bridge the gap between the popular https://github.com/pkg/errors package's error stacktrace recording and https://bugsnag/bugsnag-go package's expected error stacktrace format.

## Usage

Wrap your pkg/error `stackTracer` compatible error with `bugsnag_stacktracer.FromError(err)`. The resulting `bugsnag_stacktracer.Error` will properly pass along your stacktrace to Bugsnag.