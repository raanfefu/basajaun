 

export interface Bundle {
    policies: Policy[]
} 

export interface Policy {
    uri_resource: string
    action: string[]
    statements: Statement[]
}

export interface Statement {
    type: string
    uri: string
    func: Func
}

export interface Func {
    type: string
    term: Term
}

export interface Term {
    type: string
    uri: string
}