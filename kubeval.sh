#!/bin/sh

ls *.yml | grep -v 'values.yml$' | xargs kubeval --strict
