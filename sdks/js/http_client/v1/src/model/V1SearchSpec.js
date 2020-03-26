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

/*
 * Polyaxon SDKs and REST API specification.
 * Polyaxon SDKs and REST API specification.
 *
 * OpenAPI spec version: 1.0.7
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.4.10
 *
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['ApiClient'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'));
  } else {
    // Browser globals (root is window)
    if (!root.PolyaxonSdk) {
      root.PolyaxonSdk = {};
    }
    root.PolyaxonSdk.V1SearchSpec = factory(root.PolyaxonSdk.ApiClient);
  }
}(this, function(ApiClient) {
  'use strict';

  /**
   * The V1SearchSpec model module.
   * @module model/V1SearchSpec
   * @version 1.0.7
   */

  /**
   * Constructs a new <code>V1SearchSpec</code>.
   * @alias module:model/V1SearchSpec
   * @class
   */
  var exports = function() {
  };

  /**
   * Constructs a <code>V1SearchSpec</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/V1SearchSpec} obj Optional instance to populate.
   * @return {module:model/V1SearchSpec} The populated <code>V1SearchSpec</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();
      if (data.hasOwnProperty('query'))
        obj.query = ApiClient.convertToType(data['query'], 'String');
      if (data.hasOwnProperty('sort'))
        obj.sort = ApiClient.convertToType(data['sort'], 'String');
      if (data.hasOwnProperty('limit'))
        obj.limit = ApiClient.convertToType(data['limit'], 'Number');
      if (data.hasOwnProperty('groupby'))
        obj.groupby = ApiClient.convertToType(data['groupby'], 'String');
      if (data.hasOwnProperty('columns'))
        obj.columns = ApiClient.convertToType(data['columns'], 'String');
    }
    return obj;
  }

  /**
   * @member {String} query
   */
  exports.prototype.query = undefined;

  /**
   * @member {String} sort
   */
  exports.prototype.sort = undefined;

  /**
   * @member {Number} limit
   */
  exports.prototype.limit = undefined;

  /**
   * @member {String} groupby
   */
  exports.prototype.groupby = undefined;

  /**
   * @member {String} columns
   */
  exports.prototype.columns = undefined;

  return exports;

}));