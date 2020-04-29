#!/bin/bash

# ----------------------------------------------------------------
#   Copyright (c) ThoughtWorks, Inc.
#   Licensed under the Apache License, Version 2.0
#   See LICENSE in the project root for license information.
# ----------------------------------------------------------------*/

protoc --doc_out=markdown,gauge-proto-doc.md:doc *.proto