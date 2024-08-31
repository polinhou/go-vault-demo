package main

import (
    "context"
    "encoding/base64"
    "log"
    "os"
    "fmt"

    vault "github.com/hashicorp/vault/api"
)

const password string = "admin1234"

func main() {
    config := vault.DefaultConfig()
    config.Address = os.Getenv("VAULT_ADDR")

    client, err := vault.NewClient(config)
    if err != nil {
        log.Fatalf("Unable to initialize a Vault client: %v", err)
    }

    client.SetToken(os.Getenv("VAULT_TOKEN"))

    secretData := map[string]interface{}{
        "password": password,
    }

    ctx := context.Background()

    _, err = client.KVv2("secret").Put(ctx, "my-secret-password", secretData)
    if err != nil {
        log.Fatalf("Unable to write secret: %v to the vault", err)
    }
    log.Println("Super secret password written successfully to the vault.")

    plaintext := "my secret data"
    plaintextBase64 := base64.StdEncoding.EncodeToString([]byte(plaintext))

    // 加密数据
    encryptPath := "transit/encrypt/my-key"
    data := map[string]interface{}{
        "plaintext": plaintextBase64,
    }

    secret, err := client.Logical().Write(encryptPath, data)
    if err != nil {
        log.Fatalf("Error encrypting data: %v", err)
    }

    ciphertext := secret.Data["ciphertext"].(string)
    fmt.Printf("Encrypted data: %s\n", ciphertext)

    // 解密数据
    decryptPath := "transit/decrypt/my-key"
    data = map[string]interface{}{
        "ciphertext": ciphertext,
    }

    secret, err = client.Logical().Write(decryptPath, data)
    if err != nil {
        log.Fatalf("Error decrypting data: %v", err)
    }

    decryptedBase64 := secret.Data["plaintext"].(string)
    decryptedData, err := base64.StdEncoding.DecodeString(decryptedBase64)
    if err != nil {
        log.Fatalf("Error decoding base64 data: %v", err)
    }

    fmt.Printf("Decrypted data: %s\n", string(decryptedData))
}
