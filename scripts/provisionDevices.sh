#!/bin/sh

# Provision 10 'brewer' devices in the application targeted by simulation 'sim1'.
curl --location --request POST 'http://localhost:6001/api/simulation/sim1/provision/brewer/10'
