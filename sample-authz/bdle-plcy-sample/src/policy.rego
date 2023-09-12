package envoy.authz

import input.attributes.request.http as http_request
import future.keywords.if

default allow = false

allow if {
	input.attributes.request.http.method == "GET"
	input.attributes.request.http.path == "/health"
}