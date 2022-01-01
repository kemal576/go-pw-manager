package models

type Config struct {
	VaultAddress string `yaml:"Vault_Adress"`
	VaultPort    string `yaml:"Vault_Port"`
	VaultToken   string `yaml:"Vault_Token"`
}
