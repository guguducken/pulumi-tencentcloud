# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['OidcSsoArgs', 'OidcSso']

@pulumi.input_type
class OidcSsoArgs:
    def __init__(__self__, *,
                 authorization_endpoint: pulumi.Input[str],
                 client_id: pulumi.Input[str],
                 identity_key: pulumi.Input[str],
                 identity_url: pulumi.Input[str],
                 mapping_filed: pulumi.Input[str],
                 response_mode: pulumi.Input[str],
                 response_type: pulumi.Input[str],
                 scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a OidcSso resource.
        :param pulumi.Input[str] authorization_endpoint: Authorization request Endpoint, OpenID Connect identity provider authorization address. Corresponds to the value of the `authorization_endpoint` field in the Openid-configuration provided by the Enterprise IdP.
        :param pulumi.Input[str] client_id: Client ID, the client ID registered with the OpenID Connect identity provider.
        :param pulumi.Input[str] identity_key: The signature public key requires base64_encode. Verify the public key signed by the OpenID Connect identity provider ID Token. For the security of your account, we recommend that you rotate the signed public key regularly.
        :param pulumi.Input[str] identity_url: Identity provider URL. OpenID Connect identity provider identity.Corresponds to the value of the `issuer` field in the Openid-configuration provided by the Enterprise IdP.
        :param pulumi.Input[str] mapping_filed: Map field names. Which field in the IdP's id_token maps to the user name of the subuser, usually the sub or name field.
        :param pulumi.Input[str] response_mode: Authorize the request Forsonse mode. Authorization request return mode, form_post and frogment two optional modes, recommended to select form_post mode.
        :param pulumi.Input[str] response_type: Authorization requests The Response type, with a fixed value id_token.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] scopes: Authorize the request Scope. openid; email; profile; Authorization request information scope. The default is required openid.
        """
        pulumi.set(__self__, "authorization_endpoint", authorization_endpoint)
        pulumi.set(__self__, "client_id", client_id)
        pulumi.set(__self__, "identity_key", identity_key)
        pulumi.set(__self__, "identity_url", identity_url)
        pulumi.set(__self__, "mapping_filed", mapping_filed)
        pulumi.set(__self__, "response_mode", response_mode)
        pulumi.set(__self__, "response_type", response_type)
        if scopes is not None:
            pulumi.set(__self__, "scopes", scopes)

    @property
    @pulumi.getter(name="authorizationEndpoint")
    def authorization_endpoint(self) -> pulumi.Input[str]:
        """
        Authorization request Endpoint, OpenID Connect identity provider authorization address. Corresponds to the value of the `authorization_endpoint` field in the Openid-configuration provided by the Enterprise IdP.
        """
        return pulumi.get(self, "authorization_endpoint")

    @authorization_endpoint.setter
    def authorization_endpoint(self, value: pulumi.Input[str]):
        pulumi.set(self, "authorization_endpoint", value)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> pulumi.Input[str]:
        """
        Client ID, the client ID registered with the OpenID Connect identity provider.
        """
        return pulumi.get(self, "client_id")

    @client_id.setter
    def client_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "client_id", value)

    @property
    @pulumi.getter(name="identityKey")
    def identity_key(self) -> pulumi.Input[str]:
        """
        The signature public key requires base64_encode. Verify the public key signed by the OpenID Connect identity provider ID Token. For the security of your account, we recommend that you rotate the signed public key regularly.
        """
        return pulumi.get(self, "identity_key")

    @identity_key.setter
    def identity_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "identity_key", value)

    @property
    @pulumi.getter(name="identityUrl")
    def identity_url(self) -> pulumi.Input[str]:
        """
        Identity provider URL. OpenID Connect identity provider identity.Corresponds to the value of the `issuer` field in the Openid-configuration provided by the Enterprise IdP.
        """
        return pulumi.get(self, "identity_url")

    @identity_url.setter
    def identity_url(self, value: pulumi.Input[str]):
        pulumi.set(self, "identity_url", value)

    @property
    @pulumi.getter(name="mappingFiled")
    def mapping_filed(self) -> pulumi.Input[str]:
        """
        Map field names. Which field in the IdP's id_token maps to the user name of the subuser, usually the sub or name field.
        """
        return pulumi.get(self, "mapping_filed")

    @mapping_filed.setter
    def mapping_filed(self, value: pulumi.Input[str]):
        pulumi.set(self, "mapping_filed", value)

    @property
    @pulumi.getter(name="responseMode")
    def response_mode(self) -> pulumi.Input[str]:
        """
        Authorize the request Forsonse mode. Authorization request return mode, form_post and frogment two optional modes, recommended to select form_post mode.
        """
        return pulumi.get(self, "response_mode")

    @response_mode.setter
    def response_mode(self, value: pulumi.Input[str]):
        pulumi.set(self, "response_mode", value)

    @property
    @pulumi.getter(name="responseType")
    def response_type(self) -> pulumi.Input[str]:
        """
        Authorization requests The Response type, with a fixed value id_token.
        """
        return pulumi.get(self, "response_type")

    @response_type.setter
    def response_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "response_type", value)

    @property
    @pulumi.getter
    def scopes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Authorize the request Scope. openid; email; profile; Authorization request information scope. The default is required openid.
        """
        return pulumi.get(self, "scopes")

    @scopes.setter
    def scopes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "scopes", value)


