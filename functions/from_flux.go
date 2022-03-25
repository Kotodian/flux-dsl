package functions

import (
	"fmt"
	"strings"
)

type fromFlux struct {
	builder  strings.Builder
	bucket   string
	hosts    []string
	location string
}

func (f *fromFlux) OperatorName() string {
	return "from"
}

func (f *fromFlux) WithBucket(bucket string) *fromFlux {
	f.bucket = bucket
	return f
}

func (f *fromFlux) WithHosts(hosts ...string) *fromFlux {
	f.hosts = hosts
	return f
}

func (f *fromFlux) WithLocation(name string) *fromFlux {
	f.location = fmt.Sprintf("timezone.location(name: \"%s\")", name)
	return f
}

func (f *fromFlux) WithLocationFixed(offset string) *fromFlux {
	f.location = fmt.Sprintf("timezone.fixed(offset: \"%s\")", offset)
	return f
}

func (f *fromFlux) Build() string {
	f.builder.WriteString("import \"timezone\"")
	f.builder.WriteString("\n")
	f.builder.WriteString("option location = ")
	f.builder.WriteString(f.location)
	f.builder.WriteString("\n")
	return f.builder.String()
}
