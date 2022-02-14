// Copyright (c) 2022 Tigera, Inc. All rights reserved.

// Code generated by client-gen. DO NOT EDIT.

package v3

import (
	"context"
	"time"

	v3 "github.com/projectcalico/api/pkg/apis/projectcalico/v3"
	scheme "github.com/projectcalico/api/pkg/client/clientset_generated/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ProfilesGetter has a method to return a ProfileInterface.
// A group's client should implement this interface.
type ProfilesGetter interface {
	Profiles() ProfileInterface
}

// ProfileInterface has methods to work with Profile resources.
type ProfileInterface interface {
	Create(ctx context.Context, profile *v3.Profile, opts v1.CreateOptions) (*v3.Profile, error)
	Update(ctx context.Context, profile *v3.Profile, opts v1.UpdateOptions) (*v3.Profile, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v3.Profile, error)
	List(ctx context.Context, opts v1.ListOptions) (*v3.ProfileList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.Profile, err error)
	ProfileExpansion
}

// profiles implements ProfileInterface
type profiles struct {
	client rest.Interface
}

// newProfiles returns a Profiles
func newProfiles(c *ProjectcalicoV3Client) *profiles {
	return &profiles{
		client: c.RESTClient(),
	}
}

// Get takes name of the profile, and returns the corresponding profile object, and an error if there is any.
func (c *profiles) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.Profile, err error) {
	result = &v3.Profile{}
	err = c.client.Get().
		Resource("profiles").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Profiles that match those selectors.
func (c *profiles) List(ctx context.Context, opts v1.ListOptions) (result *v3.ProfileList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v3.ProfileList{}
	err = c.client.Get().
		Resource("profiles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested profiles.
func (c *profiles) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("profiles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a profile and creates it.  Returns the server's representation of the profile, and an error, if there is any.
func (c *profiles) Create(ctx context.Context, profile *v3.Profile, opts v1.CreateOptions) (result *v3.Profile, err error) {
	result = &v3.Profile{}
	err = c.client.Post().
		Resource("profiles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(profile).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a profile and updates it. Returns the server's representation of the profile, and an error, if there is any.
func (c *profiles) Update(ctx context.Context, profile *v3.Profile, opts v1.UpdateOptions) (result *v3.Profile, err error) {
	result = &v3.Profile{}
	err = c.client.Put().
		Resource("profiles").
		Name(profile.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(profile).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the profile and deletes it. Returns an error if one occurs.
func (c *profiles) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("profiles").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *profiles) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("profiles").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched profile.
func (c *profiles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.Profile, err error) {
	result = &v3.Profile{}
	err = c.client.Patch(pt).
		Resource("profiles").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}