@pulumi.input_type
class _OidcSsoState:
    def __init__(__self__, *,
                 authorization_endpoint: Optional[pulumi.Input[str]] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 identity_key: Optional[pulumi.Input[str]] = None,
                 identity_url: Optional[pulumi.Input[str]] = None,
                 mapping_filed: Optional[pulumi.Input[str]] = None,
                 response_mode: Optional[pulumi.Input[str]] = None,
                 response_type: Optional[pulumi.Input[str]] = None,
                 scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering OidcSso resources.
        :param pulumi.Input[str] authorization_endpoint: Authorization request Endpoint, OpenID Connect identity provider authorization address. Corresponds to the value of the `authorization_endpoint` field in the Openid-configuration provided by the Enterprise IdP.
        :param pulumi.Input[str] client_id: Client ID, the client ID registered with the OpenID Connect identity provider.
        :param pulumi.Input[str] identity_key: The signature public key requires base64_encode. Verify the public key signed by the OpenID Connect identity provider ID Token. For the security of your account, we recommend that you rotate the signed public key regularly.
        :param pulumi.Input[str] identity_url: Identity provider URL. OpenID Connect identity provider identity.Corresponds to the value of the `issuer` field in the Openid-configuration provided by the Enterprise IdP.
        :param pulumi.Input[str] mapping_filed: Map field names. Which field in the IdP's id_token maps to the user name of the subuser, usually the sub or name field.
        :param pulumi.Input[str] response_mode: Authorize the request Forsonse mode. Authorization request return mode, form_post and frogment two optional modes, recommended to select form_post mode.
        :param pulumi.Input[str] response_type: Authorization requests The Response type, with a fixed value id_token.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] scopes: Authorize the request Scope. openid; email; profile; Authorization request information scope. The default is required openid.
        """
        if authorization_endpoint is not None:
            pulumi.set(__self__, "authorization_endpoint", authorization_endpoint)
        if client_id is not None:
            pulumi.set(__self__, "client_id", client_id)
        if identity_key is not None:
            pulumi.set(__self__, "identity_key", identity_key)
        if identity_url is not None:
            pulumi.set(__self__, "identity_url", identity_url)
        if mapping_filed is not None:
            pulumi.set(__self__, "mapping_filed", mapping_filed)
        if response_mode is not None:
            pulumi.set(__self__, "response_mode", response_mode)
        if response_type is not None:
            pulumi.set(__self__, "response_type", response_type)
        if scopes is not None:
            pulumi.set(__self__, "scopes", scopes)

    @property
    @pulumi.getter(name="authorizationEndpoint")
    def authorization_endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        Authorization request Endpoint, OpenID Connect identity provider authorization address. Corresponds to the value of the `authorization_endpoint` field in the Openid-configuration provided by the Enterprise IdP.
        """
        return pulumi.get(self, "authorization_endpoint")

    @authorization_endpoint.setter
    def authorization_endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authorization_endpoint", value)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> Optional[pulumi.Input[str]]:
        """
        Client ID, the client ID registered with the OpenID Connect identity provider.
        """
        return pulumi.get(self, "client_id")

    @client_id.setter
    def client_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_id", value)

    @property
    @pulumi.getter(name="identityKey")
    def identity_key(self) -> Optional[pulumi.Input[str]]:
        """
        The signature public key requires base64_encode. Verify the public key signed by the OpenID Connect identity provider ID Token. For the security of your account, we recommend that you rotate the signed public key regularly.
        """
        return pulumi.get(self, "identity_key")

    @identity_key.setter
    def identity_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "identity_key", value)

    @property
    @pulumi.getter(name="identityUrl")
    def identity_url(self) -> Optional[pulumi.Input[str]]:
        """
        Identity provider URL. OpenID Connect identity provider identity.Corresponds to the value of the `issuer` field in the Openid-configuration provided by the Enterprise IdP.
        """
        return pulumi.get(self, "identity_url")

    @identity_url.setter
    def identity_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "identity_url", value)

    @property
    @pulumi.getter(name="mappingFiled")
    def mapping_filed(self) -> Optional[pulumi.Input[str]]:
        """
        Map field names. Which field in the IdP's id_token maps to the user name of the subuser, usually the sub or name field.
        """
        return pulumi.get(self, "mapping_filed")

    @mapping_filed.setter
    def mapping_filed(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mapping_filed", value)

    @property
    @pulumi.getter(name="responseMode")
    def response_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Authorize the request Forsonse mode. Authorization request return mode, form_post and frogment two optional modes, recommended to select form_post mode.
        """
        return pulumi.get(self, "response_mode")

    @response_mode.setter
    def response_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "response_mode", value)

    @property
    @pulumi.getter(name="responseType")
    def response_type(self) -> Optional[pulumi.Input[str]]:
        """
        Authorization requests The Response type, with a fixed value id_token.
        """
        return pulumi.get(self, "response_type")

    @response_type.setter
    def response_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "response_type", value)

    @property
    @pulumi.getter
    def scopes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Authorize the request Scope. openid; email; profile; Authorization request information scope. The default is required openid.
        """
        return pulumi.get(self, "scopes")

    @scopes.setter
    def scopes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "scopes", value)


