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

from django.conf import settings
from django.db import models

from coredb.abstracts.deleted import DeletedModel
from coredb.abstracts.describable import DescribableModel
from coredb.abstracts.diff import DiffModel
from coredb.abstracts.nameable import RequiredNameableModel
from coredb.abstracts.readme import ReadmeModel
from coredb.abstracts.tag import TagModel
from coredb.abstracts.uid import UuidModel


class Project(
    UuidModel,
    RequiredNameableModel,
    DiffModel,
    DescribableModel,
    TagModel,
    ReadmeModel,
    DeletedModel,
):
    user = models.ForeignKey(
        settings.AUTH_USER_MODEL,
        on_delete=models.SET_NULL,
        null=True,
        blank=True,
        related_name="projects",
    )
    owner = models.ForeignKey(
        "coredb.Owner", on_delete=models.CASCADE, related_name="projects"
    )
    is_public = models.BooleanField(
        default=True, help_text="If project is public or private."
    )

    class Meta:
        app_label = "coredb"
        db_table = "db_project"
        unique_together = (("owner", "name"),)
        indexes = [models.Index(fields=["name"])]