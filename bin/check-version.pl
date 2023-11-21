#!/usr/bin/env perl
use warnings;
use strict;

local $_ = $ENV{BINARY_VERSION};

die "BINARY_VERSION is not set\n" unless defined;

die "Version must be in the format X.Y.Z\n" unless /^\d+\.\d+\.\d+$/;

print "Version is in the correct format\n";
