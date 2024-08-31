ui = true

storage "file" {
  path    = "/vault/file"
}

listener "tcp" {
  address = "[::]:8200"

  tls_disable = true
}

