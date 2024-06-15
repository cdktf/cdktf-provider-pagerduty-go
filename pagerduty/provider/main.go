// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-pagerduty.provider.PagerdutyProvider",
		reflect.TypeOf((*PagerdutyProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "alias", GoGetter: "Alias"},
			_jsii_.MemberProperty{JsiiProperty: "aliasInput", GoGetter: "AliasInput"},
			_jsii_.MemberProperty{JsiiProperty: "apiUrlOverride", GoGetter: "ApiUrlOverride"},
			_jsii_.MemberProperty{JsiiProperty: "apiUrlOverrideInput", GoGetter: "ApiUrlOverrideInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "insecureTls", GoGetter: "InsecureTls"},
			_jsii_.MemberProperty{JsiiProperty: "insecureTlsInput", GoGetter: "InsecureTlsInput"},
			_jsii_.MemberProperty{JsiiProperty: "metaAttributes", GoGetter: "MetaAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAlias", GoMethod: "ResetAlias"},
			_jsii_.MemberMethod{JsiiMethod: "resetApiUrlOverride", GoMethod: "ResetApiUrlOverride"},
			_jsii_.MemberMethod{JsiiMethod: "resetInsecureTls", GoMethod: "ResetInsecureTls"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetServiceRegion", GoMethod: "ResetServiceRegion"},
			_jsii_.MemberMethod{JsiiMethod: "resetSkipCredentialsValidation", GoMethod: "ResetSkipCredentialsValidation"},
			_jsii_.MemberMethod{JsiiMethod: "resetToken", GoMethod: "ResetToken"},
			_jsii_.MemberMethod{JsiiMethod: "resetUseAppOauthScopedToken", GoMethod: "ResetUseAppOauthScopedToken"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserToken", GoMethod: "ResetUserToken"},
			_jsii_.MemberProperty{JsiiProperty: "serviceRegion", GoGetter: "ServiceRegion"},
			_jsii_.MemberProperty{JsiiProperty: "serviceRegionInput", GoGetter: "ServiceRegionInput"},
			_jsii_.MemberProperty{JsiiProperty: "skipCredentialsValidation", GoGetter: "SkipCredentialsValidation"},
			_jsii_.MemberProperty{JsiiProperty: "skipCredentialsValidationInput", GoGetter: "SkipCredentialsValidationInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformProviderSource", GoGetter: "TerraformProviderSource"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "token", GoGetter: "Token"},
			_jsii_.MemberProperty{JsiiProperty: "tokenInput", GoGetter: "TokenInput"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "useAppOauthScopedToken", GoGetter: "UseAppOauthScopedToken"},
			_jsii_.MemberProperty{JsiiProperty: "useAppOauthScopedTokenInput", GoGetter: "UseAppOauthScopedTokenInput"},
			_jsii_.MemberProperty{JsiiProperty: "userToken", GoGetter: "UserToken"},
			_jsii_.MemberProperty{JsiiProperty: "userTokenInput", GoGetter: "UserTokenInput"},
		},
		func() interface{} {
			j := jsiiProxy_PagerdutyProvider{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformProvider)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-pagerduty.provider.PagerdutyProviderConfig",
		reflect.TypeOf((*PagerdutyProviderConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-pagerduty.provider.PagerdutyProviderUseAppOauthScopedToken",
		reflect.TypeOf((*PagerdutyProviderUseAppOauthScopedToken)(nil)).Elem(),
	)
}
