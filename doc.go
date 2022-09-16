/*
Package clsgo provides some common used functions.
Use of this source code is governed by a MIT-style
license that can be found in the LICENSE file.

Installation

Install the whole package

    go get github.com/lovelacelee/clsgo

Or separately

	go get github.com/lovelacelee/clsgo/v1/config

Supported Features

Currently, clsgo supports the following quick-start functions.

    +-------------------------------------------------------------------+
    |pkg          | package                                             |
    |-------------|-----------------------------------------------------|
    |config       | github.com/lovelacelee/clsgo/v1/config              |
    |crlf         | github.com/lovelacelee/clsgo/v1/crlf                |
    |crypto       | github.com/lovelacelee/clsgo/v1/crypto              |
    |database     | github.com/lovelacelee/clsgo/v1/database            |
    |http         | github.com/lovelacelee/clsgo/v1/http                |
    |log          | github.com/lovelacelee/clsgo/v1/log                 |
    |mqtt         | github.com/lovelacelee/clsgo/v1/mqtt                |
    |net          | github.com/lovelacelee/clsgo/v1/net                 |
    |protobuf     | github.com/lovelacelee/clsgo/v1/protobuf            |
    |rabbitmq     | github.com/lovelacelee/clsgo/v1/rabbitmq            |
    |redis        | github.com/lovelacelee/clsgo/v1/redis               |
    |utils        | github.com/lovelacelee/clsgo/v1/utils               |
    |version      | github.com/lovelacelee/clsgo/v1/version             |
    |wraapper     | github.com/lovelacelee/clsgo/v1/wraapper            |
    +-------------------------------------------------------------------+

Usage

The CLSGO package is designed to function separately so that the
functionality of each subpackage is as independent as possible, if you
only need to use logging and configuration then log and config packages
are all you need.

	import (
		"github.com/lovelacelee/clsgo/v1/config"
		"github.com/lovelacelee/clsgo/v1/log"
	)

See the documentation of each pkg for more details.

*/
package clsgo

import (
	"github.com/lovelacelee/clsgo/v1/version"
)

var CLSGO_VERSION = version.Version
