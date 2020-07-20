// Copyright 2018-2020 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/* tslint:disable */
/* eslint-disable */
/**
 * Polyaxon SDKs and REST API specification.
 * Polyaxon SDKs and REST API specification.
 *
 * The version of the OpenAPI document: 1.1.4
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface V1Token
 */
export interface V1Token {
    /**
     * 
     * @type {string}
     * @memberof V1Token
     */
    uuid?: string;
    /**
     * 
     * @type {string}
     * @memberof V1Token
     */
    key?: string;
    /**
     * 
     * @type {string}
     * @memberof V1Token
     */
    name?: string;
    /**
     * 
     * @type {Array<string>}
     * @memberof V1Token
     */
    scopes?: Array<string>;
    /**
     * 
     * @type {Array<string>}
     * @memberof V1Token
     */
    services?: Array<string>;
    /**
     * 
     * @type {Date}
     * @memberof V1Token
     */
    started_at?: Date;
    /**
     * 
     * @type {Date}
     * @memberof V1Token
     */
    expires_at?: Date;
    /**
     * 
     * @type {Date}
     * @memberof V1Token
     */
    created_at?: Date;
    /**
     * 
     * @type {Date}
     * @memberof V1Token
     */
    updated_at?: Date;
}

export function V1TokenFromJSON(json: any): V1Token {
    return V1TokenFromJSONTyped(json, false);
}

export function V1TokenFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1Token {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'uuid': !exists(json, 'uuid') ? undefined : json['uuid'],
        'key': !exists(json, 'key') ? undefined : json['key'],
        'name': !exists(json, 'name') ? undefined : json['name'],
        'scopes': !exists(json, 'scopes') ? undefined : json['scopes'],
        'services': !exists(json, 'services') ? undefined : json['services'],
        'started_at': !exists(json, 'started_at') ? undefined : (new Date(json['started_at'])),
        'expires_at': !exists(json, 'expires_at') ? undefined : (new Date(json['expires_at'])),
        'created_at': !exists(json, 'created_at') ? undefined : (new Date(json['created_at'])),
        'updated_at': !exists(json, 'updated_at') ? undefined : (new Date(json['updated_at'])),
    };
}

export function V1TokenToJSON(value?: V1Token | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'uuid': value.uuid,
        'key': value.key,
        'name': value.name,
        'scopes': value.scopes,
        'services': value.services,
        'started_at': value.started_at === undefined ? undefined : (value.started_at.toISOString()),
        'expires_at': value.expires_at === undefined ? undefined : (value.expires_at.toISOString()),
        'created_at': value.created_at === undefined ? undefined : (value.created_at.toISOString()),
        'updated_at': value.updated_at === undefined ? undefined : (value.updated_at.toISOString()),
    };
}

