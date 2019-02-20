# Make Java Keystore (from PEMs)

[![Build Status](https://travis-ci.org/bsycorp/make-jks.svg?branch=master)](https://travis-ci.org/bsycorp/make-jks)

A statically linked binary to take PEM-encoded certificates (like /etc/ssl/certs) and create a Java Keystore out of it (JKS). This is useful for java processes running in docker containers where `/etc/ssl/certs` is mounted into the container, but the Java keystore isn't updated to match.

## Run

`make-jks -input /etc/ssl/certs -output $JAVA_HOME/lib/security/cacerts`
