#!/bin/bash

# Copyright 2014-2016 Amazon.com, Inc. or its affiliates. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License"). You may not
# use this file except in compliance with the License. A copy of the License
# is located at
#
#     http://aws.amazon.com/apache2.0/
#
# or in the "license" file accompanying this file. This file is distributed
# on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
# express or implied. See the License for the specific language governing
# permissions and limitations under the License.

# This script will update the information of the ecs-init during building
set -e -x

CURRENTDIR=$(dirname "${0}")
version=$(cat "${CURRENTDIR}"/../ecs-init/ECSVERSION)
git_hash=$(git rev-parse --short=8 HEAD)
git_dirty=false

if [[ "$(git status --porcelain)" != "" ]]; then
	git_dirty=true
fi

cat << EOF > "${CURRENTDIR}"/../ecs-init/version/version.go
// This is an autogenerated file and should not be edited.

// Copyright 2014-2016 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Package version contains constants to indicate the current version of the
// ecs-init. It is autogenerated.
package version

// Version is the version of the ecs-init
var Version string = "${version}"

// GitDirty indicates the cleanliness of the git repo when this ecs-init was built
var GitDirty string = "${git_dirty}"

// GitShortHash is the short hash of this ecs-init build
var GitShortHash string = "${git_hash}"
EOF

cat << EOF > "${CURRENTDIR}"/../agent/version/version.go
// This is an autogenerated file and should not be edited.

// Copyright 2014-2016 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Package version contains constants to indicate the current version of the
// ecs-agent. It is autogenerated.
package version

// Version is the version of the ecs-agent
const Version = "${version}"

// GitDirty indicates the cleanliness of the git repo when this ecs-init was built
const GitDirty = ${git_dirty}

// GitShortHash is the short hash of this ecs-agent build
const GitShortHash = "${git_hash}"
EOF
