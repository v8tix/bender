#!/usr/bin/env bash

bender -port="${1}" -con-file="${2}" -db="${3}" -port-type="${4}" >>"${5}" 2>>"${6}"