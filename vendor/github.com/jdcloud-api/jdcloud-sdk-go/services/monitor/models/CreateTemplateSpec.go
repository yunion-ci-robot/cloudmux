// Copyright 2018 JDCLOUD.COM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package models


type CreateTemplateSpec struct {

    /* 幂等性校验参数,最长36位  */
    ClientToken string `json:"clientToken"`

    /* 模板描述 (Optional) */
    Description string `json:"description"`

    /* 规则的资源类型  */
    RuleServiceCode string `json:"ruleServiceCode"`

    /* 模板的资源类型  */
    ServiceCode string `json:"serviceCode"`

    /* 模板名称,长度1-32个字符，只允许中英文、数字、''-''和"_"  */
    TemplateName string `json:"templateName"`

    /* 模板内包含的规则  */
    TemplateRules []BaseRuleT `json:"templateRules"`
}