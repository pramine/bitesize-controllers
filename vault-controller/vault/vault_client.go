package vault

import (
    "strings"
    "errors"
    "time"
    vault "github.com/hashicorp/vault/api"
    log "github.com/Sirupsen/logrus"
)

type VaultClient struct {
    Client *vault.Client
}

const initWait = "5s"

func NewVaultClient(address string, token string) (*VaultClient, error) {

    config := vault.DefaultConfig()
    config.Address = address
    if address == "" || token == "" {
        log.Errorf("Vault not configured")
        return nil, nil
    }

    client, err := vault.NewClient(config)
    if err != nil {
        log.Errorf("Vault config failed.")
        return &VaultClient{nil}, err
    }
    client.SetToken(token)
    return &VaultClient{ Client: client }, err
}

func (c *VaultClient) InitStatus() (initState bool, err error) {

    status, err := c.Client.Sys().InitStatus()
    if err != nil {
        log.Errorf("Error retrieving vault init status")
        return false, err
    } else {
        log.Debugf("InitStatus: %v", status)
        return status, err
    }
}

// Init with defaults
func (c *VaultClient) Init() (initResponse *vault.InitResponse, err error) {

    response, err := c.Client.Sys().Init(&vault.InitRequest{})
    if err != nil {
        log.Errorf("Error initializing Vault! %v", err)
        return response, err
    } else {
        log.Infof("Initialised instance %v", c.Client.Address())
        log.Debugf("InitStatus: %v", response)
        w, _ := time.ParseDuration(initWait)
        time.Sleep(w)
        return response, err
    }
}

// SealStatus returns true if vault is unsealed
func (c *VaultClient) SealStatus() (sealState bool, err error) {

    status, err := c.Client.Sys().SealStatus()
    if err != nil || status == nil {
        log.Errorf("Error retrieving vault seal status")
        return true, err
    } else {
        log.Debugf("SealStatus: %v", status)
        return status.Sealed, err
    }
}

func (c *VaultClient) Unseal(unsealKeys string) (sealState bool, err error) {

    for _, key := range strings.Split(unsealKeys, ",") {
        log.Debugf("Unseal key: %v", key)
        if len(key) <= 0 {
            continue
        }
        resp, err := c.Client.Sys().Unseal(key)
        if err != nil || resp == nil {
            log.Errorf("Error Unsealing: %v", err)
        }
        if resp.Sealed == false {
            log.Infof("Instance unsealed")
            return true, nil
        } else {
            log.Infof("Instance seal progress: %v", resp.Progress)
        }
    }
    err = errors.New("Insufficient unseal keys! Instance sealed.")
    return false, err
}
// Ready returns true if vault is unsealed
func (c *VaultClient) LeaderStatus() (leaderState bool, err error) {

    status, err := c.Client.Sys().Leader()
    if err != nil || status == nil {
        log.Errorf("Error retrieving vault leader status")
        return false, err
    } else {
        log.Debugf("LeaderStatus: %v", status)
        return status.IsSelf, err
    }

}
