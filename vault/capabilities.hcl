path "secret/users" {
    capabilities = ["create", "read", "update", "delete", "list"]
}

path "secret/users/*" {
    capabilities = ["create", "read", "update", "delete", "list"]
}