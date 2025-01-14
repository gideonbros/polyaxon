// Copyright 2018-2021 Polyaxon, Inc.
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

/**
 * Polyaxon SDKs and REST API specification.
 * Polyaxon SDKs and REST API specification.
 *
 * The version of the OpenAPI document: 1.10.1
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';

/**
 * The V1RunResources model module.
 * @module model/V1RunResources
 * @version 1.10.1
 */
class V1RunResources {
    /**
     * Constructs a new <code>V1RunResources</code>.
     * @alias module:model/V1RunResources
     */
    constructor() { 
        
        V1RunResources.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>V1RunResources</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1RunResources} obj Optional instance to populate.
     * @return {module:model/V1RunResources} The populated <code>V1RunResources</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1RunResources();

            if (data.hasOwnProperty('cpu')) {
                obj['cpu'] = ApiClient.convertToType(data['cpu'], 'Number');
            }
            if (data.hasOwnProperty('memory')) {
                obj['memory'] = ApiClient.convertToType(data['memory'], 'Number');
            }
            if (data.hasOwnProperty('gpu')) {
                obj['gpu'] = ApiClient.convertToType(data['gpu'], 'Number');
            }
            if (data.hasOwnProperty('custom')) {
                obj['custom'] = ApiClient.convertToType(data['custom'], 'Number');
            }
            if (data.hasOwnProperty('cost')) {
                obj['cost'] = ApiClient.convertToType(data['cost'], 'Number');
            }
        }
        return obj;
    }


}

/**
 * @member {Number} cpu
 */
V1RunResources.prototype['cpu'] = undefined;

/**
 * @member {Number} memory
 */
V1RunResources.prototype['memory'] = undefined;

/**
 * @member {Number} gpu
 */
V1RunResources.prototype['gpu'] = undefined;

/**
 * @member {Number} custom
 */
V1RunResources.prototype['custom'] = undefined;

/**
 * @member {Number} cost
 */
V1RunResources.prototype['cost'] = undefined;






export default V1RunResources;

