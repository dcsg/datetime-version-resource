# Date Time Version Resource

Implements a [Concourse CI](https://concourse-ci.org/) resource type to generate a date time version to be used in your pipeline. 

## Source Configuration

* `timezone`: *Optional.* The Timezone in which you want the date generated. Default is `UTC`.

* `format`: *Optional.* The Format of the output Date Time version. Default is `20060102-150405`.

## Behavior

### `check`: None.

### `in`: Generates a `version` file containing the Date Time Version.

#### Parameters

*None.*

### `out`: Generates a Date Time version in the Timezone and Format defined.

#### Parameters

*None.*

## Example

```yaml
---
resource_types:
- name: datetime-version
  type: registry-image
  source:
    repository: dcsg/datetime-version-resource

resources:
- name: datetime-version
  type: datetime-version
  source:
    timezone: 'Europe/Lisbon' # default is 'UTC'
    format: '20060102.150405' # default is '20060102-150405'

- name: master-code
  type: git
  icon: git
  source:
    uri: https://github.com/dcsg/datetime-version-resource.git
    branch: master

jobs:
  - name: notify-bugsnag
    plan:
      - put: datetime-version
      - put: master-code
        params:
          // other git config params
          tag: datetime-version/version
          only_tag: true
```

### Contributing

Please make all pull requests to the `master` branch.
