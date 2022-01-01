package secret

import (
	"fmt"

	"github.com/hashicorp/vault/api"
	"github.com/kemal576/go-pw-manager/internal/config"
)

func ReadSecrets(vaultPath string) (map[string]string, error) {
	cfg := config.ReadConfiguration("config.yaml")
	conf := &api.Config{Address: "http://" + cfg.VaultAddress + ":" + cfg.VaultPort}

	client, err := api.NewClient(conf)
	if err != nil {
		return nil, err
	}

	client.SetToken(cfg.VaultToken)

	secrets, err2 := client.Logical().Read("secret/data/" + vaultPath)
	if err2 != nil {
		return nil, err2
	}

	dataMap, _ := secrets.Data["data"].(map[string]interface{})

	dataStr := make(map[string]string)
	for k, v := range dataMap {
		dataStr[k] = fmt.Sprintf("%v", v)
	}

	return dataStr, nil
}