class OidcSso(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorization_endpoint: Optional[pulumi.Input[str]] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 identity_key: Optional[pulumi.Input[str]] = None,
                 identity_url: Optional[pulumi.Input[str]] = None,
                 mapping_filed: Optional[pulumi.Input[str]] = None,
                 response_mode: Optional[pulumi.Input[str]] = None,
                 response_type: Optional[pulumi.Input[str]] = None,
                 scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides a resource to create a CAM-OIDC-SSO.

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        foo = tencentcloud.cam.OidcSso("foo",
            authorization_endpoint="https://login.microsoftonline.com/.../oauth2/v2.0/authorize",
            client_id="...",
            identity_key="...",
            identity_url="https://login.microsoftonline.com/.../v2.0",
            mapping_filed="name",
            response_mode="form_post",
            response_type="id_token",
            scopes=[
                "openid",
                "email",
            ])
        ```

        ## Import

        CAM-OIDC-SSO can be imported using the client_id or any string which can identifier resource, e.g.

        ```sh
         $ pulumi import tencentcloud:Cam/oidcSso:OidcSso foo xxxxxxxxxxx
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] authorization_endpoint: Authorization request Endpoint, OpenID Connect identity provider authorization address. Corresponds to the value of the `authorization_endpoint` field in the Openid-configuration provided by the Enterprise IdP.
        :param pulumi.Input[str] client_id: Client ID, the client ID registered with the OpenID Connect identity provider.
        :param pulumi.Input[str] identity_key: The signature public key requires base64_encode. Verify the public key signed by the OpenID Connect identity provider ID Token. For the security of your account, we recommend that you rotate the signed public key regularly.
        :param pulumi.Input[str] identity_url: Identity provider URL. OpenID Connect identity provider identity.Corresponds to the value of the `issuer` field in the Openid-configuration provided by the Enterprise IdP.
        :param pulumi.Input[str] mapping_filed: Map field names. Which field in the IdP's id_token maps to the user name of the subuser, usually the sub or name field.
        :param pulumi.Input[str] response_mode: Authorize the request Forsonse mode. Authorization request return mode, form_post and frogment two optional modes, recommended to select form_post mode.
        :param pulumi.Input[str] response_type: Authorization requests The Response type, with a fixed value id_token.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] scopes: Authorize the request Scope. openid; email; profile; Authorization request information scope. The default is required openid.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: OidcSsoArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a CAM-OIDC-SSO.

        ## Example Usage

        ```python
        import pulumi
        import tencentcloud_iac_pulumi as tencentcloud

        foo = tencentcloud.cam.OidcSso("foo",
            authorization_endpoint="https://login.microsoftonline.com/.../oauth2/v2.0/authorize",
            client_id="...",
            identity_key="...",
            identity_url="https://login.microsoftonline.com/.../v2.0",
            mapping_filed="name",
            response_mode="form_post",
            response_type="id_token",
            scopes=[
                "openid",
                "email",
            ])
        ```

        ## Import

        CAM-OIDC-SSO can be imported using the client_id or any string which can identifier resource, e.g.

        ```sh
         $ pulumi import tencentcloud:Cam/oidcSso:OidcSso foo xxxxxxxxxxx
        ```

        :param str resource_name: The name of the resource.
        :param OidcSsoArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OidcSsoArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorization_endpoint: Optional[pulumi.Input[str]] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 identity_key: Optional[pulumi.Input[str]] = None,
                 identity_url: Optional[pulumi.Input[str]] = None,
                 mapping_filed: Optional[pulumi.Input[str]] = None,
                 response_mode: Optional[pulumi.Input[str]] = None,
                 response_type: Optional[pulumi.Input[str]] = None,
                 scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = OidcSsoArgs.__new__(OidcSsoArgs)

            if authorization_endpoint is None and not opts.urn:
                raise TypeError("Missing required property 'authorization_endpoint'")
            __props__.__dict__["authorization_endpoint"] = authorization_endpoint
            if client_id is None and not opts.urn:
                raise TypeError("Missing required property 'client_id'")
            __props__.__dict__["client_id"] = client_id
            if identity_key is None and not opts.urn:
                raise TypeError("Missing required property 'identity_key'")
            __props__.__dict__["identity_key"] = identity_key
            if identity_url is None and not opts.urn:
                raise TypeError("Missing required property 'identity_url'")
            __props__.__dict__["identity_url"] = identity_url
            if mapping_filed is None and not opts.urn:
                raise TypeError("Missing required property 'mapping_filed'")
            __props__.__dict__["mapping_filed"] = mapping_filed
            if response_mode is None and not opts.urn:
                raise TypeError("Missing required property 'response_mode'")
            __props__.__dict__["response_mode"] = response_mode
            if response_type is None and not opts.urn:
                raise TypeError("Missing required property 'response_type'")
            __props__.__dict__["response_type"] = response_type
            __props__.__dict__["scopes"] = scopes
        super(OidcSso, __self__).__init__(
            'tencentcloud:Cam/oidcSso:OidcSso',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            authorization_endpoint: Optional[pulumi.Input[str]] = None,
            client_id: Optional[pulumi.Input[str]] = None,
            identity_key: Optional[pulumi.Input[str]] = None,
            identity_url: Optional[pulumi.Input[str]] = None,
            mapping_filed: Optional[pulumi.Input[str]] = None,
            response_mode: Optional[pulumi.Input[str]] = None,
            response_type: Optional[pulumi.Input[str]] = None,
            scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'OidcSso':
        """
        Get an existing OidcSso resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] authorization_endpoint: Authorization request Endpoint, OpenID Connect identity provider authorization address. Corresponds to the value of the `authorization_endpoint` field in the Openid-configuration provided by the Enterprise IdP.
        :param pulumi.Input[str] client_id: Client ID, the client ID registered with the OpenID Connect identity provider.
        :param pulumi.Input[str] identity_key: The signature public key requires base64_encode. Verify the public key signed by the OpenID Connect identity provider ID Token. For the security of your account, we recommend that you rotate the signed public key regularly.
        :param pulumi.Input[str] identity_url: Identity provider URL. OpenID Connect identity provider identity.Corresponds to the value of the `issuer` field in the Openid-configuration provided by the Enterprise IdP.
        :param pulumi.Input[str] mapping_filed: Map field names. Which field in the IdP's id_token maps to the user name of the subuser, usually the sub or name field.
        :param pulumi.Input[str] response_mode: Authorize the request Forsonse mode. Authorization request return mode, form_post and frogment two optional modes, recommended to select form_post mode.
        :param pulumi.Input[str] response_type: Authorization requests The Response type, with a fixed value id_token.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] scopes: Authorize the request Scope. openid; email; profile; Authorization request information scope. The default is required openid.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _OidcSsoState.__new__(_OidcSsoState)

        __props__.__dict__["authorization_endpoint"] = authorization_endpoint
        __props__.__dict__["client_id"] = client_id
        __props__.__dict__["identity_key"] = identity_key
        __props__.__dict__["identity_url"] = identity_url
        __props__.__dict__["mapping_filed"] = mapping_filed
        __props__.__dict__["response_mode"] = response_mode
        __props__.__dict__["response_type"] = response_type
        __props__.__dict__["scopes"] = scopes
        return OidcSso(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="authorizationEndpoint")
    def authorization_endpoint(self) -> pulumi.Output[str]:
        """
        Authorization request Endpoint, OpenID Connect identity provider authorization address. Corresponds to the value of the `authorization_endpoint` field in the Openid-configuration provided by the Enterprise IdP.
        """
        return pulumi.get(self, "authorization_endpoint")

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> pulumi.Output[str]:
        """
        Client ID, the client ID registered with the OpenID Connect identity provider.
        """
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="identityKey")
    def identity_key(self) -> pulumi.Output[str]:
        """
        The signature public key requires base64_encode. Verify the public key signed by the OpenID Connect identity provider ID Token. For the security of your account, we recommend that you rotate the signed public key regularly.
        """
        return pulumi.get(self, "identity_key")

    @property
    @pulumi.getter(name="identityUrl")
    def identity_url(self) -> pulumi.Output[str]:
        """
        Identity provider URL. OpenID Connect identity provider identity.Corresponds to the value of the `issuer` field in the Openid-configuration provided by the Enterprise IdP.
        """
        return pulumi.get(self, "identity_url")

    @property
    @pulumi.getter(name="mappingFiled")
    def mapping_filed(self) -> pulumi.Output[str]:
        """
        Map field names. Which field in the IdP's id_token maps to the user name of the subuser, usually the sub or name field.
        """
        return pulumi.get(self, "mapping_filed")

    @property
    @pulumi.getter(name="responseMode")
    def response_mode(self) -> pulumi.Output[str]:
        """
        Authorize the request Forsonse mode. Authorization request return mode, form_post and frogment two optional modes, recommended to select form_post mode.
        """
        return pulumi.get(self, "response_mode")

    @property
    @pulumi.getter(name="responseType")
    def response_type(self) -> pulumi.Output[str]:
        """
        Authorization requests The Response type, with a fixed value id_token.
        """
        return pulumi.get(self, "response_type")

    @property
    @pulumi.getter
    def scopes(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Authorize the request Scope. openid; email; profile; Authorization request information scope. The default is required openid.
        """
        return pulumi.get(self, "scopes")

