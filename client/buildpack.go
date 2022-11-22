package client

import (
	"context"
	"github.com/cloudfoundry-community/go-cfclient/v3/internal/path"
	"github.com/cloudfoundry-community/go-cfclient/v3/resource"
	"net/url"
)

type BuildpackClient commonClient

// BuildpackListOptions list filters
type BuildpackListOptions struct {
	*ListOptions

	Names  Filter `qs:"names"`  // list of buildpack names to filter by
	Stacks Filter `qs:"stacks"` // list of stack names to filter by
}

// NewBuildpackListOptions creates new options to pass to list
func NewBuildpackListOptions() *BuildpackListOptions {
	return &BuildpackListOptions{
		ListOptions: NewListOptions(),
	}
}

func (o BuildpackListOptions) ToQueryString() url.Values {
	return o.ListOptions.ToQueryString(o)
}

// Create a new buildpack
func (c *BuildpackClient) Create(ctx context.Context, r *resource.BuildpackCreateOrUpdate) (*resource.Buildpack, error) {
	var bp resource.Buildpack
	_, err := c.client.post(ctx, "/v3/buildpacks", r, &bp)
	if err != nil {
		return nil, err
	}
	return &bp, nil
}

// Delete the specified buildpack
func (c *BuildpackClient) Delete(ctx context.Context, guid string) error {
	_, err := c.client.delete(ctx, path.Format("/v3/buildpacks/%s", guid))
	return err
}

// Get retrieves the specified buildpack
func (c *BuildpackClient) Get(ctx context.Context, guid string) (*resource.Buildpack, error) {
	var bp resource.Buildpack
	err := c.client.get(ctx, path.Format("/v3/buildpacks/%s", guid), &bp)
	if err != nil {
		return nil, err
	}
	return &bp, nil
}

// List pages all buildpacks the user has access to
func (c *BuildpackClient) List(ctx context.Context, opts *BuildpackListOptions) ([]*resource.Buildpack, *Pager, error) {
	if opts == nil {
		opts = NewBuildpackListOptions()
	}
	var res resource.BuildpackList
	err := c.client.get(ctx, path.Format("/v3/buildpacks?%s", opts.ToQueryString()), &res)
	if err != nil {
		return nil, nil, err
	}
	pager := NewPager(res.Pagination)
	return res.Resources, pager, nil
}

// ListAll retrieves all buildpacks the user has access to
func (c *BuildpackClient) ListAll(ctx context.Context, opts *BuildpackListOptions) ([]*resource.Buildpack, error) {
	if opts == nil {
		opts = NewBuildpackListOptions()
	}

	var all []*resource.Buildpack
	for {
		page, pager, err := c.List(ctx, opts)
		if err != nil {
			return nil, err
		}
		all = append(all, page...)
		if !pager.HasNextPage() {
			break
		}
		pager.NextPage(opts)
	}
	return all, nil
}

// Update the specified attributes of the buildpack
func (c *BuildpackClient) Update(ctx context.Context, guid string, r *resource.BuildpackCreateOrUpdate) (*resource.Buildpack, error) {
	var bp resource.Buildpack
	_, err := c.client.patch(ctx, path.Format("/v3/buildpacks/%s", guid), r, &bp)
	if err != nil {
		return nil, err
	}
	return &bp, nil
}
