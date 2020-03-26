#!/usr/bin/python
#
# Copyright 2018-2020 Polyaxon, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# coding: utf-8

"""
    Polyaxon SDKs and REST API specification.

    Polyaxon SDKs and REST API specification.  # noqa: E501

    OpenAPI spec version: 1.0.7
    Contact: contact@polyaxon.com
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""


import pprint
import re  # noqa: F401

import six


class V1GitType(object):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    """
    Attributes:
      swagger_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    swagger_types = {
        "url": "str",
        "revision": "bool",
        "connection": "str",
        "init": "bool",
    }

    attribute_map = {
        "url": "url",
        "revision": "revision",
        "connection": "connection",
        "init": "init",
    }

    def __init__(
        self, url=None, revision=None, connection=None, init=None
    ):  # noqa: E501
        """V1GitType - a model defined in Swagger"""  # noqa: E501

        self._url = None
        self._revision = None
        self._connection = None
        self._init = None
        self.discriminator = None

        if url is not None:
            self.url = url
        if revision is not None:
            self.revision = revision
        if connection is not None:
            self.connection = connection
        if init is not None:
            self.init = init

    @property
    def url(self):
        """Gets the url of this V1GitType.  # noqa: E501


        :return: The url of this V1GitType.  # noqa: E501
        :rtype: str
        """
        return self._url

    @url.setter
    def url(self, url):
        """Sets the url of this V1GitType.


        :param url: The url of this V1GitType.  # noqa: E501
        :type: str
        """

        self._url = url

    @property
    def revision(self):
        """Gets the revision of this V1GitType.  # noqa: E501


        :return: The revision of this V1GitType.  # noqa: E501
        :rtype: bool
        """
        return self._revision

    @revision.setter
    def revision(self, revision):
        """Sets the revision of this V1GitType.


        :param revision: The revision of this V1GitType.  # noqa: E501
        :type: bool
        """

        self._revision = revision

    @property
    def connection(self):
        """Gets the connection of this V1GitType.  # noqa: E501


        :return: The connection of this V1GitType.  # noqa: E501
        :rtype: str
        """
        return self._connection

    @connection.setter
    def connection(self, connection):
        """Sets the connection of this V1GitType.


        :param connection: The connection of this V1GitType.  # noqa: E501
        :type: str
        """

        self._connection = connection

    @property
    def init(self):
        """Gets the init of this V1GitType.  # noqa: E501


        :return: The init of this V1GitType.  # noqa: E501
        :rtype: bool
        """
        return self._init

    @init.setter
    def init(self, init):
        """Sets the init of this V1GitType.


        :param init: The init of this V1GitType.  # noqa: E501
        :type: bool
        """

        self._init = init

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.swagger_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(
                    map(lambda x: x.to_dict() if hasattr(x, "to_dict") else x, value)
                )
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(
                    map(
                        lambda item: (item[0], item[1].to_dict())
                        if hasattr(item[1], "to_dict")
                        else item,
                        value.items(),
                    )
                )
            else:
                result[attr] = value
        if issubclass(V1GitType, dict):
            for key, value in self.items():
                result[key] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, V1GitType):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other