package provider

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strconv"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

type afraidOrgClient struct {
	username  string
	password  string
	dnscookie string
}

type AfraidOrgProvider struct {
	dryRun       bool
	domainFilter DomainFilter
	client       coreDNSClient
}

func (p *afraidOrgClient) doLogin() error {
	data := url.Values{}
	data.Set("username", p.username)
	data.Set("password", p.password)
	data.Set("submit", "Login")
	data.Set("remote", "")
	data.Set("from", "")
	data.Set("action", "auth")
	resp, err := http.PostForm("https://freedns.afraid.org/zc.php?step=2", data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		err = errors.New(url +
			"\nresp.StatusCode: " + strconv.Itoa(resp.StatusCode))
		return
	}

	hClient := &http.Client{}
	resp, err := hClient.Do(req)
	if err != nil {
		return
	}
}

func (p *AfraidOrgProvider) Records(context.Context) (endpoints []*endpoint.Endpoint, err error) {
}

func (p *AfraidOrgProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
}
