/*
 * Copyright 2022 InfAI (CC SES)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package mocks

const conceptListStr = `[{"id": "mock-placeholder-concept","characteristic_ids": ["example_brightness","example_lux","example_color","example_rgb","example_hex"]},{"base_characteristic_id":"urn:infai:ses:characteristic:f6b7ee6e-423c-47f1-8b1f-c489461d3039","characteristic_ids":["urn:infai:ses:characteristic:f6b7ee6e-423c-47f1-8b1f-c489461d3039"],"creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","id":"urn:infai:ses:concept:08cde541-54dc-4d92-b2d9-b07304bd659a","name":"benzene","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"shared":false},{"base_characteristic_id":"urn:infai:ses:characteristic:7dc1bb7e-b256-408a-a6f9-044dc60fdcf5","characteristic_ids":["urn:infai:ses:characteristic:8e520c73-7ee1-4a4b-954b-1c7c55139e27","urn:infai:ses:characteristic:7621686a-56bc-402d-b4cc-5b266d39736f","urn:infai:ses:characteristic:c0353532-a8fb-4553-a00b-418cb8a80a65","urn:infai:ses:characteristic:7dc1bb7e-b256-408a-a6f9-044dc60fdcf5"],"creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","id":"urn:infai:ses:concept:ebfeabb3-50f0-44bd-b06e-95eb52df484e","name":"binary state","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"shared":false},{"base_characteristic_id":"urn:infai:ses:characteristic:72b624b5-6edc-4ec4-9ad9-fa00b39915c0","characteristic_ids":["urn:infai:ses:characteristic:72b624b5-6edc-4ec4-9ad9-fa00b39915c0"],"creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","id":"urn:infai:ses:concept:dbe4ad57-aa1d-4d24-9bee-a44a1c670d7f","name":"brightness","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"shared":false},{"base_characteristic_id":"urn:infai:ses:characteristic:7768d6da-d53f-4afb-aeb3-e0980756ec06","characteristic_ids":["urn:infai:ses:characteristic:7768d6da-d53f-4afb-aeb3-e0980756ec06"],"creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","id":"urn:infai:ses:concept:55efac99-6e5b-4ef3-8f4a-8576c1fe8614","name":"carbon dioxid","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"shared":false},{"base_characteristic_id":"urn:infai:ses:characteristic:193ff994-60a4-4ed2-90a1-21b82622cb43","characteristic_ids":["urn:infai:ses:characteristic:15af5b7b-8b17-4b18-b37e-2e91c839886b","urn:infai:ses:characteristic:193ff994-60a4-4ed2-90a1-21b82622cb43"],"creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","id":"urn:infai:ses:concept:bc7a76fc-86d0-470a-a872-00ef211904f9","name":"carbon monoxide","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"shared":false},{"base_characteristic_id":"urn:infai:ses:characteristic:0958edd7-eaae-46fe-b612-0bf4e53400af","characteristic_ids":["urn:infai:ses:characteristic:0958edd7-eaae-46fe-b612-0bf4e53400af"],"creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","id":"urn:infai:ses:concept:64106d72-d5b7-4df3-a1bf-cc369799434e","name":"cloud fraction","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"shared":false},{"base_characteristic_id":"urn:infai:ses:characteristic:5b4eea52-e8e5-4e80-9455-0382f81a1b43","characteristic_ids":["urn:infai:ses:characteristic:0fc343ce-4627-4c88-b1e0-d3ed29754af8","urn:infai:ses:characteristic:5b4eea52-e8e5-4e80-9455-0382f81a1b43","urn:infai:ses:characteristic:64928e9f-98ca-42bb-a1e5-adf2a760a2f9"],"creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","id":"urn:infai:ses:concept:8b1161d5-7878-4dd2-a36c-6f98f6b94bf8","name":"color","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"shared":false},{"base_characteristic_id":"urn:infai:ses:characteristic:5b5358c4-efd3-48a5-87b5-a8e0a34ce1ab","characteristic_ids":["urn:infai:ses:characteristic:5b5358c4-efd3-48a5-87b5-a8e0a34ce1ab"],"creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","id":"urn:infai:ses:concept:8f4a00c9-061b-41d8-901f-192b15f5a465","name":"configuration","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"shared":false},{"base_characteristic_id":"urn:infai:ses:characteristic:bc03ef2e-51d5-4034-8bec-75df78e3afee","characteristic_ids":["urn:infai:ses:characteristic:bc03ef2e-51d5-4034-8bec-75df78e3afee"],"creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","id":"urn:infai:ses:concept:85e11726-620a-4584-96a2-3a6fe4141b2d","name":"connection status","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"shared":false},{"base_characteristic_id":"urn:infai:ses:characteristic:17cfbe28-ce44-4355-bf75-15b6bf39ba1d","characteristic_ids":["urn:infai:ses:characteristic:17cfbe28-ce44-4355-bf75-15b6bf39ba1d","urn:infai:ses:characteristic:65e9a5db-348f-4888-9d95-588741d67dbf"],"creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","id":"urn:infai:ses:concept:fb85f32b-1864-456d-b99f-1540892ffd02","name":"contact","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"shared":false},{"base_characteristic_id":"urn:infai:ses:characteristic:b59c3965-c270-4053-b46a-f58b44cc980c","characteristic_ids":["urn:infai:ses:characteristic:b59c3965-c270-4053-b46a-f58b44cc980c"],"creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","id":"urn:infai:ses:concept:7ffc7509-51dc-4a62-bb42-850090996586","name":"current","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"shared":false},{"base_characteristic_id":"urn:infai:ses:characteristic:afb0ebb5-8f10-4c09-b6af-15eb291cacc1","characteristic_ids":["urn:infai:ses:characteristic:afb0ebb5-8f10-4c09-b6af-15eb291cacc1"],"creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","id":"urn:infai:ses:concept:7d28f92e-8836-42bf-bca6-342989059ddf","name":"dew point","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"shared":false},{"base_characteristic_id":"urn:infai:ses:characteristic:005ca5d3-44da-4fc8-b2b8-a88e3209e9f7","characteristic_ids":["urn:infai:ses:characteristic:005ca5d3-44da-4fc8-b2b8-a88e3209e9f7"],"creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","id":"urn:infai:ses:concept:76a1082b-7e86-4ba0-b126-7c449e581563","name":"direction","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"shared":false},{"base_characteristic_id":"urn:infai:ses:characteristic:3febed55-ba9b-43dc-8709-9c73bae3716e","characteristic_ids":["urn:infai:ses:characteristic:3febed55-ba9b-43dc-8709-9c73bae3716e"],"creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","id":"urn:infai:ses:concept:cedf4396-8e55-491d-b6f8-02ff94679b7d","name":"energy","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"shared":false},{"base_characteristic_id":"urn:infai:ses:characteristic:7aaf53b9-6cd3-4e36-b27e-9e77489f6f78","characteristic_ids":["urn:infai:ses:characteristic:7aaf53b9-6cd3-4e36-b27e-9e77489f6f78"],"creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","id":"urn:infai:ses:concept:f7fae105-b4d9-40c9-8a47-82c377899cd5","name":"fog area fraction","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"shared":false},{"base_characteristic_id":"urn:infai:ses:characteristic:0bf5bc86-dde2-465f-8a16-687ef9acba3d","characteristic_ids":["urn:infai:ses:characteristic:0bf5bc86-dde2-465f-8a16-687ef9acba3d"],"creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","id":"urn:infai:ses:concept:1c9e4d58-28e3-4e65-a0a0-48d8d519b0c9","name":"frequency","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"shared":false},{"base_characteristic_id":"urn:infai:ses:characteristic:aaf2385d-92a6-4fd0-ba4d-3e8567e733d6","characteristic_ids":["urn:infai:ses:characteristic:aaf2385d-92a6-4fd0-ba4d-3e8567e733d6"],"creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","id":"urn:infai:ses:concept:67fbf29b-fd77-4c82-b8fa-19a6000c602d","name":"humidity","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"shared":false},{"base_characteristic_id":"urn:infai:ses:characteristic:63bb46ea-64f1-4a60-aabb-67b9febdb588","characteristic_ids":["urn:infai:ses:characteristic:63bb46ea-64f1-4a60-aabb-67b9febdb588"],"creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","id":"urn:infai:ses:concept:c6d4cb5b-5ae4-489e-a664-cf02003f1d25","name":"latitude","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"shared":false},{"base_characteristic_id":"urn:infai:ses:characteristic:5caa707d-dc08-4f3b-bd9f-f08935c8dd3c","characteristic_ids":["urn:infai:ses:characteristic:5caa707d-dc08-4f3b-bd9f-f08935c8dd3c"],"creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","id":"urn:infai:ses:concept:f917e146-9192-4796-8bbe-6be14292346f","name":"lifespan","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"shared":false},{"base_characteristic_id":"urn:infai:ses:characteristic:d8a73a4d-8745-40b9-87e5-50c5d31f745a","characteristic_ids":["urn:infai:ses:characteristic:d8a73a4d-8745-40b9-87e5-50c5d31f745a"],"creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","id":"urn:infai:ses:concept:a570ecaf-2939-4524-8bd3-c5f12b1e6863","name":"longitude","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"shared":false},{"base_characteristic_id":"urn:infai:ses:characteristic:0419856b-d198-480e-9998-bf990f226844","characteristic_ids":["urn:infai:ses:characteristic:0419856b-d198-480e-9998-bf990f226844"],"creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","id":"urn:infai:ses:concept:59b02acd-aac2-4415-976e-bd4790963ce4","name":"luminiscence","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"shared":false},{"base_characteristic_id":"urn:infai:ses:characteristic:3fec34de-0832-43c7-a071-bf0a858e3292","characteristic_ids":["urn:infai:ses:characteristic:3fec34de-0832-43c7-a071-bf0a858e3292"],"creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","id":"urn:infai:ses:concept:1f60e21d-0682-483c-9b0b-0146d8a1618d","name":"nitrogen dioxide","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"shared":false},{"base_characteristic_id":"urn:infai:ses:characteristic:f17ed384-82b6-4fba-908a-d146366c9779","characteristic_ids":["urn:infai:ses:characteristic:f17ed384-82b6-4fba-908a-d146366c9779"],"creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","id":"urn:infai:ses:concept:1b926adc-8e35-4c88-bfe1-15f23e67f4f5","name":"ozone","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"shared":false},{"base_characteristic_id":"urn:infai:ses:characteristic:9673cd38-8adf-48d0-b9ff-54ab56bbdcf0","characteristic_ids":["urn:infai:ses:characteristic:de1a0d1d-85ae-45d3-844b-56c324c116cd","urn:infai:ses:characteristic:9673cd38-8adf-48d0-b9ff-54ab56bbdcf0"],"creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","id":"urn:infai:ses:concept:a63a960d-1f36-4d95-ad5b-dfb4a7fe3b5b","name":"particle amount","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"shared":false},{"base_characteristic_id":"urn:infai:ses:characteristic:b060dacf-f198-452b-87cf-740bf139e0c1","characteristic_ids":["urn:infai:ses:characteristic:b060dacf-f198-452b-87cf-740bf139e0c1"],"creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","id":"urn:infai:ses:concept:5859621b-f6b3-4430-8b4f-b25a53bae708","name":"pollen level","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"shared":false},{"base_characteristic_id":"urn:infai:ses:characteristic:107c29d3-f863-47e1-a1e0-7853a8016213","characteristic_ids":["urn:infai:ses:characteristic:107c29d3-f863-47e1-a1e0-7853a8016213"],"creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","id":"urn:infai:ses:concept:b12151d3-d187-4839-82ac-8cc18fae3dd9","name":"position","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"shared":false},{"base_characteristic_id":"urn:infai:ses:characteristic:a3931f28-664d-452f-9370-a641fadd21db","characteristic_ids":["urn:infai:ses:characteristic:a3931f28-664d-452f-9370-a641fadd21db"],"creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","id":"urn:infai:ses:concept:46de11b9-26ff-4cce-b945-e93b84f04fe6","name":"power","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"shared":false},{"base_characteristic_id":"urn:infai:ses:characteristic:e6874605-80cf-418a-9319-da3b27411af1","characteristic_ids":["urn:infai:ses:characteristic:e6874605-80cf-418a-9319-da3b27411af1"],"creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","id":"urn:infai:ses:concept:3e40dfa7-805d-4732-af15-5207c0a59c47","name":"power factor","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"shared":false},{"base_characteristic_id":"urn:infai:ses:characteristic:7424a147-fb0e-4e71-a666-6c4997928e61","characteristic_ids":["urn:infai:ses:characteristic:7424a147-fb0e-4e71-a666-6c4997928e61"],"creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","id":"urn:infai:ses:concept:0c1eaf2b-ea3d-4ca9-8446-02c55dd5d3f4","name":"precipitation","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"shared":false},{"base_characteristic_id":"urn:infai:ses:characteristic:ec4ab96e-bfc3-4a80-96c4-b9e12634611f","characteristic_ids":["urn:infai:ses:characteristic:ec4ab96e-bfc3-4a80-96c4-b9e12634611f"],"creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","id":"urn:infai:ses:concept:40e8e501-85a0-4406-9aa9-7d5fd6985d1f","name":"precipitation probability","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"shared":false},{"base_characteristic_id":"urn:infai:ses:characteristic:eb33cf65-b0a2-413d-891d-cada05be01ed","characteristic_ids":["urn:infai:ses:characteristic:2871c1c4-1c8c-4169-81f6-fb5313df08b1","urn:infai:ses:characteristic:eb33cf65-b0a2-413d-891d-cada05be01ed"],"creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","id":"urn:infai:ses:concept:9924ea1a-d96a-4882-88a2-165493956ec3","name":"pressure","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"shared":false},{"base_characteristic_id":"urn:infai:ses:characteristic:2ed0e8ca-e396-4c42-9446-6e6981e87845","characteristic_ids":["urn:infai:ses:characteristic:2ed0e8ca-e396-4c42-9446-6e6981e87845"],"creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","id":"urn:infai:ses:concept:2769a588-97e7-4bc3-94f4-b97bef591e95","name":"signal strength","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"shared":false},{"base_characteristic_id":"urn:infai:ses:characteristic:04bce19c-839d-45de-8cfb-bd470715f4cd","characteristic_ids":["urn:infai:ses:characteristic:04bce19c-839d-45de-8cfb-bd470715f4cd"],"creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","id":"urn:infai:ses:concept:95f01a66-f7ea-4c49-9304-29cf065a0c55","name":"speed","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"shared":false},{"base_characteristic_id":"urn:infai:ses:characteristic:2b4040c4-63b2-497f-8673-96378dcbc4cf","characteristic_ids":["urn:infai:ses:characteristic:ca44269e-ab5a-4ed7-8acb-ce80a6a592a5","urn:infai:ses:characteristic:2b4040c4-63b2-497f-8673-96378dcbc4cf","urn:infai:ses:characteristic:ac7ad69f-b82e-454a-8b44-5c0440cba787"],"creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","id":"urn:infai:ses:concept:a43e8989-7d4b-4d9c-97d7-766799106499","name":"speed level","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"shared":false},{"base_characteristic_id":"urn:infai:ses:characteristic:2ce5f9c7-3c44-438b-91e6-3e9901b8a3db","characteristic_ids":["urn:infai:ses:characteristic:2ce5f9c7-3c44-438b-91e6-3e9901b8a3db"],"creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","id":"urn:infai:ses:concept:c41d2566-aad6-458a-8f21-fee504225b17","name":"sulfur dioxide","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"shared":false},{"base_characteristic_id":"urn:infai:ses:characteristic:5ba31623-0ccb-4488-bfb7-f73b50e03b5a","characteristic_ids":["urn:infai:ses:characteristic:75b2d113-1d03-4ef8-977a-8dbcbb31a683","urn:infai:ses:characteristic:5ba31623-0ccb-4488-bfb7-f73b50e03b5a"],"creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","id":"urn:infai:ses:concept:0bc81398-3ed6-4e2b-a6c4-b754583aac37","name":"temperature","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"shared":false},{"base_characteristic_id":"urn:infai:ses:characteristic:02c7ea52-fc76-470e-b2b2-401653412de6","characteristic_ids":["urn:infai:ses:characteristic:02c7ea52-fc76-470e-b2b2-401653412de6"],"creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","id":"urn:infai:ses:concept:6625cc35-a8d0-4d9d-91ef-3c0cb71f0285","name":"thunder probability","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"shared":false},{"base_characteristic_id":"urn:infai:ses:characteristic:9e1024da-3b60-4531-9f29-464addccb13c","characteristic_ids":["urn:infai:ses:characteristic:9e1024da-3b60-4531-9f29-464addccb13c","urn:infai:ses:characteristic:b36eee5d-52f0-4476-a6f7-6dd03b24e0f8","urn:infai:ses:characteristic:182e6bb4-8622-453a-9a4d-bfa9c70d6c9c"],"creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","id":"urn:infai:ses:concept:58d23b4f-c315-4d7d-9230-fa989d55640a","name":"time","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"shared":false},{"base_characteristic_id":"urn:infai:ses:characteristic:6bc41b45-a9f3-4d87-9c51-dd3e11257800","characteristic_ids":["urn:infai:ses:characteristic:6bc41b45-a9f3-4d87-9c51-dd3e11257800","urn:infai:ses:characteristic:8b9411f3-bf56-4e57-8fd4-728bdcd451cd","urn:infai:ses:characteristic:c7dfcb86-2733-4917-a5ca-0a150a458eed","urn:infai:ses:characteristic:64691f8d-4909-470f-a1fa-e977ebe28684"],"creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","id":"urn:infai:ses:concept:9286d398-f49d-494f-8157-bea17947b1fa","name":"timestamp","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"shared":false},{"base_characteristic_id":"urn:infai:ses:characteristic:0a61343d-c0d1-4af8-9329-3829c30ba59f","characteristic_ids":["urn:infai:ses:characteristic:0a61343d-c0d1-4af8-9329-3829c30ba59f"],"creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","id":"urn:infai:ses:concept:a7cd1e62-72fc-448c-8492-6a8f0307bbdc","name":"ultraviolet","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"shared":false},{"base_characteristic_id":"urn:infai:ses:characteristic:a924efd4-9397-474b-ba01-6de49e0283e3","characteristic_ids":["urn:infai:ses:characteristic:a924efd4-9397-474b-ba01-6de49e0283e3"],"creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","id":"urn:infai:ses:concept:397de0f8-a79f-484f-aa13-0042e48e5648","name":"voltage","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"shared":false},{"base_characteristic_id":"urn:infai:ses:characteristic:fbfea6c7-3392-4ec2-9ec5-0d0ef6362b9b","characteristic_ids":["urn:infai:ses:characteristic:fbfea6c7-3392-4ec2-9ec5-0d0ef6362b9b"],"creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","id":"urn:infai:ses:concept:81c003ac-ed79-4b3d-8d49-6d95b61b2863","name":"volume","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"shared":false},{"base_characteristic_id":"urn:infai:ses:characteristic:3f0294a9-a53e-4650-b3d1-6a6482cf3b70","characteristic_ids":["urn:infai:ses:characteristic:3f0294a9-a53e-4650-b3d1-6a6482cf3b70"],"creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","id":"urn:infai:ses:concept:cb2608c2-b671-46be-8bad-1539e6c6e2af","name":"volume flow","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"shared":false},{"base_characteristic_id":"urn:infai:ses:characteristic:b6070dc9-0327-47b4-ba27-3287325940e5","characteristic_ids":["urn:infai:ses:characteristic:b6070dc9-0327-47b4-ba27-3287325940e5"],"creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","id":"urn:infai:ses:concept:0ec0adad-6873-4baa-931f-6544aa1dd210","name":"water level","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"shared":false}]`
const functionsListStr = `[{"mock-placeholder-function": "mock-placeholder-concept"},{"concept_id":"urn:infai:ses:concept:ebfeabb3-50f0-44bd-b06e-95eb52df484e","creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","description":"Checks if an alarm is triggered","display_name":null,"id":"urn:infai:ses:measuring-function:e59bbee4-32a4-456b-99b9-dea8c3545dfc","name":"getAlarmFunction","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction","shared":false},{"concept_id":"urn:infai:ses:concept:f917e146-9192-4796-8bbe-6be14292346f","creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","description":"Returns the current battery level","display_name":null,"id":"urn:infai:ses:measuring-function:00549f18-88b5-44c7-adb1-f558e8d53d1d","name":"getBatteryStateFunction","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction","shared":false},{"concept_id":"urn:infai:ses:concept:08cde541-54dc-4d92-b2d9-b07304bd659a","creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","description":"Measure benzene","display_name":null,"id":"urn:infai:ses:measuring-function:3395f090-9ac5-45fd-b064-81ccea7eb236","name":"getBenzeneFunction","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction","shared":false},{"concept_id":"urn:infai:ses:concept:dbe4ad57-aa1d-4d24-9bee-a44a1c670d7f","creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","description":"Read a brightness level","display_name":null,"id":"urn:infai:ses:measuring-function:c51a6ea5-90c3-4223-9052-6fe4136386cd","name":"getBrightnessFunction","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction","shared":false},{"concept_id":"urn:infai:ses:concept:55efac99-6e5b-4ef3-8f4a-8576c1fe8614","creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","description":"Collect a carbon dioxid measurement","display_name":null,"id":"urn:infai:ses:measuring-function:ad89a222-75a2-47e1-9fe3-969094320881","name":"getCarbonDioxidFunction","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction","shared":false},{"concept_id":"urn:infai:ses:concept:bc7a76fc-86d0-470a-a872-00ef211904f9","creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","description":"Get a carbon monoxide measurement","display_name":null,"id":"urn:infai:ses:measuring-function:76423a12-eac2-4f6c-8f79-0216360e2e61","name":"getCarbonMonoxideFunction","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction","shared":false},{"concept_id":"urn:infai:ses:concept:64106d72-d5b7-4df3-a1bf-cc369799434e","creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","description":"Read how much of the sky is covered in clouds","display_name":null,"id":"urn:infai:ses:measuring-function:87709727-53c7-417c-8072-6a50cc3bbf1d","name":"getCloudFractionFunction","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction","shared":false},{"concept_id":"urn:infai:ses:concept:8b1161d5-7878-4dd2-a36c-6f98f6b94bf8","creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","description":"Read a color","display_name":null,"id":"urn:infai:ses:measuring-function:bdb6a7c8-4a3d-4fe0-bab3-ce02e09b5869","name":"getColorFunction","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction","shared":false},{"concept_id":"urn:infai:ses:concept:8f4a00c9-061b-41d8-901f-192b15f5a465","creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","description":"","display_name":null,"id":"urn:infai:ses:measuring-function:507a3fc2-5ebe-47fa-84a9-1df1e5314f74","name":"getConfigFunction","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction","shared":false},{"concept_id":"urn:infai:ses:concept:85e11726-620a-4584-96a2-3a6fe4141b2d","creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","description":"Check if something is connected","display_name":null,"id":"urn:infai:ses:measuring-function:b8791b17-cf01-467f-87cf-da2271fffb6d","name":"getConnectionStatusFunction","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction","shared":false},{"concept_id":"urn:infai:ses:concept:fb85f32b-1864-456d-b99f-1540892ffd02","creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","description":"","display_name":null,"id":"urn:infai:ses:measuring-function:90f05fa4-b765-4998-a09c-448f83a5e8c5","name":"getContactFunction","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction","shared":false},{"concept_id":"urn:infai:ses:concept:7ffc7509-51dc-4a62-bb42-850090996586","creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","description":"Return the current consumption","display_name":null,"id":"urn:infai:ses:measuring-function:6a04430c-0b40-4d26-8851-2401e13edfd6","name":"getCurrentConsumptionFunction","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction","shared":false},{"concept_id":"urn:infai:ses:concept:7ffc7509-51dc-4a62-bb42-850090996586","creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","description":"Return a current measurement","display_name":null,"id":"urn:infai:ses:measuring-function:28326e5a-69fa-439b-baa9-0e63c453186f","name":"getCurrentFunction","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction","shared":false},{"concept_id":"urn:infai:ses:concept:7d28f92e-8836-42bf-bca6-342989059ddf","creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","description":"Get a dew point measurement","display_name":null,"id":"urn:infai:ses:measuring-function:48c6f5fe-2ad8-49c2-b63d-4e4b13b7af67","name":"getDewPointFunction","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction","shared":false},{"concept_id":"urn:infai:ses:concept:76a1082b-7e86-4ba0-b126-7c449e581563","creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","description":"Direction of something","display_name":null,"id":"urn:infai:ses:measuring-function:2d9bb487-7430-4830-bf09-955e8824c871","name":"getDirectionFunction","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction","shared":false},{"concept_id":"urn:infai:ses:concept:cedf4396-8e55-491d-b6f8-02ff94679b7d","creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","description":"Return the energy consumption","display_name":null,"id":"urn:infai:ses:measuring-function:57dfd369-92db-462c-aca4-a767b52c972e","name":"getEnergyConsumptionFunction","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction","shared":false},{"concept_id":"urn:infai:ses:concept:cedf4396-8e55-491d-b6f8-02ff94679b7d","creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","description":"Return the energy production","display_name":null,"id":"urn:infai:ses:measuring-function:826e5a04-71cc-4935-9fd4-92c930dc06bb","name":"getEnergyProductionFunction","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction","shared":false},{"concept_id":"urn:infai:ses:concept:f917e146-9192-4796-8bbe-6be14292346f","creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","description":"Return how much a filter has been used up","display_name":null,"id":"urn:infai:ses:measuring-function:8b994547-dbf4-4e3c-a356-e676a744c601","name":"getFilterStateFunction","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction","shared":false},{"concept_id":"urn:infai:ses:concept:f7fae105-b4d9-40c9-8a47-82c377899cd5","creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","description":"Check how much of the horizontal area is covered in fog","display_name":null,"id":"urn:infai:ses:measuring-function:68998f82-33e8-4ca7-8e3a-a3bba077f0fa","name":"getFogAreaFractionFunction","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction","shared":false},{"concept_id":"urn:infai:ses:concept:1c9e4d58-28e3-4e65-a0a0-48d8d519b0c9","creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","description":"Return a frequency measurement","display_name":null,"id":"urn:infai:ses:measuring-function:9f180351-e83b-437f-b13f-1ad030703036","name":"getFrequencyFunction","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction","shared":false},{"concept_id":"urn:infai:ses:concept:67fbf29b-fd77-4c82-b8fa-19a6000c602d","creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","description":"Read a humidity measurement","display_name":null,"id":"urn:infai:ses:measuring-function:53bb96e5-1fb1-4409-89a6-ae6c7332eae4","name":"getHumidityFunction","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction","shared":false},{"concept_id":"urn:infai:ses:concept:c6d4cb5b-5ae4-489e-a664-cf02003f1d25","creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","description":"Get a latitude measurement","display_name":null,"id":"urn:infai:ses:measuring-function:f9f2b270-8c74-4289-b822-649e96b39206","name":"getLatitudeFunction","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction","shared":false},{"concept_id":"urn:infai:ses:concept:a570ecaf-2939-4524-8bd3-c5f12b1e6863","creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","description":"Get a longitude measurement","display_name":null,"id":"urn:infai:ses:measuring-function:26508f1a-ad3c-4850-8731-dbf966b96335","name":"getLongitudeFunction","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction","shared":false},{"concept_id":"urn:infai:ses:concept:59b02acd-aac2-4415-976e-bd4790963ce4","creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","description":"Read a luminiscence measurement","display_name":null,"id":"urn:infai:ses:measuring-function:d7bd8d39-92f7-49da-9dd9-5445c45f2e27","name":"getLuminiscenceFunction","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction","shared":false},{"concept_id":"urn:infai:ses:concept:ebfeabb3-50f0-44bd-b06e-95eb52df484e","creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","description":"Check the state of something being monitored","display_name":null,"id":"urn:infai:ses:measuring-function:231da98e-f21d-47bc-9f56-6606397221a2","name":"getMonitoringStateFunction","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction","shared":false},{"concept_id":"urn:infai:ses:concept:ebfeabb3-50f0-44bd-b06e-95eb52df484e","creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","description":"Check if something is moving or has been moved","display_name":null,"id":"urn:infai:ses:measuring-function:64dfd53a-f288-40f7-a3ef-a6a7f7c8a7c3","name":"getMotionStateFunction","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction","shared":false},{"concept_id":"urn:infai:ses:concept:1f60e21d-0682-483c-9b0b-0146d8a1618d","creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","description":"Get a nitrogen dioxide measurement","display_name":null,"id":"urn:infai:ses:measuring-function:43525897-6a9c-4501-86ad-768dcbcb7c57","name":"getNitrogenDioxideFunction","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction","shared":false},{"concept_id":"urn:infai:ses:concept:ebfeabb3-50f0-44bd-b06e-95eb52df484e","creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","description":"Check if something is on or off","display_name":null,"id":"urn:infai:ses:measuring-function:20d3c1d3-77d7-4181-a9f3-b487add58cd0","name":"getOnOffStateFunction","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction","shared":false},{"concept_id":"urn:infai:ses:concept:ebfeabb3-50f0-44bd-b06e-95eb52df484e","creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","description":"Check if something is oscillating","display_name":null,"id":"urn:infai:ses:measuring-function:61b1f2db-3cff-49d8-b22d-424073269f35","name":"getOscillationStateFunction","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction","shared":false},{"concept_id":"urn:infai:ses:concept:1b926adc-8e35-4c88-bfe1-15f23e67f4f5","creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","description":"Get an ozone measurement","display_name":null,"id":"urn:infai:ses:measuring-function:56d6b5c6-9b95-4181-ad8f-d0b109cedb18","name":"getOzoneFunction","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction","shared":false},{"concept_id":"urn:infai:ses:concept:a63a960d-1f36-4d95-ad5b-dfb4a7fe3b5b","creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","description":"Read a PM1 measurement","display_name":null,"id":"urn:infai:ses:measuring-function:0e19d094-70c6-402c-8523-3aaff2ce6dd9","name":"getParticleAmountPM1","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction","shared":false},{"concept_id":"urn:infai:ses:concept:a63a960d-1f36-4d95-ad5b-dfb4a7fe3b5b","creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","description":"Read a PM10 measurement","display_name":null,"id":"urn:infai:ses:measuring-function:f2c1a22f-a49e-4549-9833-62f0994afec0","name":"getParticleAmountPM10","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction","shared":false},{"concept_id":"urn:infai:ses:concept:a63a960d-1f36-4d95-ad5b-dfb4a7fe3b5b","creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","description":"Measure arsenic in particulate matter","display_name":null,"id":"urn:infai:ses:controlling-function:8335e851-05e2-479d-8c17-1cf8174e90e9","name":"getParticleAmountPM10AS","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"rdf_type":"https://senergy.infai.org/ontology/ControllingFunction","shared":false},{"concept_id":"urn:infai:ses:concept:a63a960d-1f36-4d95-ad5b-dfb4a7fe3b5b","creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","description":"Measure benzo(a)pyrene in particulate matter","display_name":null,"id":"urn:infai:ses:measuring-function:29cb52fe-e2d9-446b-8be1-f5447ee7967f","name":"getParticleAmountPM10BAP","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction","shared":false},{"concept_id":"urn:infai:ses:concept:a63a960d-1f36-4d95-ad5b-dfb4a7fe3b5b","creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","description":"Measure cadmium in particulate matter","display_name":null,"id":"urn:infai:ses:measuring-function:cd475584-1b23-4879-ab72-e879f90ade91","name":"getParticleAmountPM10CD","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction","shared":false},{"concept_id":"urn:infai:ses:concept:a63a960d-1f36-4d95-ad5b-dfb4a7fe3b5b","creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","description":"Measure nickel in particulate matter","display_name":null,"id":"urn:infai:ses:measuring-function:cee9e218-2bdd-4bfd-bfab-2bcd6dd75bb7","name":"getParticleAmountPM10NI","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction","shared":false},{"concept_id":"urn:infai:ses:concept:a63a960d-1f36-4d95-ad5b-dfb4a7fe3b5b","creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","description":"Measure lead in particulate matter","display_name":null,"id":"urn:infai:ses:measuring-function:dd0c3b2f-3a04-430c-8b39-4e8b927cf20e","name":"getParticleAmountPM10PB","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction","shared":false},{"concept_id":"urn:infai:ses:concept:a63a960d-1f36-4d95-ad5b-dfb4a7fe3b5b","creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","description":"Read a PM2.5 measurement","display_name":null,"id":"urn:infai:ses:measuring-function:22bbaa27-595c-4f53-bdda-8b9614ecdf76","name":"getParticleAmountPM25","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction","shared":false},{"concept_id":"urn:infai:ses:concept:5859621b-f6b3-4430-8b4f-b25a53bae708","creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","description":"Get how much pollen is in the air","display_name":null,"id":"urn:infai:ses:measuring-function:3e766b2d-0ab5-421d-a4f6-de4a61ccdd62","name":"getPollenLevel","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction","shared":false},{"concept_id":"urn:infai:ses:concept:b12151d3-d187-4839-82ac-8cc18fae3dd9","creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","description":"","display_name":null,"id":"urn:infai:ses:measuring-function:97aa3103-995a-4bd2-95f0-ab8f34e9ae36","name":"getPositionFunction","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction","shared":false},{"concept_id":"urn:infai:ses:concept:46de11b9-26ff-4cce-b945-e93b84f04fe6","creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","description":"Get the power consumption","display_name":null,"id":"urn:infai:ses:measuring-function:1c7c90fb-73b6-4690-aac2-72e9735e68d0","name":"getPowerConsumptionFunction","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction","shared":false},{"concept_id":"urn:infai:ses:concept:3e40dfa7-805d-4732-af15-5207c0a59c47","creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","description":"Read a power factor measurement","display_name":null,"id":"urn:infai:ses:measuring-function:4f3cbfb4-ccb0-4233-a7c3-61760a069c02","name":"getPowerFactorFunction","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction","shared":false},{"concept_id":"urn:infai:ses:concept:0c1eaf2b-ea3d-4ca9-8446-02c55dd5d3f4","creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","description":"","display_name":null,"id":"urn:infai:ses:measuring-function:61ebe348-6abc-4e14-911c-0908bd7ba1f8","name":"getPrecipitationFunction","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction","shared":false},{"concept_id":"urn:infai:ses:concept:40e8e501-85a0-4406-9aa9-7d5fd6985d1f","creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","description":"Read how likely precipitation is","display_name":null,"id":"urn:infai:ses:measuring-function:065c7aa6-5ce3-415b-92a6-4c3a84036488","name":"getPrecipitationProbability","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction","shared":false},{"concept_id":"urn:infai:ses:concept:9924ea1a-d96a-4882-88a2-165493956ec3","creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","description":"Read a pressure measurement","display_name":null,"id":"urn:infai:ses:measuring-function:2acad12c-8983-4655-9a2d-e2359a706101","name":"getPressureFunction","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction","shared":false},{"concept_id":"","creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","description":"Receive a raw WMBus message","display_name":null,"id":"urn:infai:ses:measuring-function:f2dfef5e-757d-4a58-bcec-dafd74e83b2a","name":"getRawWMBusFunction","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction","shared":false},{"concept_id":"urn:infai:ses:concept:2769a588-97e7-4bc3-94f4-b97bef591e95","creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","description":"Read a signal strength measurement","display_name":null,"id":"urn:infai:ses:measuring-function:9377c366-2896-486e-9b28-cff237509555","name":"getSignalStrengthFunction","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction","shared":false},{"concept_id":"urn:infai:ses:concept:95f01a66-f7ea-4c49-9304-29cf065a0c55","creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","description":"Read the speed of something","display_name":null,"id":"urn:infai:ses:measuring-function:f3865739-6de2-41ce-83f7-b878d7c63bcc","name":"getSpeedFunction","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction","shared":false},{"concept_id":"urn:infai:ses:concept:a43e8989-7d4b-4d9c-97d7-766799106499","creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","description":"Read the speed of something","display_name":null,"id":"urn:infai:ses:measuring-function:f6066d39-ed16-4c69-82aa-e18bcf2be2a7","name":"getSpeedLevelFunction","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction","shared":false},{"concept_id":"urn:infai:ses:concept:c41d2566-aad6-458a-8f21-fee504225b17","creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","description":"Get a sulfur dioxide measurement","display_name":null,"id":"urn:infai:ses:measuring-function:cfefc207-44f9-4a25-b40b-162a96846599","name":"getSulfurDioxideFunction","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction","shared":false},{"concept_id":"urn:infai:ses:concept:ebfeabb3-50f0-44bd-b06e-95eb52df484e","creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","description":"Check if the device has been tampered","display_name":null,"id":"urn:infai:ses:measuring-function:531efa86-f476-4c9f-8181-8753b7fb938c","name":"getTamperStateFunction","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction","shared":false},{"concept_id":"urn:infai:ses:concept:0bc81398-3ed6-4e2b-a6c4-b754583aac37","creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","description":"Get a temperature measurement","display_name":"Temperature","id":"urn:infai:ses:measuring-function:f2769eb9-b6ad-4f7e-bd28-e4ea043d2f8b","name":"getTemperatureFunction","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"rdf_type":"","shared":false},{"concept_id":"urn:infai:ses:concept:6625cc35-a8d0-4d9d-91ef-3c0cb71f0285","creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","description":"Read how likely thunder is","display_name":null,"id":"urn:infai:ses:measuring-function:642558a6-4ba2-4b7e-9bda-1afba86a5c40","name":"getThunderProbabilityFunction","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction","shared":false},{"concept_id":"urn:infai:ses:concept:9286d398-f49d-494f-8157-bea17947b1fa","creator":"dd69ea0d-f553-4336-80f3-7f4567f85c7b","description":"Get a timestamp reading","display_name":null,"id":"urn:infai:ses:measuring-function:3b4e0766-0d67-4658-b249-295902cd3290","name":"getTimestampFunction","permission_holders":{"admin_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"execute_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"read_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"],"write_users":["dd69ea0d-f553-4336-80f3-7f4567f85c7b"]},"permissions":{"a":true,"r":true,"w":true,"x":true},"rdf_type":"https://senergy.infai.org/ontology/MeasuringFunction","shared":false}]`

const characteristicsPathMapStr = `
{
    "/characteristics/urn:infai:ses:characteristic:005ca5d3-44da-4fc8-b2b8-a88e3209e9f7": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:005ca5d3-44da-4fc8-b2b8-a88e3209e9f7",
        "max_value": 360,
        "name": "degrees",
        "sub_characteristics": null,
        "type": "https://schema.org/Float"
    },
    "/characteristics/urn:infai:ses:characteristic:02c7ea52-fc76-470e-b2b2-401653412de6": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:02c7ea52-fc76-470e-b2b2-401653412de6",
        "max_value": 100,
        "min_value": 0,
        "name": "thunder probability percentage",
        "sub_characteristics": null,
        "type": "https://schema.org/Float"
    },
    "/characteristics/urn:infai:ses:characteristic:0419856b-d198-480e-9998-bf990f226844": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:0419856b-d198-480e-9998-bf990f226844",
        "min_value": 0,
        "name": "lux",
        "sub_characteristics": null,
        "type": "https://schema.org/Float"
    },
    "/characteristics/urn:infai:ses:characteristic:04bce19c-839d-45de-8cfb-bd470715f4cd": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:04bce19c-839d-45de-8cfb-bd470715f4cd",
        "name": "meters per second",
        "sub_characteristics": null,
        "type": "https://schema.org/Float"
    },
    "/characteristics/urn:infai:ses:characteristic:0958edd7-eaae-46fe-b612-0bf4e53400af": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:0958edd7-eaae-46fe-b612-0bf4e53400af",
        "max_value": 100,
        "min_value": 0,
        "name": "cloud fraction percentage",
        "sub_characteristics": null,
        "type": "https://schema.org/Float"
    },
    "/characteristics/urn:infai:ses:characteristic:0a61343d-c0d1-4af8-9329-3829c30ba59f": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:0a61343d-c0d1-4af8-9329-3829c30ba59f",
        "min_value": 0,
        "name": "uv index",
        "sub_characteristics": null,
        "type": "https://schema.org/Integer"
    },
    "/characteristics/urn:infai:ses:characteristic:0bf5bc86-dde2-465f-8a16-687ef9acba3d": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:0bf5bc86-dde2-465f-8a16-687ef9acba3d",
        "name": "hertz",
        "sub_characteristics": null,
        "type": "https://schema.org/Float"
    },
    "/characteristics/urn:infai:ses:characteristic:0fc343ce-4627-4c88-b1e0-d3ed29754af8": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:0fc343ce-4627-4c88-b1e0-d3ed29754af8",
        "name": "hex",
        "sub_characteristics": null,
        "type": "https://schema.org/Text"
    },
    "/characteristics/urn:infai:ses:characteristic:107c29d3-f863-47e1-a1e0-7853a8016213": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:107c29d3-f863-47e1-a1e0-7853a8016213",
        "max_value": 100,
        "min_value": 0,
        "name": "position percentage",
        "sub_characteristics": null,
        "type": "https://schema.org/Integer"
    },
    "/characteristics/urn:infai:ses:characteristic:17cfbe28-ce44-4355-bf75-15b6bf39ba1d": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:17cfbe28-ce44-4355-bf75-15b6bf39ba1d",
        "name": "contact state",
        "sub_characteristics": null,
        "type": "https://schema.org/Boolean",
        "value": false
    },
    "/characteristics/urn:infai:ses:characteristic:182e6bb4-8622-453a-9a4d-bfa9c70d6c9c": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:182e6bb4-8622-453a-9a4d-bfa9c70d6c9c",
        "min_value": 0,
        "name": "millisecond",
        "sub_characteristics": null,
        "type": "https://schema.org/Float"
    },
    "/characteristics/urn:infai:ses:characteristic:193ff994-60a4-4ed2-90a1-21b82622cb43": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:193ff994-60a4-4ed2-90a1-21b82622cb43",
        "name": "carbon monoxide µg/m³",
        "sub_characteristics": null,
        "type": "https://schema.org/Float"
    },
    "/characteristics/urn:infai:ses:characteristic:2871c1c4-1c8c-4169-81f6-fb5313df08b1": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:2871c1c4-1c8c-4169-81f6-fb5313df08b1",
        "name": "bar",
        "sub_characteristics": null,
        "type": "https://schema.org/Float"
    },
    "/characteristics/urn:infai:ses:characteristic:2b4040c4-63b2-497f-8673-96378dcbc4cf": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:2b4040c4-63b2-497f-8673-96378dcbc4cf",
        "name": "revolutions per minute",
        "sub_characteristics": null,
        "type": "https://schema.org/Integer"
    },
    "/characteristics/urn:infai:ses:characteristic:2ce5f9c7-3c44-438b-91e6-3e9901b8a3db": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:2ce5f9c7-3c44-438b-91e6-3e9901b8a3db",
        "name": "sulfur dioxide µg/m³",
        "sub_characteristics": null,
        "type": "https://schema.org/Float"
    },
    "/characteristics/urn:infai:ses:characteristic:2ed0e8ca-e396-4c42-9446-6e6981e87845": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:2ed0e8ca-e396-4c42-9446-6e6981e87845",
        "max_value": 100,
        "min_value": 0,
        "name": "signal strength percentage",
        "sub_characteristics": null,
        "type": "https://schema.org/Integer"
    },
    "/characteristics/urn:infai:ses:characteristic:3f0294a9-a53e-4650-b3d1-6a6482cf3b70": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:3f0294a9-a53e-4650-b3d1-6a6482cf3b70",
        "name": "cubic meter per hour",
        "sub_characteristics": null,
        "type": "https://schema.org/Float"
    },
    "/characteristics/urn:infai:ses:characteristic:3febed55-ba9b-43dc-8709-9c73bae3716e": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:3febed55-ba9b-43dc-8709-9c73bae3716e",
        "name": "kilowatt hour",
        "sub_characteristics": null,
        "type": "https://schema.org/Float"
    },
    "/characteristics/urn:infai:ses:characteristic:3fec34de-0832-43c7-a071-bf0a858e3292": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:3fec34de-0832-43c7-a071-bf0a858e3292",
        "name": "nitrogen dioxide µg/m³",
        "sub_characteristics": null,
        "type": "https://schema.org/Float"
    },
    "/characteristics/urn:infai:ses:characteristic:5b4eea52-e8e5-4e80-9455-0382f81a1b43": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:5b4eea52-e8e5-4e80-9455-0382f81a1b43",
        "name": "RGB",
        "sub_characteristics": [
            {
                "display_unit": "",
                "id": "urn:infai:ses:characteristic:590af9ef-3a5e-4edb-abab-177cb1320b17",
                "name": "b",
                "sub_characteristics": null,
                "type": "https://schema.org/Integer"
            },
            {
                "display_unit": "",
                "id": "urn:infai:ses:characteristic:5ef27837-4aca-43ad-b8f6-4d95cf9ed99e",
                "name": "g",
                "sub_characteristics": null,
                "type": "https://schema.org/Integer"
            },
            {
                "display_unit": "",
                "id": "urn:infai:ses:characteristic:dfe6be4a-650c-4411-8d87-062916b48951",
                "name": "r",
                "sub_characteristics": null,
                "type": "https://schema.org/Integer"
            }
        ],
        "type": "https://schema.org/StructuredValue"
    },
    "/characteristics/urn:infai:ses:characteristic:5b5358c4-efd3-48a5-87b5-a8e0a34ce1ab": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:5b5358c4-efd3-48a5-87b5-a8e0a34ce1ab",
        "name": "configuration string",
        "sub_characteristics": null,
        "type": "https://schema.org/Text"
    },
    "/characteristics/urn:infai:ses:characteristic:5ba31623-0ccb-4488-bfb7-f73b50e03b5a": {
        "display_unit": "°C",
        "id": "urn:infai:ses:characteristic:5ba31623-0ccb-4488-bfb7-f73b50e03b5a",
        "min_value": -273.15,
        "name": "celsius",
        "sub_characteristics": [],
        "type": "https://schema.org/Float"
    },
    "/characteristics/urn:infai:ses:characteristic:5caa707d-dc08-4f3b-bd9f-f08935c8dd3c": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:5caa707d-dc08-4f3b-bd9f-f08935c8dd3c",
        "max_value": 100,
        "min_value": 0,
        "name": "lifespan percentage",
        "sub_characteristics": null,
        "type": "https://schema.org/Float"
    },
    "/characteristics/urn:infai:ses:characteristic:63bb46ea-64f1-4a60-aabb-67b9febdb588": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:63bb46ea-64f1-4a60-aabb-67b9febdb588",
        "name": "latitude EPSG:4326",
        "sub_characteristics": null,
        "type": "https://schema.org/Float"
    },
    "/characteristics/urn:infai:ses:characteristic:64691f8d-4909-470f-a1fa-e977ebe28684": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:64691f8d-4909-470f-a1fa-e977ebe28684",
        "name": "unix_format_milliseconds",
        "sub_characteristics": null,
        "type": "https://schema.org/Integer"
    },
    "/characteristics/urn:infai:ses:characteristic:64928e9f-98ca-42bb-a1e5-adf2a760a2f9": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:64928e9f-98ca-42bb-a1e5-adf2a760a2f9",
        "name": "HSB",
        "sub_characteristics": [
            {
                "display_unit": "",
                "id": "urn:infai:ses:characteristic:d840607c-c8f9-45d6-b9bd-2c2d444e2899",
                "max_value": 100,
                "min_value": 0,
                "name": "b",
                "sub_characteristics": null,
                "type": "https://schema.org/Integer"
            },
            {
                "display_unit": "",
                "id": "urn:infai:ses:characteristic:a66dc568-c0e0-420f-b513-18e8df405538",
                "max_value": 100,
                "min_value": 0,
                "name": "s",
                "sub_characteristics": null,
                "type": "https://schema.org/Integer"
            },
            {
                "display_unit": "",
                "id": "urn:infai:ses:characteristic:6ec70e99-8c6a-4909-8d5a-7cc12af76b9a",
                "max_value": 360,
                "min_value": 0,
                "name": "h",
                "sub_characteristics": null,
                "type": "https://schema.org/Integer"
            }
        ],
        "type": "https://schema.org/StructuredValue"
    },
    "/characteristics/urn:infai:ses:characteristic:65e9a5db-348f-4888-9d95-588741d67dbf": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:65e9a5db-348f-4888-9d95-588741d67dbf",
        "name": "zwave contact state",
        "sub_characteristics": null,
        "type": "https://schema.org/Text",
        "value": ""
    },
    "/characteristics/urn:infai:ses:characteristic:6bc41b45-a9f3-4d87-9c51-dd3e11257800": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:6bc41b45-a9f3-4d87-9c51-dd3e11257800",
        "name": "iso_format",
        "sub_characteristics": null,
        "type": "https://schema.org/Text"
    },
    "/characteristics/urn:infai:ses:characteristic:72b624b5-6edc-4ec4-9ad9-fa00b39915c0": {
        "display_unit": "%",
        "id": "urn:infai:ses:characteristic:72b624b5-6edc-4ec4-9ad9-fa00b39915c0",
        "max_value": 100,
        "name": "brightness percentage",
        "sub_characteristics": [],
        "type": "https://schema.org/Integer"
    },
    "/characteristics/urn:infai:ses:characteristic:7424a147-fb0e-4e71-a666-6c4997928e61": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:7424a147-fb0e-4e71-a666-6c4997928e61",
        "name": "precipitation mm",
        "sub_characteristics": null,
        "type": "https://schema.org/Float"
    },
    "/characteristics/urn:infai:ses:characteristic:75b2d113-1d03-4ef8-977a-8dbcbb31a683": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:75b2d113-1d03-4ef8-977a-8dbcbb31a683",
        "min_value": 0,
        "name": "kelvin",
        "sub_characteristics": null,
        "type": "https://schema.org/Integer"
    },
    "/characteristics/urn:infai:ses:characteristic:7621686a-56bc-402d-b4cc-5b266d39736f": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:7621686a-56bc-402d-b4cc-5b266d39736f",
        "name": "on_off",
        "sub_characteristics": null,
        "type": "https://schema.org/Text"
    },
    "/characteristics/urn:infai:ses:characteristic:7768d6da-d53f-4afb-aeb3-e0980756ec06": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:7768d6da-d53f-4afb-aeb3-e0980756ec06",
        "name": "CO2 ppm",
        "sub_characteristics": null,
        "type": "https://schema.org/Float"
    },
    "/characteristics/urn:infai:ses:characteristic:7aaf53b9-6cd3-4e36-b27e-9e77489f6f78": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:7aaf53b9-6cd3-4e36-b27e-9e77489f6f78",
        "max_value": 100,
        "min_value": 0,
        "name": "fog area fraction percentage",
        "sub_characteristics": null,
        "type": "https://schema.org/Float"
    },
    "/characteristics/urn:infai:ses:characteristic:7dc1bb7e-b256-408a-a6f9-044dc60fdcf5": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:7dc1bb7e-b256-408a-a6f9-044dc60fdcf5",
        "name": "boolean",
        "sub_characteristics": null,
        "type": "https://schema.org/Boolean"
    },
    "/characteristics/urn:infai:ses:characteristic:8b9411f3-bf56-4e57-8fd4-728bdcd451cd": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:8b9411f3-bf56-4e57-8fd4-728bdcd451cd",
        "name": "unix_format_seconds",
        "sub_characteristics": null,
        "type": "https://schema.org/Text"
    },
    "/characteristics/urn:infai:ses:characteristic:8e520c73-7ee1-4a4b-954b-1c7c55139e27": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:8e520c73-7ee1-4a4b-954b-1c7c55139e27",
        "name": "any_clear",
        "sub_characteristics": null,
        "type": "https://schema.org/Text"
    },
    "/characteristics/urn:infai:ses:characteristic:9673cd38-8adf-48d0-b9ff-54ab56bbdcf0": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:9673cd38-8adf-48d0-b9ff-54ab56bbdcf0",
        "name": "ppm",
        "sub_characteristics": null,
        "type": "https://schema.org/Float"
    },
    "/characteristics/urn:infai:ses:characteristic:9e1024da-3b60-4531-9f29-464addccb13c": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:9e1024da-3b60-4531-9f29-464addccb13c",
        "min_value": 0,
        "name": "second",
        "sub_characteristics": null,
        "type": "https://schema.org/Float"
    },
    "/characteristics/urn:infai:ses:characteristic:a3931f28-664d-452f-9370-a641fadd21db": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:a3931f28-664d-452f-9370-a641fadd21db",
        "name": "watt",
        "sub_characteristics": null,
        "type": "https://schema.org/Float"
    },
    "/characteristics/urn:infai:ses:characteristic:a924efd4-9397-474b-ba01-6de49e0283e3": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:a924efd4-9397-474b-ba01-6de49e0283e3",
        "name": "volt",
        "sub_characteristics": null,
        "type": "https://schema.org/Float"
    },
    "/characteristics/urn:infai:ses:characteristic:aaf2385d-92a6-4fd0-ba4d-3e8567e733d6": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:aaf2385d-92a6-4fd0-ba4d-3e8567e733d6",
        "max_value": 100,
        "min_value": 0,
        "name": "humidity percentage",
        "sub_characteristics": null,
        "type": "https://schema.org/Float"
    },
    "/characteristics/urn:infai:ses:characteristic:ac7ad69f-b82e-454a-8b44-5c0440cba787": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:ac7ad69f-b82e-454a-8b44-5c0440cba787",
        "max_value": 10,
        "min_value": 1,
        "name": "1 to 10",
        "sub_characteristics": null,
        "type": "https://schema.org/Integer"
    },
    "/characteristics/urn:infai:ses:characteristic:afb0ebb5-8f10-4c09-b6af-15eb291cacc1": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:afb0ebb5-8f10-4c09-b6af-15eb291cacc1",
        "min_value": -273.15,
        "name": "dew point celsius",
        "sub_characteristics": null,
        "type": "https://schema.org/Float"
    },
    "/characteristics/urn:infai:ses:characteristic:b060dacf-f198-452b-87cf-740bf139e0c1": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:b060dacf-f198-452b-87cf-740bf139e0c1",
        "max_value": 7,
        "min_value": 1,
        "name": "dwd pollen level",
        "sub_characteristics": null,
        "type": "https://schema.org/Integer"
    },
    "/characteristics/urn:infai:ses:characteristic:b36eee5d-52f0-4476-a6f7-6dd03b24e0f8": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:b36eee5d-52f0-4476-a6f7-6dd03b24e0f8",
        "min_value": 0,
        "name": "minute",
        "sub_characteristics": null,
        "type": "https://schema.org/Float"
    },
    "/characteristics/urn:infai:ses:characteristic:b59c3965-c270-4053-b46a-f58b44cc980c": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:b59c3965-c270-4053-b46a-f58b44cc980c",
        "name": "ampere",
        "sub_characteristics": null,
        "type": "https://schema.org/Float"
    },
    "/characteristics/urn:infai:ses:characteristic:b6070dc9-0327-47b4-ba27-3287325940e5": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:b6070dc9-0327-47b4-ba27-3287325940e5",
        "name": "water level cm",
        "sub_characteristics": null,
        "type": "https://schema.org/Float"
    },
    "/characteristics/urn:infai:ses:characteristic:bc03ef2e-51d5-4034-8bec-75df78e3afee": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:bc03ef2e-51d5-4034-8bec-75df78e3afee",
        "name": "online_offline",
        "sub_characteristics": null,
        "type": "https://schema.org/Text"
    },
    "/characteristics/urn:infai:ses:characteristic:c0353532-a8fb-4553-a00b-418cb8a80a65": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:c0353532-a8fb-4553-a00b-418cb8a80a65",
        "max_value": 1,
        "min_value": 0,
        "name": "binary status code",
        "sub_characteristics": null,
        "type": "https://schema.org/Integer"
    },
    "/characteristics/urn:infai:ses:characteristic:c7dfcb86-2733-4917-a5ca-0a150a458eed": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:c7dfcb86-2733-4917-a5ca-0a150a458eed",
        "name": "unix_format_nanoseconds",
        "sub_characteristics": null,
        "type": "https://schema.org/Text"
    },
    "/characteristics/urn:infai:ses:characteristic:ca44269e-ab5a-4ed7-8acb-ce80a6a592a5": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:ca44269e-ab5a-4ed7-8acb-ce80a6a592a5",
        "max_value": 100,
        "min_value": 0,
        "name": "speed level percentage",
        "sub_characteristics": null,
        "type": "https://schema.org/Integer"
    },
    "/characteristics/urn:infai:ses:characteristic:d8a73a4d-8745-40b9-87e5-50c5d31f745a": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:d8a73a4d-8745-40b9-87e5-50c5d31f745a",
        "name": "longitude EPSG:4326",
        "sub_characteristics": null,
        "type": "https://schema.org/Float"
    },
    "/characteristics/urn:infai:ses:characteristic:de1a0d1d-85ae-45d3-844b-56c324c116cd": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:de1a0d1d-85ae-45d3-844b-56c324c116cd",
        "name": "particle amount µg/m³",
        "sub_characteristics": null,
        "type": "https://schema.org/Float"
    },
    "/characteristics/urn:infai:ses:characteristic:e6874605-80cf-418a-9319-da3b27411af1": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:e6874605-80cf-418a-9319-da3b27411af1",
        "max_value": 1,
        "min_value": 0,
        "name": "cos Phi",
        "sub_characteristics": null,
        "type": "https://schema.org/Float"
    },
    "/characteristics/urn:infai:ses:characteristic:eb33cf65-b0a2-413d-891d-cada05be01ed": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:eb33cf65-b0a2-413d-891d-cada05be01ed",
        "name": "hPa",
        "sub_characteristics": null,
        "type": "https://schema.org/Float"
    },
    "/characteristics/urn:infai:ses:characteristic:ec4ab96e-bfc3-4a80-96c4-b9e12634611f": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:ec4ab96e-bfc3-4a80-96c4-b9e12634611f",
        "max_value": 100,
        "min_value": 0,
        "name": "precipitation probability percentage",
        "sub_characteristics": null,
        "type": "https://schema.org/Float"
    },
    "/characteristics/urn:infai:ses:characteristic:f17ed384-82b6-4fba-908a-d146366c9779": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:f17ed384-82b6-4fba-908a-d146366c9779",
        "name": "ozone µg/m³",
        "sub_characteristics": null,
        "type": "https://schema.org/Float"
    },
    "/characteristics/urn:infai:ses:characteristic:f6b7ee6e-423c-47f1-8b1f-c489461d3039": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:f6b7ee6e-423c-47f1-8b1f-c489461d3039",
        "name": "benzene µg/m³",
        "sub_characteristics": null,
        "type": "https://schema.org/Float"
    },
    "/characteristics/urn:infai:ses:characteristic:fbfea6c7-3392-4ec2-9ec5-0d0ef6362b9b": {
        "display_unit": "",
        "id": "urn:infai:ses:characteristic:fbfea6c7-3392-4ec2-9ec5-0d0ef6362b9b",
        "name": "cubic meter",
        "sub_characteristics": null,
        "type": "https://schema.org/Float"
    }
}
`

const conceptPathMapStr = `{
"mock-placeholder": {
	"id": "mock-placeholder-concept",
	"characteristic_ids": [
		"example_brightness",
		"example_lux",
		"example_color",
		"example_rgb",
		"example_hex"
	]
},
"/concepts/urn:infai:ses:concept:08cde541-54dc-4d92-b2d9-b07304bd659a": {
        "base_characteristic_id": "urn:infai:ses:characteristic:f6b7ee6e-423c-47f1-8b1f-c489461d3039",
        "characteristic_ids": [
            "urn:infai:ses:characteristic:f6b7ee6e-423c-47f1-8b1f-c489461d3039"
        ],
        "id": "urn:infai:ses:concept:08cde541-54dc-4d92-b2d9-b07304bd659a",
        "name": "benzene"
    },
    "/concepts/urn:infai:ses:concept:0bc81398-3ed6-4e2b-a6c4-b754583aac37": {
        "base_characteristic_id": "urn:infai:ses:characteristic:5ba31623-0ccb-4488-bfb7-f73b50e03b5a",
        "characteristic_ids": [
            "urn:infai:ses:characteristic:75b2d113-1d03-4ef8-977a-8dbcbb31a683",
            "urn:infai:ses:characteristic:5ba31623-0ccb-4488-bfb7-f73b50e03b5a"
        ],
        "id": "urn:infai:ses:concept:0bc81398-3ed6-4e2b-a6c4-b754583aac37",
        "name": "temperature"
    },
    "/concepts/urn:infai:ses:concept:0c1eaf2b-ea3d-4ca9-8446-02c55dd5d3f4": {
        "base_characteristic_id": "urn:infai:ses:characteristic:7424a147-fb0e-4e71-a666-6c4997928e61",
        "characteristic_ids": [
            "urn:infai:ses:characteristic:7424a147-fb0e-4e71-a666-6c4997928e61"
        ],
        "id": "urn:infai:ses:concept:0c1eaf2b-ea3d-4ca9-8446-02c55dd5d3f4",
        "name": "precipitation"
    },
    "/concepts/urn:infai:ses:concept:0ec0adad-6873-4baa-931f-6544aa1dd210": {
        "base_characteristic_id": "urn:infai:ses:characteristic:b6070dc9-0327-47b4-ba27-3287325940e5",
        "characteristic_ids": [
            "urn:infai:ses:characteristic:b6070dc9-0327-47b4-ba27-3287325940e5"
        ],
        "id": "urn:infai:ses:concept:0ec0adad-6873-4baa-931f-6544aa1dd210",
        "name": "water level"
    },
    "/concepts/urn:infai:ses:concept:1b926adc-8e35-4c88-bfe1-15f23e67f4f5": {
        "base_characteristic_id": "urn:infai:ses:characteristic:f17ed384-82b6-4fba-908a-d146366c9779",
        "characteristic_ids": [
            "urn:infai:ses:characteristic:f17ed384-82b6-4fba-908a-d146366c9779"
        ],
        "id": "urn:infai:ses:concept:1b926adc-8e35-4c88-bfe1-15f23e67f4f5",
        "name": "ozone"
    },
    "/concepts/urn:infai:ses:concept:1c9e4d58-28e3-4e65-a0a0-48d8d519b0c9": {
        "base_characteristic_id": "urn:infai:ses:characteristic:0bf5bc86-dde2-465f-8a16-687ef9acba3d",
        "characteristic_ids": [
            "urn:infai:ses:characteristic:0bf5bc86-dde2-465f-8a16-687ef9acba3d"
        ],
        "id": "urn:infai:ses:concept:1c9e4d58-28e3-4e65-a0a0-48d8d519b0c9",
        "name": "frequency"
    },
    "/concepts/urn:infai:ses:concept:1f60e21d-0682-483c-9b0b-0146d8a1618d": {
        "base_characteristic_id": "urn:infai:ses:characteristic:3fec34de-0832-43c7-a071-bf0a858e3292",
        "characteristic_ids": [
            "urn:infai:ses:characteristic:3fec34de-0832-43c7-a071-bf0a858e3292"
        ],
        "id": "urn:infai:ses:concept:1f60e21d-0682-483c-9b0b-0146d8a1618d",
        "name": "nitrogen dioxide"
    },
    "/concepts/urn:infai:ses:concept:2769a588-97e7-4bc3-94f4-b97bef591e95": {
        "base_characteristic_id": "urn:infai:ses:characteristic:2ed0e8ca-e396-4c42-9446-6e6981e87845",
        "characteristic_ids": [
            "urn:infai:ses:characteristic:2ed0e8ca-e396-4c42-9446-6e6981e87845"
        ],
        "id": "urn:infai:ses:concept:2769a588-97e7-4bc3-94f4-b97bef591e95",
        "name": "signal strength"
    },
    "/concepts/urn:infai:ses:concept:397de0f8-a79f-484f-aa13-0042e48e5648": {
        "base_characteristic_id": "urn:infai:ses:characteristic:a924efd4-9397-474b-ba01-6de49e0283e3",
        "characteristic_ids": [
            "urn:infai:ses:characteristic:a924efd4-9397-474b-ba01-6de49e0283e3"
        ],
        "id": "urn:infai:ses:concept:397de0f8-a79f-484f-aa13-0042e48e5648",
        "name": "voltage"
    },
    "/concepts/urn:infai:ses:concept:3e40dfa7-805d-4732-af15-5207c0a59c47": {
        "base_characteristic_id": "urn:infai:ses:characteristic:e6874605-80cf-418a-9319-da3b27411af1",
        "characteristic_ids": [
            "urn:infai:ses:characteristic:e6874605-80cf-418a-9319-da3b27411af1"
        ],
        "id": "urn:infai:ses:concept:3e40dfa7-805d-4732-af15-5207c0a59c47",
        "name": "power factor"
    },
    "/concepts/urn:infai:ses:concept:40e8e501-85a0-4406-9aa9-7d5fd6985d1f": {
        "base_characteristic_id": "urn:infai:ses:characteristic:ec4ab96e-bfc3-4a80-96c4-b9e12634611f",
        "characteristic_ids": [
            "urn:infai:ses:characteristic:ec4ab96e-bfc3-4a80-96c4-b9e12634611f"
        ],
        "id": "urn:infai:ses:concept:40e8e501-85a0-4406-9aa9-7d5fd6985d1f",
        "name": "precipitation probability"
    },
    "/concepts/urn:infai:ses:concept:46de11b9-26ff-4cce-b945-e93b84f04fe6": {
        "base_characteristic_id": "urn:infai:ses:characteristic:a3931f28-664d-452f-9370-a641fadd21db",
        "characteristic_ids": [
            "urn:infai:ses:characteristic:a3931f28-664d-452f-9370-a641fadd21db"
        ],
        "id": "urn:infai:ses:concept:46de11b9-26ff-4cce-b945-e93b84f04fe6",
        "name": "power"
    },
    "/concepts/urn:infai:ses:concept:55efac99-6e5b-4ef3-8f4a-8576c1fe8614": {
        "base_characteristic_id": "urn:infai:ses:characteristic:7768d6da-d53f-4afb-aeb3-e0980756ec06",
        "characteristic_ids": [
            "urn:infai:ses:characteristic:7768d6da-d53f-4afb-aeb3-e0980756ec06"
        ],
        "id": "urn:infai:ses:concept:55efac99-6e5b-4ef3-8f4a-8576c1fe8614",
        "name": "carbon dioxid"
    },
    "/concepts/urn:infai:ses:concept:5859621b-f6b3-4430-8b4f-b25a53bae708": {
        "base_characteristic_id": "urn:infai:ses:characteristic:b060dacf-f198-452b-87cf-740bf139e0c1",
        "characteristic_ids": [
            "urn:infai:ses:characteristic:b060dacf-f198-452b-87cf-740bf139e0c1"
        ],
        "id": "urn:infai:ses:concept:5859621b-f6b3-4430-8b4f-b25a53bae708",
        "name": "pollen level"
    },
    "/concepts/urn:infai:ses:concept:58d23b4f-c315-4d7d-9230-fa989d55640a": {
        "base_characteristic_id": "urn:infai:ses:characteristic:9e1024da-3b60-4531-9f29-464addccb13c",
        "characteristic_ids": [
            "urn:infai:ses:characteristic:9e1024da-3b60-4531-9f29-464addccb13c",
            "urn:infai:ses:characteristic:b36eee5d-52f0-4476-a6f7-6dd03b24e0f8",
            "urn:infai:ses:characteristic:182e6bb4-8622-453a-9a4d-bfa9c70d6c9c"
        ],
        "id": "urn:infai:ses:concept:58d23b4f-c315-4d7d-9230-fa989d55640a",
        "name": "time"
    },
    "/concepts/urn:infai:ses:concept:59b02acd-aac2-4415-976e-bd4790963ce4": {
        "base_characteristic_id": "urn:infai:ses:characteristic:0419856b-d198-480e-9998-bf990f226844",
        "characteristic_ids": [
            "urn:infai:ses:characteristic:0419856b-d198-480e-9998-bf990f226844"
        ],
        "id": "urn:infai:ses:concept:59b02acd-aac2-4415-976e-bd4790963ce4",
        "name": "luminiscence"
    },
    "/concepts/urn:infai:ses:concept:64106d72-d5b7-4df3-a1bf-cc369799434e": {
        "base_characteristic_id": "urn:infai:ses:characteristic:0958edd7-eaae-46fe-b612-0bf4e53400af",
        "characteristic_ids": [
            "urn:infai:ses:characteristic:0958edd7-eaae-46fe-b612-0bf4e53400af"
        ],
        "id": "urn:infai:ses:concept:64106d72-d5b7-4df3-a1bf-cc369799434e",
        "name": "cloud fraction"
    },
    "/concepts/urn:infai:ses:concept:6625cc35-a8d0-4d9d-91ef-3c0cb71f0285": {
        "base_characteristic_id": "urn:infai:ses:characteristic:02c7ea52-fc76-470e-b2b2-401653412de6",
        "characteristic_ids": [
            "urn:infai:ses:characteristic:02c7ea52-fc76-470e-b2b2-401653412de6"
        ],
        "id": "urn:infai:ses:concept:6625cc35-a8d0-4d9d-91ef-3c0cb71f0285",
        "name": "thunder probability"
    },
    "/concepts/urn:infai:ses:concept:67fbf29b-fd77-4c82-b8fa-19a6000c602d": {
        "base_characteristic_id": "urn:infai:ses:characteristic:aaf2385d-92a6-4fd0-ba4d-3e8567e733d6",
        "characteristic_ids": [
            "urn:infai:ses:characteristic:aaf2385d-92a6-4fd0-ba4d-3e8567e733d6"
        ],
        "id": "urn:infai:ses:concept:67fbf29b-fd77-4c82-b8fa-19a6000c602d",
        "name": "humidity"
    },
    "/concepts/urn:infai:ses:concept:76a1082b-7e86-4ba0-b126-7c449e581563": {
        "base_characteristic_id": "urn:infai:ses:characteristic:005ca5d3-44da-4fc8-b2b8-a88e3209e9f7",
        "characteristic_ids": [
            "urn:infai:ses:characteristic:005ca5d3-44da-4fc8-b2b8-a88e3209e9f7"
        ],
        "id": "urn:infai:ses:concept:76a1082b-7e86-4ba0-b126-7c449e581563",
        "name": "direction"
    },
    "/concepts/urn:infai:ses:concept:7d28f92e-8836-42bf-bca6-342989059ddf": {
        "base_characteristic_id": "urn:infai:ses:characteristic:afb0ebb5-8f10-4c09-b6af-15eb291cacc1",
        "characteristic_ids": [
            "urn:infai:ses:characteristic:afb0ebb5-8f10-4c09-b6af-15eb291cacc1"
        ],
        "id": "urn:infai:ses:concept:7d28f92e-8836-42bf-bca6-342989059ddf",
        "name": "dew point"
    },
    "/concepts/urn:infai:ses:concept:7ffc7509-51dc-4a62-bb42-850090996586": {
        "base_characteristic_id": "urn:infai:ses:characteristic:b59c3965-c270-4053-b46a-f58b44cc980c",
        "characteristic_ids": [
            "urn:infai:ses:characteristic:b59c3965-c270-4053-b46a-f58b44cc980c"
        ],
        "id": "urn:infai:ses:concept:7ffc7509-51dc-4a62-bb42-850090996586",
        "name": "current"
    },
    "/concepts/urn:infai:ses:concept:81c003ac-ed79-4b3d-8d49-6d95b61b2863": {
        "base_characteristic_id": "urn:infai:ses:characteristic:fbfea6c7-3392-4ec2-9ec5-0d0ef6362b9b",
        "characteristic_ids": [
            "urn:infai:ses:characteristic:fbfea6c7-3392-4ec2-9ec5-0d0ef6362b9b"
        ],
        "id": "urn:infai:ses:concept:81c003ac-ed79-4b3d-8d49-6d95b61b2863",
        "name": "volume"
    },
    "/concepts/urn:infai:ses:concept:85e11726-620a-4584-96a2-3a6fe4141b2d": {
        "base_characteristic_id": "urn:infai:ses:characteristic:bc03ef2e-51d5-4034-8bec-75df78e3afee",
        "characteristic_ids": [
            "urn:infai:ses:characteristic:bc03ef2e-51d5-4034-8bec-75df78e3afee"
        ],
        "id": "urn:infai:ses:concept:85e11726-620a-4584-96a2-3a6fe4141b2d",
        "name": "connection status"
    },
    "/concepts/urn:infai:ses:concept:8b1161d5-7878-4dd2-a36c-6f98f6b94bf8": {
        "base_characteristic_id": "urn:infai:ses:characteristic:5b4eea52-e8e5-4e80-9455-0382f81a1b43",
        "characteristic_ids": [
            "urn:infai:ses:characteristic:0fc343ce-4627-4c88-b1e0-d3ed29754af8",
            "urn:infai:ses:characteristic:5b4eea52-e8e5-4e80-9455-0382f81a1b43",
            "urn:infai:ses:characteristic:64928e9f-98ca-42bb-a1e5-adf2a760a2f9"
        ],
        "id": "urn:infai:ses:concept:8b1161d5-7878-4dd2-a36c-6f98f6b94bf8",
        "name": "color"
    },
    "/concepts/urn:infai:ses:concept:8f4a00c9-061b-41d8-901f-192b15f5a465": {
        "base_characteristic_id": "urn:infai:ses:characteristic:5b5358c4-efd3-48a5-87b5-a8e0a34ce1ab",
        "characteristic_ids": [
            "urn:infai:ses:characteristic:5b5358c4-efd3-48a5-87b5-a8e0a34ce1ab"
        ],
        "id": "urn:infai:ses:concept:8f4a00c9-061b-41d8-901f-192b15f5a465",
        "name": "configuration"
    },
    "/concepts/urn:infai:ses:concept:9286d398-f49d-494f-8157-bea17947b1fa": {
        "base_characteristic_id": "urn:infai:ses:characteristic:6bc41b45-a9f3-4d87-9c51-dd3e11257800",
        "characteristic_ids": [
            "urn:infai:ses:characteristic:6bc41b45-a9f3-4d87-9c51-dd3e11257800",
            "urn:infai:ses:characteristic:8b9411f3-bf56-4e57-8fd4-728bdcd451cd",
            "urn:infai:ses:characteristic:c7dfcb86-2733-4917-a5ca-0a150a458eed",
            "urn:infai:ses:characteristic:64691f8d-4909-470f-a1fa-e977ebe28684"
        ],
        "id": "urn:infai:ses:concept:9286d398-f49d-494f-8157-bea17947b1fa",
        "name": "timestamp"
    },
    "/concepts/urn:infai:ses:concept:95f01a66-f7ea-4c49-9304-29cf065a0c55": {
        "base_characteristic_id": "urn:infai:ses:characteristic:04bce19c-839d-45de-8cfb-bd470715f4cd",
        "characteristic_ids": [
            "urn:infai:ses:characteristic:04bce19c-839d-45de-8cfb-bd470715f4cd"
        ],
        "id": "urn:infai:ses:concept:95f01a66-f7ea-4c49-9304-29cf065a0c55",
        "name": "speed"
    },
    "/concepts/urn:infai:ses:concept:9924ea1a-d96a-4882-88a2-165493956ec3": {
        "base_characteristic_id": "urn:infai:ses:characteristic:eb33cf65-b0a2-413d-891d-cada05be01ed",
        "characteristic_ids": [
            "urn:infai:ses:characteristic:2871c1c4-1c8c-4169-81f6-fb5313df08b1",
            "urn:infai:ses:characteristic:eb33cf65-b0a2-413d-891d-cada05be01ed"
        ],
        "id": "urn:infai:ses:concept:9924ea1a-d96a-4882-88a2-165493956ec3",
        "name": "pressure"
    },
    "/concepts/urn:infai:ses:concept:a43e8989-7d4b-4d9c-97d7-766799106499": {
        "base_characteristic_id": "urn:infai:ses:characteristic:2b4040c4-63b2-497f-8673-96378dcbc4cf",
        "characteristic_ids": [
            "urn:infai:ses:characteristic:ca44269e-ab5a-4ed7-8acb-ce80a6a592a5",
            "urn:infai:ses:characteristic:2b4040c4-63b2-497f-8673-96378dcbc4cf",
            "urn:infai:ses:characteristic:ac7ad69f-b82e-454a-8b44-5c0440cba787"
        ],
        "id": "urn:infai:ses:concept:a43e8989-7d4b-4d9c-97d7-766799106499",
        "name": "speed level"
    },
    "/concepts/urn:infai:ses:concept:a570ecaf-2939-4524-8bd3-c5f12b1e6863": {
        "base_characteristic_id": "urn:infai:ses:characteristic:d8a73a4d-8745-40b9-87e5-50c5d31f745a",
        "characteristic_ids": [
            "urn:infai:ses:characteristic:d8a73a4d-8745-40b9-87e5-50c5d31f745a"
        ],
        "id": "urn:infai:ses:concept:a570ecaf-2939-4524-8bd3-c5f12b1e6863",
        "name": "longitude"
    },
    "/concepts/urn:infai:ses:concept:a63a960d-1f36-4d95-ad5b-dfb4a7fe3b5b": {
        "base_characteristic_id": "urn:infai:ses:characteristic:9673cd38-8adf-48d0-b9ff-54ab56bbdcf0",
        "characteristic_ids": [
            "urn:infai:ses:characteristic:de1a0d1d-85ae-45d3-844b-56c324c116cd",
            "urn:infai:ses:characteristic:9673cd38-8adf-48d0-b9ff-54ab56bbdcf0"
        ],
        "id": "urn:infai:ses:concept:a63a960d-1f36-4d95-ad5b-dfb4a7fe3b5b",
        "name": "particle amount"
    },
    "/concepts/urn:infai:ses:concept:a7cd1e62-72fc-448c-8492-6a8f0307bbdc": {
        "base_characteristic_id": "urn:infai:ses:characteristic:0a61343d-c0d1-4af8-9329-3829c30ba59f",
        "characteristic_ids": [
            "urn:infai:ses:characteristic:0a61343d-c0d1-4af8-9329-3829c30ba59f"
        ],
        "id": "urn:infai:ses:concept:a7cd1e62-72fc-448c-8492-6a8f0307bbdc",
        "name": "ultraviolet"
    },
    "/concepts/urn:infai:ses:concept:b12151d3-d187-4839-82ac-8cc18fae3dd9": {
        "base_characteristic_id": "urn:infai:ses:characteristic:107c29d3-f863-47e1-a1e0-7853a8016213",
        "characteristic_ids": [
            "urn:infai:ses:characteristic:107c29d3-f863-47e1-a1e0-7853a8016213"
        ],
        "id": "urn:infai:ses:concept:b12151d3-d187-4839-82ac-8cc18fae3dd9",
        "name": "position"
    },
    "/concepts/urn:infai:ses:concept:bc7a76fc-86d0-470a-a872-00ef211904f9": {
        "base_characteristic_id": "urn:infai:ses:characteristic:193ff994-60a4-4ed2-90a1-21b82622cb43",
        "characteristic_ids": [
            "urn:infai:ses:characteristic:15af5b7b-8b17-4b18-b37e-2e91c839886b",
            "urn:infai:ses:characteristic:193ff994-60a4-4ed2-90a1-21b82622cb43"
        ],
        "id": "urn:infai:ses:concept:bc7a76fc-86d0-470a-a872-00ef211904f9",
        "name": "carbon monoxide"
    },
    "/concepts/urn:infai:ses:concept:c41d2566-aad6-458a-8f21-fee504225b17": {
        "base_characteristic_id": "urn:infai:ses:characteristic:2ce5f9c7-3c44-438b-91e6-3e9901b8a3db",
        "characteristic_ids": [
            "urn:infai:ses:characteristic:2ce5f9c7-3c44-438b-91e6-3e9901b8a3db"
        ],
        "id": "urn:infai:ses:concept:c41d2566-aad6-458a-8f21-fee504225b17",
        "name": "sulfur dioxide"
    },
    "/concepts/urn:infai:ses:concept:c6d4cb5b-5ae4-489e-a664-cf02003f1d25": {
        "base_characteristic_id": "urn:infai:ses:characteristic:63bb46ea-64f1-4a60-aabb-67b9febdb588",
        "characteristic_ids": [
            "urn:infai:ses:characteristic:63bb46ea-64f1-4a60-aabb-67b9febdb588"
        ],
        "id": "urn:infai:ses:concept:c6d4cb5b-5ae4-489e-a664-cf02003f1d25",
        "name": "latitude"
    },
    "/concepts/urn:infai:ses:concept:cb2608c2-b671-46be-8bad-1539e6c6e2af": {
        "base_characteristic_id": "urn:infai:ses:characteristic:3f0294a9-a53e-4650-b3d1-6a6482cf3b70",
        "characteristic_ids": [
            "urn:infai:ses:characteristic:3f0294a9-a53e-4650-b3d1-6a6482cf3b70"
        ],
        "id": "urn:infai:ses:concept:cb2608c2-b671-46be-8bad-1539e6c6e2af",
        "name": "volume flow"
    },
    "/concepts/urn:infai:ses:concept:cedf4396-8e55-491d-b6f8-02ff94679b7d": {
        "base_characteristic_id": "urn:infai:ses:characteristic:3febed55-ba9b-43dc-8709-9c73bae3716e",
        "characteristic_ids": [
            "urn:infai:ses:characteristic:3febed55-ba9b-43dc-8709-9c73bae3716e"
        ],
        "id": "urn:infai:ses:concept:cedf4396-8e55-491d-b6f8-02ff94679b7d",
        "name": "energy"
    },
    "/concepts/urn:infai:ses:concept:dbe4ad57-aa1d-4d24-9bee-a44a1c670d7f": {
        "base_characteristic_id": "urn:infai:ses:characteristic:72b624b5-6edc-4ec4-9ad9-fa00b39915c0",
        "characteristic_ids": [
            "urn:infai:ses:characteristic:72b624b5-6edc-4ec4-9ad9-fa00b39915c0"
        ],
        "id": "urn:infai:ses:concept:dbe4ad57-aa1d-4d24-9bee-a44a1c670d7f",
        "name": "brightness"
    },
    "/concepts/urn:infai:ses:concept:ebfeabb3-50f0-44bd-b06e-95eb52df484e": {
        "base_characteristic_id": "urn:infai:ses:characteristic:7dc1bb7e-b256-408a-a6f9-044dc60fdcf5",
        "characteristic_ids": [
            "urn:infai:ses:characteristic:8e520c73-7ee1-4a4b-954b-1c7c55139e27",
            "urn:infai:ses:characteristic:7621686a-56bc-402d-b4cc-5b266d39736f",
            "urn:infai:ses:characteristic:c0353532-a8fb-4553-a00b-418cb8a80a65",
            "urn:infai:ses:characteristic:7dc1bb7e-b256-408a-a6f9-044dc60fdcf5"
        ],
        "id": "urn:infai:ses:concept:ebfeabb3-50f0-44bd-b06e-95eb52df484e",
        "name": "binary state"
    },
    "/concepts/urn:infai:ses:concept:f7fae105-b4d9-40c9-8a47-82c377899cd5": {
        "base_characteristic_id": "urn:infai:ses:characteristic:7aaf53b9-6cd3-4e36-b27e-9e77489f6f78",
        "characteristic_ids": [
            "urn:infai:ses:characteristic:7aaf53b9-6cd3-4e36-b27e-9e77489f6f78"
        ],
        "id": "urn:infai:ses:concept:f7fae105-b4d9-40c9-8a47-82c377899cd5",
        "name": "fog area fraction"
    },
    "/concepts/urn:infai:ses:concept:f917e146-9192-4796-8bbe-6be14292346f": {
        "base_characteristic_id": "urn:infai:ses:characteristic:5caa707d-dc08-4f3b-bd9f-f08935c8dd3c",
        "characteristic_ids": [
            "urn:infai:ses:characteristic:5caa707d-dc08-4f3b-bd9f-f08935c8dd3c"
        ],
        "id": "urn:infai:ses:concept:f917e146-9192-4796-8bbe-6be14292346f",
        "name": "lifespan"
    },
    "/concepts/urn:infai:ses:concept:fb85f32b-1864-456d-b99f-1540892ffd02": {
        "base_characteristic_id": "urn:infai:ses:characteristic:17cfbe28-ce44-4355-bf75-15b6bf39ba1d",
        "characteristic_ids": [
            "urn:infai:ses:characteristic:17cfbe28-ce44-4355-bf75-15b6bf39ba1d",
            "urn:infai:ses:characteristic:65e9a5db-348f-4888-9d95-588741d67dbf"
        ],
        "id": "urn:infai:ses:concept:fb85f32b-1864-456d-b99f-1540892ffd02",
        "name": "contact"
    }
}
`
