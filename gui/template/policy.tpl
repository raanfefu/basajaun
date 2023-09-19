package envoy.authz

import future.keywords.if
import future.keywords.in

default allow = false

{%- macro term_render(statement) -%} 
    {%- if 'body' in statement.type -%} 
         input.parsed_body.{{ statement.uri }} 
    {%- endif -%}
    {%- if 'header' in statement.type -%} 
         input.attributes.request.http.headers["{{ statement.uri }}"]
    {%- endif -%}
    {%- if 'data' in statement.type -%} 
         data.{{ statement.uri }}
    {%- endif -%}
    {%- if 'claims' in statement.type -%} 
         claims.{{ statement.uri }}
    {%- endif -%}
    {%- if 'const' in statement.type -%} 
         "{{ statement.uri }}"
    {%- endif -%}
{%- endmacro -%}

{%- macro func_render(statement) -%} 
    {%- if 'equals' == statement.func.type -%} 
        ==
    {%- endif -%}
    {%- if 'in' == statement.func.type -%} 
        in
    {%- endif -%}
    {%- if 'not-equals' == statement.func.type -%} 
        !=
    {%- endif -%}
{%- endmacro -%}
 

{%+ for policy in data.policies  +%}

allow if {
    input.attributes.request.http.method in [{%- for method in policy.action  -%}"{{ method }}",{% endfor %}]
    input.attributes.request.http.path   == "{{ policy.uri_resource }}"
 {% for statement in policy.statements  %}    
    {{ term_render(statement) }} {{ func_render(statement) }} {{ term_render(statement.func.term) }} 
 {% endfor %}
}
{% endfor %}

{%- for policy in data.policies  +%}
{% for statement in policy.statements  %}    
{%- if 'claims' == statement.type -%}

claims := payload if {
    io.jwt.verify_hs256(bearer_token, "{{signature}}")
    [_, payload, _] := io.jwt.decode(bearer_token)
}   

bearer_token := t if {
    v := input.attributes.request.http.headers.authorization
    startswith(v, "Bearer ")
    t := substring(v, count("Bearer "), -1)
}
{%- endif -%}
{% endfor %}
{% endfor %}