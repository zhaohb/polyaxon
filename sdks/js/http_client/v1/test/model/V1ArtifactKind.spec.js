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

  describe('(package)', function() {
    describe('V1ArtifactKind', function() {
      beforeEach(function() {
        instance = PolyaxonSdk.V1ArtifactKind;
      });

      it('should create an instance of V1ArtifactKind', function() {
        // TODO: update the code to test V1ArtifactKind
        expect(instance).to.be.a('object');
      });

      it('should have the property model', function() {
        expect(instance).to.have.property('model');
        expect(instance.model).to.be("model");
      });

      it('should have the property audio', function() {
        expect(instance).to.have.property('audio');
        expect(instance.audio).to.be("audio");
      });

      it('should have the property video', function() {
        expect(instance).to.have.property('video');
        expect(instance.video).to.be("video");
      });

      it('should have the property histogram', function() {
        expect(instance).to.have.property('histogram');
        expect(instance.histogram).to.be("histogram");
      });

      it('should have the property image', function() {
        expect(instance).to.have.property('image');
        expect(instance.image).to.be("image");
      });

      it('should have the property tensor', function() {
        expect(instance).to.have.property('tensor');
        expect(instance.tensor).to.be("tensor");
      });

      it('should have the property dataframe', function() {
        expect(instance).to.have.property('dataframe');
        expect(instance.dataframe).to.be("dataframe");
      });

      it('should have the property chart', function() {
        expect(instance).to.have.property('chart');
        expect(instance.chart).to.be("chart");
      });

      it('should have the property csv', function() {
        expect(instance).to.have.property('csv');
        expect(instance.csv).to.be("csv");
      });

      it('should have the property tsv', function() {
        expect(instance).to.have.property('tsv');
        expect(instance.tsv).to.be("tsv");
      });

      it('should have the property psv', function() {
        expect(instance).to.have.property('psv');
        expect(instance.psv).to.be("psv");
      });

      it('should have the property ssv', function() {
        expect(instance).to.have.property('ssv');
        expect(instance.ssv).to.be("ssv");
      });

      it('should have the property metric', function() {
        expect(instance).to.have.property('metric');
        expect(instance.metric).to.be("metric");
      });

      it('should have the property env', function() {
        expect(instance).to.have.property('env');
        expect(instance.env).to.be("env");
      });

      it('should have the property html', function() {
        expect(instance).to.have.property('html');
        expect(instance.html).to.be("html");
      });

      it('should have the property text', function() {
        expect(instance).to.have.property('text');
        expect(instance.text).to.be("text");
      });

      it('should have the property file', function() {
        expect(instance).to.have.property('file');
        expect(instance.file).to.be("file");
      });

      it('should have the property dir', function() {
        expect(instance).to.have.property('dir');
        expect(instance.dir).to.be("dir");
      });

      it('should have the property dockerfile', function() {
        expect(instance).to.have.property('dockerfile');
        expect(instance.dockerfile).to.be("dockerfile");
      });

      it('should have the property docker_image', function() {
        expect(instance).to.have.property('docker_image');
        expect(instance.docker_image).to.be("docker_image");
      });

      it('should have the property data', function() {
        expect(instance).to.have.property('data');
        expect(instance.data).to.be("data");
      });

      it('should have the property coderef', function() {
        expect(instance).to.have.property('coderef');
        expect(instance.coderef).to.be("coderef");
      });

      it('should have the property table', function() {
        expect(instance).to.have.property('table');
        expect(instance.table).to.be("table");
      });

    });
  });

}));