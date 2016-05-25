Go Wrappers for Sandstorm's API

Right now this contains:

* `capnp`, generated packages for the sandstorm API capnproto schema.
  The packages are currently generated from the schema in sandstorm
  v0.97
* `grain`, which provides a helper to connect to the sandstorm API.
* `websession`, a *very* WIP set of helpers around WebSession, most
  notably code to glue it to `net/http`.

Since the schema names are not all legal package names, the following
changes have been made:

* Schema with dashes in their names have had the dashes removed.
* The schema `package` has been mapped to the package `spk` (since
  `package` is a go reserved word).

This depends on a change to `zombiezen.com/go/capnproto2` that at the
time of writing has not been merged:

<https://github.com/zombiezen/go-capnproto2/pull/31>

To build these packages you will have to manually check out that branch
on my fork and put it in the appropriate place in `$GOPATH`.

# Licensing

Apache 2.0, both my stuff and the bits pulled from upstream:

    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.

A copy of the license is in the file `license.txt`. Individual source
files contain relevant copyright notices.
