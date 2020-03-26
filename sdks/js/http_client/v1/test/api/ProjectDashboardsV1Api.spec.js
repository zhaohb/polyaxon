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
    // AMD.
    define(['expect.js', '../../src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require('../../src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.PolyaxonSdk);
  }
}(this, function(expect, PolyaxonSdk) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new PolyaxonSdk.ProjectDashboardsV1Api();
  });

  describe('(package)', function() {
    describe('ProjectDashboardsV1Api', function() {
      describe('createProjectDashboard', function() {
        it('should call createProjectDashboard successfully', function(done) {
          // TODO: uncomment, update parameter values for createProjectDashboard call and complete the assertions
          /*
          var owner = "owner_example";
          var project = "project_example";
          var body = new PolyaxonSdk.V1Dashboard();
          body.uuid = "";
          body.name = "";
          body.description = "";
          body.tags = [""];
          body.disabled = false;
          body.deleted = false;
          body.run_view = false;
          body.search = new PolyaxonSdk.V1SearchSpec();
          body.search.query = "";
          body.search.sort = "";
          body.search.limit = 0;
          body.search.groupby = "";
          body.search.columns = "";
          body.spec = ;
          body.created_at = new Date();
          body.updated_at = new Date();

          instance.createProjectDashboard(owner, project, body, function(error, data, response) {
            if (error) {
              done(error);
              return;
            }
            // TODO: update response assertions
            expect(data).to.be.a(PolyaxonSdk.V1Dashboard);
            expect(data.uuid).to.be.a('string');
            expect(data.uuid).to.be("");
            expect(data.name).to.be.a('string');
            expect(data.name).to.be("");
            expect(data.description).to.be.a('string');
            expect(data.description).to.be("");
            {
              let dataCtr = data.tags;
              expect(dataCtr).to.be.an(Array);
              expect(dataCtr).to.not.be.empty();
              for (let p in dataCtr) {
                let data = dataCtr[p];
                expect(data).to.be.a('string');
                expect(data).to.be("");
              }
            }
            expect(data.disabled).to.be.a('boolean');
            expect(data.disabled).to.be(false);
            expect(data.deleted).to.be.a('boolean');
            expect(data.deleted).to.be(false);
            expect(data.run_view).to.be.a('boolean');
            expect(data.run_view).to.be(false);
            expect(data.search).to.be.a(PolyaxonSdk.V1SearchSpec);
                  expect(data.search.query).to.be.a('string');
              expect(data.search.query).to.be("");
              expect(data.search.sort).to.be.a('string');
              expect(data.search.sort).to.be("");
              expect(data.search.limit).to.be.a('number');
              expect(data.search.limit).to.be(0);
              expect(data.search.groupby).to.be.a('string');
              expect(data.search.groupby).to.be("");
              expect(data.search.columns).to.be.a('string');
              expect(data.search.columns).to.be("");
            expect(data.spec).to.be.a(Object);
            expect(data.spec).to.be();
            expect(data.created_at).to.be.a(Date);
            expect(data.created_at).to.be(new Date());
            expect(data.updated_at).to.be.a(Date);
            expect(data.updated_at).to.be(new Date());

            done();
          });
          */
          // TODO: uncomment and complete method invocation above, then delete this line and the next:
          done();
        });
      });
      describe('deleteProjectDashboard', function() {
        it('should call deleteProjectDashboard successfully', function(done) {
          // TODO: uncomment, update parameter values for deleteProjectDashboard call
          /*
          var owner = "owner_example";
          var project = "project_example";
          var uuid = "uuid_example";

          instance.deleteProjectDashboard(owner, project, uuid, function(error, data, response) {
            if (error) {
              done(error);
              return;
            }

            done();
          });
          */
          // TODO: uncomment and complete method invocation above, then delete this line and the next:
          done();
        });
      });
      describe('getProjectDashboard', function() {
        it('should call getProjectDashboard successfully', function(done) {
          // TODO: uncomment, update parameter values for getProjectDashboard call and complete the assertions
          /*
          var owner = "owner_example";
          var project = "project_example";
          var uuid = "uuid_example";

          instance.getProjectDashboard(owner, project, uuid, function(error, data, response) {
            if (error) {
              done(error);
              return;
            }
            // TODO: update response assertions
            expect(data).to.be.a(PolyaxonSdk.V1Dashboard);
            expect(data.uuid).to.be.a('string');
            expect(data.uuid).to.be("");
            expect(data.name).to.be.a('string');
            expect(data.name).to.be("");
            expect(data.description).to.be.a('string');
            expect(data.description).to.be("");
            {
              let dataCtr = data.tags;
              expect(dataCtr).to.be.an(Array);
              expect(dataCtr).to.not.be.empty();
              for (let p in dataCtr) {
                let data = dataCtr[p];
                expect(data).to.be.a('string');
                expect(data).to.be("");
              }
            }
            expect(data.disabled).to.be.a('boolean');
            expect(data.disabled).to.be(false);
            expect(data.deleted).to.be.a('boolean');
            expect(data.deleted).to.be(false);
            expect(data.run_view).to.be.a('boolean');
            expect(data.run_view).to.be(false);
            expect(data.search).to.be.a(PolyaxonSdk.V1SearchSpec);
                  expect(data.search.query).to.be.a('string');
              expect(data.search.query).to.be("");
              expect(data.search.sort).to.be.a('string');
              expect(data.search.sort).to.be("");
              expect(data.search.limit).to.be.a('number');
              expect(data.search.limit).to.be(0);
              expect(data.search.groupby).to.be.a('string');
              expect(data.search.groupby).to.be("");
              expect(data.search.columns).to.be.a('string');
              expect(data.search.columns).to.be("");
            expect(data.spec).to.be.a(Object);
            expect(data.spec).to.be();
            expect(data.created_at).to.be.a(Date);
            expect(data.created_at).to.be(new Date());
            expect(data.updated_at).to.be.a(Date);
            expect(data.updated_at).to.be(new Date());

            done();
          });
          */
          // TODO: uncomment and complete method invocation above, then delete this line and the next:
          done();
        });
      });
      describe('listProjectDashboardNames', function() {
        it('should call listProjectDashboardNames successfully', function(done) {
          // TODO: uncomment, update parameter values for listProjectDashboardNames call and complete the assertions
          /*
          var owner = "owner_example";
          var project = "project_example";
          var opts = {};
          opts.offset = 56;
          opts.limit = 56;
          opts.sort = "sort_example";
          opts.query = "query_example";

          instance.listProjectDashboardNames(owner, project, opts, function(error, data, response) {
            if (error) {
              done(error);
              return;
            }
            // TODO: update response assertions
            expect(data).to.be.a(PolyaxonSdk.V1ListDashboardsResponse);
            expect(data.count).to.be.a('number');
            expect(data.count).to.be(0);
            {
              let dataCtr = data.results;
              expect(dataCtr).to.be.an(Array);
              expect(dataCtr).to.not.be.empty();
              for (let p in dataCtr) {
                let data = dataCtr[p];
                expect(data).to.be.a(PolyaxonSdk.V1Dashboard);
                expect(data.uuid).to.be.a('string');
                expect(data.uuid).to.be("");
                expect(data.name).to.be.a('string');
                expect(data.name).to.be("");
                expect(data.description).to.be.a('string');
                expect(data.description).to.be("");
                {
                  let dataCtr = data.tags;
                  expect(dataCtr).to.be.an(Array);
                  expect(dataCtr).to.not.be.empty();
                  for (let p in dataCtr) {
                    let data = dataCtr[p];
                    expect(data).to.be.a('string');
                    expect(data).to.be("");
                  }
                }
                expect(data.disabled).to.be.a('boolean');
                expect(data.disabled).to.be(false);
                expect(data.deleted).to.be.a('boolean');
                expect(data.deleted).to.be(false);
                expect(data.run_view).to.be.a('boolean');
                expect(data.run_view).to.be(false);
                expect(data.search).to.be.a(PolyaxonSdk.V1SearchSpec);
                      expect(data.search.query).to.be.a('string');
                  expect(data.search.query).to.be("");
                  expect(data.search.sort).to.be.a('string');
                  expect(data.search.sort).to.be("");
                  expect(data.search.limit).to.be.a('number');
                  expect(data.search.limit).to.be(0);
                  expect(data.search.groupby).to.be.a('string');
                  expect(data.search.groupby).to.be("");
                  expect(data.search.columns).to.be.a('string');
                  expect(data.search.columns).to.be("");
                expect(data.spec).to.be.a(Object);
                expect(data.spec).to.be();
                expect(data.created_at).to.be.a(Date);
                expect(data.created_at).to.be(new Date());
                expect(data.updated_at).to.be.a(Date);
                expect(data.updated_at).to.be(new Date());
              }
            }
            expect(data.previous).to.be.a('string');
            expect(data.previous).to.be("");
            expect(data.next).to.be.a('string');
            expect(data.next).to.be("");

            done();
          });
          */
          // TODO: uncomment and complete method invocation above, then delete this line and the next:
          done();
        });
      });
      describe('listProjectDashboards', function() {
        it('should call listProjectDashboards successfully', function(done) {
          // TODO: uncomment, update parameter values for listProjectDashboards call and complete the assertions
          /*
          var owner = "owner_example";
          var project = "project_example";
          var opts = {};
          opts.offset = 56;
          opts.limit = 56;
          opts.sort = "sort_example";
          opts.query = "query_example";

          instance.listProjectDashboards(owner, project, opts, function(error, data, response) {
            if (error) {
              done(error);
              return;
            }
            // TODO: update response assertions
            expect(data).to.be.a(PolyaxonSdk.V1ListDashboardsResponse);
            expect(data.count).to.be.a('number');
            expect(data.count).to.be(0);
            {
              let dataCtr = data.results;
              expect(dataCtr).to.be.an(Array);
              expect(dataCtr).to.not.be.empty();
              for (let p in dataCtr) {
                let data = dataCtr[p];
                expect(data).to.be.a(PolyaxonSdk.V1Dashboard);
                expect(data.uuid).to.be.a('string');
                expect(data.uuid).to.be("");
                expect(data.name).to.be.a('string');
                expect(data.name).to.be("");
                expect(data.description).to.be.a('string');
                expect(data.description).to.be("");
                {
                  let dataCtr = data.tags;
                  expect(dataCtr).to.be.an(Array);
                  expect(dataCtr).to.not.be.empty();
                  for (let p in dataCtr) {
                    let data = dataCtr[p];
                    expect(data).to.be.a('string');
                    expect(data).to.be("");
                  }
                }
                expect(data.disabled).to.be.a('boolean');
                expect(data.disabled).to.be(false);
                expect(data.deleted).to.be.a('boolean');
                expect(data.deleted).to.be(false);
                expect(data.run_view).to.be.a('boolean');
                expect(data.run_view).to.be(false);
                expect(data.search).to.be.a(PolyaxonSdk.V1SearchSpec);
                      expect(data.search.query).to.be.a('string');
                  expect(data.search.query).to.be("");
                  expect(data.search.sort).to.be.a('string');
                  expect(data.search.sort).to.be("");
                  expect(data.search.limit).to.be.a('number');
                  expect(data.search.limit).to.be(0);
                  expect(data.search.groupby).to.be.a('string');
                  expect(data.search.groupby).to.be("");
                  expect(data.search.columns).to.be.a('string');
                  expect(data.search.columns).to.be("");
                expect(data.spec).to.be.a(Object);
                expect(data.spec).to.be();
                expect(data.created_at).to.be.a(Date);
                expect(data.created_at).to.be(new Date());
                expect(data.updated_at).to.be.a(Date);
                expect(data.updated_at).to.be(new Date());
              }
            }
            expect(data.previous).to.be.a('string');
            expect(data.previous).to.be("");
            expect(data.next).to.be.a('string');
            expect(data.next).to.be("");

            done();
          });
          */
          // TODO: uncomment and complete method invocation above, then delete this line and the next:
          done();
        });
      });
      describe('patchProjectDashboard', function() {
        it('should call patchProjectDashboard successfully', function(done) {
          // TODO: uncomment, update parameter values for patchProjectDashboard call and complete the assertions
          /*
          var owner = "owner_example";
          var project = "project_example";
          var dashboard_uuid = "dashboard_uuid_example";
          var body = new PolyaxonSdk.V1Dashboard();
          body.uuid = "";
          body.name = "";
          body.description = "";
          body.tags = [""];
          body.disabled = false;
          body.deleted = false;
          body.run_view = false;
          body.search = new PolyaxonSdk.V1SearchSpec();
          body.search.query = "";
          body.search.sort = "";
          body.search.limit = 0;
          body.search.groupby = "";
          body.search.columns = "";
          body.spec = ;
          body.created_at = new Date();
          body.updated_at = new Date();

          instance.patchProjectDashboard(owner, project, dashboard_uuid, body, function(error, data, response) {
            if (error) {
              done(error);
              return;
            }
            // TODO: update response assertions
            expect(data).to.be.a(PolyaxonSdk.V1Dashboard);
            expect(data.uuid).to.be.a('string');
            expect(data.uuid).to.be("");
            expect(data.name).to.be.a('string');
            expect(data.name).to.be("");
            expect(data.description).to.be.a('string');
            expect(data.description).to.be("");
            {
              let dataCtr = data.tags;
              expect(dataCtr).to.be.an(Array);
              expect(dataCtr).to.not.be.empty();
              for (let p in dataCtr) {
                let data = dataCtr[p];
                expect(data).to.be.a('string');
                expect(data).to.be("");
              }
            }
            expect(data.disabled).to.be.a('boolean');
            expect(data.disabled).to.be(false);
            expect(data.deleted).to.be.a('boolean');
            expect(data.deleted).to.be(false);
            expect(data.run_view).to.be.a('boolean');
            expect(data.run_view).to.be(false);
            expect(data.search).to.be.a(PolyaxonSdk.V1SearchSpec);
                  expect(data.search.query).to.be.a('string');
              expect(data.search.query).to.be("");
              expect(data.search.sort).to.be.a('string');
              expect(data.search.sort).to.be("");
              expect(data.search.limit).to.be.a('number');
              expect(data.search.limit).to.be(0);
              expect(data.search.groupby).to.be.a('string');
              expect(data.search.groupby).to.be("");
              expect(data.search.columns).to.be.a('string');
              expect(data.search.columns).to.be("");
            expect(data.spec).to.be.a(Object);
            expect(data.spec).to.be();
            expect(data.created_at).to.be.a(Date);
            expect(data.created_at).to.be(new Date());
            expect(data.updated_at).to.be.a(Date);
            expect(data.updated_at).to.be(new Date());

            done();
          });
          */
          // TODO: uncomment and complete method invocation above, then delete this line and the next:
          done();
        });
      });
      describe('promoteProjectDashboard', function() {
        it('should call promoteProjectDashboard successfully', function(done) {
          // TODO: uncomment, update parameter values for promoteProjectDashboard call and complete the assertions
          /*
          var owner = "owner_example";
          var project = "project_example";
          var dashboard_uuid = "dashboard_uuid_example";

          instance.promoteProjectDashboard(owner, project, dashboard_uuid, function(error, data, response) {
            if (error) {
              done(error);
              return;
            }
            // TODO: update response assertions
            expect(data).to.be.a(PolyaxonSdk.V1Dashboard);
            expect(data.uuid).to.be.a('string');
            expect(data.uuid).to.be("");
            expect(data.name).to.be.a('string');
            expect(data.name).to.be("");
            expect(data.description).to.be.a('string');
            expect(data.description).to.be("");
            {
              let dataCtr = data.tags;
              expect(dataCtr).to.be.an(Array);
              expect(dataCtr).to.not.be.empty();
              for (let p in dataCtr) {
                let data = dataCtr[p];
                expect(data).to.be.a('string');
                expect(data).to.be("");
              }
            }
            expect(data.disabled).to.be.a('boolean');
            expect(data.disabled).to.be(false);
            expect(data.deleted).to.be.a('boolean');
            expect(data.deleted).to.be(false);
            expect(data.run_view).to.be.a('boolean');
            expect(data.run_view).to.be(false);
            expect(data.search).to.be.a(PolyaxonSdk.V1SearchSpec);
                  expect(data.search.query).to.be.a('string');
              expect(data.search.query).to.be("");
              expect(data.search.sort).to.be.a('string');
              expect(data.search.sort).to.be("");
              expect(data.search.limit).to.be.a('number');
              expect(data.search.limit).to.be(0);
              expect(data.search.groupby).to.be.a('string');
              expect(data.search.groupby).to.be("");
              expect(data.search.columns).to.be.a('string');
              expect(data.search.columns).to.be("");
            expect(data.spec).to.be.a(Object);
            expect(data.spec).to.be();
            expect(data.created_at).to.be.a(Date);
            expect(data.created_at).to.be(new Date());
            expect(data.updated_at).to.be.a(Date);
            expect(data.updated_at).to.be(new Date());

            done();
          });
          */
          // TODO: uncomment and complete method invocation above, then delete this line and the next:
          done();
        });
      });
      describe('updateProjectDashboard', function() {
        it('should call updateProjectDashboard successfully', function(done) {
          // TODO: uncomment, update parameter values for updateProjectDashboard call and complete the assertions
          /*
          var owner = "owner_example";
          var project = "project_example";
          var dashboard_uuid = "dashboard_uuid_example";
          var body = new PolyaxonSdk.V1Dashboard();
          body.uuid = "";
          body.name = "";
          body.description = "";
          body.tags = [""];
          body.disabled = false;
          body.deleted = false;
          body.run_view = false;
          body.search = new PolyaxonSdk.V1SearchSpec();
          body.search.query = "";
          body.search.sort = "";
          body.search.limit = 0;
          body.search.groupby = "";
          body.search.columns = "";
          body.spec = ;
          body.created_at = new Date();
          body.updated_at = new Date();

          instance.updateProjectDashboard(owner, project, dashboard_uuid, body, function(error, data, response) {
            if (error) {
              done(error);
              return;
            }
            // TODO: update response assertions
            expect(data).to.be.a(PolyaxonSdk.V1Dashboard);
            expect(data.uuid).to.be.a('string');
            expect(data.uuid).to.be("");
            expect(data.name).to.be.a('string');
            expect(data.name).to.be("");
            expect(data.description).to.be.a('string');
            expect(data.description).to.be("");
            {
              let dataCtr = data.tags;
              expect(dataCtr).to.be.an(Array);
              expect(dataCtr).to.not.be.empty();
              for (let p in dataCtr) {
                let data = dataCtr[p];
                expect(data).to.be.a('string');
                expect(data).to.be("");
              }
            }
            expect(data.disabled).to.be.a('boolean');
            expect(data.disabled).to.be(false);
            expect(data.deleted).to.be.a('boolean');
            expect(data.deleted).to.be(false);
            expect(data.run_view).to.be.a('boolean');
            expect(data.run_view).to.be(false);
            expect(data.search).to.be.a(PolyaxonSdk.V1SearchSpec);
                  expect(data.search.query).to.be.a('string');
              expect(data.search.query).to.be("");
              expect(data.search.sort).to.be.a('string');
              expect(data.search.sort).to.be("");
              expect(data.search.limit).to.be.a('number');
              expect(data.search.limit).to.be(0);
              expect(data.search.groupby).to.be.a('string');
              expect(data.search.groupby).to.be("");
              expect(data.search.columns).to.be.a('string');
              expect(data.search.columns).to.be("");
            expect(data.spec).to.be.a(Object);
            expect(data.spec).to.be();
            expect(data.created_at).to.be.a(Date);
            expect(data.created_at).to.be(new Date());
            expect(data.updated_at).to.be.a(Date);
            expect(data.updated_at).to.be(new Date());

            done();
          });
          */
          // TODO: uncomment and complete method invocation above, then delete this line and the next:
          done();
        });
      });
    });
  });

}));